 // go run 5sliceAppendCopy.go
 package main

import ("fmt"
		"os"
)
func main() {
slice1:= []int{1,2,3,4,5}
slice2:=append(slice1,66,21)
fmt.Println(slice1,slice2)


// func main() {
// slice1 := []int{1,2,3}
// slice2 := append(slice1, 4, 5)
// fmt.Println(slice1, slice2)




//copy in slice
ff:=[]int{1,2,4,5,6}
fk:=make([]int,2)
copy(fk,ff)
fmt.Println(ff,fk)



}
