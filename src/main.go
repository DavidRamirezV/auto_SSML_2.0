package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {

	text := "Los signos de puntuación delimitan las frases y parrafos, a modo de dar una estructura al texto. Son once los signos de puntuación: punto; coma; punto y coma; dos puntos; comillas (simples o dobles); paréntesis; signos de interrogación; signos de exclamación; puntos suspensivos; guión y raya. Del mismo modo existen diferentes formas para utilizarlos por ejemplo... podrian realizar puntos . suspensivos , o bien puntos finales. Cuidado con los correos electrincos como por ejemplo david.ramirez.v@usach.cl."

	words := strings.Fields(text)
	short_min := 50
	short_max := 125
	long_min := 150
	long_max := 250
	extra_long_min := 250
	extra_long_max := 350
	rand.Seed(time.Now().UnixNano())
	var range_aux float64

	fmt.Println(words, len(words))
	for index, element := range words {

		if strings.Contains(element, ",") {
			range_aux = float64(rand.Intn(short_max-short_min)+short_min) / 100
			fmt.Println(index, element, range_aux)
			words[index] = strings.Replace(element, ",", " <break time=\""+fmt.Sprintf("%.2f", range_aux)+"s\"/>", -1)

		}
		if strings.Contains(element, ":") {
			//divisiones??
			range_aux = float64(rand.Intn(long_max-long_min)+long_min) / 100
			fmt.Println(index, element, range_aux)
		}
		if strings.Contains(element, ";") {
			range_aux = float64(rand.Intn(long_max-long_min)+long_min) / 100
			fmt.Println(index, element, range_aux)
			words[index] = strings.Replace(element, ";", " <break time=\""+fmt.Sprintf("%.2f", range_aux)+"s\"/>", -1)
		}
		if strings.Contains(element, ".") {
			if strings.Contains(element, "...") {
				range_aux = float64(rand.Intn(long_max-long_min)+long_min) / 100
				//punto suspensivo
				fmt.Println(index, element, range_aux)
				words[index] = strings.Replace(element, "...", " <break time=\""+fmt.Sprintf("%.2f", range_aux)+"s\"/>", -1)
			} else {
				range_aux = float64(rand.Intn(extra_long_max-extra_long_min)+extra_long_min) / 100
				//punto simple
				fmt.Println(index, element, range_aux)
			}
		}

		//strings.Replace(words[index], ",", "</voice></prosody>", -1)

	}

	fmt.Println(words, len(words))

	//aux := functions.ChangeVoiceLanguageENG(text)

	//fmt.Printf("%s\n", aux)

	/* Big example
	ssml := "<speak><prosody rate=\"fast\"><break time=\"500ms\"/>¿Sabes que es <say-as interpret-as=\"characters\"> SSML</say-as>?,<break time=\"1s\"/> En primer lugar,<say-as interpret-as=\"characters\"> SSML</say-as> viene del inglés <voice language=\"en-US\" gender=\"female\"><break time=\"500ms\"/> Speech Synthesis Markup Languages.</voice> </prosody></speak>"
	*/

	//example and test
	//functions.Synthethize(aux, "lang")

	//API listener
	//api.SynthethizeService()

}
