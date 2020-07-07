// go run 7switchcase.go
package main
import(
	"fmt"
	"time"
)
func main(){
	i:=2
	switch i{
	case 1:
		fmt.Print("value of i = 1")
	case 2:

		fmt.Print("value of i = 2")

	case 3:

		fmt.Print("value of i = 3")
	}



	switch time.Now().Weekday(){
	case time.Friday,time.Thursday :
		fmt.Println("weekend")
	default :
		fmt.Println("weekday")
	}


	t:=time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("its Noon")
	
	default :
		fmt.Println("after Noon")
	}

	getType:=func(i interface{}){
		switch fd:=i.(type){
		case bool:
			fmt.Println("its boolean")
		case int:
			fmt.Println("its int")
		default:
			fmt.Printf("default value is %T",fd)
		}
	}
	getType(123)
	getType(true)
	getType("its a string")

}