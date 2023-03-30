package unit_test

import (
	"auto-ssml/src/functions"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestBreakBeginEnd(t *testing.T) {
	esperado := "<speak> <break time=\"500ms\"/> Lorem ipsum text for testing code <break time=\"500ms\"/> </speak>"
	result := functions.BreakBeginEnd("Lorem ipsum text for testing code")
	t.Log("Revisando etiquetas SSML...")
	if result != esperado {
		t.Errorf("La funcion BreakBeginEnd() salio mal \n Se obtuvo: %s \n Se esperaba: %s", result, esperado)
		t.Fail()
	} else {

		t.Log("OK!")
		t.Log("Test BreakBeginEnd() finalizado correctamente.")
	}

}

func TestChangeSpeed(t *testing.T) {
	esperado := "<prosody rate=\"118%\"> Lorem ipsum text for testing code </prosody>"
	result := functions.ChangeSpeed("Lorem ipsum text for testing code")

	t.Log("Revisando etiquetas SSML...")

	if result != esperado {
		t.Errorf("La funcion ChangeSpeed() salio mal \n Se obtuvo: %s \n Se esperaba: %s", result, esperado)
		t.Fail()
	} else {

		t.Log("OK!")
		t.Log("Test ChangeSpeed() finalizado correctamente.")
	}

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
	} else {

		t.Log("OK!")
	}
	t.Log("Test mayor cantidad de eng que de xeng...")
	if result_2 != esperado_2 {
		t.Errorf("La funcion ChangeVoiceLanguageENG() salio mal \n Se obtuvo: %s \n Se esperaba: %s", result_2, esperado_2)
		t.Fail()
	} else {

		t.Log("OK!")
	}

	t.Log("Test mayor cantidad de xeng que de eng...")
	if result_3 != esperado_3 {
		t.Errorf("La funcion ChangeVoiceLanguageENG() salio mal \n Se obtuvo: %s \n Se esperaba: %s", result_3, esperado_3)
		t.Fail()
	} else {

		t.Log("OK!")
		t.Log("Test ChangeVoiceLanguageENG() finalizado correctamente.")
	}
}

func TestIsEmailValid(t *testing.T) {
	result := functions.IsEmailValid("12_Aa44@sgf.rng")
	result_2 := functions.IsEmailValid("asdfafs@gasgg")
	result_3 := functions.IsEmailValid("uu38nfan_sadfgj")

	t.Log("Revisando correo valido...")

	if result != true {
		t.Errorf("La funcion isEmailValid() salio mal")
		t.Fail()
	} else {

		t.Log("Retorna true!")
		t.Log("OK!")

	}
	t.Log("Revisando correo no valido...")
	if result_2 != false {
		t.Errorf("La funcion isEmailValid() salio mal")
		t.Fail()
	} else {

		t.Log("Retorna false!")
		t.Log("OK!")
	}
	t.Log("Revisando correo no valido...")
	if result_3 != false {
		t.Errorf("La funcion isEmailValid() salio mal")
		t.Fail()
	} else {

		t.Log("Retorna false!")
		t.Log("OK!")
		t.Log("Test ChangeSpeed() finalizado correctamente.")
	}

}

