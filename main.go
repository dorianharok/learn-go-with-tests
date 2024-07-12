package main

const englishPrefixHello = "Hello, "
const spanishPrefixHello = "Hola, "
const frenchPrefixHello = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishPrefixHello

	switch language {
	case "Spanish":
		prefix = spanishPrefixHello
	case "French":
		prefix = frenchPrefixHello
	}

	return prefix + name
}

func main() {

}
