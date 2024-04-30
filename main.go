package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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
	for i := 1; i < len(arr); i++ {
		if arr[i] == "(hex)" && i > 0 {
			if num, err := strconv.ParseInt(arr[i-1], 16, 64); err == nil {
				arr[i-1] = fmt.Sprintf("%d", num)
				arr = append(arr[:i], arr[i+1:]...)
			}
		} else if arr[i] == "(bin)" && i > 0 {
			if num, err := strconv.ParseInt(arr[i-1], 2, 64); err == nil {
				arr[i-1] = fmt.Sprintf("%d", num)
				arr = append(arr[:i], arr[i+1:]...)
			}
		}
	}
	return arr
}

func applyFormatBasedOnDirective(arr []string, i int) []string {
	switch {
	case arr[i] == "(up)" && i > 0:
		arr[i-1] = strings.ToUpper(arr[i-1])
	case arr[i] == "(low)" && i > 0:
		arr[i-1] = strings.ToLower(arr[i-1])
	case arr[i] == "(cap)" && i > 0:
		arr[i-1] = capitalize(arr[i-1])
	case strings.HasPrefix(arr[i], "(up,") || strings.HasPrefix(arr[i], "(low,") || strings.HasPrefix(arr[i], "(cap,"):
		if i+1 < len(arr) {
			arr = applyMultiFormat(arr, i)
		}
	}
	return arr
}

func capitalize(s string) string {
	if len(s) > 0 {
		return strings.ToUpper(string(s[0])) + s[1:]
	}
	return s
}

func applyMultiFormat(arr []string, i int) []string {
	counter, err := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(arr[i], "(up,"), ")"))
	if err != nil {
		counter, err = strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(arr[i], "(low,"), ")"))
	}
	if err != nil {
		counter, err = strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(arr[i], "(cap,"), ")"))
	}

	if err == nil {
		for j := 0; j < counter && (i-j-1) >= 0; j++ {
			switch {
			case strings.HasPrefix(arr[i], "(up,"):
				arr[i-j-1] = strings.ToUpper(arr[i-j-1])
			case strings.HasPrefix(arr[i], "(low,"):
				arr[i-j-1] = strings.ToLower(arr[i-j-1])
			case strings.HasPrefix(arr[i], "(cap,"):
				arr[i-j-1] = capitalize(arr[i-j-1])
			}
		}
		return append(arr[:i], arr[i+2:]...)
	}
	return arr
}

func processDirectives(arr []string) []string {
	i := 0
	for i < len(arr) {
		if strings.HasPrefix(arr[i], "(") && i > 0 {
			arr = applyFormatBasedOnDirective(arr, i)
			if strings.ContainsAny(arr[i], ",") {
				i -= 2 // adjust for changes in array length
			} else {
				i-- // adjust for the removed element
			}
		}
		i++
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

func mergeQuotedStrings(arr []string) []string {
	n := len(arr)
	if n == 0 {
		return arr
	}

	var result []string
	i := 0

	for i < n {
		if len(arr[i]) > 0 && arr[i][0] == '\'' {
			// Accumulate parts of the string within quotes
			temp := []string{}
			j := i
			inQuotes := false

			for ; j < n; j++ {
				current := arr[j]
				startsQuote := current[0] == '\''
				endsQuote := current[len(current)-1] == '\''

				if startsQuote && endsQuote && len(current) == 1 {
					// Current string is just "'", handle specially
					if inQuotes {
						// Close the current quote
						inQuotes = false
					} else {
						// Open a new quote
						inQuotes = true
					}
					continue
				}

				if startsQuote {
					// Strip leading quote if not already inside quotes
					if !inQuotes {
						current = current[1:]
					}
					inQuotes = true
				}

				if endsQuote {
					// Strip trailing quote if it ends the quote
					current = current[:len(current)-1]
					inQuotes = false
				}

				// Append the current processed string part
				temp = append(temp, current)

				if !inQuotes {
					// Quote closed, stop processing further
					break
				}
			}

			// Join all parts gathered within the quotes, add to result
			result = append(result, "'"+strings.Join(temp, " ")+"'")
			i = j // Move the index past the processed segment
		} else {
			result = append(result, arr[i])
		}
		i++
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
	arr = mergeQuotedStrings(arr)
	fmt.Println(strings.Join(arr, " "))
}
