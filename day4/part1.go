package day4

import (
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"log"
	"strings"
)

type Passport struct {
	BirthYear      string `mapstructure:"byr"`
	IssueYear      string `mapstructure:"iyr"`
	ExpirationYear string `mapstructure:"eyr"`
	Height         string `mapstructure:"hgt"`
	HairColor      string `mapstructure:"hcl"`
	EyeColor       string `mapstructure:"ecl"`
	PassportId     string `mapstructure:"pid"`
	CountryId      string `mapstructure:"cid"`
}

func NewPassport(buffer []string) *Passport {
	mymap := make(map[string]interface{})
	for _, line := range buffer {
		for _, kv_text := range strings.Split(line, " ") {
			kv := strings.Split(kv_text, ":")
			if len(kv)!=2 {
				log.Fatalf("Input Format error: expect key:value found [%s]", kv)
			}
			mymap[kv[0]] = kv[1]
		}
	}

	var passport Passport
	err := mapstructure.Decode(mymap, &passport)
	if err!=nil {
		log.Fatal(err)
	}
	return &passport
}

func (p Passport) isValid() bool {
	return p.BirthYear!="" && p.ExpirationYear!="" &&
		p.EyeColor!="" && p.HairColor!="" && p.Height!="" &&
		p.IssueYear!="" && p.PassportId!=""
}

func readFile(filename string) []string  {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(bytes), "\n")
}

func ReadPassports(filename string) []Passport {
	var passports = make([]Passport, 0, 10)
	var buffer []string
	for _, line := range readFile(filename) {
		if line == "" {
			passports = append(passports, *NewPassport(buffer))
			buffer = nil
		} else {
			buffer = append(buffer, line)
		}
	}
	if buffer!=nil {
		passports = append(passports, *NewPassport(buffer))
	}
	return passports
}
