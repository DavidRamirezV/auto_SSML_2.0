package api

import (
	"auto-ssml/src/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Probando Servicio GET",
	})

}

func TextPost(c *gin.Context) {
	body := models.Body{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	//Do something
	//tts.Espeak_speak(body.Text)
	fmt.Printf("\n{\n	\"texto\" : \"%s\"\n}\n\n", body.Text)

	c.JSON(http.StatusAccepted, &body)

}
