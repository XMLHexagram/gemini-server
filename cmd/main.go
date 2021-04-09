package main

import "go-gemini-server/pkg/gemini"

func main() {
	engine, err := gemini.New("MyCertificate.crt", "MyKey.key")
	if err != nil {
		panic(err)
	}

	engine.Handle("/", func(c *gemini.Context) {
		c.Render(20, "aaa\n#hello gemini!!!")
	})

	engine.Handle("/app", func(c *gemini.Context) {
		c.Render(20, "gemini server write by hexagram")
	})

	engine.HandleDir("/dir/*", "dir")

	engine.HandleFile("/file", "test.gemini")
	err = engine.Run(":1965")
	if err != nil {
		panic(err)
	}

}
