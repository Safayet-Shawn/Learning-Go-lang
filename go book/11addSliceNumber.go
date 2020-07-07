// go run 11addSliceNumber.go
package main
import"fmt"
func main(){
	var total float64=0.00
	slxc:=[]float64{12.87,34.11,56.00,98.16,66.32}
	for _,i:=range slxc{

		total +=float64(i)
	} 
	fmt.Println(total)
}
// func main(){

// x :=[]float64 {33.22,33.54,12.34,48.00,31.88}
// var total float64 = 0
// for i := 0; i < 5; i++ {
// total +=  x[i]
// }
//  fmt.Println(total / float64(len(x)))
// }