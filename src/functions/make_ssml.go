package functions

import (
	"fmt"
)

func MakeSSML(text string) string {

	output := BreakBeginingEnd(text)

	//print final ssml
	fmt.Printf("\n%v\n\n", output)

	return output
}
