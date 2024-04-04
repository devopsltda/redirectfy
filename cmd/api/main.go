package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"redirectify/internal/server"
	"redirectify/internal/services/database"
	"syscall"
	"time"
)

func main() {

	server := server.NewServer()

	go func() {
		if err := server.ListenAndServe(); err != nil {
			panic(fmt.Sprintf("O servidor não pôde ser iniciado: %s", err))
		}
	}()

	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT, os.Interrupt, os.Kill)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println("Parando...")

	if err := server.Shutdown(ctx); err != nil {
		panic(fmt.Sprintf("O servidor não pôde ser finalizado graciosamente: %s", err))
	}

	_ = os.RemoveAll(database.TempDir)
	fmt.Println("O servidor parou.")
}
