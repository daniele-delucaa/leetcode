package main

// conto le occorrenze della stringa s grazie a una mappa,
// itero sulla stringa t, se trovo che un carattere
// della stringa t è presente anche in s (grazie all'uso
// della mappa), decremento l'occorrenza sulla mappa
// altrimenti trovo il carattere da restituire.
// Il decremento sulla mappa serve per forza,
// se abbiamo un caso limite in cui s= "a", t ="aa",
// se l'if invece di rilevare se la chiave sia presente
// rilevasse se non sia presente, il programma non
// funzionerebbe, non trova che la seconda 'a', della stringa
// t, è il carattere in più.
// Restituirebbe il valore predefinito di b (byte)
func findTheDifference(s string, t string) byte {
	var b byte
	occS := make(map[byte]int)
	for i := range s {
		occS[s[i]]++
	}
	for i := range t {
		if _, ok := occS[t[i]]; ok {
			occS[t[i]]--
			if occS[t[i]] == 0 {
				delete(occS, t[i])
			}
		} else {
			b = t[i]
		}
	}
	return b
}

// fa le due somme dei valori in byte di ogni carattere
// delle stringhe, se facciamo la differenza delle due somme
// totali troviamo il valore del carattere in più sulla stringa
// t.
func findTheDifference(s string, t string) byte {
	sumS, sumT := 0, 0
	for _, i := range s {
		sumS += int(i)
	}

	for _, i := range t {
		sumT += int(i)
	}

	diff := sumT - sumS

	return byte(diff)
}
