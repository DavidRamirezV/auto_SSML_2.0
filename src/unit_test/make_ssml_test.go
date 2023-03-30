package unit_test

import (
	"auto-ssml/src/functions"
	"testing"
)

func TestMakeSSML(t *testing.T) {
	esperado := "<speak> <break time=\"500ms\"/> <prosody rate=\"118%\"> Un ejemplo de verbo irregular en inglés es el verbo ir que en su forma base es <prosody rate=\"80%\"><voice language=\"en-US\"> Go </voice></prosody> o <prosody rate=\"80%\"><voice language=\"en-US\"> Goes </voice></prosody> en <prosody volume=\"loud\" rate=\"85%\" pitch=\"+1st\"> pasado simple es </prosody> <prosody volume=\"x-loud\" rate=\"80%\" pitch=\"+2st\"> <prosody rate=\"80%\"><voice language=\"en-US\"> Went </voice></prosody> </prosody> <prosody volume=\"loud\" rate=\"90%\" pitch=\"+2st\"> y en pasado participio es </prosody> <prosody volume=\"default\" rate=\"80%\" pitch=\"+1st\"> <prosody rate=\"80%\"><voice language=\"en-US\"> Gone </voice></prosody> </prosody> </prosody> <break time=\"500ms\"/> </speak>"
	result := functions.MakeSSML("Un ejemplo de verbo irregular en inglés es el verbo ir que en su forma base es ##eng Go ##xeng o ##eng Goes ##xeng en ##enfasis pasado simple es ##eng Went ##xeng y en pasado participio es ##eng Gone ##xeng ##xenfasis")

	t.Log("Revisando integración de las diferentes funciones...")

	if result != esperado {
		t.Errorf("La funcion MakeSSML() salio mal \n Se obtuvo: %s \n Se esperaba: %s", result, esperado)
		t.Fail()
	} else {

		t.Log("OK!")
		t.Log("Test MakeSSML() finalizado correctamente.")
	}

}
