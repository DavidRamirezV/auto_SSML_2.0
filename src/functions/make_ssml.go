package functions

import (
	"fmt"
)

/*
ENTRADA: String texto al que se le aplicara el SSML
SALIDA: String output, es el texto de entrada con un SSML completo aplicado
DESCRIPCION: Obtiene el texto de entrada y le aplica las diferentes etiquetas necesarios para
	simular un discurso oral explicativo, las funciones que aplican el SSML se encuentran en el
	archivo "ssml_logic.go".
*/
func MakeSSML(text string) string {

	pauses := Pauses(text)
	lang_eng := ChangeVoiceLanguageENG(pauses)
	emphasis := AddEmphasis(lang_eng) //move to 1
	global_speed := ChangeSpeed(emphasis)
	output := BreakBeginEnd(global_speed)

	//print final ssml
	fmt.Printf("\n%v\n\n", output)

	return output
}
