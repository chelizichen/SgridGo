package utils

import (
	"errors"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

var AbortWithError = func(c *gin.Context, err string) {
	c.AbortWithStatusJSON(http.StatusOK, &gin.H{
		"status":  0,
		"message": err,
		"data":    nil,
	})
}

// Done
var AbortWithSucc = func(c *gin.Context, data interface{}) {
	c.AbortWithStatusJSON(http.StatusOK, &gin.H{
		"status":  200,
		"message": "ok",
		"result": map[string]interface{}{
			"total": 0,
			"data":  data,
		},
	})
}

// List
var AbortWithSuccList = func(c *gin.Context, data interface{}, total int64) {
	c.AbortWithStatusJSON(http.StatusOK, &gin.H{
		"status":  200,
		"message": "ok",
		"result": map[string]interface{}{
			"total": total,
			"data":  data,
		},
	})
}

func ConvertFiled(from interface{}, to interface{}) error {
	fromVal := reflect.ValueOf(from).Elem()
	toVal := reflect.ValueOf(to).Elem()

	// 检查两个结构体类型是否相同，或者具有相同字段
	if fromVal.Type() != toVal.Type() {
		for i := 0; i < fromVal.NumField(); i++ {
			fromField := fromVal.Field(i)
			toField := toVal.FieldByName(fromVal.Type().Field(i).Name)

			// 如果目标结构体里没有相应字段则跳过
			if !toField.IsValid() || !toField.CanSet() {
				continue
			}

			// 将字段赋值
			toField.Set(fromField)
		}
		return nil
	}
	return errors.New("incompatible types")
}
