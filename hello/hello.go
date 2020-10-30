package main

import (
	"fmt"

	"example.com/greetings"
	"github.com/fatih/color"
	log "github.com/koding/logging"
)

var sayi = 13

const aciklama = "örnek"

func hesapla(x, y, z int) (int, bool) {
	var a int = x + y + z
	b := 20 < a
	return a, b
}

func main() {
	// Get a greeting message and print it.
	color.Red("Hello")
	message := greetings.Hello("Gladys")
	fmt.Println(message)
	log.Info("app started")

	var empty int
	number := 8

	toplam, buyuk := hesapla(empty, number, sayi)
	fmt.Println(toplam, buyuk)

}

//private
func add(x, y int) int {
	return x + y
}

//Public
func Multiply(x, y int) int {
	return x * y
}
