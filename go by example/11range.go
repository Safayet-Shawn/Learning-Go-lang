// go run 11range.go
package main
import"fmt"
func main(){
	a:=[]float64{199.01,16.97,11.10,121.11}
	add:=0.0;
	for _,num:=range a{
		add +=num
	}


	fmt.Println("total : ",add)
	s:=map[string]string{
		"a":"apple",
		"b":"beg",
		"c":"car",
	}

	for key,value:=range s{
		fmt.Printf(" %s ->  %s\n",key,value)
	}
	for _,value:=range s{
		fmt.Println(value)
	}

	for keeyy:=range s{
		fmt.Println(keeyy)
	}

	aa:=[]int{12,11,23,43,154}
	for i,v:=range aa{
		if v==154{
			fmt.Println("index of 11 = ",i)
		}
	}


	for index,cha:=range "go lang"{
		fmt.Println(index,cha)
	}

}