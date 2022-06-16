package functions

func BreakBeginingEnd(text string) string {
	output := "<speak> <break time=\"500ms\"/> " + text + " <break time=\"500ms\"/> </speak>"
	return output
}
