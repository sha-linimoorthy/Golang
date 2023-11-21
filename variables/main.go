package main
import "fmt"

func main(){
	var usrname string="Shalini"
	fmt.Println(usrname)
	fmt.Printf("Variable of type is %T\n",usrname)

	var isLogged bool=false
	fmt.Println(isLogged)
	fmt.Printf("Variable of type is %T\n",isLogged)

	var smllval uint8= 255 //256 is out of bound
	fmt.Println(smllval)
	fmt.Printf("Variable of type is %T\n",smllval)

		
}