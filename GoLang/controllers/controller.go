package controllers

import (
	"GoLang/repository"
	"encoding/json"
	"fmt"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, World! => Hitted")
	err := json.NewEncoder(w).Encode(repository.Response{Msg: "Hello, World!"})
	if err != nil {
		return
	}
}
