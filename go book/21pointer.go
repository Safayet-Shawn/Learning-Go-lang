// go run 21pointer.go
package main
import"fmt"
func zero(xpnt *int){
	*xpnt=122
}
func main(){
	x:=5
	zero(&x)
	fmt.Println(x)
}