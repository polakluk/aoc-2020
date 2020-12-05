package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	passport string
	expected bool
}

func TestPreprocess(t *testing.T) {
	cases := []testCase{
		{"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm", true},
		{"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884 hcl:#cfa07d byr:1929", false},
		{"hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm", true},
		{"hcl:#cfa07d eyr:2025 pid:166559648 iyr:2011 ecl:brn hgt:59in", false},
	}
	for _, c := range cases {
		result, _ := preprocessPassport(c.passport)
		if c.expected != result {
			t.Errorf("Test case %s ==> Expected %t, got %t", c.passport, c.expected, result)
		}
		fmt.Println(c)
	}
}
func TestCheck(t *testing.T) {
	cases := []testCase{
		{"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm", true},
		{"hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm", true},
	}
	for _, c := range cases {
		_, mapFields := preprocessPassport(c.passport)

		result := checkPassport(mapFields)
		if c.expected != result {
			t.Errorf("Test case %s ==> Expected %t, got %t", c.passport, c.expected, result)
		}
		fmt.Println(c)
	}
}
