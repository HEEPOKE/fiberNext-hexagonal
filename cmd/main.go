package main

import (
	"fmt"
	"log"

	"github.com/HEEPOKE/fiber-hexagonal/internals/server"
	"github.com/HEEPOKE/fiber-hexagonal/pkg/configs"
	"github.com/HEEPOKE/fiber-hexagonal/pkg/database"
)

// @title Swagger Example API
// @title Go Fiber Hexagonal API
// @version 1.0
// @description This is a Go Fiber Hexagonal API server.
// @contact.name Heepoke
// @contact.url https://github.com/HEEPOKE
// @contact.email Damon1FX@gmail.com
// @host localhost:6476
// @BasePath /apis
// @schemes http https
func main() {
	_, err := configs.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	_, err = database.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	address := fmt.Sprintf(":%s", configs.Cfg.PORT)
	route := server.NewServer()
	route.Init(address)
}
