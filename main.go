package main

const englishPrefixHello = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}

	return englishPrefixHello + name
}

func main() {

}
