# postfix

math postfix

#### postfix.Postfix(string) []string :

example :

postfix.Postfix("1+2*3") // returns [1 2 3 * +]

#### postfix.Calc([]string) float64 :

example :

postfix.Calc([]string{"1", "2", "3", "*", "+"}) // returns 7

## example:
```
package main

import (
	"fmt"

	"github.com/suhaib68/postfix/v1"
)

func main() {
	P := postfix.Postfix("1+2*3")
	ans, _ := postfix.Calc(P)
	fmt.Println(ans)
}
```
