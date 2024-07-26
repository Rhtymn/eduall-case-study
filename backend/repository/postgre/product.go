package postgre

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/Rhtymn/eduall-case-study/backend/apperror"
	"github.com/Rhtymn/eduall-case-study/backend/domain"
	"github.com/jackc/pgx/v5"
)

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *productRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) GetAll(ctx context.Context, q domain.ProductQuery) ([]domain.Product, error) {
	var sb strings.Builder
	sb.WriteString(`SELECT ` + ProductField + ` 
					FROM products `)
	args := pgx.NamedArgs{}

	if q.SearchTerm != "" {
		sb.WriteString(` WHERE brand ILIKE '%' || @search || '%' OR model ILIKE '%' || @search || '%' `)
		args["search"] = q.SearchTerm
	}

	offset := (q.Page - 1) * q.Limit
	if q.Limit != 0 {
		fmt.Fprintf(&sb, " OFFSET %d LIMIT %d ", offset, q.Limit)
	}

	rows, err := r.db.QueryContext(ctx, sb.String(), args)
	if err != nil {
		return nil, apperror.Wrap(err)
	}
	defer rows.Close()

	products := []domain.Product{}
	for rows.Next() {
		np := domain.NullableProduct{}
		p := domain.Product{}
		err := rows.Scan(&p.ID,
			&p.Brand,
			&np.Model,
			&p.ScreenSize,
			&np.Color,
			&np.HardDisk,
			&np.CPU,
			&np.RAM,
			&np.OS,
			&np.SpecialFeatures,
			&np.Graphics,
			&np.GraphicCoprocessor,
			&np.CpuSpeed,
			&np.Rating,
			&np.Price,
		)
		if err != nil {
			return nil, apperror.Wrap(err)
		}
		products = append(products, domain.MergeProduct(p, np))
	}

	return products, nil
}

func (r *productRepository) GetPageInfo(ctx context.Context, q domain.ProductQuery) (domain.PageInfo, error) {
	var sb strings.Builder
	sb.WriteString(`SELECT COUNT(id) 
					FROM products `)
	args := pgx.NamedArgs{}

	if q.SearchTerm != "" {
		sb.WriteString(` WHERE brand ILIKE '%' || @search || '%' OR model ILIKE '%' || @search || '%' `)
		args["search"] = q.SearchTerm
	}

	var totalData int64
	err := r.db.
		QueryRowContext(ctx, sb.String(), args).
		Scan(&totalData)
	if err != nil {
		return domain.PageInfo{}, apperror.Wrap(err)
	}

	return domain.PageInfo{
		CurrentPage: int(q.Page),
		ItemCount:   totalData,
	}, nil
}
