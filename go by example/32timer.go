// go run 32timer.go
package main
import(
	"fmt"
	"time"
)
func main(){
	timer1:=time.NewTimer(2*time.Second)
	<-timer1.C
	fmt.Println("timer1 run")
	 timer2:=time.NewTimer(time.Second)
	 go func(){
	 	<-timer2.C
	 	fmt.Println("timer2 run",<-timer2.C)
	 }()
	 stop:=timer2.Stop()
	 if stop{
	 	fmt.Println("timer2 stopped")
	 }

}