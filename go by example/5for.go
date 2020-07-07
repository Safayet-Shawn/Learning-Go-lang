// go run 5for.go
package main
import"fmt"
func main(){
i:=0
for i<3{
	fmt.Println(i)
	i++
}
for j:=3;j<8;j++{
	fmt.Println(j)
}
for k:=0;k<20;k++{
	if k%2==0{
		fmt.Println("even number between 20 = ",k)
	}
}
for{
	fmt.Println("bool")
	break
}
}