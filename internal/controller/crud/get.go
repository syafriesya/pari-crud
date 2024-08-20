package crud

import (
	"net/http"
	"pari/internal/domain"

	"github.com/gin-gonic/gin"
)

func (cc CrudController) GetAll(c *gin.Context) {
	ctx := c.Request.Context()

	result, err := cc.crudUsecase.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Fail",
			"message": err,
			"data":    result,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Get All Acquired.",
		"data":    result,
	})
}

func (cc CrudController) GetById(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.GetUint("id")

	result, err := cc.crudUsecase.GetById(ctx, domain.RequestGet{Id: id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Fail",
			"message": err,
			"data":    result,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Get All Acquired.",
		"data":    result,
	})
}
