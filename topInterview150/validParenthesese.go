func isValid(s string) bool {
	var stack []string
	m := make(map[string]string) // mappare ")" con la corrispondente aperta "(", e lo stesso con le altre
	m[")"] = "("
	m["]"] = "["
	m["}"] = "{"
	for i := range s {
		if val, ok := m[string(s[i])]; ok { // se la parentesi è una chiusa, controlliamo con la mappa
			if len(stack) > 0 && stack[len(stack)-1] == val { // se la parentesi prima della chiusa, è la corrispondente aperta, le due parentesi "si matchano"
				stack = stack[:len(stack)-1] // pop
			} else {
				return false
			}
		} else { // se la parentesi è una aperta
			stack = append(stack, string(s[i])) // push
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}