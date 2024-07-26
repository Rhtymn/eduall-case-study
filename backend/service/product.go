package service

import (
	"context"

	"github.com/Rhtymn/eduall-case-study/backend/apperror"
	"github.com/Rhtymn/eduall-case-study/backend/domain"
)

type ProductServiceOpts struct {
	Product domain.ProductRepository
}

type productService struct {
	productRepository domain.ProductRepository
}

func NewProductService(opts ProductServiceOpts) *productService {
	return &productService{
		productRepository: opts.Product,
	}
}

func (s *productService) GetAll(ctx context.Context, q domain.ProductQuery) ([]domain.Product, domain.PageInfo, error) {
	products, err := s.productRepository.GetAll(ctx, q)
	if err != nil {
		return nil, domain.PageInfo{}, apperror.Wrap(err)
	}

	pageInfo, err := s.productRepository.GetPageInfo(ctx, q)
	if err != nil {
		return nil, domain.PageInfo{}, apperror.Wrap(err)
	}

	pageInfo.ItemsPerPage = q.Limit
	if q.Limit == 0 {
		pageInfo.ItemsPerPage = len(products)
	}

	if pageInfo.ItemsPerPage == 0 {
		pageInfo.PageCount = 0
	} else {
		pageInfo.PageCount = (int(pageInfo.ItemCount) + pageInfo.ItemsPerPage - 1) / pageInfo.ItemsPerPage
	}

	return products, pageInfo, nil
}
