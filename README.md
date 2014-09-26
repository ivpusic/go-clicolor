go-clicolor
===========

Package provides ability to print colored text on stdout

## Installation
```
go get github.com/ivpusic/go-clicolor/clicolor
```

Package provides ``Print`` function followed by two functions for actual outputing string to ``stdout``.
You can use ``In`` function by providing one of supported color names. Whole text will be printed in one color.
If you want to print multiple colors for some text, you can use ``InFormat`` function. Function is not accepting 
any arguments, but you need to provide proper formatted string to ``Print`` function. All text after ``{somecolor}`` will
be printed in that color. You can repeat this multiple times. Let's see example:

## Example of usage:
```Go
package main

import cli "github.com/ivpusic/go-clicolor/clicolor"

func main() {
	cli.Print("this is text").In("green")
	cli.Print("{red}[ERROR]: {white}Some text in white. {default}Some text in default color").InFormat()
}
```

## Supported colors
- black
- red
- green
- yellow
- blue
- magenta
- cyan
- white
- default
