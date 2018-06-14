package jellyfish

import (
    "strings"
    "GoJellyfish/pkg/util"
)

func Levenshtein_distance(s1 , s2 string) int {
    if strings.Compare(s1,s2) == 0 {
        return 0
    } else {
        rows := len(s1) + 1
        cols := len(s2) + 1

        if s1 == "" {
            return cols - 1
        }

        if s2 == "" {
            return rows - 1
        }

        cur := util.Generate_arr(0,cols)
        for r:=1; r<rows; r++ {
            prev := make([]int, cols)
            util.Memset(prev, 0)
            prev[0] = r

            for c:=1; c<cols; c++ {
                deletion := prev[c] + 1
                insertion := cur[c-1] + 1
                edit := prev[c-1]
                if s1[r-1] != s2[c-1]{
                    edit +=1
                }
                cur[c] = util.Min(edit, deletion, insertion) 
            }

        }

        return cur[len(cur)-1]

    }
}

func Damerau_levenshtein_distance(s1 , s2 string) int{
    len1 := len(s1)
    len2 := len(s2)
    infinite := len1 + len2
    da := make(map[byte]int)
    score := make([][]int,len1+2)
    //Distance matrix
    for i := range score {
        score[i] = make([]int, len2+2)
        util.Memset(score[i], 0)
    }

    score[0][0] = infinite

    for i:=0; i<len1+1; i++{
        score[i+1][0] = infinite
        score[i+1][1] = i
    }

    for i:=0; i<len2+1; i++{
        score[0][i+1] = infinite
        score[1][i+1] = i
    }

    for i:=1; i<len1+1; i++{
        db:=0
        for j:=1; j<len2+1; j++{
            i1 := da[s2[j-1]]
            j1 := db
            cost := 1
            if s1[i-1] == s2[j-1]{
                cost = 0
                db = j
            }

            score[i+1][j+1] = util.Min(score[i][j] + cost, score[i+1][j] +1, score[i][j+1] +1, score[i1][j1] + (i-i1-1) + 1 + (j-j1-1))
        }
        da[s1[i-1]] = i
    }

    return score[len1+1][len2+1]
}



