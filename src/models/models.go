package models

/*
Estructura para el cuerpo del json de entrada
*/
type Body struct {
	Text string `json:"text" binding:"required"`
}
