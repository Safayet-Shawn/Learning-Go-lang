// go run 19method.go
package main
import"fmt"
type rect struct{
	width,height int
}
func (r rect)area()int{
	return r.width * r.height
}
func (r *rect) perim()int{
	return 2*r.width + 2*r.height
}
func main(){
	rr:=rect{width:10,height:5}
	fmt.Println("area: ",rr.area())
	fmt.Println("perim: ",rr.perim())
	

	rrr:=rect{width:100,height:50}

	rrpoint:=&rrr

	fmt.Println("area: ",rrpoint.area())
	fmt.Println("perim: ",rrpoint.perim())


}