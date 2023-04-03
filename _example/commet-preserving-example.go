package main

import (
	"log"
	"os"

	"github.com/GuanceCloud/toml"
)

type ReleaseHistory []struct {
	Date     string   `toml:"date"`
	Version  string   `toml:"version"`
	Features []string `toml:"features"`
}

type Author struct {
	Name          string     `toml:"name"`
	Age           int        `toml:"age"`
	Gender        string     `toml:"gender"`
	MaritalStatus bool       `toml:"marital status"`
	Hobbies       [][]string `toml:"hobbies"`
}

type Country struct {
	Name       string   `toml:"name"`
	Code       string   `toml:"code"`
	Continent  string   `toml:"continent"`
	MainCities []string `toml:"main_cities"`
	State      []State  `toml:"state"`
}

type State struct {
	Name       string    `toml:"name"`
	Capital    string    `toml:"capital"`
	Area       int       `toml:"area"`
	Population int       `toml:"population"`
	Latitude   [2]string `toml:"latitude"`
	Longitude  [2]string `toml:"longitude"`
}

type tomlStruct struct {
	Version        string         `toml:"version"`
	ReleaseHistory ReleaseHistory `toml:"release_history"`
	Author         Author         `toml:"author"`
	Country        []Country      `toml:"country"`
}

func main() {
	var ts tomlStruct

	meta, err := toml.DecodeFile("./input.toml", &ts)
	if err != nil {
		log.Fatal(err)
	}

	ts.Author.Hobbies = append(ts.Author.Hobbies, []string{"volleyball"})
	ts.Version = "v1.2.3"
	ts.Author.Name = "GuanceCloud"
	ts.Country[1].State[0].Capital = "Boston"

	ts.Country[0].State = append(ts.Country[0].State, State{
		Name:       "Anhui",
		Capital:    "Hefei",
		Area:       140100,
		Population: 61270000,
		Latitude:   [2]string{"29°41′ N", "34°38′ N"},
		Longitude:  [2]string{"114°54′ E", "119°37′ E"},
	})

	enc := toml.NewEncoder(os.Stdout)
	if err := enc.EncodeWithComments(ts, meta); err != nil {
		log.Fatal(err)
	}
}
