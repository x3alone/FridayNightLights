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

func major(arr []string) []string {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == "(up)" && i > 0 {
			arr[i-1] = strings.ToUpper(arr[i-1])
			arr[i-1] = fmt.Sprintf("%s", arr[i-1])
			arr = append(arr[:i], arr[i+1:]...)
		} else if arr[i] == "(low)" && i > 0 {
			arr[i-1] = strings.ToLower(arr[i-1])
			arr[i-1] = fmt.Sprintf("%s", arr[i-1])
			arr = append(arr[:i], arr[i+1:]...)
		} else if arr[i] == "(cap)" && i > 0 {
			if len(arr[i-1]) > 0 {
				FirstChar := strings.ToUpper(string(arr[i-1][0]))
				if len(arr[i-1]) > 1 {
					arr[i-1] = FirstChar + arr[i-1][1:]
				} else {
					arr[i-1] = FirstChar
				}
			}
			arr[i-1] = fmt.Sprintf("%s", arr[i-1])
			arr = append(arr[:i], arr[i+1:]...)
		} else if strings.ContainsAny(arr[i], "(up,") || strings.ContainsAny(arr[i], "(low,") || strings.ContainsAny(arr[i], "(cap,") {
			counter, err := strconv.Atoi(strings.Trim(arr[i+1], ")"))
			o := i - 1
			if err == nil {
				for j := 0; j < counter; j++ {
					if arr[i] == "(up," {
						arr[o] = strings.ToUpper(arr[o])
						o--
					} else if arr[i] == "(low," {
						arr[o] = strings.ToLower(arr[o])
						o--
					} else if arr[i] == "(cap," {
						Fchar := strings.ToUpper(string(arr[o][0]))
						if len(arr[o]) > 1 {
							arr[o] = Fchar + arr[o][1:]
						} else {
							arr[o] = Fchar
						}
						o--
					}
				}
				arr = append(arr[:i], arr[i+2:]...)
			}
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

func punctuation(arr [] string)[]string {
	for i := 0; i < len(arr)-1; i++{
		if len(arr[i+1]) > 0 {
			 
		}
	}
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
	fmt.Println(strings.Join(arr, " "))
}
