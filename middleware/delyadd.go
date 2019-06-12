package middleware

import (
	"delycarapi/repository"
	_ "delycarapi/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCarsNearByG(c *gin.Context) {

	var filter = repository.Filter{}
	if err := c.ShouldBindQuery(&filter); err != nil {
		println(err.Error())
	}
	car, err := repository.GetCarsNearBy(filter)
	if err != nil {
		println(err.Error())
	}
	c.JSON(http.StatusOK, car)
}
