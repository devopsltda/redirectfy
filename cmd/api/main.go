package main

import (
	"fmt"
	"github.com/TheDevOpsCorp/redirectify/internal/server"
)

func main() {

	server := server.NewServer()

	err := server.ListenAndServe()

	if err != nil {
		panic(fmt.Sprintf("O servidor não pôde ser iniciado: %s", err))
	}
}
