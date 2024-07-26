package util

import (
	"strconv"
	"strings"

	"github.com/Rhtymn/eduall-case-study/seeder/entity"
)

func ToProducts(data [][]string) ([]entity.Product, error) {
	products := make([]entity.Product, len(data))

	for i := 0; i < len(data); i++ {
		p := entity.Product{}
		p.Brand = data[i][0]
		p.Model = data[i][1]
		p.ScreenSize = data[i][2]
		p.Color = data[i][3]
		p.HardDisk = data[i][4]
		p.CPU = data[i][5]
		p.RAM = data[i][6]
		p.OS = data[i][7]
		p.SpecialFeatures = data[i][8]
		p.Graphics = data[i][9]
		p.GraphicCoprocessor = data[i][10]
		p.CpuSpeed = data[i][11]

		var err error
		var rating float64
		if data[i][12] != "" {
			rating, err = strconv.ParseFloat(strings.Replace(data[i][12], ",", "", -1), 64)
			if err != nil {
				return nil, err
			}
		}
		p.Rating = rating

		var price float64
		if data[i][13] != "" {
			price, err = strconv.ParseFloat(strings.TrimSpace(strings.Replace(data[i][13][1:], ",", "", -1)), 64)
			if err != nil {
				return nil, err
			}
		}
		p.Price = price

		products[i] = p
	}

	return products, nil
}
