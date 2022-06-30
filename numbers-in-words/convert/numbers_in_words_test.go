package convert

import "testing"

func TestZero(t *testing.T) {
	want := "zero"
	got := ConvertNumberToWords(0)
	if got != want {
		t.Errorf("ConvertNumberToWords() = %v, want %v", got, want)
	}
}

func TestOneToTwenty(t *testing.T) {
	var entry = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten",
		"eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen", "twenty"}
	for i, numberWord := range entry {
		want := numberWord
		got := ConvertNumberToWords(i + 1)
		if got != want {
			t.Errorf("ConvertNumberToWords() = %v, want %v", got, want)
		}
	}
}

func TestTwentyOne(t *testing.T) {
	want := "twentyone"
	got := ConvertNumberToWords(21)
	if got != want {
		t.Errorf("ConvertNumberToWords() = %v, want %v", got, want)
	}
}

func TestUnkonwn(t *testing.T) {
	want := "error"
	got := ConvertNumberToWords(-1)
	if got != want {
		t.Errorf("ConvertNumberToWords() = %v, want %v", got, want)
	}
}
