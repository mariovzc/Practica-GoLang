package main 
import(
  "fmt"
  "bufio"
  "os"
)
func main() {
	//edad := 25
	//var bander bool
	//precio := 14.5
	/*var edad int
	fmt.Println("Coloca tu Edad:")
	fmt.Scanf("%d\n", &edad)
	fmt.Printf("Mi edad es = %d \n",edad)*/
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("ingresa tu nombre")
	nombre,err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Hola "+ nombre)
	}
}