package gemini

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"github.com/kr/pretty"
	"net"
	"net/url"
	"strconv"
)

func New(certFile, keyFile string) (engine *Engine, err error) {
	cert, err := tls.LoadX509KeyPair("MyCertificate.crt", "MyKey.key")
	if err != nil {
		return
	}

	engine = &Engine{
		Cert:      cert,
		RouterMap: make(map[string]HandlerFunc),
	}
	return
}

func (e *Engine) Handle(router string, f HandlerFunc) {
	e.RouterMap[router] = f
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
	s := bufio.NewScanner(conn)
	s.Scan()
	u, err := url.Parse(s.Text())
	if err != nil {
		panic(err)
	}
	pretty.Println(u)

	c := &Context{
		Conn: conn,
		URL:  u,
		Keys: make(map[string]interface{}),
	}
	e.handleRequest(c)
}

func (e *Engine) handleRequest(c *Context) {
	f := e.RouterMap[c.URL.Path]
	f(c)
}

type Engine struct {
	Cert      tls.Certificate
	RouterMap map[string]HandlerFunc
}

type HandlerFunc func(*Context)

type Context struct {
	Conn net.Conn
	URL  *url.URL
	Keys map[string]interface{}
}

func (c *Context) Render(code int, body string) {
	writeResponseHeader(code, c.Conn)
	c.Conn.Write([]byte(body))
	c.Conn.Close()
}

func writeResponseHeader(code int, conn net.Conn) {
	meta := "text/gemini; lang=en; charset=utf-8"
	_, err := conn.Write([]byte(strconv.Itoa(code) + " " + meta + "\r\n"))
	if err != nil {
		panic(err)
	}
}
