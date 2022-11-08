package main

import "fmt"

type GoEnum = int

const (
	Foo GoEnum = iota
	Bar
	Baz
)

func main() {
	fmt.Printf("test")
}

/**

func ReturnsError(value int) error {
	return fmt.Errorf("erreur de qualité prénium %v", value)
}

type Foo struct {}
func (f *Foo) thisisforfoo() error {
	return fmt.Errorf("erreur de qualité prénium")
}

func main (){
	fmt.Println("Hello World")
	err := ReturnsError(42)
	fmt.Println(err)
	err2 := Foo{}
	fmt.Println(err2.thisisforfoo())
}
*/
