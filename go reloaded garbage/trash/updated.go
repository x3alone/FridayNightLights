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
func hasAlphabetical(s string) bool {
    for _, ch := range s {
        if unicode.IsLetter(ch) {
            return true
        }
    }
    return false
}

func major(arr []string) []string {
    i := 0
    for i < len(arr) {
        if strings.HasPrefix(arr[i], "(") {
            if i == 0 || !hasAlphabetical(arr[i-1]) {
                // If no alphabetical word before the keyword, skip it
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
                    arr = append(arr[:i], arr[i+1:]...)
                }
                // In all cases above, the array is modified, so the index does not need to increment
            default:
                i++ // Increment if not a recognized command
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

