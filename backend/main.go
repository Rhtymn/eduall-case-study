package main

import (
	"fmt"
	"os"

	"github.com/Rhtymn/eduall-case-study/backend/config"
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

	fmt.Println(conf)
}
