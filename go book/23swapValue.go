// go run 23swapValue.go
package main
import"fmt"
func swap(xx,yy *int){
*xx=18
*yy=12
}
func main(){
x:=12
y:=18
swap(&x,&y)
fmt.Println(x,y)
}