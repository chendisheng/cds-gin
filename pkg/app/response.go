package app

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"cds-gin/pkg/e"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": httpCode,
		"msg":  e.GetMsg(errCode),
		"data": data,
	})

	return
}

func (g *Gin) ResponseOk(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": httpCode,
		"msg":  e.GetMsg(errCode),
		"data": data,
	})

	return
}

func (g *Gin) MsResponse(httpCode int, code BaseMicroServiceCode, data interface{}) {
	g.C.JSON(httpCode, BaseMicroServiceResponse{
		Code: code,
		Bo:   data,
	})
	return
}

func (g *Gin) MsResponseOk(code BaseMicroServiceCode, data interface{}) {
	g.MsResponse(http.StatusOK, code, data)
	return
}

//func (g *Gin) MsResponseFail(code BaseMicroServiceCode, data interface{}) {
//	g.MsResponse(http.StatusOK,code,data)
//	return
//}

func (g *Gin) ResponseWithOther(httpCode int, code BaseMicroServiceCode, data interface{}, other interface{}) {
	g.C.JSON(httpCode, BaseMicroServiceResponse{
		Code:  code,
		Bo:    data,
		Other: other,
	})
	return
}

func (g *Gin) ResponseOkWithOther(code BaseMicroServiceCode, data interface{}, other interface{}) {
	g.ResponseWithOther(http.StatusOK, code, data, other)
	return
}

//分页输出对象
type PageOutPut struct {
	Current string `json:"current"`
	Rows    string `json:"rows"`
	Total   string `json:"total"`
}

//微服务返回对象
type BaseMicroServiceResponse struct {
	Code  BaseMicroServiceCode `json:"code"`
	Bo    interface{}          `json:"bo"`
	Other interface{}          `json:"other"`
}

type BaseMicroServiceCode struct {
	Code  string `json:"code"`
	MsgId string `json:"mstId"`
	Msg   string `json:"msg"`
}
