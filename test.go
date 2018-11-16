package main

import (
	"log"
	"net/http"

	"github.com/claimh-solais/go-sessions/session"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// gc, err := config.Init()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	router := httprouter.New()

	options := &session.MiddlewareOptions{
		IsResave:            session.DISABLE_RESAVE,
		IsSaveUninitialized: session.DISABLE_SAVE_UNINITIALIZED,
		Secret:              &[]string{"mysecret"},
		// Cookie: &session.HTTPCookie{
		// 	Name:  "jshdkajs",
		// 	Value: "",
		// },
	}
	middleware, _ := session.NewMiddleware(router, options)
	// router.GET("/health", core.ReportHealth())

	log.Fatal(http.ListenAndServe(":18080", middleware))
}
