package main

import "auto-ssml/src/api"

func main() {

	//text := "<speak>Te recuerdo que el comparativo se usa para comparar dos cosas por ejemplo: ##eng bigger, better, more expensive. ##xeng Por otro lado el superlativo se usa para comparar una cosa con toda una categoría, por ejemplo: ##eng the biggest, the best, the most expensive. ##xeng</speak>"

	//aux := functions.ChangeVoiceLanguageENG(text)

	//fmt.Printf("%s\n", aux)
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
	//functions.Synthethize(aux, "lang")

	//API listener
	api.SynthethizeService()

}
