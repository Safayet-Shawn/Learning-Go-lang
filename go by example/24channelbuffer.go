// go run 24channelbuffer.go
package main
import"fmt"
func main(){
	mychanl:=make(chan string,2)
	mychanl<-" i m shawn"
	mychanl<-" learning go"

	sss:= <-mychanl
	ss:= <-mychanl
	fmt.Println(sss)
	fmt.Println(ss)
}