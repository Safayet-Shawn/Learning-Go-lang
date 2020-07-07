// go run 18struct.go
package main
import"fmt"
	type person struct{
		name string
		age int 
	}

func newperson(getname string) *person{
	p:=person{name:getname}
	p.age=42
	return &p
}

func main(){
fmt.Println(person{name:"shawn",age:25})
fmt.Println(person{name:"sojib",age:55})
fmt.Println(&person{name:"sojib",age:55})
fmt.Println(person{name:"shawn"})
s:=person{age:33}
fmt.Println(s.age)
s.name="farukh"
fmt.Println(s.name)
fmt.Println(newperson("sandip"))


}