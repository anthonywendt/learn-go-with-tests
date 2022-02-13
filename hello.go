package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjor, "
const defaultName = "world"

func Hello(name string, language string) string {
	if len(name) == 0 {
		name = defaultName
	}

	return grettingPrefix(language) + name
}

func grettingPrefix(language string) (prefix string) {
	switch language {
	case "french":
		prefix = frenchHelloPrefix
	case "spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Anthony", ""))
}
