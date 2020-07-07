// go run largenumVarid.go
package main
import"fmt"
func largenumVarid(gnum ...int)int{
	var great int=0
	for _,num:=range gnum{
		if great<num{
			great =num
		}
	}

		return great
}
func main(){
	x:=[]int{11,76,452,90,124,65}
	fmt.Println(largenumVarid(x ...))
}