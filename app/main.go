package main

import (
	"github.com/jackc/pgx/v4"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"	
	"context"
	"fmt"
	"os"
)

type User struct {
	id uint64
	name string
}

func main() {
	db, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close(context.Background())

	query := `
CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL 
);
	`

	_, err = db.Exec(context.Background(), query)

// 	query = `
// INSERT INTO users(name) VALUES
// ('Jhon'),
// ('Jack'),
// ('Kate');
// 	`

// 	_, err = db.Exec(context.Background(), query)

	router := routing.New()
	router.Get("/", func (c *routing.Context) error {
		var ids int32
		var names string

		err := db.QueryRow(context.Background(), "SELECT * FROM users").Scan(&ids, &names)
		if err != nil {
			return err
		}

		fmt.Fprintf(c, "ids: %v; names: %v", ids, names)
		return nil
	})
	fmt.Println(">>> Server is on 8080")
	panic(fasthttp.ListenAndServe("0.0.0.0:8080", router.HandleRequest))
}
