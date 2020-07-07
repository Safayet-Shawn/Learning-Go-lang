// go run 23channel.go
package main
import"fmt"
func main(){
	message:=make(chan string)
	go func(){
		message<- "sending message to channel"
	}()
	msg:= <-message
	fmt.Println(msg)
}
// package main
// import (
// "fmt"
// "time"
// )
// func pinger(c chan string) {
// for i := 0; i<3; i++ {
// c <- "ping"
// }
// }
// func printer(c chan string) {
// for {
// msg := <- c
// fmt.Println(msg)
// time.Sleep(time.Second * 1)
// }
// }
// func main() {
// var c chan string = make(chan string)
// go pinger(c)
// go printer(c)
// var input string
// fmt.Scanln(&input)
// }