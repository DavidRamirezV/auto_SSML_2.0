package main

import "auto-ssml/src/api"

func main() {

	// text := "El énfasis es la acción de resaltar algo, con el objetivo de destacar una característica sobre las demás, ##enfasis cabe destacar que este recurso del habla ##xenfasis incrementa el volumen y la entonación de la voz."
	// aux := functions.AddEmphasis(text)
	// fmt.Println(aux)

	/* Big example
	ssml := "<speak><prosody rate=\"fast\"><break time=\"500ms\"/>¿Sabes que es <say-as interpret-as=\"characters\"> SSML</say-as>?,<break time=\"1s\"/> En primer lugar,<say-as interpret-as=\"characters\"> SSML</say-as> viene del inglés <voice language=\"en-US\" gender=\"female\"><break time=\"500ms\"/> Speech Synthesis Markup Languages.</voice> </prosody></speak>"
	*/

	//example and test
	//functions.Synthethize(aux, "lang")

	//API listener
	api.SynthethizeService()

}
