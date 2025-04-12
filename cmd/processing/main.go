package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"govtech-hackaton-backend/cmd/processing/db"
	"govtech-hackaton-backend/cmd/processing/validation"
	"log"
)

func main() {
	log.Println("Initializing...")

	db.Init()

	log.Printf("Server started...")

	err := fasthttp.ListenAndServe(":8080", requestHandler)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/":
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBody([]byte("The server is working"))
		return
	case "/process-request-bovine":
		validation.HandlerProcessRequestBovine(ctx)
		return
	}
}
