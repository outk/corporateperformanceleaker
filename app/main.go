package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/outk/corporateperformanceleaker/app/controller"
)

func main() {
	r := gin.Default()

	r.GET("", func(ctx *gin.Context) {
		ctx.Writer.WriteString("Welcome to CorporatePerformanceLeaker!")
	})

	corporates := r.Group("corporates")
	{
		corporates.GET("", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "About All Corporates")
		})
		corporates.POST("", func(ctx *gin.Context) {
			var createCorporatesRequest controller.CreateCorporatesRequest
			err := ctx.ShouldBindJSON(&createCorporatesRequest)
			if err != nil {
				ctx.String(http.StatusBadRequest, err.Error())
				return
			}

			request, err := json.Marshal(createCorporatesRequest)
			if err != nil {
				ctx.String(http.StatusBadRequest, err.Error())
				return
			}

			ctx.String(http.StatusOK, "Received New Corporates: "+string(request))
		})
		corporates.GET(":id", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "About #"+ctx.Param("id")+" Corporate")
		})

		performances := corporates.Group("performances")
		{
			performances.GET("", func(ctx *gin.Context) {
				ctx.Writer.WriteString("Performances")
			})
		}
	}

	r.Run("localhost:8080")
}
