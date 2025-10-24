package main

import (
	"errors"
	"fmt"
	"net/http"

	"go-adv-demo/internal/auth"
	"go-adv-demo/internal/config"
	"go-adv-demo/internal/link"
	"go-adv-demo/pkg/db"
	"go-adv-demo/pkg/middleware"
)

func main() {
	router := http.NewServeMux()

	conf := config.LoadConfig()
	database := db.NewDb(conf)

	// Repositories
	linkRepo := link.NewLinkRepository(database)

	// Handlers
	auth.NewAuthHandler(router)
	link.NewLinkHandler(
		router,
		linkRepo,
	)

	// Middlewares
	corsMiddleware := func(next http.Handler) http.Handler {
		return middleware.CORS(next, conf.Cors.AllowedOrigins, conf.Cors.AllowCredentials)
	}
	stack := middleware.Chain(
		corsMiddleware,
		middleware.Logging,
	)

	port := 8080
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: stack(router),
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
