package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
)

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

func main() {
	port := flag.Int("port", 3000, "the port to start the CYOA web application on")
	filename := flag.String("file", "gopher.json", "the JSON file with the CYOA")
	flag.Parse()
	fmt.Printf("using the story in %s. \n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	fmt.Print(f)

	story, err := JsonStory(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)

	h := NewHandler(story)
	fmt.Println("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Printf(":%d", *port), handler))
}

func JsonStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

func init() {

	tpl = template.Must(template.New("").Parse(defaultHandlerTemp))
}

var tpl *template.Template

func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

func (h handler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, h.s["intro"])
	if err != nil {
		panic(err)
	}
}

var defaultHandlerTemp = `
  <!DOCTYPE html>
<head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
</head>
<body>
    <h1>{{.Title}}</h1>
        {{range .Paragraphs}}
            <p>{{.}}</p>
        {{end}}
    <ul>
        {{range .Options}}
            <li> <a href="/{{.Arc}}">{{.Text}}</a></li>
        {{end}}
    </ul>
</body>
  `
