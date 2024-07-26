package dto

import (
	"fmt"

	"github.com/Rhtymn/eduall-case-study/backend/domain"
)

type ProductQuery struct {
	SearchTerm string `form:"search_term" binding:"omitempty"`
	Page       int    `form:"page" binding:"omitempty,numeric,min=1"`
	Limit      int    `form:"limit" binding:"omitempty,numeric,min=1"`
}

type ProductResponse struct {
	ID                 int64   `json:"id"`
	Brand              string  `json:"brand"`
	Model              string  `json:"model"`
	ScreenSize         string  `json:"screen_size"`
	Color              string  `json:"color"`
	HardDisk           string  `json:"harddisk"`
	CPU                string  `json:"cpu"`
	RAM                string  `json:"ram"`
	OS                 string  `json:"os"`
	SpecialFeatures    string  `json:"special_features"`
	Graphics           string  `json:"graphics"`
	GraphicCoprocessor string  `json:"graphic_coprocessor"`
	CpuSpeed           string  `json:"cpu_speed"`
	Rating             float64 `json:"rating"`
	Price              string  `json:"price"`
}

type ProductResponseWithMeta struct {
	Products []ProductResponse `json:"products"`
	Meta     PageInfo          `json:"meta"`
}

func NewProductResponse(p domain.Product) ProductResponse {
	return ProductResponse{
		ID:                 p.ID,
		Brand:              p.Brand,
		Model:              p.Model,
		ScreenSize:         p.ScreenSize,
		Color:              p.Color,
		HardDisk:           p.HardDisk,
		CPU:                p.CPU,
		RAM:                p.RAM,
		OS:                 p.OS,
		SpecialFeatures:    p.SpecialFeatures,
		Graphics:           p.Graphics,
		GraphicCoprocessor: p.GraphicCoprocessor,
		CpuSpeed:           p.CpuSpeed,
		Rating:             p.Rating,
		Price:              fmt.Sprintf("$%.2f", p.Price),
	}
}

func NewProductsResponse(products []domain.Product) []ProductResponse {
	res := make([]ProductResponse, len(products))
	for i := 0; i < len(products); i++ {
		res[i] = NewProductResponse(products[i])
	}
	return res
}

func NewProductResponseWithMeta(products []domain.Product, pageInfo domain.PageInfo) ProductResponseWithMeta {
	return ProductResponseWithMeta{
		Products: NewProductsResponse(products),
		Meta:     NewPageInfoResponse(pageInfo),
	}
}

func (q *ProductQuery) ToProductQuery() domain.ProductQuery {
	page := q.Page
	if page == 0 || q.Limit == 0 {
		page = 1
	}

	return domain.ProductQuery{
		SearchTerm: q.SearchTerm,
		Page:       page,
		Limit:      q.Limit,
	}
}
