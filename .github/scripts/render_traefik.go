// Render public dynamic Traefik templates with synthetic values only.
package main

import (
    "fmt"
    "os"
    "regexp"
    "strconv"
    "strings"
    "text/template"
)

func main() {
    if len(os.Args) != 2 { panic("Expected one public template path") }
    data, err := os.ReadFile(os.Args[1])
    if err != nil { panic("Cannot read public template") }
    validKey := regexp.MustCompile(`^[A-Z][A-Z0-9_]*$`)
    functions := template.FuncMap{
        "env": func(key string) (string, error) {
            if !validKey.MatchString(key) { return "", fmt.Errorf("Invalid input key") }
            if strings.Contains(key, "LOCAL_") { return "192.0.2.16", nil }
            return "example.invalid", nil
        },
        "quote": func(value interface{}) string { return strconv.Quote(fmt.Sprint(value)) },
        "fail": func(message string) (string, error) { return "", fmt.Errorf("Template refused synthetic inputs") },
    }
    parsed, err := template.New("public").Funcs(functions).Option("missingkey=error").Parse(string(data))
    if err != nil { panic("Cannot parse public Go template") }
    if err := parsed.Execute(os.Stdout, nil); err != nil { panic("Cannot render public Go template") }
}
