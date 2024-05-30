package utils

import (
	"strings"

	"github.com/rpstvs/steamscraper-go/internals/steamapi"
)

func ParseResults(results steamapi.SearchResult) {

	for _, result := range results.Results {

		name := strings.Split(" ", result.HashName)

	}
}
