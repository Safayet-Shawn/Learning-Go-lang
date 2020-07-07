// go run 13VariadicFunctions.go
package main
import"fmt"
func add(num ...int)int{
	var total int = 0
for _,v:=range num{
	total+=v
}
return total

}
func main(){
	
	xx:=[]int{1,5,7,8,23,65,87}
	fmt.Println(add(xx...))
	// or just use

	// fmt.Println(add(1,66,7))

}
