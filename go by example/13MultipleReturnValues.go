// go run 13MultipleReturnValues.go
package main
import"fmt"
const pi int=3
func mvalue(r int)(int,int){
	return 2*pi*r*r, 2*pi
}
func addMul(aa int)(int ,int){
	return aa*pi,aa+pi
}
func main(){
a,b:=mvalue(7)
fmt.Println(a,b)
fmt.Println(a)
fmt.Println(b)

_,aaa:=addMul(212)
fmt.Println(aaa)

vv,_:=addMul(212)
fmt.Println(vv)
}