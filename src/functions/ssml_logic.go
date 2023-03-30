package functions

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

/*
ENTRADA: String texto al que se le aplicara el SSML
SALIDA: String output, es el texto de entrada con un SSML completo aplicado
DESCRIPCION: Añade las etiquetas <speak> y <break> al inicio y termino del texto de entrada,
	<speak> para que el texto pueda ser utilzado por cualquier TTS que acepte SSML y <break>
	para que el audio de salida tenga un tiempo muerto tanto al inicio como al final y se
	logre escuchar el audio completo en cualquier reproductor de audio.
*/
func BreakBeginEnd(text string) string {

	output := "<speak> <break time=\"500ms\"/> " + text + " <break time=\"500ms\"/> </speak>"

	return output
}

/*
ENTRADA: String texto al que se le aplicara el SSML
SALIDA: String output, es el texto de entrada con un SSML completo aplicado
DESCRIPCION: Añade las etiquetas <prosody rate="n%"> al inicio y termino del texto de entrada,
	para que el audio de salida tenga una velocidad n% respecto a la velocidad por defecto y se pueda
	incrementar o reducir la velocidad del texto de entrada.
*/
func ChangeSpeed(text string) string {

	/*
	   son 185 wpm promedio por defecto en español estadounidense - voz mujer
	   son 188 wpm promedio por defecto en español estadounidense - voz hombre

	   tomando el menor valor,

	   185 -> 100
	   219 -> X

	   219*100/185=118
	*/

	output := "<prosody rate=\"118%\"> " + text + " </prosody>"

	return output

}

/*
ENTRADA:  String text al que se le aplicara el SSML
SALIDA:  String output, es el texto de entrada con el SSML decambio de idioma aplicado
DESCRIPCION:
*/
func ChangeVoiceLanguageENG(text string) string {

	//codigos que marcan el inicio y el final del cambio de idioma
	begin_code := "##eng"
	end_code := "##xeng"

	eng_count := strings.Count(text, begin_code)
	xeng_count := strings.Count(text, end_code)
	output := ""

	//Si la cantidad de simbolos ##eng y ##xeng es diferente, se debe reemplazar lo la cantidad menor que se encuentre,
	//para no agregar etiquetas que no se abren o cierran
	if eng_count > xeng_count {
		str_aux := strings.Replace(text, begin_code, "<prosody rate=\"80%\"><voice language=\"en-US\">", xeng_count)
		output = strings.Replace(str_aux, end_code, "</voice></prosody>", -1)
	} else if eng_count < xeng_count {
		str_aux := strings.Replace(text, begin_code, "<prosody rate=\"80%\"><voice language=\"en-US\">", -1)
		output = strings.Replace(str_aux, end_code, "</voice></prosody>", eng_count)
	} else {
		//si eng_count=0 y xeng_count=0, no existe la marca en el texto de entrada y no se reemplazara nada
		str_aux := strings.Replace(text, begin_code, "<prosody rate=\"80%\"><voice language=\"en-US\">", -1)
		output = strings.Replace(str_aux, end_code, "</voice></prosody>", -1)
	}

	//fmt.Printf("\n %s\n", output)
	return output

}

