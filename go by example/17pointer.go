// go run 17pointer.go
package main
import"fmt"
func value(i int){
	i=100
	// fmt.Println("value of i ",i)
}
func pointvl(i *int){
	*i=100
	// fmt.Println("value of i in pointvl",i)
}
func main(){
i:=10
value(i)
fmt.Println("i value in value :",i)
pointvl(&i)
fmt.Println("i pointvalue in value :",i)

fmt.Println("memoryaddress of i : ",&i)

}