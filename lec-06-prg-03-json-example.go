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
	jsonFile, err := os.ReadFile("lec-06-prg-03-json-example.json")
	if err != nil {
		fmt.Printf("파일 읽기 오류: %v\n", err)
		return
	}

	var superHeroes SuperHeroSquad
	if err := json.Unmarshal(jsonFile, &superHeroes); err != nil {
		fmt.Printf("JSON 파싱 오류: %v\n", err)
		return
	}

	fmt.Println(superHeroes.HomeTown)
	fmt.Println(superHeroes.Active)
	fmt.Println(superHeroes.Members[1].Powers[2])
}