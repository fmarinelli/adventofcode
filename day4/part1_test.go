package day4

import "testing"

func TestHello(t *testing.T) {
	passports := ReadPassports("input_test.txt")
	const want_passports = 4
	if len(passports)!= want_passports {
		t.Errorf("Hello() = %q, want %q", len(passports), want_passports)
	}
	valid := 0
	for _, v := range passports {
		if v.isValid() {
			valid++
		}
	}

	const want_valid = 2
	if valid!= want_valid {
		t.Errorf("Hello() = %q, want %q", valid, want_valid)
	}
}