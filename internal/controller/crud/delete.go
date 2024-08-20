package crud

import (
	"net/http"
	"pari/internal/domain"

	"github.com/gin-gonic/gin"
)

func (cc CrudController) Delete(c *gin.Context) {
	id := c.GetUint("id")

	ctx := c.Request.Context()
	if err := cc.crudUsecase.DeleteItem(ctx, domain.RequestGet{Id: id}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Fail",
			"message": "Failed to delete item",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Item deleted successfully",
	})
}
