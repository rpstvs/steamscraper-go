package utils

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/rpstvs/steamscraper-go/internals/steamapi"
)

func ParseResults(results steamapi.SearchResult) {

	//for _, result := range results.Results {

	//	name := strings.Split(" ", result.HashName)

	//}
}

func ParseSticker(sticker []string) Sticker {
	if len(sticker) == 0 {
		log.Fatal("no input")
		return Sticker{}
	}

	re := regexp.MustCompile(`\(([^)]+)\)`)
	matches := re.FindStringSubmatch(sticker[1])

	tmp := strings.Split(sticker[1], "(")
	name := strings.TrimSpace(tmp[0])
	condition := matches[1]
	tournamentName := "none"
	if len(sticker) > 2 {
		tournamentName = sticker[2]
	}

	fmt.Println(name, tournamentName, condition)

	return Sticker{
		Name:       name,
		Tournament: tournamentName,
		Condition:  condition,
	}
}
