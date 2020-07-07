// go run map2.go
package main
import "fmt"
func main(){
	element:=map[string]string{
		"Shawn":"learning go",
		"Sojib":"learning java",
		"Sujon":"learning php",
		"Rakib":"learning python",
	}
	if el,ok:=element["Shawn"];ok{
		fmt.Println(el,ok)
		fmt.Println(el)
	}
}