package main

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const englishWorldPostfix = "World"
const spanishWorldPostfix = "Mundo"
const frenchWorldPostfix = "Monde"

func Hello(name, language string) string {
	prefix := englishHelloPrefix
	postfix := englishWorldPostfix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
		postfix = spanishWorldPostfix
	case french:
		prefix = frenchHelloPrefix
		postfix = frenchWorldPostfix
	}

	if name != "" {
		postfix = name
	}

	return prefix + postfix
}

func main() {
	//fmt.Println(Hello("Nikolas"))
}
