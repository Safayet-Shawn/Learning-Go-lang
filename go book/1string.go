// go run 1string.go
package main

import "fmt"
func main() {
// fmt.Println(len("Hello World"))
// fmt.Println("Hello World"[3])
// fmt.Println("Hello " + "World")
for i:=0;i<=20;i++{
	if i%2==0 {
		fmt.Printf("\n%d is even\n",i)

	}else{
		fmt.Printf("\n%d is odd\n\n",i)
	}
}

}