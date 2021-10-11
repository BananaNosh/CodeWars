package main

import (
	"fmt"
	"regexp"
	"strings"
)


func findUnitLength(bits string) int {
	trimmed := strings.Trim(bits, "0")
	for i := range trimmed {
		reg, _ := regexp.Compile(fmt.Sprintf(`((1|\b)0{%d}(1|\b))|((0|\b)1{%d}(0|\b))`, i+1, i+1))
		if reg.MatchString(trimmed) {
			return i+1
		}
	}
	return 0
}

func DecodeBits(bits string) string {
	unit := findUnitLength(bits)
	regMap := []struct {replVal string; regex string}{
		{"   ", fmt.Sprintf(`0{%d}`, unit * 7)},
		{"   ", fmt.Sprintf(`0{%d}`, unit * 7)},
		{" ", fmt.Sprintf(`0{%d}`, unit * 3)},
		{"-", fmt.Sprintf(`1{%d}`, unit * 3)},
		{".", fmt.Sprintf(`1{%d}`, unit)},
		{"", fmt.Sprintf(`0{%d}`, 1)},
	}
	for _, regStruct := range regMap {
		wordSepReg, _ := regexp.Compile(regStruct.regex)
		bits = wordSepReg.ReplaceAllString(bits, regStruct.replVal)
	}
	return strings.TrimSpace(bits)
}


func DecodeMorse(morseCode string) string {
	var result []string
	for _, word := range strings.Split(strings.TrimSpace(morseCode), "   ") {
		var word_result []string
		for _, letter := range strings.Split(word, " ") {
			//word_result = append(word_result, MORSE_CODE[letter])
			word_result = append(word_result, letter)
		}
		result = append(result, strings.Join(word_result, ""))
	}
	return strings.Join(result, " ")
}


func main() {
	bits := "01110"
	fmt.Println(findUnitLength(bits))
	fmt.Println(DecodeMorse(".... . -.--   .--- ..- -.. ."))
	fmt.Println(DecodeBits(bits))
}