package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"redirectfy/internal/server"
	"syscall"
	"time"
)

// Essa função serve de bootstrap para a aplicação, iniciando o servidor e
// finalizando ele graciosamente caso ele seja parado de alguma forma.
//
// O processo de finalização graciosa espera 10 segundo para a finalização
// dos processos existentes, e então remove diretórios temporários, como
// os que contém as réplicas do banco de dados, e fecha as conexões
// existentes.
func main() {
	server := server.NewServer()

	slog.Info("Servidor iniciado", slog.String("addr", "http://localhost"+server.Addr))

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

	fmt.Println("O servidor parou.")
}
