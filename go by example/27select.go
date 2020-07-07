// go run 27select.go 
package main
import(
	"fmt"
	"time"
)
func main(){
	c1:=make(chan string)
	c2:=make(chan string)

	go func(){
		time.Sleep(1*time.Second)
		c1<-"first"
	}()
	go func(){
		time.Sleep(1*time.Second)
		c2<-"Second"
	}()

		select {
		case msg1:= <-c1:
			fmt.Println("receive from c1 ",msg1)
		
		case msg2:= <-c2:
			fmt.Println("receive from c2 ",msg2)
		}
		
	}


