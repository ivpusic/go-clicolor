package clicolor_test

import cli "github.com/ivpusic/go-clicolor/clicolor"

func ExamplePrint_in() {
	cli.Print("this is green text").In("green")
	// prints: this is green text
}

func ExamplePrint_inFormat() {
	cli.Print("{red}[ERROR]: {blue}Some text in blue. {white}Some text in white. {default}Some error!!!").InFormat()
	// prints: [ERROR]: Some text in blue. Some text in white. Some text in default color
}
