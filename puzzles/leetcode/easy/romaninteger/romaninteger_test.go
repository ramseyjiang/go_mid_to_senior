package romaninteger

import (
	"log"
	"testing"
)

func TestRomanToInt11(t *testing.T) {
	want := 3
	expected := romanToInt1("III")
	if want != expected {
		log.Println(expected)
		t.Error("Wrong")
	}
}

func TestRomanToInt12(t *testing.T) {
	want := 58
	expected := romanToInt1("LVIII")
	if want != expected {
		log.Println(expected)
		t.Error("Wrong")
	}
}

func TestRomanToInt13(t *testing.T) {
	want := 1994
	expected := romanToInt1("MCMXCIV")
	if want != expected {
		log.Println(expected)
		t.Error("Wrong")
	}
}

func TestRomanToInt21(t *testing.T) {
	want := 3
	expected := romanToInt2("III")
	if want != expected {
		log.Println(expected)
		t.Error("Wrong")
	}
}

func TestRomanToInt22(t *testing.T) {
	want := 58
	expected := romanToInt2("LVIII")
	if want != expected {
		log.Println(expected)
		t.Error("Wrong")
	}
}

func TestRomanToInt23(t *testing.T) {
	want := 1994
	expected := romanToInt2("MCMXCIV")
	if want != expected {
		log.Println(expected)
		t.Error("Wrong")
	}
}
