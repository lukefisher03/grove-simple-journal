package main

import (
	"context"
	"fmt"
	"os"
	"github.com/jackc/pgx/v5"
)

func showError(err error) {
	fmt.Println("Uh oh, could not connect")
	fmt.Fprintf(os.Stderr, "Unable to connect to db %v\n", err)
	os.Exit(1)
}

func main() {
	con, err := pgx.Connect(context.Background(), "postgres://lukefisher:postgres@localhost:5432/grove")
	if err != nil {
		showError(err)
	}

	defer con.Close(context.Background())
	err = pgx.BeginFunc(context.Background(), con, func(tx pgx.Tx) error {
		_, err := tx.Exec(context.Background(), "INSERT INTO notes (content, title) VALUES('This is so cool bro', 'This is a sample!');")
		return err
	})

	if err != nil {
		showError(err)
	}
}