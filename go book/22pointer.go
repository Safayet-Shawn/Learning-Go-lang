// go run 22pointer.go
package main
import"fmt"
func sqr(x *float64){
	*x=3
		*x=*x**x
}
func main(){
	wx:=1.5
	sqr(&wx)
	fmt.Println(wx)
}