// go run 12average.go
package main
import"fmt"

func avgnum(avg []float64)float64{

var total float64=0.00
for _,v:=range avg{
	total+=v
}
return total/float64(len(avg))
}
func main(){
x:=[]float64{23.33,44.33,4467.87,55.33,223.44,44.43}
fmt.Println(avgnum(x))
}