package jellyfish

import (
	"GoJellyfish/pkg/util"
	"strings"
)

func Levenshtein_distance(s1, s2 string) int {
	if strings.Compare(s1, s2) == 0 {
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

		cur := util.Generate_arr(0, cols)
		for r := 1; r < rows; r++ {
			prev := make([]int, cols)
			util.Memset(prev, 0)
			prev[0] = r

			for c := 1; c < cols; c++ {
				deletion := prev[c] + 1
				insertion := cur[c-1] + 1
				edit := prev[c-1]
				if s1[r-1] != s2[c-1] {
					edit += 1
				}
				cur[c] = util.Min(edit, deletion, insertion)
			}

		}

		return cur[len(cur)-1]

	}
}

func Damerau_levenshtein_distance(s1, s2 string) int {
	len1 := len(s1)
	len2 := len(s2)
	infinite := len1 + len2
	da := make(map[byte]int)
	score := make([][]int, len1+2)
	//Distance matrix
	for i := range score {
		score[i] = make([]int, len2+2)
		util.Memset(score[i], 0)
	}

	score[0][0] = infinite

	for i := 0; i < len1+1; i++ {
		score[i+1][0] = infinite
		score[i+1][1] = i
	}

	for i := 0; i < len2+1; i++ {
		score[0][i+1] = infinite
		score[1][i+1] = i
	}

	for i := 1; i < len1+1; i++ {
		db := 0
		for j := 1; j < len2+1; j++ {
			i1 := da[s2[j-1]]
			j1 := db
			cost := 1
			if s1[i-1] == s2[j-1] {
				cost = 0
				db = j
			}

			score[i+1][j+1] = util.Min(score[i][j]+cost, score[i+1][j]+1, score[i][j+1]+1, score[i1][j1]+(i-i1-1)+1+(j-j1-1))
		}
		da[s1[i-1]] = i
	}

	return score[len1+1][len2+1]
}

func Jaro_distance(s1, s2 string) float64 {
	return _jaro_winkler(s1, s2, false, false)
}

func _jaro_winkler(ying, yang string, long_tolerance, winkelerize bool) float64 {
	ying_len := len(ying)
	yang_len := len(yang)

	if ying_len == 0 || yang_len == 0 {
		return 0.0
	}

	min_len := util.Max(ying_len, yang_len)
	search_range := (min_len / 2) - 1

	if search_range < 0 {
		search_range = 0
	}

	ying_flags := make([]int, ying_len)
	util.Memset(ying_flags, 0)
	yang_flags := make([]int, yang_len)
	util.Memset(yang_flags, 0)

	//looking only within search range, count & flag matched pairs
	common_chars := 0
	for i, ying_ch := range ying {
		ying_ch := byte(ying_ch)
		low := i
		if i > search_range {
			low -= search_range
		}
		hi := i
		if i+search_range < yang_len {
			hi += search_range
		} else {
			hi += yang_len - 1
		}

		for j := low; j <= hi; j++ {
			if yang_flags[j] == 0 && yang[j] == ying_ch {
				ying_flags[i] = 1
				yang_flags[j] = 1
				common_chars += 1
				break
			}
		}
	}

	//short circuit if no characters matched
	if common_chars == 0 {
		return 0.0
	}

	k, trans_count := 0, float64(0)
	for i, ying_f := range ying_flags {
		if ying_f == 1 {
			temp := 0
			for j := k; j < yang_len; j++ {
				temp = j
				if yang_flags[j] == 1 {
					k = j + 1
					break
				}
			}

			if ying[i] != yang[temp] {
				trans_count += 1
			}
		}
	}

	trans_count /= 2

	var weight float64
    common_chars_f := float64(common_chars)
    ying_len_f := float64(ying_len)
    yang_len_f := float64(yang_len)
	weight = (common_chars_f/ying_len_f + common_chars_f/yang_len_f +
		                (common_chars_f-trans_count)/common_chars_f) / 3

	if winkelerize && weight > 0.7 && ying_len > 3 && yang_len > 3 {
		j := util.Min(min_len, 4)
		i := 0
		for i < j && ying[i] == yang[i] && string(yang[i]) != "" {
			i += 1
		}

		if i != 0 {
			weight += float64(i) * 0.1 * (1.0 - weight)
		}

		if long_tolerance && min_len > 4 && common_chars_f > float64(i+1) && 
                                    2*common_chars_f >= float64(min_len+1) {
			weight += ((1.0 - weight) * (common_chars_f - float64(i) - 1) / 
                                    (ying_len_f + yang_len_f - float64(i)*2 + 2))
		}
	}

	return weight
}

func  Jaro_winkler(s1, s2 string, long_tolerance ...bool) float64 {
    if len(long_tolerance) > 0 {
        return _jaro_winkler(s1, s2, long_tolerance[0], true)
    } else {
        return _jaro_winkler(s1, s2, true, true)
    }
}

func Hamming_distance(s1, s2 string) int{
    if len(s2) > len(s1){
        temp := s1
        s1 = s2
        s2 = temp
    }

    distance := len(s1) - len(s2)
    for i:=0; i<len(s2); i++ {
        if s2[i] != s1[i]{
            distance+=1
        }
    }
    return distance
}
