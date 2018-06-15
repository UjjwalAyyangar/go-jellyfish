package jellyfish

import (
    "strings"
    "golang.org/x/text/unicode/norm"
)

func Soundex(s string) string {
    if len(s) == 0 {
        return ""
    }
    s = string(norm.NFKD.Bytes([]byte(s)))
    s = strings.ToUpper(s)

    replacements := map[string]string{
        "BFPV":"1",
        "CGJKQSXZ":"2",
        "DT":"3",
        "L":"4",
        "MN":"5",
        "R":"6",
    }

    var result []string
    result = append(result,string(s[0]))
    count := 1
    last := ""
    for  lset,sub := range replacements {
        if strings.Contains(lset, string(s[0])){
            last = sub
            break
        }
    }

    break_flag := false
    for i:= range s[1:]{
        letter := s[i]
        for lset,sub := range replacements{
            if strings.Contains(lset, string(letter)){
                if sub!=last{
                    result = append(result, sub)
                    count+=1
                }
                last = sub
                break_flag = true
                break
            }
        }
        if !break_flag{
            last = ""
        }
        
        if count ==4 {
            break
        }
    
    }
    
    for i:=0; i<4-count; i++ {
        result = append(result, "0")
    }
    return strings.Join(result,"")
}
