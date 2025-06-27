package main

import (
	"auth/internal/configs"
	"auth/internal/handlers"
	"auth/internal/repository"
	"fmt"
	"net/http"
)

func main() {
	cfg := configs.GetConfig()
	db, err := configs.InitDb(cfg)
	if err != nil {
		fmt.Errorf("error initializing db %v", err)
	}
	psqlClient, err := repository.NewPostgreSQLCartRepository(db)
	if err != nil {
		fmt.Errorf("error initializing db %v", err)
	}
	r := handlers.Manager(nil, cfg, psqlClient)
	port := ":3000"
	fmt.Printf("Server listening on port %s\n", port)
	err = http.ListenAndServe(port, r)
	if err != nil {
		fmt.Errorf("error initializing server %v", err)
	}

}
