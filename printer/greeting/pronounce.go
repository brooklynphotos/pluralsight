package greeting

import "fmt"

type Pronounceable interface {
	Pronounce(phrase string)
}

type Child struct {
	Name string
	Age  int8
}

func (child Child) Pronounce(phrase string) {
	fmt.Printf("%s of age %d says %s\n", child.Name, child.Age, phrase)
}

type MyInt int

func (mi MyInt) Pronounce(phrase string) {
	fmt.Printf("%d = %s\n", mi, phrase)
}
