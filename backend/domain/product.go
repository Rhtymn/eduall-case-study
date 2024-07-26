package domain

import "context"

type Product struct {
	ID                 int64
	Brand              string
	Model              *string
	ScreenSize         string
	Color              *string
	HardDisk           *string
	CPU                *string
	RAM                *string
	OS                 *string
	SpecialFeatures    *string
	Graphics           *string
	GraphicCoprocessor *string
	CpuSpeed           *string
	Rating             *float64
	Price              *float64
}

type ProductQuery struct {
	SearchTerm string
}

type ProductRepository interface {
	GetAll(ctx context.Context, q ProductQuery) ([]Product, error)
}

type ProductService interface {
	GetAll(ctx context.Context, q ProductQuery) ([]Product, error)
}
