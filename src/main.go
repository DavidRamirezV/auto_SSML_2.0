package main

import (
	"auto-ssml/src/api"
)

func main() {

	//text := "Los signos de puntuación delimitan las frases y parrafos, a modo de dar una estructura al texto. Son once los signos de puntuación: punto; coma; punto y coma; dos puntos; comillas (simples o dobles); paréntesis; signos de interrogación; signos de exclamación; puntos suspensivos; guión y raya.Del mismo modo existen diferentes modo para utilizarlos por ejemplo... podrian realizar puntos suspensivos, o bien puntos finales."
	/*
		result := ""

		for i := 0; i < len(text); i++ {
			char := string(text[i])
			if char == "." || char == "," || char == ";" || char == ":" {
				result = result + "##"
			} else {
				result = result + string(text[i])
			}

		}

		r, err := charset.NewReader(strings.NewReader(result), "latin1")
		if err != nil {
			log.Fatal(err)
		}
		kasd, err := ioutil.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", kasd)

	*/
	/* Big example
	ssml := "<speak><prosody rate=\"fast\"><break time=\"500ms\"/>¿Sabes que es <say-as interpret-as=\"characters\"> SSML</say-as>?,<break time=\"1s\"/> En primer lugar,<say-as interpret-as=\"characters\"> SSML</say-as> viene del inglés <voice language=\"en-US\" gender=\"female\"><break time=\"500ms\"/> Speech Synthesis Markup Languages.</voice> </prosody></speak>"
	*/

	//example and test
	//functions.Synthethize(ssml, "test")

	//API listener
	api.SynthethizeService()

}
