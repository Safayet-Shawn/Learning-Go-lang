// go run 8fourthlemnt.go
package main
import"fmt"
func main(){
	// x:=[]int{12,34,234,123}
	// a:=x[3]
	// fmt.Println(a)
	x := [6]string{"a","b","c","d","e","f"}
	a:=x[2:5]
	fmt.Println(a)
	fmt.Println("c,d,e")
}
