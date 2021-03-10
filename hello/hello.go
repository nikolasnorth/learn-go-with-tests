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
	postfix := getPostfix(language)
	if name != "" {
		postfix = name
	}

	return getPrefix(language) + postfix
}

func getPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func getPostfix(language string) (postfix string) {
	switch language {
	case spanish:
		postfix = spanishWorldPostfix
	case french:
		postfix = frenchWorldPostfix
	default:
		postfix = englishWorldPostfix
	}
	return
}

func main() {
	//fmt.Println(Hello("Nikolas"))
}
