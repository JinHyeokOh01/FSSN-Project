package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type SuperHeroSquad struct {
	SquadName  string   `json:"squadName"`
	HomeTown   string   `json:"homeTown"`
	Formed     int      `json:"formed"`
	SecretBase string   `json:"secretBase"`
	Active     bool     `json:"active"`
	Members    []Member `json:"members"`
}

type Member struct {
	Name           string   `json:"name"`
	Age            int      `json:"age"`
	SecretIdentity string   `json:"secretIdentity"`
	Powers         []string `json:"powers"`
}

func main() {
	superHeroes := SuperHeroSquad{
		SquadName:  "Super hero squad",
		HomeTown:   "Metro City",
		Formed:     2016,
		SecretBase: "Super tower",
		Active:     true,
		Members: []Member{
			{
				Name:           "Molecule Man",
				Age:            29,
				SecretIdentity: "Dan Jukes",
				Powers:         []string{"Radiation resistance", "Turning tiny", "Radiation blast"},
			},
			{
				Name:           "Madame Uppercut",
				Age:            39,
				SecretIdentity: "Jane Wilson",
				Powers:         []string{"Million tonne punch", "Damage resistance", "Superhuman reflexes"},
			},
			{
				Name:           "Eternal Flame",
				Age:            1000000,
				SecretIdentity: "Unknown",
				Powers:         []string{"Immortality", "Heat Immunity", "Inferno", "Teleportation", "Interdimensional travel"},
			},
		},
	}

	fmt.Println(superHeroes.HomeTown)
	fmt.Println(superHeroes.Active)
	fmt.Println(superHeroes.Members[1].Powers[2])

	jsonData, err := json.MarshalIndent(superHeroes, "", "\t")
	if err != nil {
		fmt.Printf("JSON 변환 오류: %v\n", err)
		return
	}

	err = os.WriteFile("lec-06-prg-04-json-example.json", jsonData, 0644)
	if err != nil {
		fmt.Printf("파일 쓰기 오류: %v\n", err)
		return
	}
}