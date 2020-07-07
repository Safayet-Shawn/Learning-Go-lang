// go run 19makeOddGeneratop.go
package main
import"fmt"
func makeOdd()func()int{
	i:=int(-1)
 return func()(odd int){
 	odd =i

 	i+=2
 	return i
 }
}

func main(){
	oddnum:=makeOdd()
	fmt.Println(oddnum())
	fmt.Println(oddnum())
	fmt.Println(oddnum())

}