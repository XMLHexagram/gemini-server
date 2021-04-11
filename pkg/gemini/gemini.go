package gemini

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"go-gemini-server/pkg/debug"
	"go-gemini-server/pkg/statusCode"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

func New(certFile, keyFile, DefaultLang string) (engine *Engine, err error) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return
	}

	engine = &Engine{
		Cert:        cert,
		RouterMap:   make(map[string]HandlerFunc),
		DefaultLang: DefaultLang,
	}
	return
}

func (e *Engine) Handle(router string, f HandlerFunc) {
	e.RouterMap[router] = f
	debug.PrintRoute(router)
}

func (e *Engine) HandleDir(router string, dirPath string, Index string) {
	a := func(c *Context) {
		lenRouter := len(router)
		path_ := c.URL.Path[lenRouter:]
		fmt.Println(path_)
		file, err := ioutil.ReadFile(path.Join(dirPath, path_))
		if err != nil {
			panic(err)
		}
		c.Render(20, string(file))
	}
	if Index != "" {
		_, err := os.Stat(path.Join(dirPath, Index))
		if err != nil {
			panic(err)
		} else {
			e.HandleFile(router, path.Join(dirPath, Index))
		}
	}
	e.RouterMap[router+"/*"] = a
	debug.PrintLoadDir(router+"/*", dirPath)
}

func (e *Engine) HandleFile(router string, filePath string) {
	a := func(c *Context) {
		lenRouter := len(router)
		path_ := c.URL.Path[lenRouter-1:]
		fmt.Println(path_)
		file, err := ioutil.ReadFile(filePath)
		if err != nil {
			panic(err)
		}
		c.Render(20, string(file))
	}
	e.RouterMap[router] = a
	debug.PrintLoadFile(router, filePath)
}

func (e *Engine) HandleProxy(router string, url string) {
	a := func(c *Context) {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			panic(err)
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		bodyMap := make(map[string]interface{})
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(body, &bodyMap)
		if err != nil {
			panic(err)
		}
		c.Render(20, bodyMap["body"].(string))
	}
	e.RouterMap[router] = a
	debug.PrintProxy(router, url)
}

func (e *Engine) Run(addr string) (err error) {
	tlscfg := &tls.Config{
		Certificates: []tls.Certificate{e.Cert},
		MinVersion:   tls.VersionTLS13,
	}

	listener, err := tls.Listen("tcp", addr, tlscfg)
	if err != nil {
		return fmt.Errorf("tls.Listen:%w", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		go e.ServeGemini(conn)
	}

}

func (e *Engine) ServeGemini(conn net.Conn) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("[panic] ", r)
			writeErrorResponseHeader(int(statusCode.PermanentFailure), "unhandled crash error", conn)
		}
	}()
	s := bufio.NewScanner(conn)
	s.Scan()
	u, err := url.Parse(s.Text())
	if err != nil {
		writeErrorResponseHeader(int(statusCode.BadRequest), "invalid url", conn)
	}
	//pretty.Println(u)

	c := &Context{
		Conn: conn,
		URL:  u,
		Keys: make(map[string]interface{}),
		Lang: e.DefaultLang,
	}
	e.handleRequest(c)
}

func (e *Engine) handleRequest(c *Context) {
	var F HandlerFunc
	for k, f := range e.RouterMap {
		if k == c.URL.Path {
			F = f
			break
		}
		if isMatch, _ := filepath.Match(k, c.URL.Path); isMatch {
			F = f
		}
	}
	if F != nil {
		F(c)
	}
	if e.AutoRedirect {
		writeErrorResponseHeader(30, e.AutoRedirectUrl, c.Conn)
	}
	writeErrorResponseHeader(51, "", c.Conn)

	return
}

type Engine struct {
	Cert            tls.Certificate
	RouterMap       map[string]HandlerFunc
	DefaultLang     string
	AutoRedirect    bool
	AutoRedirectUrl string
}

type HandlerFunc func(*Context)

type Context struct {
	Conn net.Conn
	URL  *url.URL
	Keys map[string]interface{}
	Lang string
}

func writeErrorResponseHeader(code int, meta string, conn net.Conn) {
	conn.Write([]byte(strconv.Itoa(code) + " " + meta + "\r\n"))
	conn.Close()
}

func (c *Context) Render(code int, body string) {
	defer c.Conn.Close()
	writeResponseHeader(code, c.Lang, c.Conn)
	c.Conn.Write([]byte(body))
}

func writeResponseHeader(code int, lang string, conn net.Conn) {
	meta := "text/gemini; lang=" + lang + "; charset=utf-8"
	_, err := conn.Write([]byte(strconv.Itoa(code) + " " + meta + "\r\n"))
	if err != nil {
		panic(err)
	}
}
