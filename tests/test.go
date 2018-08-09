package main

import (
    //"testing"
    "github.com/UjjwalAyyangar/go-jellyfish/pkg/jellyfish"
    "encoding/csv"
    "os"
    "bufio"
    "fmt"
    "io"
    "strconv"
)

type line []string

func ReadFile(filePath string) []line{
    f, _ := os.Open(filePath)

    r := csv.NewReader(bufio.NewReader(f))
    var data []line
    for {
        record, err := r.Read()
        if err == io.EOF {
            break
        }
        data = append(data,record)
    }

    return data
}

func TestLevenshtein_distance(){
    data := ReadFile("./test_data/levenshtein_distance.csv")
    l := len(data)
    for i:=0; i< l; i++ {

        inp1,inp2 := data[i][0],data[i][1]
        output, _ := strconv.Atoi(data[i][2])
        calc_output := jellyfish.Levenshtein_distance(inp1,inp2)

        if calc_output == output{
            fmt.Println("TRUE")
            fmt.Println(calc_output,output)
        }else {
            fmt.Println("False")
            fmt.Println(inp1,inp2,output,calc_output)
        }

    }

}


func TestJaro_distance(){
    data := ReadFile("./test_data/jaro_distance.csv")
    l := len(data)
    for i:=0; i< l; i++ {

        inp1,inp2 := data[i][0],data[i][1]
        output, _ := strconv.ParseFloat(data[i][2],64)
        calc_output := jellyfish.Jaro_distance(inp1,inp2)

        if calc_output == output{
            fmt.Println("TRUE")
            fmt.Println(calc_output,output)
        }else {
            fmt.Println("False")
            fmt.Println(inp1,inp2,output,calc_output)
        }

    }

}

func TestJaro_winkler(){
    data := ReadFile("./test_data/jaro_winkler.csv")
    l := len(data)
    for i:=0; i< l; i++ {

        inp1,inp2 := data[i][0],data[i][1]
        output, _ := strconv.ParseFloat(data[i][2],64)
        calc_output := jellyfish.Jaro_winkler(inp1,inp2)

        if calc_output == output{
            fmt.Println("TRUE")
            fmt.Println(calc_output,output)
        }else {
            fmt.Println("False")
            fmt.Println(inp1,inp2,output,calc_output)
        }
    }

}

func TestHamming_distance(){
    data := ReadFile("./test_data/hamming_distance.csv")
    l := len(data)
    for i:=0; i< l; i++ {

        inp1,inp2 := data[i][0],data[i][1]
        output, _ := strconv.ParseFloat(data[i][2],64)
        calc_output := jellyfish.Hamming_distance(inp1,inp2)

        if calc_output == int(output){
            fmt.Println("TRUE")
            fmt.Println(calc_output,output)
        }else {
            fmt.Println("False")
            fmt.Println(inp1,inp2,output,calc_output)
        }
    }

}

func TestMatch_rating_comparison(){
    data := ReadFile("./test_data/match_rating.csv")
    l := len(data)
    for i:=0; i< l; i++ {

        inp1,inp2 := data[i][0],data[i][1]
        output, _ := strconv.ParseBool(data[i][2])
        _,calc_output := jellyfish.Match_rating_comparison(inp1,inp2)

        if calc_output == output{
            fmt.Println("TRUE")
            fmt.Println(calc_output,output)
        }else {
            fmt.Println("FALSE")
            fmt.Println(inp1,inp2,output,calc_output)
        }
    }

}

func TestDamerau_levenshtein_distance(){
    data := ReadFile("./test_data/damerau_levenshtein_distance.csv")
    l := len(data)
    for i:=0; i< l; i++ {

        inp1,inp2 := data[i][0],data[i][1]
        output, _ := strconv.Atoi(data[i][2])
        calc_output := jellyfish.Damerau_levenshtein_distance(inp1,inp2)
        if calc_output == output{
            fmt.Println("TRUE")
            fmt.Println(calc_output,output)
        }else {
            fmt.Println("FALSE")
            fmt.Println(inp1,inp2,output,calc_output)
        }
    }

}




func main(){
    //ReadFile("./test_data/levenshtein_distance.csv")
    //TestLevenshtein_distance()
    TestDamerau_levenshtein_distance()
    //TestJaro_distance()
    //TestJaro_winkler()
    //TestHamming_distance()
    //TestMatch_rating_comparison()

}
