package domain

import (
	"context"
	"database/sql"
)

type Product struct {
	ID                 int64
	Brand              string
	Model              string
	ScreenSize         string
	Color              string
	HardDisk           string
	CPU                string
	RAM                string
	OS                 string
	SpecialFeatures    string
	Graphics           string
	GraphicCoprocessor string
	CpuSpeed           string
	Rating             float64
	Price              float64
}

type NullableProduct struct {
	Model              sql.NullString
	ScreenSize         sql.NullString
	Color              sql.NullString
	HardDisk           sql.NullString
	CPU                sql.NullString
	RAM                sql.NullString
	OS                 sql.NullString
	SpecialFeatures    sql.NullString
	Graphics           sql.NullString
	GraphicCoprocessor sql.NullString
	CpuSpeed           sql.NullString
	Rating             sql.NullFloat64
	Price              sql.NullFloat64
}

func MergeProduct(p Product, np NullableProduct) Product {
	if np.Model.Valid {
		p.Model = np.Model.String
	}

	if np.ScreenSize.Valid {
		p.ScreenSize = np.ScreenSize.String
	}

	if np.Color.Valid {
		p.Color = np.Color.String
	}

	if np.HardDisk.Valid {
		p.HardDisk = np.HardDisk.String
	}

	if np.CPU.Valid {
		p.CPU = np.CPU.String
	}

	if np.RAM.Valid {
		p.RAM = np.RAM.String
	}

	if np.OS.Valid {
		p.OS = np.OS.String
	}

	if np.SpecialFeatures.Valid {
		p.SpecialFeatures = np.SpecialFeatures.String
	}

	if np.Graphics.Valid {
		p.Graphics = np.Graphics.String
	}

	if np.CpuSpeed.Valid {
		p.CpuSpeed = np.CpuSpeed.String
	}

	if np.Rating.Valid {
		p.Rating = np.Rating.Float64
	}

	if np.Price.Valid {
		p.Price = np.Price.Float64
	}

	return p
}

type ProductQuery struct {
	SearchTerm string
	Page       int
	Limit      int
}

type ProductRepository interface {
	GetAll(ctx context.Context, q ProductQuery) ([]Product, error)
	GetPageInfo(ctx context.Context, q ProductQuery) (PageInfo, error)
}

type ProductService interface {
	GetAll(ctx context.Context, q ProductQuery) ([]Product, PageInfo, error)
}
