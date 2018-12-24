package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/UjjwalAyyangar/go-jellyfish/pkg/jellyfish"
	"io"
	"os"
	"strconv"
	"testing"
)

type line []string

func ReadFile(filePath string) []line {
	f, _ := os.Open(filePath)

	r := csv.NewReader(bufio.NewReader(f))
	var data []line
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		data = append(data, record)
	}

	return data
}

/*

TESTS FOR DISTANCES

*/

func TestLevenshtein_distance(t *testing.T) {
	data := ReadFile("./test_data/levenshtein_distance.csv")
	l := len(data)
	for i := 0; i < l; i++ {

		inp1, inp2 := data[i][0], data[i][1]
		output, _ := strconv.Atoi(data[i][2])
		calc_output := jellyfish.Levenshtein_distance(inp1, inp2)

		if calc_output != output {
			t.Errorf("Levenshtein distance failed")
		}

	}

}

func TestJaro_distance(t *testing.T) {
	data := ReadFile("./test_data/jaro_distance.csv")
	l := len(data)
	for i := 0; i < l; i++ {

		inp1, inp2 := data[i][0], data[i][1]
		output, _ := strconv.ParseFloat(data[i][2], 64)
		calc_output := jellyfish.Jaro_distance(inp1, inp2)

		if calc_output != output {
			t.Errorf("Jaro distance failed")
		}

	}

}

func TestJaro_winkler(t *testing.T) {
	data := ReadFile("./test_data/jaro_winkler.csv")
	l := len(data)
	for i := 0; i < l; i++ {

		inp1, inp2 := data[i][0], data[i][1]
		output, _ := strconv.ParseFloat(data[i][2], 64)
		calc_output := jellyfish.Jaro_winkler(inp1, inp2)

		if calc_output != output {
			t.Errorf("Jaro winkler failed")
		}
	}

}

func TestHamming_distance(t *testing.T) {
	data := ReadFile("./test_data/hamming_distance.csv")
	l := len(data)
	for i := 0; i < l; i++ {

		inp1, inp2 := data[i][0], data[i][1]
		output, _ := strconv.ParseFloat(data[i][2], 64)
		calc_output := jellyfish.Hamming_distance(inp1, inp2)

		if calc_output != int(output) {
			t.Errorf("Hamming distance failed")
		}
	}

}

func TestMatch_rating_comparison(t *testing.T) {
	data := ReadFile("./test_data/match_rating.csv")
	l := len(data)
	for i := 0; i < l; i++ {

		inp1, inp2 := data[i][0], data[i][1]
		output, _ := strconv.ParseBool(data[i][2])
		_, calc_output := jellyfish.Match_rating_comparison(inp1, inp2)

		if calc_output != output {
			t.Log("Input: ", inp1, ",", inp2, "Expected:", calc_output, " Got:", output)
			t.Errorf("Match rating failed")
		}
	}

}

func TestDamerau_levenshtein_distance(t *testing.T) {
	data := ReadFile("./test_data/damerau_levenshtein_distance.csv")
	l := len(data)
	for i := 0; i < l; i++ {

		inp1, inp2 := data[i][0], data[i][1]
		output, _ := strconv.Atoi(data[i][2])
		calc_output := jellyfish.Damerau_levenshtein_distance(inp1, inp2)
		if calc_output != output {
			t.Errorf("Damerau levenshtein distance failed")
		}
	}

}

/*

Testing phonetics

*/

func TestMetaphone(t *testing.T) {
	data := ReadFile("./test_data/metaphone.csv")
	l := len(data)
	for i := 0; i < l; i++ {
		inp, output := data[i][0], data[i][1]
		calc_output := jellyfish.Metaphone(inp)
		if calc_output != output {
			t.Errorf("Failed to computer the correct metaphone")
		}
	}

}

func TestSoundex(t *testing.T) {
	data := ReadFile("./test_data/soundex.csv")
	l := len(data)
	for i := 0; i < l; i++ {
		inp, output := data[i][0], data[i][1]
		fmt.Println(inp, output)
		calc_output := jellyfish.Soundex(inp)
		if calc_output != output {
			t.Log("Input: ", inp, "Expected:", calc_output, " Got:", output)
			t.Errorf("Failed to compute the correct soundex")
		}
	}

}

//func main(){
//ReadFile("./test_data/levenshtein_distance.csv")
//TestLevenshtein_distance()
//TestDamerau_levenshtein_distance()
//TestJaro_distance()
//TestJaro_winkler()
//TestHamming_distance()
//TestMatch_rating_comparison()

// Phonetics

//TestMetaphone()
//TestSoundex()

//}
