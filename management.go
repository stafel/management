package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strings"
)

// world save file
const WORLD_PATH string = "world.json"

// management sim single character
type MgCharacter struct {
	Name string  `json:"name"`
	Age  float64 `json:"age"`
	Mood string  `json:"mood"`
}

// management sim world contains all data
type MgWorld struct {
	Characters []*MgCharacter `json:"characters"`
}

// saves management sim world in file as json
func saveWorld(world *MgWorld, outpath string) {
	jsonData, err := json.Marshal(world)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(outpath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.Write(jsonData)

	log.Printf("Saved world in %s", outpath)
}

// loads management sim world from json file
func loadWorld(inpath string) *MgWorld {

	log.Printf("Loading world from %s", inpath)

	data, err := os.ReadFile(inpath)
	if err != nil {
		log.Fatal(err)
	}

	w := &MgWorld{}

	err = json.Unmarshal(data, w)
	if err != nil {
		log.Fatal(err)
	}

	return w
}

// initialize new world
func newWorld() *MgWorld {
	log.Print("Creating new world")
	return &MgWorld{} // just init
}

// loads existing json or creates new world if not existing
func loadWorldOrNew(inpath string) *MgWorld {
	if _, err := os.Stat(inpath); err != nil {
		// file does not exist
		return newWorld()
	}
	return loadWorld(inpath)
}

// returns input line without CR
func readLine() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.Replace(text, "\n", "", 1), nil // remove the CR from text
}

func saveTestworld() {
	c := &MgCharacter{
		Name: "Lex",
		Age:  22,
		Mood: "Hungry",
	}

	w := &MgWorld{}
	w.Characters = append(w.Characters, c)

	saveWorld(w, WORLD_PATH)
}

func main() {
	log.Print("Welcome to management")

	w := loadWorldOrNew(WORLD_PATH)

	log.Printf("Our first character is called %s and feels %s", w.Characters[0].Name, w.Characters[0].Mood)

	log.Print("New mood: ")
	w.Characters[0].Mood, _ = readLine()

	log.Printf("Our first character is called %s and feels %s", w.Characters[0].Name, w.Characters[0].Mood)

	saveWorld(w, WORLD_PATH)
}
