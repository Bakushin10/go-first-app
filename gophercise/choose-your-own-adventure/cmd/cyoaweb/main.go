package chooseyourownadventure

import (
	"encoding/json"
	"flag"
	"fmt"
	// "github.com/gophercises/cyoa"
	"choose-your-own-adventure/story"
	"os"
)

func main() {
	filename := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("using the story in %s. \n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	fmt.Print(f)

	d := json.NewDecoder(f)
	var story story.Story
	if err := d.Decode(&story); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)
}
