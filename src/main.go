package main

import (
	"auto-ssml/src/functions"
	"fmt"
)

func main() {

	text := "Con respecto a la reforma de la ley, ##enfasis cabe destacar que,  años atrás, hemos  padecido las consecuencias de una ##xenfasis reforma  similar que ha llevado a un retraso en la calidad de vida de los ciudadanos."
	aux := functions.AddEmphasis(text)
	fmt.Println(aux)

	/* Big example
	ssml := "<speak><prosody rate=\"fast\"><break time=\"500ms\"/>¿Sabes que es <say-as interpret-as=\"characters\"> SSML</say-as>?,<break time=\"1s\"/> En primer lugar,<say-as interpret-as=\"characters\"> SSML</say-as> viene del inglés <voice language=\"en-US\" gender=\"female\"><break time=\"500ms\"/> Speech Synthesis Markup Languages.</voice> </prosody></speak>"
	*/

	//example and test
	//functions.Synthethize(aux, "lang")

	//API listener
	//api.SynthethizeService()

}
