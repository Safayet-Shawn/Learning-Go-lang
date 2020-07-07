// go run 20fibnacciNum.go
package main
import"fmt"
func fabno(x uint)uint{
	if x==0{
		return 0
	}else if x==1{
		return 1
	}else{
		return fabno(x-1) + fabno(x-2)
	}
}
func main(){
	fmt.Println(fabno(15))
}