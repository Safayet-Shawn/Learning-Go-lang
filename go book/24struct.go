// go run 24struct.go
package main
import("fmt"
		"math")
type circle struct{

	r float64
}
func circleArea(c *circle)float64{
 return math.Pi * c.r*c.r
}
func main(){
 cs:=circle{2}
 fmt.Println(circleArea(&cs))
}