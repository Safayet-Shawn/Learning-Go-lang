// go run map.go

package main
import"fmt"
func main(){
	x:=make(map[string]int)
	x["shawn"]=100
	x["sojib"]=121
	x["sonjoy"]=40
	x["roni"]=77
	fmt.Println(x)

	delete(x,"sojib")
	fmt.Println(x)
}