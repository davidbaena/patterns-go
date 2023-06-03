package main

import "fmt"

func main() {
	nikeFactory, _ := ProvideSportFactory(NikeFactory)
	adidasFactory, _ := ProvideSportFactory(AdidasFactory)

	nikeShirt := nikeFactory.CreateShirt()
	nikeShoe := nikeFactory.CreateShoe()

	printShirtDetails(nikeShirt)
	printShoeDetails(nikeShoe)

	adidasShirt := adidasFactory.CreateShirt()
	adidasShoe := adidasFactory.CreateShoe()

	printShirtDetails(adidasShirt)
	printShoeDetails(adidasShoe)
}

func printShoeDetails(s IShoe) {
	fmt.Println("Shoe:")
	fmt.Printf("Size: %d\n", s.Size())
	fmt.Printf("Color: %s\n", s.Color())
	fmt.Println()
}
func printShirtDetails(sh IShirt) {
	fmt.Println("Shirt:")
	fmt.Printf("Size: %d\n", sh.Size())
	fmt.Printf("Has Sleeves: %t\n", sh.HasSleeves())
	fmt.Println()

}
