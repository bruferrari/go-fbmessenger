package main

import macaron "gopkg.in/macaron.v1"
import "github.com/bruferrari/go-fbmessenger/fblib"

func macaronSetUp() {
	m := macaron.Classic()

	m.Get("/", func() string {
		return "1951683431"
	})

	m.Run()
}

func main() {
	fblib.DefineAccessToken("")
	macaronSetUp()
}
