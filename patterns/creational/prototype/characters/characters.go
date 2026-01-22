package main

import "fmt"

type Character interface {
	clone() Character
}

type Warrior struct {
	weapon string
}

func (w Warrior) clone() Character {
	return Warrior{weapon: w.weapon}
}

type Mage struct {
	spell string
}

func (m Mage) clone() Character {
	return Mage{spell: m.spell}
}

type Thief struct {
	stealth string
}

func (t Thief) clone() Character {
	return Thief{stealth: t.stealth}
}

type CharacterRegistry struct {
	characters map[string]Character
}

func NewCharacterRegistry() *CharacterRegistry {
	return &CharacterRegistry{
		characters: make(map[string]Character),
	}
}
func (cr *CharacterRegistry) Register(name string, character Character) {
	cr.characters[name] = character
}

func (cr *CharacterRegistry) CreateCharacter(name string) Character {
	character, ok := cr.characters[name]
	if !ok {
		fmt.Println("Character not found")
		return nil
	}
	return character.clone()
}

func main() {
	warrior := Warrior{weapon: "sword"}
	mage := Mage{spell: "fireball"}
	thief := Thief{stealth: "cloak"}

	characterRegistry := NewCharacterRegistry()
	characterRegistry.Register("warrior", warrior)
	characterRegistry.Register("mage", mage)
	characterRegistry.Register("thief", thief)

	warriorClone := characterRegistry.CreateCharacter("warrior")
	mageClone := characterRegistry.CreateCharacter("mage")
	thiefClone := characterRegistry.CreateCharacter("thief")

	fmt.Println(warriorClone)
	fmt.Println(mageClone)
	fmt.Println(thiefClone)
}
