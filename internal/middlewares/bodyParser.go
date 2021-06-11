package middlewares

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

// BodyParser parses json body
func BodyParser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.ContentType() == "application/json" {
			data, err := ctx.GetRawData()
			if err != nil {
				fmt.Println("error")
				fmt.Println(err.Error())
			}

			ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		}

		ctx.Next()
	}

}
