package main

import "testing"

var input string = "omama"
var vowels int = 2
var consonants int = 2

func TestCountChar(t *testing.T) {
	countC, countV := CountCharModel(input)
	t.Logf("Word : %s\n", input)
	t.Logf("Huruf Mati : %d\n", countC)
	t.Logf("Huruf Hidup : %d", countV)

	if countC != consonants {
		t.Errorf("Error! expected (%d) got (%d)", consonants, countC)
	}
	if countV != vowels {
		t.Errorf("Error! expected (%d) got (%d)", vowels, countV)
	}

}

func BenchmarkCountChar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountCharModel(input)
	}
}
