package main

import (
	"auto-ssml/src/api"
	"auto-ssml/src/functions"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin" //Documentacion https://github.com/gin-gonic/gin
)

func main() {
	/* Big example
	ssml := "<speak><prosody rate=\"fast\"><break time=\"500ms\"/>¿Sabes que es <say-as interpret-as=\"characters\"> SSML</say-as>?,<break time=\"1s\"/> En primer lugar,<say-as interpret-as=\"characters\"> SSML</say-as> viene del inglés <voice language=\"en-US\" gender=\"female\"><break time=\"500ms\"/> Speech Synthesis Markup Languages.</voice> </prosody></speak>"
	*/

	//
	route := gin.Default()
	route.GET("/", api.HomeGet)
	route.POST("/synthesize", api.TextPost)
	route.Run()

	text := "Hola output. mp3"
	filename := "voice"
	ssml := functions.MakeSSML(text)

	//reemplazar nombe de archivo para evitar errores de puntuación
	filename2 := strings.Replace(filename, ".", "_", -1)
	filename_final := filename2 + ".mp3"

	fmt.Printf("\n%v\n%v\n", filename_final, ssml)

	//functions.Synthethize(ssml, filename_final)

	fmt.Printf("\nArchivo '%v' guardado en la carpeta /src/files/\n\n", filename_final)

}
