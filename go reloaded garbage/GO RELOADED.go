package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func ReadFile(name string) (string, error) {

	data, err := os.ReadFile(name)
	if err != nil {
	}
	return string(data), nil
}

func SplitWhiteSpaces(s string) []string {
	var result []string
	var word string
	for i, c := range s {
		if c != ' ' && c != '\t' && c != '\n' {
			word += string(c)
		}
		if ((c == ' ' || c == '\t' || c == '\n') && word != "") || i == len(s)-1 {
			result = append(result, word)
			word = ""
		}
	}
	return result
}

func nums(arr []string) []string {
	i := 0
	for i < len(arr) {
		if i > 0 {
			switch arr[i] {
			case "(hex)":
				if num, err := strconv.ParseInt(arr[i-1], 16, 64); err == nil {
					arr[i-1] = fmt.Sprintf("%d", num)
					arr = append(arr[:i], arr[i+1:]...)
					i--
				} else {
					i++
				}
			case "(bin)":
				if num, err := strconv.ParseInt(arr[i-1], 2, 64); err == nil {
					arr[i-1] = fmt.Sprintf("%d", num)
					arr = append(arr[:i], arr[i+1:]...)
					i--
				} else {
					i++
				}
			default:
				i++
			}
		} else {
			i++
		}
	}
	return arr
}

// func hex(arr []string) []string {
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] == "(hex)" && i > 0 {
// 			if num, err := strconv.ParseInt(arr[i-1], 16, 64); err == nil {
// 				arr[i-1] = fmt.Sprintf("%d", num)
// 				arr = append(arr[:i], arr[i+1:]...)
// 			}
// 		}
// 	}
// 	return arr
// }

//	func bin(arr []string, index) []string {
//		for i := 0; i < len(arr); i++ {
//			if arr[i] == "(bin)" && i > 0 {
//				if num, err := strconv.ParseInt(arr[i-1], 2, 64); err == nil {
//					arr[i-1] = fmt.Sprintf("%d", num)
//					arr = append(arr[:i], arr[i+1:]...)
//				}
//			}
//		}
//		return arr
//	}
func IsNumeric(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}
	return false
}

func applyToUpper(arr []string, index int) []string {
	if index > 0 && index-1 < len(arr) {
		s := arr[index-1]
		arr[index-1] = FindAlpha(strings.ToUpper, s)
	}
	return append(arr[:index], arr[index+1:]...) // delete after modif
}

func applyToLower(arr []string, index int) []string {
	if index > 0 && index-1 < len(arr) {
		s := arr[index-1]
		arr[index-1] = FindAlpha(strings.ToLower, s)
	}
	return append(arr[:index], arr[index+1:]...) // ""
}

func applyToCapitalize(arr []string, index int) []string {
	if index > 0 && index-1 < len(arr) {
		s := arr[index-1]
		arr[index-1] = FindAlpha(strings.ToLower, s)
		arr[index-1] = capitalizeFirstAlphabetical(s)
	}
	return append(arr[:index], arr[index+1:]...) // ""
}

// lower the string then finds the first letter and caps it
func capitalizeFirstAlphabetical(s string) string {
	lowered := strings.ToLower(s)
	for i, ch := range lowered {
		if unicode.IsLetter(ch) {
			return lowered[:i] + strings.ToUpper(string(ch)) + lowered[i+1:]
		}
	}
	return lowered // retruni lowered if no word is found
}

func capitalizeWords(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		words[i] = strings.ToUpper(word[:1]) + word[1:]
	}
	return strings.Join(words, " ")
}

func FindAlpha(transform func(string) string, s string) string {
	transformed := []rune(s)
	for i, ch := range transformed {
		if unicode.IsLetter(ch) {
			transformed[i] = rune(transform(string(ch))[0])
		}
	}
	return string(transformed)
}

func AlphaStringsCount(arr []string) int { // uses IsAlpha to count the alpha strings so we can match the counter in major func
	count := 0
	for _, s := range arr {
		if IsAlpha(s) {
			count++
		}
	}
	return count
}

func IsAlpha(s string) bool { //checks if the string contains alpha by skipping non alpha; ps i used this to skip "" or other examples
	for _, ch := range s {
		if unicode.IsLetter(ch) {
			return true
		}
	}
	return false
}

func applyRepeatedFormat(arr []string, directive string, index int, counter int) []string {
	if counter <= 0 || index-1 < 0 || index-1 >= len(arr) || index-1-counter < 0 {
		fmt.Println("Error: Invalid index or counter.")
		return arr // return unmodified array on error
	}

	// Process elements from index-1 and move backwards up to counter elements
	for i := 0; i < counter && (index-1-i) >= 0; i++ {
		currentIndex := index - 1 - i
		switch directive {
		case "(up,":
			arr[currentIndex] = FindAlpha(strings.ToUpper, arr[currentIndex])
		case "(low,":
			arr[currentIndex] = FindAlpha(strings.ToLower, arr[currentIndex])
		case "(cap,":
			arr[currentIndex] = capitalizeWords(arr[currentIndex])
		}
	}

	// Optionally remove elements after processing, here we keep all elements
	return arr
}

