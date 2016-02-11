package main 
import (
  "fmt"
)
func main() {
	/*
	 == igual a
	 != diferente de
	 < menor que
	 > mayor que
	 >= mayor o igual
	 <= menor o igual
	 && AND
	 || OR

	*/
	x := 10
	y := 12
  if x > y {
  	fmt.Printf("%d es mayor que  %d \n",x,y)
  }else if x < y{
  	fmt.Printf("%d es mayor que  %d \n",y,x)
  }else{
  	fmt.Println("iguales")
  }	
}