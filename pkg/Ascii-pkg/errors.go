package ascii_art

import (
	"fmt"
	"os"
)

func ErrorExit()  {
	os.Exit(0)
}

func ErrorError()  {
	fmt.Println("ERROR: error.")
}

func NoArgument()  {
	fmt.Println("ERROR: no argument.")
}

func FormatFileUnknown()  {
	fmt.Println("ERROR: format file unknown.")
}

func NotEnoughArguments()  {
	fmt.Println("ERROR: not enough arguments.")
}

func TooManyArguments()  {
	fmt.Println("ERROR: too many arguments.")
}

func incorrectSyntax()  {
	fmt.Println("ERROR: incorrect syntax.")
}

func unknownColor()  {
	fmt.Println("ERROR: unknown color.")
}

func unknownAlign()  {
	fmt.Println("ERROR: unknown alignment.")
}

func ClassicUsage()  {
	fmt.Println("USAGE: ./classic.exe \"text\"")
}

func FsUsage()  {
	fmt.Println("USAGE: ./fs.exe \"text\" <format>")
	fmt.Println("LIST:  standard shadow thinkertoy doom")
}

func ColorUsage()  {
	fmt.Println("USAGE: ./color.exe \"text\" --color=<color>")
	fmt.Println("LIST:  \u001B[38;2;255;255;255mwhite\u001B[0m \u001B[38;2;255;127;0morange\u001B[0m \u001B[38;2;255;0;255mmagenta\u001B[0m \u001B[38;2;63;191;255mlightblue\u001B[0m \u001B[38;2;255;255;0myellow\u001B[0m \u001B[38;2;63;255;0mlime\u001B[0m \u001B[38;2;255;127;127mpink\u001B[0m \u001B[38;2;63;63;63mgray\u001B[0m \u001b[38;2;127;127;127mlightgray\u001B[0m \u001B[38;2;0;255;255mcyan\u001B[0m \u001B[38;2;127;0;191mpurple\u001B[0m \u001B[38;2;0;0;255mblue\u001B[0m \u001B[38;2;127;63;0mbrown\u001B[0m \u001B[38;2;0;255;0mgreen\u001B[0m \u001B[38;2;255;0;0mred\u001B[0m \u001B[38;2;0;0;0mblack\u001B[0m")
}

func OutputUsage()  {
	fmt.Println("USAGE: ./output.exe \"text\" <format> --output=<output>.txt")
	fmt.Println("LIST:  standard shadow thinkertoy doom")
}

func JustifyUsage()  {
	fmt.Println("USAGE: ./justify.exe \"text\" <format> --align=<align>")
	fmt.Println("LIST:  standard shadow thinkertoy doom")
	fmt.Println("LIST:  right left center justify")
}