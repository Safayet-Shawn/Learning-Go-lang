// go run 9slice.go
package main
import"fmt"
func main(){
	s:=make([]int ,3)
	fmt.Println(s)

s[0]=12
s[1]=51
s[2]=752
fmt.Println(s)


ss:=[]string{"shawn","suvo","rakib","rokey","dines"}
for i:=0;i<len(ss);i++{
	fmt.Println(ss[i])
}
l:=ss[:4]
fmt.Println(l)
ll:=ss[1:]
fmt.Println(ll)

l=ss[1:4]
fmt.Println(l)
ss=append(ss,"sujon","prokash")
ss=append(ss,"dipok vai")
fmt.Println(ss)
cs:=make([]string,len(ss))
copy(cs,ss)
fmt.Println(cs)
cs=append(cs,"kasem","kobir")
fmt.Println(cs)
}