// go run 12function.go
package main
import"fmt"


func add( a ,b int)int{
	return a+b
}
func addThree(a,b,c float64)float64{
	return a+b+c
}
func main(){
fmt.Println("addition of 11,23 number",add(11,23))
fmt.Println("addition of 65.32,78.32,98.00 number : ",addThree(65.32,78.32,98.00))

}