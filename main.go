package main
import (
    "GoJellyfish/pkg/jellyfish"
    "fmt"
)

func main(){
    fmt.Println(jellyfish.Levenshtein_distance("ABC","ABD"))
}
