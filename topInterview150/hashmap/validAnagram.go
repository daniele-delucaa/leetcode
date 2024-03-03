func isAnagram(s string, t string) bool {
	letters1 := make(map[byte]int)
	letters2 := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		letters1[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		letters2[t[i]]++
	}
	return reflect.DeepEqual(letters1, letters2)
}