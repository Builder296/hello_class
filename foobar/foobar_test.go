package foobar_test

import (
	"testing"

	"github.com/Builder296/hello_class/foobar"
)

func TestSayAnyGivenOneWantOne(t *testing.T) {
	given := "1"
	want := "1"

	result := foobar.SayAny(given)

	if result != want {
		t.Errorf("Say(%s) = %s; want %s", given, result, want)
	}
}

func TestSayAnyGivenThreeWantFoo(t *testing.T) {
	given := "3"
	want := "foo"

	result := foobar.SayAny(given)

	if result != want {
		t.Errorf("Say(%s) = %s; want %s", given, result, want)
	}
}

func TestGivenOneWantOne(t *testing.T) {
	given := 1
	want := "1"

	result := foobar.Say(given)

	if result != want {
		t.Errorf("Say(%d) = %s; want %s", given, result, want)
	}
}

func TestGivenTwoWantTwo(t *testing.T) {
	given := 2
	want := "2"

	result := foobar.Say(given)

	if result != want {
		t.Errorf("Say(%d) = %s; want %s", given, result, want)
	}
}

func TestGivenThreeWantfoo(t *testing.T) {
	given := 3
	want := "foo"

	result := foobar.Say(given)

	if result != want {
		t.Errorf("Say(%d) = %s; want %s", given, result, want)
	}
}

func TestGivenFourWantFour(t *testing.T) {
	given := 4
	want := "4"

	result := foobar.Say(given)

	if result != want {
		t.Errorf("Say(%d) = %s; want %s", given, result, want)
	}
}

func TestGivenFiveWantBar(t *testing.T) {
	given := 5
	want := "bar"

	result := foobar.Say(given)

	if result != want {
		t.Errorf("Say(%d) = %s; want %s", given, result, want)
	}
}

func TestGivenSixWantFoo(t *testing.T) {
	given := 6
	want := "foo"

	result := foobar.Say(given)

	if result != want {
		t.Errorf("Say(%d) = %s; want %s", given, result, want)
	}
}

func TestGivenSevenWantSeven(t *testing.T) {
	given := 6
	want := "foo"

	result := foobar.Say(given)

	if result != want {
		t.Errorf("Say(%d) = %s; want %s", given, result, want)
	}
}
