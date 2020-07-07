// go run 15factorial.go
package main
import"fmt"
func fac(x uint)uint{
	if x==0{
		return 1
	}
	return x * fac(x-1)
}
func main(){
	fmt.Println(fac(5))
}