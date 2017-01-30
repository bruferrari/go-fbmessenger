package main

import (
	"log"

	"github.com/bruferrari/go-fbmessenger/fblib"
	macaron "gopkg.in/macaron.v1"
)

func macaronSetUp() {
	m := macaron.Classic()

	m.Get("/", func(ctx *macaron.Context) {
		url := ctx.Req.URL
		params := url.Query()

		log.Println(url)
		log.Println(params["hub.mode"])

		if params["hub.mode"][0] == "subscribe" {
			log.Println("HELL YEAH!")
		}
		log.Println(":(")
	})

	m.Run()
}

func main() {
	fblib.DefineAccessToken("")
	macaronSetUp()
}
