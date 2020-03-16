package 面试题_01_06

import "strconv"

func compressString(S string) string {
	if len(S) == 0 {
		return S
	}
	countWords := 0
	nowWords := int32(S[0])
	var resWords string
	for _, i := range S {
		if nowWords != i {
			resWords += string(nowWords) + strconv.Itoa(countWords)
			nowWords = i
			countWords = 0
		}
		countWords++
	}
	resWords += string(nowWords) + strconv.Itoa(countWords)

	if len(resWords) >= len(S) {
		return S
	} else {
		return resWords
	}
}
