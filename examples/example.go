package main
import (
    "github.com/UjjwalAyyangar/go-jellyfish/pkg/jellyfish"
    "fmt"
)

func main(){
    fmt.Println(jellyfish.Levenshtein_distance("Hello","World"))
    fmt.Println(jellyfish.Damerau_levenshtein_distance("ABC","ABD"))
    fmt.Println(jellyfish.Jaro_distance("Hello","World"))
    fmt.Println(jellyfish.Jaro_winkler("ABC","ABD"))
    fmt.Println(jellyfish.Hamming_distance("ABC","ABD"))
    fmt.Println(jellyfish.Match_rating_comparison("hello","yellow"))
    fmt.Println(jellyfish.Match_rating_codex("Jellyfish"))
    fmt.Println(jellyfish.Soundex("Jellyfish"))
    fmt.Println(jellyfish.Nysiis("JellyFish"))
    fmt.Println(jellyfish.Metaphone("Jellyfish"))
}
