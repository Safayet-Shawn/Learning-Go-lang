// go run 20interface.go
package main
import(
	"fmt"
	"math"
)
type geomatric interface{
	area()float64
	perim()float64
}
type rect struct{
	width float64
	height float64
}
type circle struct{
	radius float64
}
func (r rect)area()float64{
	return r.width*r.height
}
func (r rect)perim()float64{
	return 2*r.width + 2*r.height
}
func(c circle)area()float64{
	return math.Pi*c.radius*c.radius
}
func(c circle)perim()float64{
	return 2*math.Pi*c.radius
}
func measure(g geomatric){
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
func main(){
	r:=rect{width:3.00,height:4.00}
	c:=circle{radius:5.00}
	measure(r)
	measure(c)
}