package main
type Language int

func (lang Language) String () string {

	if int(lang) < len(Language_Display){
		return Language_Display[lang]
	} else {
		return "Not a Language"
	}
}

var Language_Display []string = []string{
	"Go",
	"Typescript",
	"JavaScript",
	"C#",
	"C++",
	"Java",
	"Scala",
	"BrainFuck",
	"Erlang",
	"Node.JS",
	"C",
	"PHP",
}

const (
	GO Language = iota
	TYPESCRIPT
	JAVASCRIPT
	C_SHARP
	CPP
	JAVA
	SCALA
	BRAINFUCK
	ERLANG
	NODE
	C
	PHP
	NotALanguage

)
