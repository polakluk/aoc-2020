package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Rule holds information about a rule
type Rule struct {
	lower1 int
	upper1 int
	lower2 int
	upper2 int
}

//validRule checked whether the value matches the rule
func validRule(val int, rule Rule) bool {
	return (rule.lower1 <= val && rule.upper1 >= val) || (rule.lower2 <= val && rule.upper2 >= val)
}

//allTicketValidField checkes whether a field in all tickets matches the given rule
func allTicketValidField(tickets [][]int, fieldIdx int, rule Rule) bool {
	for _, ticket := range tickets {
		if !validRule(ticket[fieldIdx], rule) {
			// there is a field which breaks the rule
			return false
		}
	}
	return true
}

func determineRules(tickets [][]int, rules map[string]Rule) map[string]int {
	orderedRules := map[string]int{}

	compatibleRules := map[int]map[string]bool{}
	for pos := 0; pos < len(rules); pos++ {
		compatibleRules[pos] = map[string]bool{}
	}

	for ruleName, rule := range rules {
		for fieldIdx := 0; fieldIdx < len(rules); fieldIdx++ {
			if allTicketValidField(tickets, fieldIdx, rule) {
				compatibleRules[fieldIdx][ruleName] = true
			}
		}
	}

	for len(orderedRules) < len(rules) {
		for idx, compatibility := range compatibleRules {
			if len(compatibility) == 1 {
				// found field with exactly one matching rule
				for rule := range compatibility {
					// mark this rule
					orderedRules[rule] = idx

					// remove this field from compatibility map
					for pos := 0; pos < len(rules); pos++ {
						delete(compatibleRules[pos], rule)
					}
				}
				break
			}
		}
	}
	return orderedRules
}

// checkValidTicket iterates over all rules and returns sum of the error fields (0 means no error)
func checkValidTicket(inp []int, rules map[string]Rule) (bool, int) {
	errors := 0

	invalid := false
	for _, field := range inp {
		foundRule := false
		for _, rule := range rules {
			foundRule = foundRule || validRule(field, rule)
		}

		if !foundRule {
			// no matching rule was found so this is an error
			errors += field
			invalid = true
		}
	}

	return invalid, errors
}

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)

	rules := map[string]Rule{}

	var lower1, upper1, lower2, upper2 int
	keepReading := true
	// read rules
	for keepReading {
		scanner.Scan()
		line := scanner.Text()
		if line == "" {
			keepReading = false
			continue
		}
		mainParts := strings.Split(line, ": ")
		fmt.Sscanf(mainParts[1], "%d-%d or %d-%d", &lower1, &upper1, &lower2, &upper2)
		rules[mainParts[0]] = Rule{lower1, upper1, lower2, upper2}
	}

	// read your ticket
	scanner.Scan()
	scanner.Scan()
	line := scanner.Text()
	yourTicketParts := strings.Split(line, ",")
	yourTicket := []int{}
	for _, part := range yourTicketParts {
		val, _ := strconv.Atoi(part)
		yourTicket = append(yourTicket, val)
	}

	// read nearby tickets
	validNearByTickets := [][]int{}
	scanner.Scan()
	scanner.Scan()

	errorRate := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		actTicket := []int{}
		for _, part := range parts {
			val, _ := strconv.Atoi(part)
			actTicket = append(actTicket, val)
		}
		invalid, actErrorRate := checkValidTicket(actTicket, rules)
		if invalid {
			// this is invalid ticket - discard it
			errorRate += actErrorRate
		} else {
			// this seems to be a valid ticket
			validNearByTickets = append(validNearByTickets, actTicket)
		}
	}

	// identify rule names
	orderedRules := determineRules(validNearByTickets, rules)

	// multiply "departure" fields
	mult := 1
	for rule, idx := range orderedRules {
		if len(rule) > 9 && rule[:9] == "departure" {
			mult *= yourTicket[idx]
		}
	}
	fmt.Printf("Ticket scanning error rate - %d\n", errorRate)
	fmt.Printf("Your ticket value - %d\n", mult)
}
