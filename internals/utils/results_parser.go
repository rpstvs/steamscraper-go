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

	pcondition := parseCondition(sticker[1])

	tmp := strings.Split(sticker[1], "(")
	name := strings.TrimSpace(tmp[0])
	condition := pcondition[1]
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

func parseSkin(skin []string) Skin {

	if len(skin) == 0 {
		log.Fatal("no input")
		return Skin{}
	}

	gunName := skin[0]

	pCondition := parseCondition(skin[1])
	tmp := strings.Split(skin[1], "(")
	skinName := strings.TrimSpace(tmp[0])
	condition := pCondition[1]

	return Skin{
		GunName:   gunName,
		SkinName:  skinName,
		Condition: condition,
	}
}

func parseCondition(condition string) []string {
	re := regexp.MustCompile(`\(([^)]+)\)`)
	matches := re.FindStringSubmatch(condition)

	return matches
}
