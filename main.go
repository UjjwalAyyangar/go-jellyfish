package main
import (
    "GoJellyfish/pkg/jellyfish"
    "fmt"
)

func main(){
    fmt.Println(jellyfish.Levenshtein_distance("ABC","ABD"))
    fmt.Println(jellyfish.Damerau_levenshtein_distance("ABC","ABD"))
    fmt.Println(jellyfish.Jaro_distance("ABC","ABD"))
    
}
