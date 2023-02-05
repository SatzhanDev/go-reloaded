package pkg

import "strings"

func ToFixPunct(fixedSpace []string) []string {
	var fixedPunt string
	char := []string{",", ".", ":", ";", "?", "!", "'", "(", ")"}

	for i, v := range fixedSpace {
		if len(fixedSpace) < 2 {
			fixedPunt += fixedSpace[i]
		} else if i != len(fixedSpace)-1 {
			if v == char[0] || v == char[1] || v == char[2] || v == char[3] || v == char[4] || v == char[5] {
				if v == char[5] && fixedSpace[i+1] == char[4] {
					fixedPunt += fixedSpace[i]
				} else if (v == char[1] && fixedSpace[i+1] == char[1]) || (v == char[1] && fixedSpace[i+1] == char[6]) {
					fixedPunt += fixedSpace[i]
				} else {
					fixedPunt += fixedSpace[i]
					fixedPunt += " "
				}
			} else if v == char[6] {
				count := 0
				for _, l := range fixedPunt {
					if string(l) == char[6] {
						count++
					}
				}
				if count%2 == 0 {
					fixedPunt += fixedSpace[i]
				} else {
					fixedPunt += fixedSpace[i]
					fixedPunt += " "
				}
			} else if v == char[7] {
				fixedPunt += fixedSpace[i]
			} else if v == char[8] {
				if fixedSpace[i+1] == "," || fixedSpace[i+1] == "." || fixedSpace[i+1] == ";" || fixedSpace[i+1] == ":" || fixedSpace[i+1] == ")" || fixedSpace[i+1] == "!" || fixedSpace[i+1] == "?" {
					fixedPunt += fixedSpace[i]
				} else {
					fixedPunt += fixedSpace[i]
					fixedPunt += " "
				}
			} else {
				yesChar := false
				for _, k := range char[:6] {
					if fixedSpace[i+1] == k {
						fixedPunt += fixedSpace[i]
						yesChar = true
					}
				}
				if yesChar == false {
					yesDigit := false
					for _, p := range v {
						if p > 47 && p < 58 && fixedSpace[i+1] == ")" {
							fixedPunt += fixedSpace[i]
							yesDigit = true
							break
						}
					}
					if yesDigit == false {
						if fixedSpace[i+1] == "," || fixedSpace[i+1] == "." || fixedSpace[i+1] == ";" || fixedSpace[i+1] == ":" || fixedSpace[i+1] == ")" || fixedSpace[i+1] == "!" || fixedSpace[i+1] == "?" {
							fixedPunt += fixedSpace[i]
						} else {
							fixedPunt += fixedSpace[i]
							fixedPunt += " "
						}
					}
				}
			}
		} else if i == len(fixedSpace)-1 {
			fixedPunt += fixedSpace[i]
		}
	}
	var fixedDontStr string
	for i, v := range fixedPunt {
		if len(fixedPunt) > 1 {
			if v == ' ' && fixedPunt[i+1] == '\'' && ((fixedPunt[i+2] == 't' && fixedPunt[i-1] == 'n') || fixedPunt[i+2] == 's') {
				continue
			} else if v == ' ' && (v != rune(fixedPunt[len(fixedPunt)-2]) || v != rune(fixedPunt[len(fixedPunt)-1])) {
				if v == ' ' && ((fixedPunt[i+1] == '.' && fixedPunt[i+2] == '.') || (fixedPunt[i+1] == '!' && fixedPunt[i+2] == '?')) {
					continue
				} else {
					fixedDontStr += string(fixedPunt[i])
				}
			} else {
				fixedDontStr += string(fixedPunt[i])
			}
		} else {
			fixedDontStr += string(fixedPunt[i])
		}
	}
	fixedPunct2 := strings.Fields(fixedDontStr)
	return fixedPunct2
}
