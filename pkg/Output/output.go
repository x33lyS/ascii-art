package ascii_art

import (
	"../Ascii-pkg"
	"fmt"
	"os"
	"strings"
)

func Output(args []string) {

	if len(args) < 2 {

		ascii_art.NoArgument()
		ascii_art.OutputUsage()
		ascii_art.ErrorExit()

	} else if len(args) < 4 {

		ascii_art.NotEnoughArguments()
		ascii_art.OutputUsage()
		ascii_art.ErrorExit()

	} else if len(args) > 4 {

		ascii_art.TooManyArguments()
		ascii_art.OutputUsage()
		ascii_art.ErrorExit()

	}

	if ascii_art.FormatCheck(args) == false {

		ascii_art.FormatFileUnknown()
		ascii_art.OutputUsage()
		ascii_art.ErrorExit()

	}

	var text = args[1]
	var format = args[2]
	var output = ascii_art.OutputFlag(args)
	var result = ""

	var file, err = os.Create(output)

	if err != nil {

		ascii_art.ErrorError()
		ascii_art.ErrorExit()

	}

	defer file.Close()

	if ascii_art.NewlineCheck(text) { // == true

		var splitText = strings.Split(text, "\\n")

		for _, v := range splitText {

			for i := 0; i < 8; i++ {

				for _, v2 := range v {

					result += ascii_art.Scanner(format, 1+int(v2-32)*9+i)

				}

				fmt.Fprintln(file, result)
				result = ""

			}

		}

	} else { // if newlineCheck(text) == false

		for i := 0; i < 8; i++ {

			for _, v := range text {

				result += ascii_art.Scanner(format, 1+int(v-32)*9+i)

			}

			fmt.Fprintln(file, result)
			result = ""

		}

	}

}
