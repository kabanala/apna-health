package main
import "fmt"
func main() {
	fmt.Println("testing the vim go")
	result := add(1,2)
	fmt.Println(result)
	res := subtract(1,2)
	fmt.Println(res)
}

func add(i,j int) int {
  return i + j
}

func subtract(i,j int) int {
	return i + j
}

	
