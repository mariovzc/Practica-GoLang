package main 
import "fmt"
import "strconv"
func main() {
	edad := 25
	edad2,_ := strconv.Atoi("25")
	//convertir integer en string
	fmt.Println("mi Edad es = " + strconv.Itoa(edad))
	fmt.Println (10 * edad2)
}