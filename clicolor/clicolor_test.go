package clicolor_test

import cli "github.com/ivpusic/go-clicolor/clicolor"

func ExamplePrint_in() {
	cli.Print("this is green text").In("green")
	// prints: this is green text
}

func ExamplePrint_inFormat() {
	cli.Print("{red}Some text in red. {white}Some text in white. {default}Some text in default color.").InFormat()
	// prints: Some text in red. Some text in white. Some text in default color.
}
