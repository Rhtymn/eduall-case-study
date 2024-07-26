package main

import (
	"fmt"
	"os"

	"github.com/Rhtymn/eduall-case-study/backend/config"
	"github.com/Rhtymn/eduall-case-study/backend/database"
)

func main() {
	err := config.Init()
	if err != nil {
		fmt.Printf("Error env initialization : %s", err.Error())
		os.Exit(1)
	}

	conf, err := config.Load()
	if err != nil {
		fmt.Printf("Error load env : %s", err.Error())
		os.Exit(1)
	}

	db, err := database.ConnectPostgreDB(conf.PostgreURL)
	if err != nil {
		fmt.Printf("Error connecting to DB : %s", err.Error())
		os.Exit(1)
	}

	fmt.Println(db)
}
