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
	for i := 0; i < len(arr); i++ {
		if arr[i] == "(hex)" && i > 0 {
			// if arr[i-1] != strconv.Atoi(arr[i-1]){
			// 	return arr
			// }
			if num, err := strconv.ParseInt(arr[i-1], 16, 64); err == nil {
				arr[i-1] = fmt.Sprintf("%d", num)
				arr = append(arr[:i], arr[i+1:]...)
			}
		} else if arr[i] == "(bin)" && i > 0 {
			// if arr[i-1] != strconv.Atoi(arr[i-1]){
			// 	return arr
			// }
			if num, err := strconv.ParseInt(arr[i-1], 2, 64); err == nil {
				arr[i-1] = fmt.Sprintf("%d", num)
				arr = append(arr[:i], arr[i+1:]...)
			}
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
	if index > 0 && index-1 < len(arr) {
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

func mapFirstAlphabetical(transform func(string) string, s string) string { //finds first letter
	for i, ch := range s {
		if unicode.IsLetter(ch) {
			return s[:i] + transform(string(ch)) + s[i+1:] //this part we sliced the s string (substring) until to specify the whers is the first letter
		} //ch is the first alpha in the string by using isletter func, transform passes the provided info the map, so the other func can modify
	}
	return s //retruni string lwl if no letter is found
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
	var currentIndex int = index - 1
	// if counter <= 0 || counter > len(arr){
	// 	fmt.Println("err")
	// 	os.Exit(1)
	// }
	for j := 0; j < counter && currentIndex >= 0; j++ {
		if directive == "(up," {
			arr[currentIndex] = strings.ToUpper(arr[currentIndex])
		} else if directive == "(low," {
			arr[currentIndex] = strings.ToLower(arr[currentIndex])
		} else if directive == "(cap," {
			if len(arr[currentIndex]) > 0 {
				arr[currentIndex] = strings.ToUpper(string(arr[currentIndex][0])) + arr[currentIndex][1:]
			}
		}
		currentIndex--
	}
	return append(arr[:index], arr[index+2:]...)
}

func major(arr []string) []string {
	i := 0
	alphabeticalCount := AlphaStringsCount(arr) // Count alphabetical strings once

	for i < len(arr) {
		if strings.HasPrefix(arr[i], "(") {
			if i == 0 || !IsAlpha(arr[i-1]) {
				// Skip if no alphabetical word is before the directive
				i++
				continue
			}

			switch arr[i] {
			case "(up)", "(low)", "(cap)", "(hex)", "(bin)":
				switch arr[i] {
				case "(up)":
					arr = applyToUpper(arr, i)
				case "(low)":
					arr = applyToLower(arr, i)
				case "(cap)":
					arr = applyToCapitalize(arr, i)
				case "(hex)", "(bin)":
					arr = append(arr[:i], arr[i+1:]...) // Remove the directive
				}
				// all cases the arr is changed, therefore i doesn't increment 
			default:
				if strings.HasSuffix(arr[i], ",") && i+1 < len(arr) {
					counter, err := strconv.Atoi(strings.Trim(arr[i+1], ")"))
					if err == nil && counter > 0 && counter <= alphabeticalCount {
						arr = applyRepeatedFormat(arr, arr[i], i, counter)
						i -= 2 // Adjust for removed elements
					} else {
						i++
					}
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
			FirstLetter := strings.ToLower(string(arr[i+1][0]))
			if (arr[i] == "a" || arr[i] == "A") && strings.ContainsAny(FirstLetter, "aeiouhAEIOUH") {
				if arr[i] == "A" {
					arr[i] = "An"
				} else if arr[i] == "a" {
					arr[i] = "an"
				}
			} else if (arr[i] == "an" || arr[i] == "An") && !strings.ContainsAny(FirstLetter, "aeiouhAEIOUH") {
				if arr[i] == "An" {
					arr[i] = "A"
				} else if arr[i] == "an" {
					arr[i] = "a"
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
	fmt.Println(strings.Join(arr, " "))
}

