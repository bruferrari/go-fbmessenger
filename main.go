package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bruferrari/go-fbmessenger/fblib"
	macaron "gopkg.in/macaron.v1"
)

func macaronSetUp() {
	m := macaron.Classic()

	m.Get("/", func(ctx *macaron.Context) (int, string) {
		url := ctx.Req.URL
		params := url.Query()

		log.Println(url)
		log.Println(params["hub.mode"])

		if params["hub.mode"][0] == "subscribe" {
			if params["hub.verify_token"][0] != os.Getenv("VERIFY_TOKEN") {
				log.Println("403", "Verification token mismatch")
				return 403, "Verification token mismatch"
			}
			return http.StatusOK, params["hub.challenge"][0]
		}
		return http.StatusBadRequest, "Bad request"
	})

	m.Run()
}

func main() {
	fblib.DefineAccessToken("")
	macaronSetUp()
}
