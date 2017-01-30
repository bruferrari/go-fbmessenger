package main

import (
	"github.com/bruferrari/go-fbmessenger/fblib"
	macaron "gopkg.in/macaron.v1"
)

const accessToken = "EAAIN6p1IjksBAHewejE1JillvIC7dU3C8hReGJxu4hZAtNt3syBBMoD3pRYpbrsZANnzcsvqbBUWtIJA2AdMp3J7WnNxAWZAoHIxUWDUq0FZBhl2RHMqGoBmemUz3iJqF37JTvDD8BcEkSZCKjluYqSlftZAfhc7Apk8cZCIFnxZCQZDZD"

func macaronSetUp() {
	m := macaron.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Run()
}

func main() {
	fblib.DefineAccessToken(accessToken)
	macaronSetUp()
}
