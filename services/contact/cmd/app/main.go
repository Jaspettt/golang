package main

import (
	"architecture_go/pkg/store/postgres"
	"architecture_go/services/contact/internal"
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	conn, err := postgres.New(postgres.Settings{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "1488",
		DBName:   "cleanArc",
	})
	if err != nil {
		panic(err)
	}
	defer func(Pool *sql.DB) {
		err := Pool.Close()
		if err != nil {

		}
	}(conn.Pool)

	fmt.Println("Connected to postgres database")

	contactRepo := internal.NewRepository(conn.Pool)
	contactUseCase := internal.NewUseCase(contactRepo)
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	contactDelivery := internal.NewDelivery(contactUseCase, logger)

	http.Handle("/contacts", contactDelivery)
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal("HTTP server error: ", err)
		}
	}()

	fmt.Println("HTTP server started on port 8080")

	select {}
}
