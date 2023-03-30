package api

import (
	"auto-ssml/src/functions"
	"auto-ssml/src/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SynthethizeService() {
	route := gin.Default()
	route.GET("/", HomeGet)
	route.POST("/ssml", ApplySSML)
	route.POST("/synthesize", TextPost)
	route.POST("/synthesize_normal", SinthesisNoSSML)
	route.Run()
}

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

	//Si funciona bien:
	//Do something
	fmt.Printf("\n{\n	\"texto\" : \"%s\"\n	\"filename\" : \"%s\"\n}\n\n", body.Text, body.Filename)
	txt_with_ssml := functions.MakeSSML(body.Text)
	functions.Synthethize(txt_with_ssml, body.Filename)

	c.JSON(http.StatusAccepted, &body)

}

func SinthesisNoSSML(c *gin.Context) {
	body := models.Body{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	//Si funciona bien:
	//Do something
	fmt.Printf("\n{\n	\"texto\" : \"%s\"\n	\"filename\" : \"%s\"\n}\n\n", body.Text, body.Filename)
	txt_with_ssml := functions.BreakBeginEnd(body.Text)
	functions.Synthethize(txt_with_ssml, body.Filename)

	c.JSON(http.StatusAccepted, &body)

}

func ApplySSML(c *gin.Context) {
	body := models.Body{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	//Si funciona bien:
	//Do something
	fmt.Printf("\n{\n	\"texto\" : \"%s\"\n	\"filename\" : \"%s\"\n}\n\n", body.Text, body.Filename)
	txt_with_ssml := functions.MakeSSML(body.Text)

	c.JSON(200, gin.H{"ssml": txt_with_ssml})

}
