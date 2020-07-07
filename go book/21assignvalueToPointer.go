// go run 21assignvalueToPointer.go
package main
import"fmt"
func ppp(x *int){
	*x=121
}

func main(){
	xs:=14
	ys:=&xs
	ppp(ys)
	fmt.Println(xs)
}