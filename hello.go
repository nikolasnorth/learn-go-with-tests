package main

const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name, language string) string {
	if language == spanish {
		if name == "" {
			name = "Mundo"
		}
		return spanishHelloPrefix + name
	} else {
		if name == "" {
			name = "World"
		}
		return englishHelloPrefix + name
	}
}

func main() {
	//fmt.Println(Hello("Nikolas"))
}
