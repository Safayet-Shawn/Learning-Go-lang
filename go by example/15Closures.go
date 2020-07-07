// go run 15Closures.go
package main
import"fmt"
func cls() func()int{
	i:=10
	return func()int{
		i++
		return i
	}
}
func main(){
	nextf:=cls()
	fmt.Println(nextf())
	fmt.Println(nextf())
	fmt.Println(nextf())

	nxtf:=cls()
	fmt.Println(nxtf())
}