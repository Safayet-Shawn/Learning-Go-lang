// go run 14variadicFunctions.go
package main
import"fmt"
func sum(x ...int){
	fmt.Println(x,"")
	ss:=0
for _,s:=range x{
	ss+=s
}
fmt.Println(ss)
}
func main(){
sum(1,4,877)

sl:=[]int{12,456,23,98,23}
sum(sl ...)
}