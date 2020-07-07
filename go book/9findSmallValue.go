// go run 9findSmallValue.go
package main
import "fmt"
func main(){
	x:=[]int{23,43,12,11,87,32}
	min:=x[0]
	for _,i:=range x{
		if i<min{
			min=i
		}
	}
	fmt.Println(x)
	fmt.Println("minimum value = ",min)
}