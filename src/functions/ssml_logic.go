package functions

/*
ENTRADA: String texto al que se le aplicara el SSML
SALIDA: String output, es el texto de entrada con un SSML completo aplicado
DESCRIPCION: Añade las etiquetas <speak> y <break> al inicio y termino del texto de entrada,
	<speak> para que el texto pueda ser utilzado por cualquier TTS que acepte SSML y <break>
	para que el audio de salida tenga un tiempo muerto tanto al inicio como al final y se
	logre escuchar el audio completo en cualquier reproductor de audio.
*/
func BreakBeginingEnd(text string) string {

	output := "<speak> <break time=\"500ms\"/> " + text + " <break time=\"500ms\"/> </speak>"

	return output
}

func ChangeSpeed(text string) string {

	output := "<prosody rate=\"100%\"/> " + text + "</prosody>"

	return output
}

/*
son 185 wpm por defecto en español - voz mujer
son 188 wpm por defecto en español - voz hombre
*/
