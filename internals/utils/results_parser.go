package utils

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func ParseResults(results SearchResult) {

	for _, result := range results.Results {
		//fmt.Println(result.HashName)
		name := strings.Split(result.HashName, "|")

		if len(name) == 1 {
			caixa := parseCase(name)
			fmt.Println(caixa)
			continue
		}

		if name[0] == "Sticker" {
			//fmt.Printf("Estou a entrar no if dos das stickers %s", name)
			sticker := ParseSticker(name)
			fmt.Println(sticker)
		} else if strings.Contains(result.AssetDescription.Type, "Agente") {
			//fmt.Printf("Estou a entrar no if dos das armas %s", name)
			agent := parseAgent(name)
			fmt.Println(agent)

		} else {
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

	condition := "none"
	tournamentName := "none"
	var pcondition []string

	tmp := strings.Split(sticker[1], "(")

	if len(tmp) == 2 {
		pcondition = parseCondition(sticker[1])
	}

	name := strings.TrimSpace(tmp[0])

	if len(sticker) > 2 {
		tournamentName = sticker[2]
	}

	if pcondition != nil {
		condition = pcondition[1]
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

	condition := "none"

	gunName := skin[0]

	tmp := strings.Split(skin[1], "(")
	skinName := strings.TrimSpace(tmp[0])

	pCondition := parseCondition(skin[1])

	if pCondition != nil {
		condition = pCondition[1]
	}

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

	if len(matches) == 0 {
		return nil
	}

	return matches
}

func parseAgent(agent []string) Agent {

	return Agent{
		Name:  agent[0],
		Group: agent[1],
	}
}

func priceConverter(priceText string) float64 {

	priceNum, err := strconv.ParseFloat(priceText, 32)
	if err != nil {
		return 0.00
	}
	price := priceNum / 100

	return price
}
