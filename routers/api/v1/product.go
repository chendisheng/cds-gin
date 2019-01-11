package v1

import (
	"cds-gin/pkg/app"
	"cds-gin/services/product_service"

	"github.com/gin-gonic/gin"
	"net/http"
)

type AddProductForm struct {
	Code  string  `form:"code" valid:"Required"`
	Price float32 `form:"price" valid:"Required"`
}

// @Summary 新增产品
// @Produce  json
// @Param code query string true "Code"
// @Param price query float32 true "Price"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/product [post]
func AddProduct(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddProductForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != http.StatusOK {
		appG.MsResponse(httpCode, app.FAIL, nil)
		return
	}

	productService := product_service.ProductService{Code: form.Code, Price: form.Price}
	if err := productService.Add(); err != nil {
		appG.MsResponseOk(app.FAIL, nil)
		return
	}

	appG.MsResponseOk(app.SUCCESS, nil)
}
