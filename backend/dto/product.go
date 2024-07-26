package dto

import "github.com/Rhtymn/eduall-case-study/backend/domain"

type ProductQuery struct {
	SearchTerm string `form:"search_term"`
}

func (q *ProductQuery) ToProductQuery() domain.ProductQuery {
	return domain.ProductQuery{
		SearchTerm: q.SearchTerm,
	}
}
