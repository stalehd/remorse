# remorse - it's all in the name

Launch with `go build && ./remorse`. This will print out the source code of the
program with all the functions, and variables in morse code. It's a bit broken
and not ready for prime.

Running it with the `fx` style makes HAM radio enthusiasts happy.

```golang
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
        var ditditdiddyditditditditdiddyditditditdiddyditditdiddydiddydiddydit, ditditditdiddydiddyditdiddydiddyditdiddyditditdit string
        flag.StringVar(&ditditdiddyditditditditdiddyditditditdiddyditditdiddydiddydiddydit, "file", "main.go", "Source")
        flag.StringVar(&ditditditdiddydiddyditdiddydiddyditdiddyditditdit, "style", "classic", "Style")
        flag.Parse()

        ditditdiddyditditditdit := token.NewFileSet()

        ditditdiddydit, ditditdiddyditditdiddydit := parser.ParseFile(ditditdiddyditditditdit, ditditdiddyditditditditdiddyditditditdiddyditditdiddydiddydiddydit, nil, parser.AllErrors)
        if ditditdiddyditditdiddydit != nil {
                fmt.Println("Error parsing: ", ditditdiddyditditdiddydit)
                return
        }
        ast.Inspect(ditditdiddydit, func(diddydit ast.Node) bool {
                switch diddy := diddydit.(type) {
                case *ast.Ident:
                        if diddy.Obj != nil {
                                if diddy.Name != "main" {
                                        diddy.Name = ditditditdiddyditdiddyditditditdiddyditdiddydiddyditdiddydiddydiddydiddydiddydiddydiddydiddydiddyditdiddyditditditditdit(diddy.Name, ditditditdiddydiddyditdiddydiddyditdiddyditditdit)
                                }
                        }

                default:

                }
                return true
        })
        printer.Fprint(os.Stdout, ditditdiddyditditditdit, ditditdiddydit)
}

type ditditditdiddydiddyditdiddydiddyditdiddyditditditditditdit struct {
        diddyditditdiddydiddydiddydiddy                 string
        diddyditditditdiddyditditditditditditdit        string
}

var diddyditditditditditdiddydiddyditdiddydiddyditdiddyditditdit = map[string]ditditditdiddydiddyditdiddydiddyditdiddyditditditditditdit{
        "classic":      {"dot", "dash"},
        "fx":           {"dit", "diddy"},
        "chinese":      {"点", "短跑"},
}

var ditdiddyditdiddyditditditdiddydiddyditditditditditditdiddydiddyditditditditdiddy = map[rune]string{
        'a':    ".-",
        'b':    "-...",
        'c':    "-.-.",
        'd':    "-..",
        'e':    ".",
        'f':    "..-.",
        'g':    "--.",
        'h':    "....",
        'i':    "..",
        'j':    ".---",
        'k':    "-.-",
        'l':    ".-..",
        'm':    "--",
        'n':    "-.",
        'o':    "---",
        'p':    ".--.",
        'q':    "--.-",
        'r':    ".-.",
        's':    "...",
        't':    "-",
        'u':    "..-",
        'v':    "...-",
        'w':    ".--",
        'x':    "-..-",
        'y':    "-.--",
        'z':    "--..",
        '1':    ".----",
        '2':    "..---",
        '3':    "...--",
        '4':    "....-",
        '5':    ".....",
        '6':    "-....",
        '7':    "--...",
        '8':    "---..",
        '9':    "----.",
        '0':    "------",
}

func ditditditdiddyditdiddyditditditdiddyditdiddydiddyditdiddydiddydiddydiddydiddydiddydiddydiddydiddyditdiddyditditditditdit(ditditdit, ditditditdiddydiddyditdiddydiddyditdiddyditditdit string) string {
        ditdiddyditditdiddy := ""
        for _, ditditditdiddy := range strings.ToLower(ditditdit) {
                diddyditdiddyditdiddydiddydiddydiddyditditdit, diddydiddydiddydiddyditdiddy := ditdiddyditdiddyditditditdiddydiddyditditditditditditdiddydiddyditditditditdiddy[ditditditdiddy]
                if !diddydiddydiddydiddyditdiddy {
                        ditdiddyditditdiddy += string(ditditditdiddy)
                        continue
                }
                for _, ditditdiddyditdit := range diddyditdiddyditdiddydiddydiddydiddyditditdit {
                        if ditditdiddyditdit == '-' {
                                ditdiddyditditdiddy += diddyditditditditditdiddydiddyditdiddydiddyditdiddyditditdit[ditditditdiddydiddyditdiddydiddyditdiddyditditdit].diddyditditditdiddyditditditditditditdit
                                continue
                        }
                        ditdiddyditditdiddy += diddyditditditditditdiddydiddyditdiddydiddyditdiddyditditdit[ditditditdiddydiddyditdiddydiddyditdiddyditditdit].diddyditditdiddydiddydiddydiddy
                }
        }
        return ditdiddyditditdiddy
}
```
