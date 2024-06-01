package utils

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func ParseResults(results SearchResult) {

	for _, result := range results.Results {

		name := strings.Split(result.HashName, "|")

		if len(name) == 1 {
			fmt.Printf("Estou a entrar no if dos das caixas %s", name)
			caixa := parseCase(name)
			fmt.Println(caixa)
			continue
		}

		if name[1] == "Sticker" {
			fmt.Printf("Estou a entrar no if dos das stickers %s", name)
			sticker := ParseSticker(name)
			fmt.Println(sticker)
		} else {
			fmt.Printf("Estou a entrar no if dos das armas %s", name)
			skin := parseSkin(name)
			fmt.Println(skin)
		}

	}
}

func ParseSticker(sticker []string) Sticker {
	if len(sticker) == 0 {
		log.Fatal("no input")
		return Sticker{}
	}

	println(sticker)

	pcondition := parseCondition(sticker[1])

	tmp := strings.Split(sticker[1], "(")
	name := strings.TrimSpace(tmp[0])
	condition := pcondition[1]
	tournamentName := "none"
	if len(sticker) > 2 {
		tournamentName = sticker[2]
	}

	//fmt.Println(name, tournamentName, condition)

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

func parseCase(name []string) Case {
	return Case{
		CaseName: name[0],
	}
}

func parseCondition(condition string) []string {
	re := regexp.MustCompile(`\(([^)]+)\)`)
	matches := re.FindStringSubmatch(condition)

	return matches
}
