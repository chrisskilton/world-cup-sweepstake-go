package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	peopleJSONPathPtr := flag.String("people", "", "Path to JSON array of names")
	teamsJSONPathPtr := flag.String("teams", "", "Path to JSON array of teams")

	flag.Parse()

	if len(*peopleJSONPathPtr) == 0 {
		panic("must supply --people option. run sweeps --help if you are stuck")
	}

	if len(*teamsJSONPathPtr) == 0 {
		panic("must supply --teams option. run sweeps --help if you are stuck")
	}

	peopleJSON, err := ioutil.ReadFile(*peopleJSONPathPtr)
	checkError(err)
	teamsJSON, err := ioutil.ReadFile(*teamsJSONPathPtr)
	checkError(err)

	var people []string
	var teams []string
	var slicedPeople []string

	json.Unmarshal(peopleJSON, &people)
	json.Unmarshal(teamsJSON, &teams)

	picks := make(map[string][]string)

	for i := 0; i < (len(teams) - len(teams)%len(people)); i++ {
		if len(slicedPeople) == 0 {
			slicedPeople = make([]string, len(people))
			copy(slicedPeople, people)
		}

		team := teams[i]
		personIndex := rand.Intn(len(slicedPeople))
		person := slicedPeople[personIndex]

		fmt.Print("assigning team " + team + " to person " + person + "\n")

		picks[person] = append(picks[person], team)
		slicedPeople = append(slicedPeople[:personIndex], slicedPeople[personIndex+1:]...)
	}

	outputJson, err := json.MarshalIndent(picks, "", " ")
	checkError(err)
	fmt.Print(string(outputJson))
}
