package main

import "go-gemini-server/pkg/gemini"

//type Handler interface {
//	ServeGemini(*Response, *Request)
//}
//
//type Server struct {
//	Addr string
//	//Handler      Handler
//	ReadTimeout  time.Duration
//	WriteTimeout time.Duration
//	listener     net.Listener
//	Log          io.Writer
//}

func main() {
	engine, err := gemini.New("MyCertificate.crt", "MyKey.key")
	if err != nil {
		panic(err)
	}

	engine.Handle("/", func(c *gemini.Context) {
		c.Render(20,"aaa\n#hello gemini!!!")
	})

	engine.Handle("/app", func(c *gemini.Context) {
		c.Render(20,"gemini server write by hexagram")
	})

	err = engine.Run(":1965")
	if err != nil {
		panic(err)
	}
	//s := new(Server)
	//s.Addr = ":12210"
	//cert, err := tls.LoadX509KeyPair("MyCertificate.crt", "MyKey.key")
	//if err != nil {
	//	panic(err)
	//	return
	//}
	//
	////fmt.Printf("%+v\n",cert)
	//tlscfg := &tls.Config{
	//	Certificates: []tls.Certificate{cert},
	//	MinVersion:   tls.VersionTLS13,
	//}
	//
	//listener, err := tls.Listen("tcp", s.Addr, tlscfg)
	//if err != nil {
	//	panic(err)
	//	return
	//}
	//defer listener.Close()
	//s.listener = listener
	//
	//for {
	//	conn, err := listener.Accept()
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//
	//	fmt.Println(conn.LocalAddr(),conn.RemoteAddr())
	//	meta := "text/gemini; lang=en; charset=utf-8"
	//	header := "20 " + meta + "\r\n"
	//	_, err = conn.Write([]byte(header))
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	_, err = conn.Write([]byte("#hello world!\r\n#hello gemini!\r\n"))
	//	if err != nil {
	//		panic(err)
	//	}
	//	s:=bufio.NewScanner(conn)
	//	s.Scan()
	//	if err != nil {
	//		return
	//	}
	//	fmt.Println(s.Text())
	//
	//	conn.Close()
	//}
	//service.Init()
	//service.Run()
}
