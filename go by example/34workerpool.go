// go run 34workerpool.go
package main
import(
	"fmt"
	"time"
)
func worker(id int,jobs<-chan int,results chan<- int){
	for j:=range jobs{
		fmt.Println("worker ",id,"started job ",j)
		time.Sleep(2*time.Second)
		fmt.Println("worker ",id,"finished job ",j)
		results<-j*2
	}
}

func main(){
	const numjob int =5
	jobs:=make(chan int,numjob)
	results:=make(chan int,numjob)
	for w:=1;w<=3;w++{
		go worker(w,jobs,results)
	}
	for j:=1;j<=numjob;j++{
		jobs<-j
	}
	close(jobs)
	for a:=1;a<=numjob;a++{
		<-results
	}

}