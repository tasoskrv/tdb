package main

import (
	"log"

	"../api"
)

func main() {
	con, f, err := InitDatabase()

	if err != nil {
		log.Fatal(err)
	}

	api.Start(con.client, con.db)

	f()
}
