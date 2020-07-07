// go run 8array.go
package main
import"fmt"
func main(){
	 a :=[5]int{}
	// var a [6]int
	fmt.Println(a)
	a[3]=54
	a[1]=76
	fmt.Println(a)
	fmt.Println(a[1],a[3])
	fmt.Println("length of array a = ",len(a))

	var two [2][3]int
	for i:=0;i<2;i++{
		for j:=0;j<3;j++{
			two[i][j]=i+j
		}
	}
	fmt.Println("2d: ",two)

}