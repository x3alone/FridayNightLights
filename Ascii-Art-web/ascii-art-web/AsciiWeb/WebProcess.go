package web

import (
	"fmt"
	"net/http"
)

/* the function that handles the printing of the specified banner*/
func fs_art(str, banner string) (string, error) {
	if check_input(str) {
		return "", fmt.Errorf("non supported charachters")
	}
	new_lines, trigger := check_newline(str)
	if trigger {

		return new_lines, nil
	}
	if banner == "standard" || banner == "thinkertoy" || banner == "shadow" {

		data := read_file(banner + ".txt")

		shape := array_2d(data)

		result := print_shapes(shape, str)

		return result, nil
	}
	return "", fmt.Errorf("wrong banner")
}

func Collect_info(req *http.Request) (string, error) {
	args := []string{req.PostFormValue("input"), req.PostFormValue("banners")}

	data, err := fs_art(args[0], args[1])

	if err != nil {
		return "", fmt.Errorf("500")
	} else {
		return data, err
	}
}
