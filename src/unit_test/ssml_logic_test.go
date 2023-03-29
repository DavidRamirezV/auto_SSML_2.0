package unit_test

import (
	"auto-ssml/src/functions"
	"testing"
)

func TestBreakBeginEnd(t *testing.T) {
	esperado := "<speak> <break time=\"500ms\"/> Lorem ipsum text for testing code <break time=\"500ms\"/> </speak>"
	result := functions.BreakBeginEnd("Lorem ipsum text for testing code")
	t.Log("Revisando etiquetas SSML...")
	if result != esperado {
		t.Errorf("La funcion BreakBeginEnd() salio mal \n Se obtuvo: %s \n Se esperaba: %s", result, esperado)
		t.Fail()
	}
	t.Log("OK!")
	t.Log("Test BreakBeginEnd() finalizado correctamente.")

}

func TestChangeSpeed(t *testing.T) {
	esperado := "<prosody rate=\"118%\"> Lorem ipsum text for testing code </prosody>"
	result := functions.ChangeSpeed("Lorem ipsum text for testing code")

	t.Log("Revisando etiquetas SSML...")

	if result != esperado {
		t.Errorf("La funcion ChangeSpeed() salio mal \n Se obtuvo: %s \n Se esperaba: %s", result, esperado)
		t.Fail()
	}
	t.Log("OK!")
	t.Log("Test ChangeSpeed() finalizado correctamente.")

}

func TestChangeVoiceLanguageENG(t *testing.T) {
	esperado_con := "Un ejemplo de verbo irregular en inglés es el verbo ir, que en su forma base es <prosody rate=\"80%\"><voice language=\"en-US\"> Go </voice></prosody> o <prosody rate=\"80%\"><voice language=\"en-US\"> Goes </voice></prosody>, en pasado simple es <prosody rate=\"80%\"><voice language=\"en-US\"> Went </voice></prosody> y en pasado participio es <prosody rate=\"80%\"><voice language=\"en-US\"> Gone </voice></prosody>"

	result := functions.ChangeVoiceLanguageENG("Un ejemplo de verbo irregular en inglés es el verbo ir, que en su forma base es ##eng Go ##xeng o ##eng Goes ##xeng, en pasado simple es ##eng Went ##xeng y en pasado participio es ##eng Gone ##xeng")

	esperado_2 := "Un ejemplo de verbo irregular en inglés es el verbo ir, que en su forma base es <prosody rate=\"80%\"><voice language=\"en-US\"> Go </voice></prosody> o ##eng Goes, en pasado simple es Went y en pasado participio es Gone"

	result_2 := functions.ChangeVoiceLanguageENG("Un ejemplo de verbo irregular en inglés es el verbo ir, que en su forma base es ##eng Go ##xeng o ##eng Goes, en pasado simple es Went y en pasado participio es Gone")

	esperado_3 := "Un ejemplo de verbo irregular en inglés es el verbo ir, que en su forma base es <prosody rate=\"80%\"><voice language=\"en-US\"> Go </voice></prosody> o Goes ##xeng, en pasado simple es Went y en pasado participio es Gone"

	result_3 := functions.ChangeVoiceLanguageENG("Un ejemplo de verbo irregular en inglés es el verbo ir, que en su forma base es ##eng Go ##xeng o Goes ##xeng, en pasado simple es Went y en pasado participio es Gone")

	t.Log("Test misma cantidad de eng y xeng...")
	if result != esperado_con {
		t.Errorf("La funcion ChangeVoiceLanguageENG() salio mal \n Se obtuvo: %s \n Se esperaba: %s", result, esperado_con)
		t.Fail()
	}
	t.Log("OK!")
	t.Log("Test mayor cantidad de eng que de xeng...")
	if result_2 != esperado_2 {
		t.Errorf("La funcion ChangeVoiceLanguageENG() salio mal \n Se obtuvo: %s \n Se esperaba: %s", result_2, esperado_2)
		t.Fail()
	}
	t.Log("OK!")

	t.Log("Test mayor cantidad de xeng que de eng...")
	if result_3 != esperado_3 {
		t.Errorf("La funcion ChangeVoiceLanguageENG() salio mal \n Se obtuvo: %s \n Se esperaba: %s", result_3, esperado_3)
		t.Fail()
	}
	t.Log("OK!")
	t.Log("Test ChangeVoiceLanguageENG() finalizado correctamente.")
}

func TestIsEmailValid(t *testing.T) {
	result := functions.IsEmailValid("12_Aa44@sgf.com")
	result_2 := functions.IsEmailValid("asdfafs@gasgg")
	result_3 := functions.IsEmailValid("uu38nfan_sadfgj")

	t.Log("Revisando correo valido...")

	if result != true {
		t.Errorf("La funcion isEmailValid() salio mal")
		t.Fail()
	}
	t.Log("OK!")
	t.Log("Revisando correo no valido...")
	if result_2 != false {
		t.Errorf("La funcion isEmailValid() salio mal")
		t.Fail()
	}
	t.Log("OK!")
	t.Log("Revisando correo no valido...")
	if result_3 != false {
		t.Errorf("La funcion isEmailValid() salio mal")
		t.Fail()
	}
	t.Log("OK!")
	t.Log("Test ChangeSpeed() finalizado correctamente.")

}
