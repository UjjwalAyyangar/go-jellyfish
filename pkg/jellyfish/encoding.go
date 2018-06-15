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

func Nysiis(s string) string {
    if len(s) == 0 {
        return ""
    }
    
    var key []string
    s = strings.ToUpper(s)
    
    if strings.HasPrefix(s, "MAC") {
        s = "MCC" + s[3:]
    } else if strings.HasPrefix(s, "KN") {
        s = s[1:]
    } else if strings.HasPrefix(s, "K") {
        s = "C" + s[1:]
    } else if strings.HasPrefix(s, "PH") || strings.HasPrefix(s, "PF") {
        s = "FF" + s[2:]
    } else if strings.HasPrefix(s, "SCH") {
        s = "SSS" + s[3:]
    }
    
    if strings.HasSuffix(s,"IE") || strings.HasSuffix(s, "EE"){
        s = s[:len(s)-2] + "Y"
    } else if strings.HasSuffix(s,"DT") || strings.HasSuffix(s,"RT") || 
    strings.HasSuffix(s,"RD") || strings.HasSuffix(s,"NT") ||
    strings.HasSuffix(s, "ND") {
        s = s[:len(s)-2] + "D"
    }
    
    key = append(key,string(s[0]))

    i:=1
    len_s := len(s)
    for i<len_s {
        ch:= string(s[i])
        if ch == "E" && i+1 < len_s && string(s[i+1]) == "V" {
            ch = "AF"
            i+=1
        } else if strings.Contains("AEIOU",ch) {
            ch = "A"
        } else if ch == "Q" {
            ch = "G"
        } else if ch =="Z" {
            ch = "S"
        } else if ch == "M" {
            ch = "N"
        } else if ch == "K" {
            if i+1 <len(s) && string(s[i+1]) == "N" {
                ch = "N"
            } else {
                ch = "C"
            }
        } else if ch == "S" && i+3<len_s && s[i+1:i+3] == "CH" {
            ch ="SS"
            i+=2
        } else if ch =="P" && i+1<len(s) && string(s[i+1]) == "H"{
            ch = "F"
            i+=1
        } else if ch=="H" && (!strings.Contains("AEIOU",string(s[i-1])) || 
        (i+1 < len(s) &&  !strings.Contains("AEIOU",string(s[i+1])))){
            if strings.Contains("AEIOU",string(s[i-1])){
                ch = "A"
            } else {
                ch = string(s[i-1])
            }
        } else if ch=="W" && strings.Contains("AEIOU",string(s[i-1])) {
            ch = string(s[i-1])
        }

        temp := key[len(key)-1]
        temp = string(temp[len(temp)-1])
        if string(ch[len(ch)-1]) != temp{
            key = append(key,ch)
        }

        i+=1

    }

    key_s := strings.Join(key,"")

    if strings.HasSuffix(key_s,"S") && key_s !="S"{
        key_s = key_s[:len(key_s)-1]
    }

    if strings.HasSuffix(key_s, "AY") {
        key_s = key_s[:len(key_s)-2] + "Y"
    }

    if strings.HasSuffix(key_s, "A") && key_s !="A"{
        key_s = key_s[:len(key_s)-1]
    
    }

    return key_s

}
