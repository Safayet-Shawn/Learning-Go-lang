// go run 18makeHalf.go
package main
import"fmt"
func makeHalf(x int)(int,bool){
	if x%2 !=0{
		return x/2,false
	}else{
		return x/2,true
	}
	
}
func main(){
	mh,okd:= makeHalf(1861)
	fmt.Println(mh,okd)
}
