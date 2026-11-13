package main

import (
	"context"
	"fmt"
	"os"

	"grapevine.boobles.cloud/database"
	"grapevine.boobles.cloud/startup"
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

	secretmanager, err := startup.ConfigureSecretManagerWitPw(ctx)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	secretmanager.SetupCache(ctx)
}
