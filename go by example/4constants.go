// go run 4constants.go
package main
import(
		"fmt"
		"math"
	)

const s string="Safayet Shawn"
func main(){
	fmt.Println(s)
	const n=50000
	const d=3e20/n
	fmt.Println(d)
	fmt.Println(int(d))
	fmt.Println(math.Sin(n))
}