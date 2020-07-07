// go run 26channel-directions.go
package main
import"fmt"

func ping(pings chan<- string,msg string){
	pings<-msg
}
func pong(pings <-chan string,pongs chan<- string){
	msg:= <-pings
	pongs <-msg

}

func main(){
	pings:= make(chan string,1)
	pongs:= make(chan string,1)
	ping(pings,"receive from pings channel")
	pong(pings,pongs)
	fmt.Println(<-pongs)
	fmt.Println(<-pings)


}