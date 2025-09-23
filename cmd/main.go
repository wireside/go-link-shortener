package main

import (
	"errors"
	"fmt"
	"net/http"

	"go-adv-demo/configs"
	"go-adv-demo/internal/auth"
	"go-adv-demo/internal/link"
	"go-adv-demo/pkg/db"
)

func main() {
	router := http.NewServeMux()

	conf := configs.LoadConfig()
	_ = db.NewDb(conf)

	// Handlers
	auth.NewAuthHandler(
		router, &auth.AuthHandlerDeps{
			Config: conf,
		},
	)
	link.NewLinkHandler(
		router, &link.LinkHandlerDeps{
			Config: conf,
		},
	)

	port := 8080
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	fmt.Printf("Server is listening on port %d\n", port)
	err := server.ListenAndServe()
	if err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Println("Server is closed")
			return
		}
		fmt.Println(err.Error())
	}
}
