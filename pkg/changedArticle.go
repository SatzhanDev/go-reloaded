package pkg

func ChangeVowelForward(fixedPunct []string) []string {
	for i, v := range fixedPunct {
		if len(fixedPunct) > 4 {
			if v == "a" && (fixedPunct[i+1][0] == 'a' || fixedPunct[i+1][0] == 'e' || fixedPunct[i+1][0] == 'i' || fixedPunct[i+1][0] == 'o' || fixedPunct[i+1][0] == 'u' || fixedPunct[i+1][0] == 'h') {
				fixedPunct[i] = "an"
			} else if v == "a" && (fixedPunct[i+1][0] == 'A' || fixedPunct[i+1][0] == 'E' || fixedPunct[i+1][0] == 'I' || fixedPunct[i+1][0] == 'O' || fixedPunct[i+1][0] == 'U' || fixedPunct[i+1][0] == 'H') {
				fixedPunct[i] = "an"
			} else if v == "A" && (fixedPunct[i+1][0] == 'a' || fixedPunct[i+1][0] == 'e' || fixedPunct[i+1][0] == 'i' || fixedPunct[i+1][0] == 'o' || fixedPunct[i+1][0] == 'u' || fixedPunct[i+1][0] == 'h') {
				fixedPunct[i] = "An"
			} else if v == "A" && (fixedPunct[i+1][0] == 'A' || fixedPunct[i+1][0] == 'E' || fixedPunct[i+1][0] == 'I' || fixedPunct[i+1][0] == 'O' || fixedPunct[i+1][0] == 'U' || fixedPunct[i+1][0] == 'H') {
				fixedPunct[i] = "AN"
			} else if v == "an" && fixedPunct[i+1][0] != 'a' && fixedPunct[i+1][0] != 'e' && fixedPunct[i+1][0] != 'i' && fixedPunct[i+1][0] != 'o' && fixedPunct[i+1][0] != 'u' && fixedPunct[i+1][0] != 'h' && fixedPunct[i+1][0] >= 95 && fixedPunct[i+1][0] <= 122 {
				fixedPunct[i] = "a"
			} else if v == "an" && fixedPunct[i+1][0] != 'A' && fixedPunct[i+1][0] != 'E' && fixedPunct[i+1][0] != 'I' && fixedPunct[i+1][0] != 'O' && fixedPunct[i+1][0] != 'U' && fixedPunct[i+1][0] != 'H' && fixedPunct[i+1][0] >= 65 && fixedPunct[i+1][0] <= 90 {
				fixedPunct[i] = "a"
			} else if v == "An" && fixedPunct[i+1][0] != 'a' && fixedPunct[i+1][0] != 'e' && fixedPunct[i+1][0] != 'i' && fixedPunct[i+1][0] != 'o' && fixedPunct[i+1][0] != 'u' && fixedPunct[i+1][0] != 'h' && fixedPunct[i+1][0] >= 95 && fixedPunct[i+1][0] <= 122 {
				fixedPunct[i] = "A"
			} else if v == "AN" && fixedPunct[i+1][0] != 'A' && fixedPunct[i+1][0] != 'E' && fixedPunct[i+1][0] != 'I' && fixedPunct[i+1][0] != 'O' && fixedPunct[i+1][0] != 'U' && fixedPunct[i+1][0] != 'H' && fixedPunct[i+1][0] >= 65 && fixedPunct[i+1][0] <= 90 {
				fixedPunct[i] = "A"
			}
		} else {
			break
		}
	}
	return fixedPunct
}
