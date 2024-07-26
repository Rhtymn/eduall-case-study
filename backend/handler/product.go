package handler

import (
	"net/http"

	"github.com/Rhtymn/eduall-case-study/backend/apperror"
	"github.com/Rhtymn/eduall-case-study/backend/domain"
	"github.com/Rhtymn/eduall-case-study/backend/dto"
	"github.com/gin-gonic/gin"
)

type ProductHandlerOpts struct {
	ProductSrv domain.ProductService
	Domain     string
}

type ProductHandler struct {
	productSrv domain.ProductService
	domain     string
}

func NewProductHandler(opts ProductHandlerOpts) *ProductHandler {
	return &ProductHandler{
		productSrv: opts.ProductSrv,
		domain:     opts.Domain,
	}
}

func (h *ProductHandler) GetAll(ctx *gin.Context) {
	var q dto.ProductQuery
	err := ctx.ShouldBindQuery(&q)
	if err != nil {
		ctx.Error(apperror.NewInvalidValidation(err))
		ctx.Abort()
		return
	}

	products, pageInfo, err := h.productSrv.GetAll(ctx, q.ToProductQuery())
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseOK(dto.NewProductResponseWithMeta(products, pageInfo)))
}
