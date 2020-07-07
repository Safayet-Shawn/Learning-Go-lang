// go run 29non-blocking-channel-operations.go
package main
import"fmt"
func main(){
	messages:=make(chan string)
	signals:=make(chan bool)
	select{
	case rcv:=<-messages:
		fmt.Println("message received",rcv)
	default:
		fmt.Println("no message received")
	}
	msg:="hi"
	select{
	case messages<-msg:
		fmt.Println("sent message",msg)
	default:
		fmt.Println("message not sent")
	}
	select{
	case msg:=<-messages:
		fmt.Println("messages received",msg)
	case signal:=<-signals:
		fmt.Println("signal received",signal)
	default:
		fmt.Println("no activity")
	}
}
