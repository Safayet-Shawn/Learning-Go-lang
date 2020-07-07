// go run 25Method.go
package main
import"fmt"
type area struct{
	height float64
	width float64
}
func (a *area)area()float64{
	return a.height*a.width
}
func main(){
 ar :=  area{12.22,10.18}
 fmt.Println(ar.area())
}