package functions

import (
	"fmt"
	"strings"
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

	output := "<prosody rate=\"118%\"> " + text + "</prosody>"

	return output

}

/*
ENTRADA:  String text al que se le aplicara el SSML
SALIDA:  String output, es el texto de entrada con el SSML decambio de idioma aplicado
DESCRIPCION:
*/
func ChangeVoiceLanguageENG(text string) string {

	eng_count := strings.Count(text, "##eng")
	xeng_count := strings.Count(text, "##xeng")
	output := ""

	//Si la cantidad de simbolos ##eng y ##xeng es diferente, se debe reemplazar lo la cantidad menor que se encuentre,
	//para no agregar etiquetas que no se abren o cierran
	if eng_count > xeng_count {
		str_aux := strings.Replace(text, "##eng", "<prosody rate=\"80%\"><voice language=\"en-US\">", -1)
		output = strings.Replace(str_aux, "##xeng", "</voice></prosody>", xeng_count)
	} else if eng_count < xeng_count {
		str_aux := strings.Replace(text, "##eng", "<prosody rate=\"80%\"><voice language=\"en-US\">", eng_count)
		output = strings.Replace(str_aux, "##xeng", "</voice></prosody>", -1)
	} else {
		//si eng_count=0 y xeng_count=0, no existe la marca en el texto de entrada y no se reemplazara nada
		str_aux := strings.Replace(text, "##eng", "<prosody rate=\"80%\"><voice language=\"en-US\">", -1)
		output = strings.Replace(str_aux, "##xeng", "</voice></prosody>", -1)
	}

	fmt.Printf("\n %s\n", output)
	return output

}

/*
ENTRADA:  String text al que se le aplicara el SSML
SALIDA:  String output, es el texto de entrada con el SSML decambio de idioma aplicado
DESCRIPCION:
*/
func Pauses(text string) string {

	output := ""

	return output

}
