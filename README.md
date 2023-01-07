# postfix

math postfix

#### postfix.Calc(string) float64 :

example :

postfix.Calc("1 + 2 * 3") // returns 7

## example:
```
package main

import (
	"fmt"

	"github.com/suhaib68/postfix/v1"
)

func main() {
	ans := postfix.Calc("1+2*3")
	fmt.Println(ans)
}
```
