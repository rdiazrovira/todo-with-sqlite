package main

import (
	"fmt"
	"net/http"
	"os"

	"todox-with-sqlite/internal"

	"github.com/leapkit/core/envor"
	"github.com/leapkit/core/server"
)

func main() {
	s := server.New(
		server.WithHost(envor.Get("HOST", "0.0.0.0")),
		server.WithPort(envor.Get("PORT", "3000")),
	)

	if err := internal.AddServices(s); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := internal.AddRoutes(s); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Server started at", s.Addr())
	err := http.ListenAndServe(s.Addr(), s.Handler())
	if err != nil {
		fmt.Println(err)
	}
}
