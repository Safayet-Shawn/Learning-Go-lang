// go run 7recover.go
package main
import "fmt"
func main() {
defer func() {
str := recover()
fmt.Println(str)
}()
panic("PANIC")
}

x :=[]float64 {33.22,33.54,12.34,48.00,31.88}
var total float64 = 0
for i := 0; i < 5; i++ {
total +=  x[i]
}
 fmt.Println(total / float64(len(x)))