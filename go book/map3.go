// go run map3.go
package main
import"fmt"
func main(){
	element:=map[string]map[string]string{
		"Shawn":map[string]string{
			"phn":"01726681903",
			"address":"132/3,Dashpara ,Narsingdi",
		},
		"sujon":map[string]string{
			"phn":"016553212",
			"address":"33/2,Palbari",
		},
		"rijbi":map[string]string{
			"phn":"01937052288",
			"address":"Dashpara 121",
		},
		"himel":map[string]string{
			"phn":"016212321",
			"address":"12/e Palbari",
		},
	}
	if Bio,statusok:=element["rijbi"];statusok{
		fmt.Printf("phn number of rijbi : %s\n",Bio["phn"])
		fmt.Printf("address of  rijbi: %s\n",Bio["address"])
		fmt.Printf("Status :%d",statusok)
	}
}