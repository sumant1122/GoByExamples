package main

import "fmt"

func main() {

	var creature string = "shark"
	var pointer *string = &creature

	fmt.Println("creature: ", creature)
	fmt.Println("pointer: ", pointer)
	fmt.Println("*pointer: ", *pointer)

	*pointer = "jellyfish"

	fmt.Println("*pointer: ", *pointer)

}
