// go run 6.closure.go
package main
import "fmt"
func main() {
// x := 0
// increment := func() int {
// x++
// return x
// }
// fmt.Println(increment())
// fmt.Println(increment())
add:=func (xx,xy int)int{
	return xx+xy
}
fmt.Println(add(1,2))
}