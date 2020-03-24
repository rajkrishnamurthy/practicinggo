package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// const filepattern = `^[a-z]+\[[0-9]+\]$`
	const filepattern = `^rule([0-9]+)_`
	// const filepattern = `^([0-9]*)-([0-9]*)$`
	var ruleregex = regexp.MustCompile(filepattern)
	matchStrings := []string{
		"Rule202_Dummy",
		"Rule209_EnsureImageSprawl6.1",
		"Rule2_Accounting-ChartOfAccounts",
		"Rule10028_TestingSomething",
		"Rule6ITInvestmentsBudget",
		"200-10092",
		"-10002",
		"100-",
	}

	for _, matchStr := range matchStrings {
		source := fmt.Sprintf("%s", strings.ToLower(matchStr))
		submatches := ruleregex.FindAllStringSubmatch(source, -1)
		isMatch := ruleregex.MatchString(source)
		fmt.Printf("Match?%v\t%v\t%d\n", isMatch, submatches, len(submatches))

	}
}
