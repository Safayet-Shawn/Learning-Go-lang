// go run 14Closure.go
package main
import"fmt"
func makeEven() func()uint{
	i:=uint(0)
	return func()(retea uint){
		retea=i
		i+=2
		return 
	}
}
func main(){
odnum:=makeEven()
fmt.Println(odnum())
fmt.Println(odnum())
fmt.Println(odnum())
}