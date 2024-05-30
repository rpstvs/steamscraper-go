package utils

import (
	"fmt"
	"strings"

	"github.com/rpstvs/steamscraper-go/internals/steamapi"
)

func ParseResults(results steamapi.SearchResult) {

	//for _, result := range results.Results {

	//	name := strings.Split(" ", result.HashName)

	//}
}

func ParseSticker(sticker []string) Sticker {

	tmp := strings.Split(sticker[1], "(")
	name := tmp[0]
	tournamentName := sticker[2]
	condition := tmp[1]

	fmt.Println(name, tournamentName, condition)

	return Sticker{
		Name:       name,
		Tournament: tournamentName,
		Condition:  condition,
	}
}
