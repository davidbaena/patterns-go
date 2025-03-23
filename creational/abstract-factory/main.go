package main

import "fmt"

func main() {
	nikeFactory, _ := ProvideSportFactory(NikeFactory)
	adidasFactory, _ := ProvideSportFactory(AdidasFactory)

	nikeShirt := nikeFactory.CreateShirt()
	nikeShoe := nikeFactory.CreateShoe()
	nikeHat := nikeFactory.CreateHat() // Create Nike hat

	printShirtDetails(nikeShirt)
	printShoeDetails(nikeShoe)
	printHatDetails(nikeHat) // Print Nike hat details

	adidasShirt := adidasFactory.CreateShirt()
	adidasShoe := adidasFactory.CreateShoe()
	adidasHat := adidasFactory.CreateHat() // Create Adidas hat

	printShirtDetails(adidasShirt)
	printShoeDetails(adidasShoe)
	printHatDetails(adidasHat) // Print Adidas hat details
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

func printHatDetails(h IHat) {
	fmt.Println("Hat:")
	fmt.Printf("Size: %d\n", h.Size())
	fmt.Printf("Color: %s\n", h.Color())
	fmt.Println()
}
