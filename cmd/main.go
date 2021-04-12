package main

import "github.com/lmx-Hexagram/gemini-server/internal/service"

func main() {
	//engine, err := gemini.New("MyCertificate.crt", "MyKey.key")
	//if err != nil {
	//	panic(err)
	//}
	//
	//engine.Handle("/dir/test", func(c *gemini.Context) {
	//	c.Render(20, "aaa\n#hello gemini!!!")
	//})
	//
	//engine.Handle("/app", func(c *gemini.Context) {
	//	c.Render(20, "gemini server write by hexagram")
	//})
	//
	//engine.HandleDir("/dir", "dir")
	//
	//engine.HandleFile("/file", "test.gemini")
	//
	//engine.HandleProxy("/url","http://127.0.0.1:12222")
	//
	//err = engine.Run(":1965")
	//if err != nil {
	//	panic(err)
	//}
	service.Init()
	service.Run()
}
