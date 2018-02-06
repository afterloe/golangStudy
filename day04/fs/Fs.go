package fs

import "fmt"

var MIN_LINE int

func ReadFile(name string) string{
	fmt.Printf("try to read file -> %s \n", name)

	return "done"
}

func sayHello()  {
	fmt.Println("say a word")
}

func init() {
	MIN_LINE = 123
	sayHello()
}