package main

import (
	"fmt"

	"github.com/feud72/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	word := "hello"
	def := "Greeting"
	err := dictionary.Add(word, def)
	if err != nil {
		fmt.Println(err)
	}
	definition, err := dictionary.Search(word)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(definition)
	err = dictionary.Add(word, def)
	if err != nil {
		fmt.Println(err)
	}
}
