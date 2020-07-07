// go run 17passSlice.go
package main
import"fmt"
func add(x []int)int{
	var total int =0
   for _,v:=range x{
   	total+=v
   }
   return total
}
func main(){
	fmt.Println(add([]int{11,65,43}))
}