// go run 30closing-channels.go
package main
import(
	"fmt"
	"time"
)
func main(){
	jobs:=make(chan int,5)
	done:=make(chan bool)
	go func(){
		for{
			j,m:=<-jobs
			if m{
				fmt.Println("received job no : ",j,m)
			} else{
				fmt.Println("received all jobs",m)
				done<-true
				return
			}
		}
	}()
	for j:=1;j<=3;j++{
		jobs<-j
		fmt.Println("sending job no : ",j)
		time.Sleep(3*time.Second)
		}
		close(jobs)
		fmt.Println("all job sents ")
		fmt.Println(<-done)
	}

