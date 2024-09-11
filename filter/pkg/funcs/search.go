package funcs

import (
	"strconv"
	"strings"
)

var FoundArtists []Artist

func Search(keyword string) {
	var found []Artist
	keyword = strings.ToLower(keyword)
	for i, v := range All.Artists {
		if strings.Contains(strings.ToLower(v.FirstAlbum), keyword) {
			found = append(found, All.Artists[i])
		}
		if strings.Contains(strings.ToLower(v.Name), keyword) {
			found = append(found, All.Artists[i])
		}
		if strings.Contains(strconv.Itoa(v.CreationDate), keyword) {
			found = append(found, All.Artists[i])
		}
		for _, k := range v.Members {
			if strings.Contains(strings.ToLower(k), keyword) {
				found = append(found, All.Artists[i])
			}
		}
		_, ok := v.Rel[keyword]
		if ok {
			found = append(found, All.Artists[i])
		}
		if !ok {
			for j := range All.Artists[i].Rel {
				if strings.Contains(strings.ToLower(j), keyword) {
					found = append(found, All.Artists[i])
					break
				}
			}
		}
	}
	FoundArtists = Unique(found)
}

func Unique(found []Artist) []Artist {
	newlist := []Artist{}
	for _, v := range found {
		if !FoundIn(newlist, v.ID) {
			newlist = append(newlist, v)
		}
	}
	return newlist
}

func FoundIn(newlist []Artist, ID int) bool {
	for _, v := range newlist {
		if v.ID == ID {
			return true
		}
	}
	return false
}
