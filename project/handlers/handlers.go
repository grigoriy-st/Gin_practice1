package handlers

import (
	"first/models"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.XML(200, models.Person {
		FirstName: "Grigoriy",
		LastName:  "Starostin",
	})
}
