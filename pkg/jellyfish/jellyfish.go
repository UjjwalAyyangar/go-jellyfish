package jellyfish

import (
    "strings"
    "GoJellyfish/pkg/util"
)

func Levenshtein_distance(s1 string, s2 string) int {
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

