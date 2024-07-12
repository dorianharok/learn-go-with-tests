package main

const englishPrefixHello = "Hello, "
const spanishPrefixHello = "Hola, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "Spanish" {
		return spanishPrefixHello + name
	}

	return englishPrefixHello + name
}

func main() {

}