func TestPauses(t *testing.T) {
	short_min := 0.10
	short_max := 0.20
	long_min := 0.40
	long_max := 0.65
	result := functions.Pauses("Frutas, como por ejemplo: manzana, pera, platano, durazno, etc.")
	words := strings.Fields(result)
	var float_numbers []float64
	var extracted string

	t.Log("Revisando numeros generados aleatoriamente en comas (\",\")...")
	for _, element := range words {
		if strings.Contains(element, "time=\"") {
			extracted = strings.Replace(strings.Replace(element, "time=\"", "", 1), "s\"/>", "", 1)
			extracted_float, _ := strconv.ParseFloat(extracted, 64)
			float_numbers = append(float_numbers, extracted_float)
		}
	}

	esperado := functions.Pauses("Frutas <break strength=\"weak\" time=\"" + fmt.Sprintf("%.2f", float_numbers[0]) + "s\"/> como por ejemplo <break strength=\"medium\" time=\"" + fmt.Sprintf("%.2f", float_numbers[1]) + "s\"/> manzana <break strength=\"weak\" time=\"" + fmt.Sprintf("%.2f", float_numbers[2]) + "s\"/> pera <break strength=\"weak\" time=\"" + fmt.Sprintf("%.2f", float_numbers[3]) + "s\"/> platano <break strength=\"weak\" time=\"" + fmt.Sprintf("%.2f", float_numbers[4]) + "s\"/> durazno <break strength=\"weak\" time=\"" + fmt.Sprintf("%.2f", float_numbers[5]) + "s\"/> etc.")

	//prueba unitaria fija para el ejemplo en result, considerando valores aleatorios para las pausas
	if short_min <= float_numbers[0] && float_numbers[0] <= short_max &&
		short_min <= float_numbers[2] && float_numbers[2] <= short_max &&
		short_min <= float_numbers[3] && float_numbers[3] <= short_max &&
		short_min <= float_numbers[4] && float_numbers[4] <= short_max &&
		short_min <= float_numbers[5] && float_numbers[5] <= short_max {

		t.Log("OK!")
	} else {
		t.Errorf("La funcion Pauses() salio mal: tiempo de alguna \",\" fuera de su rango")
		t.Fail()
	}
	t.Log("Revisando numeros generados aleatoriamente en dos puntos (\":\")...")
	if long_min <= float_numbers[1] && float_numbers[1] <= long_max {

		t.Log("OK!")
	} else {
		t.Errorf("La funcion Pauses() salio mal: tiempo de algun \":\" fuera de su rango")
		t.Fail()
	}
	t.Log("Revisando etiquetas SSML...")
	if result != esperado {
		t.Errorf("La funcion ChangeSpeed() salio mal \n Se obtuvo: %s \n Se esperaba: %s", result, esperado)
		t.Fail()
	} else {

		t.Log("OK!")
		t.Log("Test Pauses() finalizado correctamente.")
	}

}

func TestStatusEnfasis(t *testing.T) {
	var testarr_1 []int
	var testarr_2 []int
	result := functions.StatusEnfasis(0, 0, testarr_1, testarr_2)
	t.Log("Comprobando condiciones de StatusEnfasis...")
	t.Log("Caso sin enfasis...")
	if result != true {
		t.Errorf("La funcion StatusEnfasis() salio mal \n Se obtuvo: %t \n Se esperaba: %t", result, true)
		t.Fail()
	} else {
		t.Log("Retorna true!")
		t.Log("OK!")
	}

	t.Log("Caso con igual cantidad de enfasis...")
	testarr_1 = append(testarr_1, 1)
	testarr_2 = append(testarr_2, 7)
	result = functions.StatusEnfasis(1, 1, testarr_1, testarr_2)
	if result != true {
		t.Errorf("La funcion StatusEnfasis() salio mal \n Se obtuvo: %t \n Se esperaba: %t", result, true)
		t.Fail()
	} else {
		t.Log("Retorna true!")
		t.Log("OK!")
	}

	t.Log("Mayor cantidad de xenfasis...")
	testarr_2 = append(testarr_2, 10)
	result = functions.StatusEnfasis(1, 2, testarr_1, testarr_2)
	if result != false {
		t.Errorf("La funcion StatusEnfasis() salio mal \n Se obtuvo: %t \n Se esperaba: %t", result, false)
		t.Fail()
	} else {
		t.Log("Retorna false!")
		t.Log("OK!")
	}

	t.Log("Mayor cantidad de enfasis...")
	testarr_1 = append(testarr_1, 5)
	testarr_2 = testarr_2[:len(testarr_2)-1]
	result = functions.StatusEnfasis(2, 1, testarr_1, testarr_2)
	if result != false {
		t.Errorf("La funcion StatusEnfasis() salio mal \n Se obtuvo: %t \n Se esperaba: %t", result, false)
		t.Fail()
	} else {
		t.Log("Retorna false!")
		t.Log("OK!")
	}

	t.Log("Probando igual cantidad de enfasis y xenfasis, con posiciones erradas...")
	testarr_1 = append(testarr_1[:len(testarr_1)-1], 5)
	testarr_2 = append(testarr_2, 15)

	result = functions.StatusEnfasis(2, 2, testarr_1, testarr_2)
	if result != false {
		t.Errorf("La funcion StatusEnfasis() salio mal \n Se obtuvo: %t \n Se esperaba: %t", result, false)
		t.Fail()
	} else {
		t.Log("Retorna false!")
		t.Log("OK!")
	}

}

