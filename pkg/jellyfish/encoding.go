package jellyfish

import (
    "strings"
    "golang.org/x/text/unicode/norm"
    "github.com/UjjwalAyyangar/go-jellyfish/pkg/util"
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
    
    if util.HasSuffix(s,"IE", "EE"){
        s = s[:len(s)-2] + "Y"
    } else if util.HasSuffix(s,"DT","RT","RD","NT","ND") {
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

    if util.HasSuffix(key_s,"S") && key_s !="S"{
        key_s = key_s[:len(key_s)-1]
    }

    if util.HasSuffix(key_s, "AY") {
        key_s = key_s[:len(key_s)-2] + "Y"
    }

    if util.HasSuffix(key_s, "A") && key_s !="A"{
        key_s = key_s[:len(key_s)-1]
    
    }

    return key_s

}

func Metaphone(s string) string {
    var result []string 
    s = string(norm.NFKD.Bytes([]byte(s)))
    s = strings.ToLower(s)

    if util.HasPrefix(s, "kn","gn","pn","ac","wr","ae"){
        s = s[1:]
    }
    i:=0
    
    for i< len(s){
        c := string(s[i])
        var next string
        if i < len(s)-1 {
            next = string(s[i+1])
        } else {
            next = "*****"
        }

        var nextnext string
        
        if i < len(s)-2{
            nextnext = string(s[i+2])
        } else {
            nextnext = "*****"
        }

        if c==next && c !="c" {
            i+=1
            continue
        }

        if strings.Contains("aeiou",c){
            if i==0 || string(s[i-1]) == " "{
                result = append(result, c)
            } 
        } else if c == "b" {
            if !(i !=0 && string(s[i-1]) == "m")  || len(next)!=0{
                result = append(result, "b")
            }
        } else if c == "c" {
            if next =="i" && nextnext == "a" || next == "h" {
                result = append(result, "x")
                i+=1
            } else if strings.Contains("iey",next) {
                result = append(result,"s")
            } else {
                result = append(result, "k")
            }
        } else if c=="d" {
            if next == "g" && strings.Contains("iey",nextnext){
                result = append(result, "j")
                i += 2
            } else {
                result = append(result , "t")
            }
        } else if strings.Contains("fjlmnr",c){
            result = append(result, c)
        } else if c == "g" {
            if strings.Contains("iey",next){
                result = append(result, "j")
            } else if !strings.Contains("hn", next) {
                result = append(result , "k")
            } else if next == "h" && len(nextnext)!=0 && !strings.Contains("aeiou",nextnext){
                i+=1
            }
        } else if c == "h" {
            if i==0 || strings.Contains("aeiou",next) || !strings.Contains("aeiou",string(s[i-1])){
                result = append(result,"h")
            }
        } else if c=="k" {
            if i==0 || string(s[i-1])!= "c" {
                result = append(result,"k")
            }
        } else if c=="p" {
            if next == "h" {
                result = append(result , "f")
                i += 1
            } else {
                result = append(result, "p")
            }
        } else if c == "q" {
            result = append(result, "k")
        } else if c == "s" {
            if next == "h" {
                result = append(result, "x")
                i += 1
            } else if next == "i" && strings.Contains("oa",nextnext) {
                result = append(result,"x")
                i += 2
            } else {
                result = append(result , "s")
            }
        } else if c =="t"{
            if next == "i" && strings.Contains("oa",nextnext) {
                result = append(result, "x")
            } else if next == "h" {
                result = append(result, "0")
            } else if next!="c" || nextnext !="h" {
                result = append(result, "t")
            }
        } else if c == "v" {
            result = append(result,"f")
        } else if c == "w" {
            if i==0 && next =="h"{
                i+=1
                if strings.Contains("aeiou",nextnext) || nextnext=="*****" {
                    result = append(result,"w")
                }
            } else if strings.Contains("aeiou",next) || next == "*****" {
                result = append(result, "w")
            }
        } else if c=="x"{
            if i==0{
                if next =="h" || (next == "i" && strings.Contains("oa",nextnext)){
                    result = append(result, "x")
                } else {
                    result = append(result, "s")
                }

            } else {
                result = append(result, "k")
                result = append(result, "s")
            }
        } else if c =="y" {
            if strings.Contains("aeiou", next){
                result = append(result, "y")
            }
        } else if c == "z" {
            result = append(result, "s")
        } else if c == " "{
            if len(result) > 0 && result[len(result)-1] != " "{
                result = append(result," ")
            }
        }

        i +=1

    }
    
    return strings.ToUpper(strings.Join(result,""))

}
