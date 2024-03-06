package main

func mergeAlternately(word1 string, word2 string) string {
	var res string
	i := 0
	j := len(word1)
	k := len(word2)
	for i < j && i < k {
		res += string(word1[0]) + string(word2[0])
		word1 = word1[1:]
		word2 = word2[1:]
		i++
	}
	for i < j {
		res += string(word1[0])
		word1 = word1[1:]
		i++
	}
	for i < k {
		res += string(word2[0])
		word2 = word2[1:]
		i++
	}
	return res
}
