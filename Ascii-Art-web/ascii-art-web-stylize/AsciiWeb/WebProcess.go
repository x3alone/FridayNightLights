package web

import (
	"fmt"
	"net/http"
	"strings"
)

/* the function that handles the printing of the specified banner*/
func fs_art(str, banner string) (string, error) {
	if check_input(str) {
		return "", fmt.Errorf("400")
	}

	if len(str)-strings.Count(str, "\n") > 250 {
		return "", fmt.Errorf("400")
	}

	new_lines, trigger := check_newline(str)
	if trigger {
		return new_lines, nil
	}

	if banner != "standard" && banner != "thinkertoy" && banner != "shadow" {
		return "", fmt.Errorf("400")
	}
	data := read_file(banner + ".txt")
	shape := array_2d(data)
	result := print_shapes(shape, str)
	return result, nil
}

func Collect_info(req *http.Request) (string, error) {
	args := []string{req.PostFormValue("input"), req.PostFormValue("banners")}

	data, err := fs_art(args[0], args[1])

	if err != nil {
		return "", fmt.Errorf("400")
	}
	return data, err
}
