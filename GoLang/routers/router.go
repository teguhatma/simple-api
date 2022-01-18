package routers

import (
	"GoLang/controllers"
	"log"
	"net/http"
)

func HandleRequests() {
	http.HandleFunc(
		"/api/hello-world",
		controllers.HelloWorld,
	)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
