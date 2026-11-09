package main

import (
	"context"
	"fmt"
	"os"

	"grapevine.boobles.cloud/database"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	dbHandler, err := database.CreateNewDbHandler()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dbService := database.CreateNewDbService(dbHandler)

}
