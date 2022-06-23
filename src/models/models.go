package models

/*
Estructura para el cuerpo del json de entrada
*/
type Body struct {
	Filename string `json:"filename,omitempty"`
	Text     string `json:"text" validate:"required"`
}
