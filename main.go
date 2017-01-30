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

	m.Get("/", rootHandler)

	m.Run()
}

func rootHandler(ctx *macaron.Context) (int, string) {
	var resp string
	if ok, resp := webhook(ctx); ok != false {
		return http.StatusOK, resp
	}
	return http.StatusUnauthorized, resp
}

func webhook(ctx *macaron.Context) (bool, string) {
	url := ctx.Req.URL
	params := url.Query()

	log.Println(url)
	log.Println(params["hub.mode"])

	if params["hub.mode"][0] == "subscribe" {
		if params["hub.verify_token"][0] == os.Getenv("VERIFY_TOKEN") {
			log.Println("Verification token mismatch")
			return true, params["hub.challenge"][0]
		}
	}
	return false, ""
}

func main() {
	fblib.DefineAccessToken("")
	macaronSetUp()
}
