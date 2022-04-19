package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {
	name string
}

type spanishBot struct {
	name string
}

//////////////   REAL DIFFERENT   ////////////////
func (b englishBot) getGreeting() string {
	return fmt.Sprintf("%s says: Hello!", b.name)
}

func (b spanishBot) getGreeting() string {
	return fmt.Sprintf("%s says: Hola!", b.name)
}

/////////////////////////////////////////////////

//////////////   REAL SIMILAR   ////////////////
// func printGreeting(b englishBot) {
// 	fmt.Println(b.getGreeting())
// }

// func printGreeting(b spanishBot) {
// 	fmt.Println(b.getGreeting())
// }

///////////////////////////////////////////////

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{name: "William"}
	sb := spanishBot{name: "Jose"}

	printGreeting(eb)
	printGreeting(sb)
}
