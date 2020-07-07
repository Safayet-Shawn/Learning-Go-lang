// go run 33ticker.go
package main
import(
	"fmt"
	"time"
)
func main(){
	ticker:=time.NewTicker(500*time.Millisecond)
	done:=make(chan bool)
	go func(){
		for{
		select{
		case tk:= <-ticker.C:
			fmt.Println("Ticker running at : ",tk)
		case <-done:
			return
		}
	}
	}()
	time.Sleep(2100*time.Millisecond)
	ticker.Stop()
	fmt.Println("ticker stoppted")

}