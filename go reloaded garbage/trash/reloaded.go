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

func applyToUpper(arr []string, index int) []string {
	if index > 0 && index-1 < len(arr) {
		s := arr[index-1]
		arr[index-1] = mapFirstAlphabetical(strings.ToUpper, s)
	}
	return append(arr[:index], arr[index+1:]...) // delete after modif
}

func applyToLower(arr []string, index int) []string {
	if index > 0 && index-1 < len(arr) {
		s := arr[index-1]
		arr[index-1] = mapFirstAlphabetical(strings.ToLower, s)
	}
	return append(arr[:index], arr[index+1:]...) // ""
}

func applyToCapitalize(arr []string, index int) []string {
	if index >= 0 && index < len(arr) {
		s := arr[index-1]
		arr[index-1] = mapFirstAlphabetical(strings.ToLower, s)
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

func mapFirstAlphabetical(transform func(string) string, s string) string { // finds first letter
	for i, ch := range s {
		if unicode.IsLetter(ch) {
			return s[:i] + transform(string(ch)) + s[i+1:] // this part we sliced the s string (substring) until to specify the whers is the first letter
		} // ch is the first alpha in the string by using isletter func, transform passes the provided info the map, so the other func can modify
	}
	return s // retruni string lwl if no letter is found
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

func IsAlpha(s string) bool { // checks if the string contains alpha by skipping non alpha; ps i used this to skip "" or other examples
	for _, ch := range s {
		if unicode.IsLetter(ch) {
			return true
		}
	}
	return false
}

func applyRepeatedFormat(arr []string, directive string, index int, counter int) []string {
	var currentIndex int = index - 1
	for j := 0; j < counter && currentIndex >= 0; j++ {
		if directive == "(up," {
			arr[currentIndex] = strings.ToUpper(arr[currentIndex])
		} else if directive == "(low," {
			arr[currentIndex] = strings.ToLower(arr[currentIndex])
		} else if directive == "(cap," {
			if len(arr[currentIndex]) > 0 {
				arr[currentIndex] = strings.ToUpper(string(arr[currentIndex][0])) + arr[currentIndex][1:]
			}
		} else {
			return arr
		}
		currentIndex--
	}
	return append(arr[:index], arr[index+2:]...)
}

func countAlphabeticalBefore(arr []string, index int) int {
	count := 0
	for i := 0; i < index; i++ {
		if IsAlpha(arr[i]) { // Assume IsAlpha checks if the string is purely alphabetical.
			count++
		}
	}
	return count
}

func major(arr []string) []string {
	i := 0
	for i < len(arr) {
		if strings.HasPrefix(arr[i], "(") && i > 0 && IsAlpha(arr[i-1]) {
			// Handling each correct directive format
			validDirective := false
			var count int
			var err error

			if strings.HasSuffix(arr[i], ")") {
				switch arr[i] {
				case "(up)":
					arr = applyToUpper(arr, i)
					validDirective = true
				case "(low)":
					arr = applyToLower(arr, i)
					validDirective = true
				case "(cap)":
					arr = applyToCapitalize(arr, i)
					validDirective = true
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
				}
			} else if idx := strings.Index(arr[i], ","); idx != -1 && i+1 < len(arr) {
				// Handling repeated commands (up, low, cap)
				directive := arr[i][:idx+1] // includes the comma
				count, err = strconv.Atoi(strings.Trim(arr[i+1], ")"))
				if err == nil && count > 0 && count <= countAlphabeticalBefore(arr, i) {
					arr = applyRepeatedFormat(arr, directive, i, count)
					validDirective = true
					i++ // Additional increment to skip the number part
				}
			}

			if validDirective {
				// Successful directive means we need to adjust index due to the removal of a directive
				i-- // The next element shifts into the current position
			}
		}
		i++ // Move to the next element
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

func Quotaions(arr string) string {
	res := ""
	firstQuote := true
	for i := 0; i < len(arr); i++ {
		if arr[i] == '\'' {
			if i == 0 && firstQuote {
				if arr[i+1] == ' ' {
					arr = arr[:i+1] + arr[i+2:]
					res += string(arr[i])
					firstQuote = false
					continue
				} else {
					firstQuote = false
					res += string(arr[i])
					continue
				}
			}
			if i != len(arr)-1 && i > 0 && (arr[i+1] == ' ' || arr[i-1] == ' ') && firstQuote {
				if arr[i+1] == ' ' {
					arr = arr[:i+1] + arr[i+2:]
					res += string(arr[i])
					firstQuote = false
					continue
				} else {
					firstQuote = false
					res += string(arr[i])
					continue
				}
			}
			if !firstQuote {
				if i != len(arr)-1 && i >= 0 && (arr[i-1] == ' ' || arr[i+1] == ' ') {
					if arr[i-1] == ' ' {
						res = res[:len(res)-1]
						res += string(arr[i])
						firstQuote = true
						continue
					} else {
						firstQuote = true
						res += string(arr[i])
						continue
					}
				}
			}
			if i == len(arr)-1 && !firstQuote {
				if arr[i-1] == ' ' {
					res = res[:len(res)-1]
					res += string(arr[i])
					firstQuote = true
					continue
				} else {
					firstQuote = false
					res += string(arr[i])
					continue
				}
			}
		}
		if arr[i] != ' ' || (i > 0 && arr[i-1] != ' ') {
			res += string(arr[i])
		}
	}
	return res
}

func check(e error) {
	if e != nil {
		fmt.Println("error: not enough arguments")
	}
}

func main() {
	if len(os.Args) != 3 {
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
	new = Quotaions(new)
	result := strings.Trim(new, " ")
	err = os.WriteFile(os.Args[2], []byte(result), 0o644)
	check(err)
}

