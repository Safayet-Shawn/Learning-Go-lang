// go run 6elseif.go
package main
import"fmt"
func main(){
	s:=9;
	if s<0{
		fmt.Println("S is a negative value")
	}else if s==9{
		fmt.Println("Value of s = 9,and digit num is 1")
	}else{
		fmt.Println("cannt get the value")
	}

	if 9%2==0{
		fmt.Println("it's an even number")

	}else{

		fmt.Print("it's an odd (",s ,") number")
	}
}