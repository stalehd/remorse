package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"strings"
)

func main() {
	var fileName, style string
	flag.StringVar(&fileName, "file", "main.go", "Source")
	flag.StringVar(&style, "style", "classic", "Style")
	flag.Parse()

	fs := token.NewFileSet()
	//var v visitor

	f, err := parser.ParseFile(fs, fileName, nil, parser.AllErrors)
	if err != nil {
		fmt.Println("Error parsing: ", err)
		return
	}
	ast.Inspect(f, func(n ast.Node) bool {
		switch t := n.(type) {
		case *ast.Ident:
			if t.Obj != nil {
				if t.Name != "main" {
					t.Name = stringToMorse(t.Name, style)
				}
			}

		default:
			// ignore
		}
		return true
	})
	printer.Fprint(os.Stdout, fs, f)
}

type styles struct {
	dot  string
	dash string
}

var dstyle = map[string]styles{
	"classic": {"dot", "dash"},
	"fx":      {"dit", "diddy"},
	"chinese": {"点", "短跑"},
}

var alphabet = map[rune]string{
	'a': ".-",
	'b': "-...",
	'c': "-.-.",
	'd': "-..",
	'e': ".",
	'f': "..-.",
	'g': "--.",
	'h': "....",
	'i': "..",
	'j': ".---",
	'k': "-.-",
	'l': ".-..",
	'm': "--",
	'n': "-.",
	'o': "---",
	'p': ".--.",
	'q': "--.-",
	'r': ".-.",
	's': "...",
	't': "-",
	'u': "..-",
	'v': "...-",
	'w': ".--",
	'x': "-..-",
	'y': "-.--",
	'z': "--..",
	'1': ".----",
	'2': "..---",
	'3': "...--",
	'4': "....-",
	'5': ".....",
	'6': "-....",
	'7': "--...",
	'8': "---..",
	'9': "----.",
	'0': "------",
}

func stringToMorse(s, style string) string {
	ret := ""
	for _, v := range strings.ToLower(s) {
		code, ok := alphabet[v]
		if !ok {
			ret += string(v)
			continue
		}
		for _, id := range code {
			if id == '-' {
				ret += dstyle[style].dash
				continue
			}
			ret += dstyle[style].dot
		}
	}
	return ret
}
