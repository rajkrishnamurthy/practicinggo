package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	example2()
}

func example2() {
	const filepattern = `\([a-z,\s]+\)`
	var ruleregex = regexp.MustCompile(filepattern)

	matchStrings := []string{
		`INSERT INTO schema.table (id,name,age) values (1,2,3)`,
		`INSERT INTO schema.table (id,name,age) values (?,?,?)`,
		`SELECT * FROM schema.table`,
		`SELECT id,name,age FROM schema.table WHERE 1=1`,
	}
	for _, matchStr := range matchStrings {
		source := strings.ToLower(matchStr)
		// submatches := ruleregex.FindAllStringSubmatch(source, -1)
		submatches := ruleregex.FindStringSubmatch(source)
		isMatch := ruleregex.MatchString(source)
		fmt.Printf("Match?%v\t%v\t%d\n", isMatch, submatches, len(submatches))

	}

}

func example1() { // const filepattern = `^[a-z]+\[[0-9]+\]$`
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
