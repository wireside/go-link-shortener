package main

import (
	"errors"
	"fmt"
	"net/http"

	"go-adv-demo/configs"
	"go-adv-demo/internal/auth"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	auth.NewAuthHandler(
		router, &auth.AuthHandlerDeps{
			Config: conf,
		},
	)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Printf("Server is listening on port %d\n", 8080)
	err := server.ListenAndServe()
	if err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Println("Server is closed")
			return
		}
		fmt.Println(err.Error())
	}
}
