package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Rhtymn/eduall-case-study/seeder/entity"
)

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *productRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) BatchSave(products []entity.Product) error {
	var sb strings.Builder
	sb.WriteString(`INSERT INTO 
		products(brand, model, screen_size, color, harddisk, cpu, ram, os, special_features, graphics, graphics_coprocessor, cpu_speed, rating, price)
	VALUES `)

	args := make([]interface{}, 14*len(products))

	for i := 0; i < len(products); i++ {
		p := products[i]

		sb.WriteString(
			fmt.Sprintf(
				`($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)`,
				1+(i*14), 2+(i*14), 3+(i*14), 4+(i*14),
				5+(i*14), 6+(i*14), 7+(i*14), 8+(i*14),
				9+(i*14), 10+(i*14), 11+(i*14), 12+(i*14),
				13+(i*14), 14+(i*14),
			),
		)

		if i != len(products)-1 {
			sb.WriteString(",")
		}

		args[0+(i*14)] = p.Brand
		args[1+(i*14)] = p.Model
		args[2+(i*14)] = p.ScreenSize
		args[3+(i*14)] = p.Color
		args[4+(i*14)] = p.HardDisk
		args[5+(i*14)] = p.CPU
		args[6+(i*14)] = p.RAM
		args[7+(i*14)] = p.OS
		args[8+(i*14)] = p.SpecialFeatures
		args[9+(i*14)] = p.Graphics
		args[10+(i*14)] = p.GraphicCoprocessor
		args[11+(i*14)] = p.CpuSpeed
		args[12+(i*14)] = p.Rating
		args[13+(i*14)] = p.Price
	}

	_, err := r.db.Exec(sb.String(), args...)
	if err != nil {
		return err
	}

	return nil
}
