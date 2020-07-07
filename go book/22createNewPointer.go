// go run 22createNewPointer.go
package main
import "fmt"
func one(xPtr *int) {
*xPtr = 12
}
func main() {
xPtrs := new(int)
one(xPtrs)
fmt.Println(*xPtrs) // x is 1
}
