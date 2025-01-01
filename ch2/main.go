package main

import (
	"flag"
	"fmt"
)

func main() {
	var lang string
	flag.StringVar(&lang, "lang", "en", "language")
	flag.Parse()
	greeting := greet(language(lang))
	fmt.Println(greeting)
}

type language string

var phrasebook = map[language]string{
	"en": "Hello world",       // English
	"fr": "Bonjour le monde",  // French
	"he": "שלום עולם",         // Hebrew
	"ur": "ہیلو دنیا",         // Urdu
	"vi": "Xin chào Thế Giới", // Vietnamese
}

func greet(l language) string {
	greeting, ok := phrasebook[l]
	if !ok {
		greeting = fmt.Sprintf("unsupported language: %q", l)
	}
	return greeting
}