func major(arr []string) []string {
	i := 0
	alphabeticalCount := AlphaStringsCount(arr) // Count alphabetical strings once

	for i < len(arr) {
		if strings.HasPrefix(arr[i], "(") {
			if i == 0 || (!IsAlpha(arr[i-1]) && !IsNumeric(arr[i-1])) {
				// Skip if no alphabetical or numeric word is before the directive
				i++
				continue
			}

			switch arr[i] {
			case "(up)":
				arr = applyToUpper(arr, i)
			case "(low)":
				arr = applyToLower(arr, i)
			case "(cap)":
				arr = applyToCapitalize(arr, i)
			case "(hex)":
				if num, err := strconv.ParseInt(arr[i-1], 16, 64); err == nil {
					arr[i-1] = fmt.Sprintf("%d", num)
					arr = append(arr[:i], arr[i+1:]...)
					i-- // Move back to reevaluate the position with possibly new directives
				}
			case "(bin)":
				if num, err := strconv.ParseInt(arr[i-1], 2, 64); err == nil {
					arr[i-1] = fmt.Sprintf("%d", num)
					arr = append(arr[:i], arr[i+1:]...)
					i-- // Move back to reevaluate the position with possibly new directives
				}
			default:
				if strings.HasSuffix(arr[i], ",") && i+1 < len(arr) {
					counter, err := strconv.Atoi(strings.Trim(arr[i+1], ")"))
					if err == nil && counter > 0 && counter <= alphabeticalCount {
						arr = applyRepeatedFormat(arr, arr[i], i, counter)
						i -= 2 // Adjust for removed elements
					} else {
						i++
					}
				} else {
					i++
				}
			}
		} else {
			i++
		}
	}
	return arr
}

func vowels(arr []string) []string {
	for i := 0; i < len(arr)-1; i++ {
		if len(arr[i+1]) > 0 {
			firstLetter := strings.ToLower(string(arr[i+1][0]))
			switch arr[i] {
			case "a", "A", "'a", "'A":
				if strings.ContainsAny(firstLetter, "aeiouhAEIOUH") {
					switch arr[i] {
					case "A":
						arr[i] = "An"
					case "a":
						arr[i] = "an"
					case "'a":
						arr[i] = "'an"
					case "'A":
						arr[i] = "'An"
					}
				}
			case "an", "An", "'an", "'An":
				if !strings.ContainsAny(firstLetter, "aeiouhAEIOUH") {
					switch arr[i] {
					case "An":
						arr[i] = "A"
					case "an":
						arr[i] = "a"
					case "'An":
						arr[i] = "'A"
					case "'an":
						arr[i] = "'a"
					}
				}
			}
		}
	}
	return arr
}

func punctuationsHandler(arr []string) []string {
	i := 1
	for i < len(arr) {
		if arr[i] != "" {
			if arr[i][0] == ',' || arr[i][0] == '.' || arr[i][0] == ':' || arr[i][0] == '!' || arr[i][0] == '?' || arr[i][0] == ';' {
				arr[i-1] += arr[i][:1]
				if len(arr[i]) > 1 {
					arr[i] = arr[i][1:]
				} else {
					arr = append(arr[:i], arr[i+1:]...)
				}
				continue
			}
		}
		i++
	}
	return arr
}

func FormatQuotes(text string) string {
	result := ""
	firstQuote := true
	for i := 0; i < len(text); i++ {
		if text[i] == '\'' {
			if i == 0 && firstQuote {
				if text[i+1] == ' ' {
					text = text[:i+1] + text[i+2:]
					result += string(text[i])
					firstQuote = false
					continue
				} else {
					firstQuote = false
					result += string(text[i])
					continue
				}
			}
			if i != len(text)-1 && i > 0 && (text[i+1] == ' ' || text[i-1] == ' ') && firstQuote {
				if text[i+1] == ' ' {
					text = text[:i+1] + text[i+2:]
					result += string(text[i])
					firstQuote = false
					continue
				} else {
					firstQuote = false
					result += string(text[i])
					continue
				}
			}
			if !firstQuote {
				if i != len(text)-1 && i >= 0 && (text[i-1] == ' ' || text[i+1] == ' ') {
					if text[i-1] == ' ' {
						result = result[:len(result)-1]
						result += string(text[i])
						firstQuote = true
						continue
					} else {
						firstQuote = true
						result += string(text[i])
						continue
					}
				}
			}
			if i == len(text)-1 && !firstQuote {
				if text[i-1] == ' ' {
					result = result[:len(result)-1]
					result += string(text[i])
					firstQuote = true
					continue
				} else {
					firstQuote = false
					result += string(text[i])
					continue
				}
			}
		}
		if text[i] != ' ' || (i > 0 && text[i-1] != ' ') {
			result += string(text[i])
		}
	}
	return result
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("file name is not provided")
		return
	}
	file, err := ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("error reading file : %s\n", err)
		return
	}
	arr := SplitWhiteSpaces(file)
	arr = nums(arr)
	arr = major(arr)
	arr = vowels(arr)
	arr = punctuationsHandler(arr)
	new := strings.Join(arr, " ")
	new = FormatQuotes(new)
	// fmt.Println(strings.Join(arr, " "))
	fmt.Print(new)
}