/*
ENTRADA:  String
SALIDA:  Booleando que detecta si un correo electronico es correcto
DESCRIPCION:
*/
func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))`)
	return emailRegex.MatchString(e)
}

/*
ENTRADA:  String text al que se le aplicara el SSML
SALIDA:  String output, es el texto de entrada con el SSML decambio de idioma aplicado
DESCRIPCION:
*/
func Pauses(text string) string {

	words := strings.Fields(text)
	short_min := 10
	short_max := 20
	long_min := 40
	long_max := 65
	extra_long_min := 90
	extra_long_max := 120
	rand.Seed(time.Now().UnixNano())
	var range_aux float64

	//fmt.Println(words, len(words))
	for index, element := range words {

		if strings.Contains(element, ",") {
			range_aux = float64(rand.Intn(short_max-short_min)+short_min) / 100
			//fmt.Println(index, element, range_aux)
			words[index] = strings.Replace(element, ",", " <break strength=\"weak\" time=\""+fmt.Sprintf("%.2f", range_aux)+"s\"/>", -1)

		}
		if strings.Contains(element, ":") {

			range_aux = float64(rand.Intn(long_max-long_min)+long_min) / 100
			//fmt.Println(index, element, range_aux)
			words[index] = strings.Replace(element, ":", " <break strength=\"medium\" time=\""+fmt.Sprintf("%.2f", range_aux)+"s\"/>", -1)
		}
		if strings.Contains(element, ";") {
			range_aux = float64(rand.Intn(long_max-long_min)+long_min) / 100
			//fmt.Println(index, element, range_aux)
			words[index] = strings.Replace(element, ";", " <break strength=\"medium\" time=\""+fmt.Sprintf("%.2f", range_aux)+"s\"/>", -1)
		}
		if strings.Contains(element, ".") {
			if strings.Contains(element, "...") {
				range_aux = float64(rand.Intn(long_max-long_min)+long_min) / 100
				//punto suspensivo
				//fmt.Println(index, element, range_aux)
				words[index] = strings.Replace(element, "...", " <break strength=\"strong\" time=\""+fmt.Sprintf("%.2f", range_aux)+"s\"/>", -1)
			} else {
				range_aux = float64(rand.Intn(extra_long_max-extra_long_min)+extra_long_min) / 100
				//punto simple
				if !(IsEmailValid(element)) {

					//fmt.Println(index, element, range_aux)
					//words[index] = strings.Replace(element, ".", " <break strength=\"x-strong\" time=\""+fmt.Sprintf("%.2f", range_aux)+"s\"/>", -1)
				}

			}
		}
	}
	//fmt.Println(strings.Join(words, " "))
	return strings.Join(words, " ")

}

/*EMPHASIS LOGIC*/
/*
ENTRADA: int enfasis count cantidad de enfasis, int enfasis count cantidad de enfasis, arreglo con indices de los enfasis y arreglo con indices de los xenfasis
SALIDA: true si no hay problemas en el uso de enfasis, false si las marcas de enfasis son incongruentes con el formato presentado
DESCRIPCION: Permite determinar si el texto de entrada posee el formato correcto para el uso de enfasis
*/
func StatusEnfasis(enfasis_count int, xenfasis_count int, index_enfasis []int, index_xenfasis []int) bool {

	i := 0
	j := 0
	//fmt.Println("enfasis_count", enfasis_count)
	if enfasis_count != xenfasis_count {
		fmt.Println("Diferente cantidad de ##enfasis y ##xenfasis", "Hay", enfasis_count, "##enfasis", "- Hay", xenfasis_count, "##xenfasis")
		return false
	}
	for i < enfasis_count {
		//fmt.Println("i", i, "j", j)

		if i == enfasis_count-1 {
			if index_enfasis[i] < index_xenfasis[i] {
				j++
			} else {
				fmt.Println("Enfasis dentro de enfasis")
				return false
			}
		} else {
			if index_enfasis[i] < index_xenfasis[i] && index_xenfasis[i] < index_enfasis[i+1] {
				j++
			} else {
				fmt.Println("Enfasis dentro de enfasis")

				return false
			}

		}
		i++
	}
	return true
}

/*
ENTRADA:  int n numero
SALIDA:  bool
DESCRIPCION: Paridad de un numero
*/
func IsEven(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}

/*
ENTRADA:  String text al que se le aplicara el SSML
SALIDA:  String output, es el texto de entrada con el SSML decambio de idioma aplicado
DESCRIPCION:
*/
func AddEmphasis(text string) string {
	output := ""
	tag_enfasis := "##enfasis"
	tag_xenfasis := "##xenfasis"

	words := strings.Fields(text)                       //arreglo de las palabras en el texto de entrada separadas respecto a los espacios en blanco " "
	enfasis_count := strings.Count(text, tag_enfasis)   //cantidad de ocurrencias de enfasis en el texto de entrada
	xenfasis_count := strings.Count(text, tag_xenfasis) //cantidad de ocurrencias de xenfasis en el texto de entrada
	index_enfasis := make([]int, enfasis_count)         //arreglo que guardara los indices de las marcas de enfasis del arreglo words
	index_xenfasis := make([]int, xenfasis_count)       //arreglo que guardara los indices de las marcas de xenfasis del arreglo words
	i_e := 0                                            //auxiliar para recorrer el arreglo de indices de enfasis
	i_x := 0                                            //auxiliar para recorrer el arreglo de indices de xenfasis

	for i, element := range words {
		if strings.Contains(element, tag_enfasis) {
			index_enfasis[i_e] = i
			i_e = i_e + 1
		}
		if strings.Contains(element, tag_xenfasis) {
			index_xenfasis[i_x] = i
			i_x = i_x + 1
		}
		//fmt.Println(i, element)
	}
	//fmt.Println(words)
	//fmt.Println("indices de los enfasis: ", index_enfasis, "indices de los xenfasis: ", index_xenfasis)
	if StatusEnfasis(enfasis_count, xenfasis_count, index_enfasis, index_xenfasis) {
		//en este punto la cantidad de marcas de enfasis y xenfasis es la misma y no existen enfasis dentro de otros enfasis
		for i := enfasis_count - 1; 0 <= i; i-- {
			between_words := index_xenfasis[i] - index_enfasis[i] - 1 //-1 debido a indexacion
			//fmt.Println("betweeen_words:", between_words)
			words[index_xenfasis[i]] = strings.Replace(tag_xenfasis, words[index_xenfasis[i]], "</prosody>", 1)
			if between_words <= 3 {
				//fmt.Println("Caso menor que 3")
				words[index_enfasis[i]] = strings.Replace(tag_enfasis, words[index_enfasis[i]], "<prosody volume=\"loud\" rate=\"80%\" pitch=\"+1st\">", 1)
			} else if 4 <= between_words && between_words <= 7 {
				if between_words == 4 || between_words == 5 {
					//fmt.Println("Caso entre 4 y 5")
					words[index_xenfasis[i]-2] = words[index_xenfasis[i]-2] + " </prosody><prosody volume=\"loud\" rate=\"85%\" pitch=\"+1st\">"
					words[index_enfasis[i]+1] = words[index_enfasis[i]+1] + " </prosody><prosody volume=\"x-loud\" rate=\"70%\" pitch=\"+1st\">"
					words[index_enfasis[i]] = strings.Replace(tag_enfasis, words[index_enfasis[i]], "<prosody volume=\"loud\" rate=\"85%\" pitch=\"+1st\">", 1)
				} else {
					//fmt.Println("Caso entre 6 y 7")
					words[index_xenfasis[i]-3] = words[index_xenfasis[i]-3] + " </prosody><prosody volume=\"loud\" rate=\"85%\" pitch=\"+1st\">"
					words[index_enfasis[i]+2] = words[index_enfasis[i]+2] + " </prosody><prosody volume=\"x-loud\" rate=\"70%\" pitch=\"+1st\">"
					words[index_enfasis[i]] = strings.Replace(tag_enfasis, words[index_enfasis[i]], "<prosody volume=\"loud\" rate=\"85%\" pitch=\"+1st\">", 1)
				}
			} else if 8 <= between_words && between_words <= 30 {
				n := 0
				if IsEven(between_words) {
					n = (between_words) / 4
				} else {
					n = (between_words - 1) / 4
				}
				//fmt.Println("Caso entre 8 y 30, n:", n)
				words[index_xenfasis[i]-n-1] = words[index_xenfasis[i]-n-1] + " </prosody> <prosody volume=\"default\" rate=\"80%\" pitch=\"+1st\">"
				words[index_enfasis[i]+2*n] = words[index_enfasis[i]+2*n] + " </prosody> <prosody volume=\"loud\" rate=\"90%\" pitch=\"+2st\">"
				words[index_enfasis[i]+n] = words[index_enfasis[i]+n] + " </prosody> <prosody volume=\"x-loud\" rate=\"80%\" pitch=\"+2st\">"
				words[index_enfasis[i]] = strings.Replace(tag_enfasis, words[index_enfasis[i]], "<prosody volume=\"loud\" rate=\"85%\" pitch=\"+1st\">", 1)
			} else {
				words[index_enfasis[i]] = strings.Replace(tag_enfasis, words[index_enfasis[i]], "<emphasis level=\"moderate\">", 1)
				words[index_xenfasis[i]] = strings.Replace(tag_xenfasis, words[index_xenfasis[i]], "</emphasis>", 1)
			}
		}
		//fmt.Println(words)

	}

	output = strings.Join(words, " ")
	return output

}

/*
<speak>cabe destacar que, <prosody volume="loud" rate="80%" pitch="+1st"> años atrás </prosody> <prosody volume="x-loud" rate="70%" pitch="+5st"> hemos sufrido las</prosody>
 <prosody  volume="medium" rate="80%" pitch="+1st"> consecuencias de una reforma similar </prosody>  que ha llevado un retraso en la calidad de vida de los ciudadanos. </speak>
*/
