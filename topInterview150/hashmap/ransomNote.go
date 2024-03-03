package main

func canConstruct(ransomNote string, magazine string) bool {
	m1 := make(map[byte]int)
	m2 := make(map[byte]int)
	for i := range ransomNote {
		m1[ransomNote[i]]++
	}
	for i := range magazine {
		m2[magazine[i]]++
	}
	for _, ch := range ransomNote {
		if _, ok := m2[byte(ch)]; ok {
			if m2[byte(ch)] == 0 {
				return false
			}
			m2[byte(ch)]--
		} else {
			return false
		}
	}
	return true
}
