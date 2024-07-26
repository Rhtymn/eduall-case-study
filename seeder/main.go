package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/Rhtymn/eduall-case-study/seeder/config"
	"github.com/Rhtymn/eduall-case-study/seeder/database"
	"github.com/Rhtymn/eduall-case-study/seeder/repository"
	"github.com/Rhtymn/eduall-case-study/seeder/util"
)

func main() {
	err := config.Init()
	if err != nil {
		fmt.Printf("Error env initialization : %s\n", err.Error())
		os.Exit(1)
	}

	conf, err := config.Load()
	if err != nil {
		fmt.Printf("Error loading .env : %s\n", err.Error())
		os.Exit(1)
	}

	db, err := database.ConnectPostgreDB(conf.PostgreURL)
	if err != nil {
		fmt.Printf("Error connecting to DB : %s\n", err.Error())
		os.Exit(1)
	}

	productRepository := repository.NewProductRepository(db)

	file, err := os.Open("amazon_laptop_prices_v01.csv")
	if err != nil {
		fmt.Printf("Error opening dataset : %s\n", err.Error())
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("error reading records")
		os.Exit(1)
	}

	products, err := util.ToProducts(records[1:])
	if err != nil {
		fmt.Printf("error converting record to product entity : %s\n", err.Error())
		os.Exit(1)
	}

	err = productRepository.BatchSave(products)
	if err != nil {
		fmt.Printf("error inserting products : %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("successfully seeding products %d rows", len(products))
}
