// go run 31range-over-channels.go
package main
import"fmt"
func main(){
	number:=make(chan string,2)
	number<-"one"
	number<-"two"
	close(number)
	for num:=range number{
		fmt.Println(num)
	}
}