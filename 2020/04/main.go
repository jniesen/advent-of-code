package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	BirthYear      string `json:"byr,omitempty"`
	IssueYear      string `json:"iyr,omitempty"`
	ExpirationYear string `json:"eyr,omitempty"`
	Height         string `json:"hgt,omitempty"`
	HairColor      string `json:"hcl,omitempty"`
	EyeColor       string `json:"ecl,omitempty"`
	PassportID     string `json:"pid,omitempty"`
	CountryID      string `json:"cid,omitempty"`
}

func NewPassport(fields map[string]string) (Passport, bool) {
	var p Passport

	jsonBytes, err := json.Marshal(fields)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(jsonBytes, &p); err != nil {
		panic(err)
	}

	return p, p.Valid()
}

type empty struct{}

func (p Passport) Valid() bool {
	v := map[bool]empty{}
	v[valid("byr", p.BirthYear)] = empty{}
	v[valid("iyr", p.IssueYear)] = empty{}
	v[valid("eyr", p.ExpirationYear)] = empty{}
	v[valid("hgt", p.Height)] = empty{}
	v[valid("hcl", p.HairColor)] = empty{}
	v[valid("ecl", p.EyeColor)] = empty{}
	v[valid("pid", p.PassportID)] = empty{}

	if _, invalidField := v[false]; invalidField {
		return false
	}

	return true
}

func validRange(min, max, val int) bool {
	return (val >= min && val <= max)
}

func validYear(min, max int, yr string) bool {
	if len(yr) != 4 {
		return false
	}

	yrInt, err := strconv.ParseInt(yr, 10, 0)
	if err != nil {
		panic(err)
	}

	return validRange(min, max, int(yrInt))
}

func validHeight(min, max int, hgt string) bool {
	hgtInt, err := strconv.ParseInt(hgt, 10, 0)
	if err != nil {
		panic(err)
	}

	return validRange(min, max, int(hgtInt))
}

func validRegex(ptrn, val string) bool {
	match, err := regexp.Match(ptrn, []byte(val))
	if err != nil {
		panic(err)
	}
	return match
}

func valid(field, val string) bool {
	v := false

	switch field {
	case "byr":
		v = validYear(1920, 2002, val)
	case "iyr":
		v = validYear(2010, 2020, val)
	case "eyr":
		v = validYear(2020, 2030, val)
	case "hgt":
		if strings.HasSuffix(val, "cm") {
			v = validHeight(150, 193, strings.TrimSuffix(val, "cm"))
		}

		if strings.HasSuffix(val, "in") {
			v = validHeight(59, 76, strings.TrimSuffix(val, "in"))
		}
	case "hcl":
		if len(val) != 7 {
			break
		} else {
			v = validRegex(`#[0-9a-f]*`, val)
		}
	case "ecl":
		validColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		for _, c := range validColors {
			if val == c {
				v = true
				break
			}
		}
	case "pid":
		if len(val) != 9 {
			break
		} else {
			v = validRegex(`[0-9]*`, val)
		}
	}

	return v
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	passportData := [][]string{}
	currentPassport := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			passportData = append(passportData, currentPassport)
			currentPassport = []string{}
			continue
		}

		fields := strings.Split(scanner.Text(), " ")
		currentPassport = append(currentPassport, fields...)
	}

	passportData = append(passportData, currentPassport)

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	validPassports := []Passport{}
	invalidPassports := []Passport{}
	for _, d := range passportData {
		dataMap := map[string]string{}

		for _, field := range d {
			kv := strings.Split(field, ":")
			dataMap[kv[0]] = kv[1]
		}

		passport, valid := NewPassport(dataMap)
		if valid {
			validPassports = append(validPassports, passport)
		} else {
			invalidPassports = append(invalidPassports, passport)
		}
	}

	fmt.Printf("Invalid Passports: %v, Valid Passports: %v\n", len(invalidPassports), len(validPassports))
}
