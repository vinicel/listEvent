package main

import (
	"github.com/vinicel/listEvent/internal/server"
)

func main() {
	router := server.NewServer()
	router.Start()
}
