package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// regexp matching key-value pairs
var keyValRegExp, _ = regexp.Compile("([a-z]{3}):([a-z0-9#]+)")

// regexp matching hair color
var hairColorRegExp, _ = regexp.Compile("^#[a-f0-9]{6}$")

// regexp matching eye color
var eyeColorRegExp, _ = regexp.Compile("^(amb|blu|brn|gry|grn|hzl|oth)$")

// regexp matching passport ID
var passIDRegExp, _ = regexp.Compile("^[0-9]{9}$")

// required fields on passport
var reqFields = [7]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

// preprocessPassport verifies that passport contains all required fields and returns it as a map
func preprocessPassport(passport string) (bool, map[string]string) {
	keyMap := make(map[string]string)

	// find all key-val groups
	matchedFields := keyValRegExp.FindAllStringSubmatch(passport, -1)

	// make sure there's enough fields (less than # of required fields means it is always invalid passport)
	if len(matchedFields) >= len(reqFields) {
		// transform group of fields to map
		for _, group := range matchedFields {
			keyMap[group[1]] = group[2]
		}

		// check that each required field is present
		for _, field := range reqFields {
			_, ok := keyMap[field]

			if !ok {
				return false, keyMap
			}
		}
	} else {
		// not enough fields on the paspport
		return false, keyMap
	}

	return true, keyMap
}

func validateYear(year string, min int, max int) bool {
	realYear, ok := strconv.Atoi(year)

	return ok == nil && realYear >= min && realYear <= max
}

// checkPassport verifies that all fields contain valid values
func checkPassport(mapFields map[string]string) bool {
	byr, _ := mapFields["byr"]
	iyr, _ := mapFields["iyr"]
	eyr, _ := mapFields["eyr"]
	hgt, _ := mapFields["hgt"]
	hcl, _ := mapFields["hcl"]
	ecl, _ := mapFields["ecl"]
	pid, _ := mapFields["pid"]

	// validate byr
	if !validateYear(byr, 1920, 2002) {
		return false
	}

	// validate iyr
	if !validateYear(iyr, 2010, 2020) {
		return false
	}

	// validate eyr
	if !validateYear(eyr, 2020, 2030) {
		return false
	}

	// validate hgt
	hgtValue, _ := strconv.Atoi(hgt[0 : len(hgt)-2])
	if hgt[len(hgt)-2] == 'c' {
		// height in cm
		if hgtValue < 150 || hgtValue > 193 {
			return false
		}
	} else {
		// height in inches
		if hgtValue < 59 || hgtValue > 76 {
			return false
		}
	}

	// validate hcl
	if !(hairColorRegExp.MatchString(hcl)) {
		return false
	}

	// validate ecl
	if !(eyeColorRegExp.MatchString(ecl)) {
		return false
	}

	// validate pid
	if !(passIDRegExp.MatchString(pid)) {
		return false
	}
	return true
}

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)

	// current passport
	passport := ""
	// number of passports with required fields
	validFieldsPassports := 0

	// number of valid fields
	validPassports := 0

	for scanner.Scan() {
		// read the line
		line := scanner.Text()
		if line == "" {
			// preprocess the passport
			isValid, mapFields := preprocessPassport(passport)
			if isValid {
				validFieldsPassports++

				// check whether the required fields contain valid values
				if checkPassport(mapFields) {
					validPassports++
				}
			}

			// prepare search for a new passport
			passport = ""
		} else {
			// keep reading this passport
			passport += " " + line
		}
	}

	// check the last passport
	isValid, mapFields := preprocessPassport(passport)

	if isValid {
		validFieldsPassports++

		// check whether the required fields contain valid values
		if checkPassport(mapFields) {
			validPassports++
		}
	}

	fmt.Printf("Passport with required fields - %d\n", validFieldsPassports)
	fmt.Printf("Passport with valid fields - %d\n", validPassports)
}
