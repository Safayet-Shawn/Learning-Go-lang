// go run 10findMaxNumber.go
package main
import "fmt"
func main(){
	x:=[]int{12,44,11,567,90,21}
	max:=x[0]
	for _,value:=range x{
		if value>max{
			max=value
		}
	}
	fmt.Println(x)
	fmt.Println("max number is ",max)
}