func TestIsEven(t *testing.T) {
	result := functions.IsEven(28458)
	result_1 := functions.IsEven(317)
	result_2 := functions.IsEven(0)

	t.Log("Revisando paridad del numero 28458...")

	if result != true {
		t.Errorf("La funcion IsEven() salio mal \n Se obtuvo: %t \n Se esperaba: %t", result, true)
		t.Fail()
	} else {

		t.Log("Retorna true!")
		t.Log("OK!")
	}

	t.Log("Revisando paridad del numero 317...")
	if result_1 != false {
		t.Errorf("La funcion IsEven() salio mal \n Se obtuvo: %t\n Se esperaba: %t", result, false)
		t.Fail()
	} else {

		t.Log("Retorna false!")
		t.Log("OK!")
	}

	t.Log("Revisando paridad del numero 0...")
	if result_2 != true {
		t.Errorf("La funcion IsEven() salio mal \n Se obtuvo: %t \n Se esperaba: %t", result, true)
		t.Fail()
	} else {

		t.Log("Retorna true!")
		t.Log("OK!")
		t.Log("Test IsEven() finalizado correctamente.")
	}

}

func TestAddEmphasis(t *testing.T) {
	esperado := "Lorem ipsum text for testing code"
	t.Log("Reemplazando ##enfasis y ##xenfasis por etiquetas SSML...")
	t.Log("Caso sin marcas...")
	result := functions.AddEmphasis("Lorem ipsum text for testing code")
	if result != esperado {
		t.Errorf("La funcion TestAddEmphasis() salio mal \n Se obtuvo: %s \n Se esperaba: %s", result, esperado)
		t.Fail()
	} else {
		t.Log("OK!")
	}

	esperado_2 := "La hipotenusa <prosody volume=\"loud\" rate=\"80%\" pitch=\"+1st\"> se encuentra </prosody> ubicada al lado opuesto del ángulo recto en un triángulo rectángulo"
	t.Log("Caso menor o igual a 3 palabras...")
	result_2 := functions.AddEmphasis("La hipotenusa ##enfasis se encuentra ##xenfasis ubicada al lado opuesto del ángulo recto en un triángulo rectángulo")
	if result_2 != esperado_2 {
		t.Errorf("La funcion TestAddEmphasis() salio mal \n Se obtuvo: %s \n Se esperaba: %s", result_2, esperado_2)
		t.Fail()
	} else {
		t.Log("OK!")
	}

	esperado_1 := "La hipotenusa <prosody volume=\"loud\" rate=\"85%\" pitch=\"+1st\"> se encuentra </prosody><prosody volume=\"x-loud\" rate=\"70%\" pitch=\"+1st\"> ubicada al </prosody><prosody volume=\"loud\" rate=\"85%\" pitch=\"+1st\"> lado opuesto </prosody> del ángulo recto en un triángulo rectángulo"
	t.Log("Caso entre 4 y 7 palabras...")
	result_1 := functions.AddEmphasis("La hipotenusa ##enfasis se encuentra ubicada al lado opuesto ##xenfasis del ángulo recto en un triángulo rectángulo")
	if result_1 != esperado_1 {
		t.Errorf("La funcion TestAddEmphasis() salio mal \n Se obtuvo: %s \n Se esperaba: %s", result_1, esperado_1)
		t.Fail()
	} else {
		t.Log("OK!")
	}

	esperado_3 := "<prosody volume=\"loud\" rate=\"85%\" pitch=\"+1st\"> La hipotenusa se </prosody> <prosody volume=\"x-loud\" rate=\"80%\" pitch=\"+2st\"> encuentra ubicada al </prosody> <prosody volume=\"loud\" rate=\"90%\" pitch=\"+2st\"> lado opuesto del ángulo recto en </prosody> <prosody volume=\"default\" rate=\"80%\" pitch=\"+1st\"> un triángulo rectángulo </prosody>"
	t.Log("Caso entre 8 y 30 palabras...")
	result_3 := functions.AddEmphasis("##enfasis La hipotenusa se encuentra ubicada al lado opuesto del ángulo recto en un triángulo rectángulo ##xenfasis")
	if result_3 != esperado_3 {
		t.Errorf("La funcion TestAddEmphasis() salio mal \n Se obtuvo: %s \n Se esperaba: %s", result_3, esperado_3)
		t.Fail()
	} else {
		t.Log("OK!")
		t.Log("Test TestAddEmphasis() finalizado correctamente.")
	}
}
