# colorize

![](https://img.shields.io/github/issues/michalswi/color)
![](https://img.shields.io/github/forks/michalswi/color)
![](https://img.shields.io/github/stars/michalswi/color)
![](https://img.shields.io/github/last-commit/michalswi/color)

```
go get -u -v github.com/michalswi/color
```

### \# example
```
package main

import "github.com/michalswi/color"

func main() {
	println(color.Format(color.RED, "red"))
	println(color.Format(color.BLUE, "blue"))
	println(color.Format(color.GREEN, "green"))
	println(color.Format(color.YELLOW, "yellow"))
	println(color.Format(color.PURPLE, "purple"))
	println(color.Format(color.PURPLE, "purple"))
	println(color.Format(color.MINGLE, "mingle"))
	println(color.Format(color.WHITE, "white"))
}
```