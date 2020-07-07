// go run 10map.go
package main
import"fmt"
func main(){
	m:=make(map[string]int)
	m["shawn"]=1
	m["sujon"]=12
	m["roy"]=11
	m["hisam"]=100
	fmt.Println("map value = ",m)

	value:=m["hisam"]
	fmt.Println(value)
	delete(m,"shawn")
	fmt.Println(m)

	_,presnt:=m["shawn"]
	fmt.Println("presnt",presnt)

	n:=map[string]string{"shawn":"learning go","sujon":"learning java","kobir":"learning python"}
	fmt.Println("key value of map n : ",n)

	delete(n,"shawn")
	fmt.Println(n)
	fmt.Println(n["sujon"])

}