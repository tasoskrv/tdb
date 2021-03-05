package main

import (
	"log"

	"../api"
)

func main() {

	con, f, err := InitDatabase()
	defer f()

	if err != nil {
		log.Fatal(err)
	}

	api.Start(con.client, con.db)
}
