package main

import (
	"auto-ssml/src/api"
)

func main() {
	/* Big example
	ssml := "<speak><prosody rate=\"fast\"><break time=\"500ms\"/>¿Sabes que es <say-as interpret-as=\"characters\"> SSML</say-as>?,<break time=\"1s\"/> En primer lugar,<say-as interpret-as=\"characters\"> SSML</say-as> viene del inglés <voice language=\"en-US\" gender=\"female\"><break time=\"500ms\"/> Speech Synthesis Markup Languages.</voice> </prosody></speak>"
	*/
	api.SynthethizeService()

	//functions.Synthethize(ssml, filename_final)

}
