package main

import "fmt"

type County struct {
	name     string
	state    string
	id       string
	numBeds  int
	timeline [95]int
}

func convertToCode(given string) string {
	switch given {
	case "alabama":
		return "al"
	case "alaska":
		return "ak"
	case "american samoa":
		return "as"
	case "arkansas":
		return "ar"
	case "california":
		return "ca"
	case "colorado":
		return "co"
	case "delaware":
		return "de"
	case "district of columbia":
		return "dc"
	case "florida":
		return "fl"
	case "georgia":
		return "ga"
	case "guam":
		return "gu"
	case "hawaii":
		return "hi"
	case "idaho":
		return "id"
	case "illinois":
		return "il"
	case "indiana":
		return "in"
	case "iowa":
		return "ia"
	case "kansas":
		return "ks"
	case "kentucky":
		return "ky"
	case "louisiana":
		return "la"
	case "maine":
		return "me"
	case "maryland":
		return "md"
	case "massachusetts":
		return "ma"
	case "michigan":
		return "mi"
	case "minnesota":
		return "mn"
	case "mississippi":
		return "ms"
	case "missouri":
		return "mo"
	case "montana":
		return "mt"
	case "nebraska":
		return "ne"
	case "nevada":
		return "nv"
	case "new hampshire":
		return "nh"
	case "new jersey":
		return "nj"
	case "new mexico":
		return "nm"
	case "new york":
		return "ny"
	case "north carolina":
		return "nc"
	case "north dakota":
		return "nd"
	case "northern mariana is":
		return "mp"
	case "ohio":
		return "oh"
	case "oklahoma":
		return "ok"
	case "oregon":
		return "or"
	case "pennsylvania":
		return "pa"
	case "puerto rico":
		return "pr"
	case "rhode island":
		return "ri"
	case "south carolina":
		return "sc"
	case "south dakota":
		return "sd"
	case "tennessee":
		return "tn"
	case "texas":
		return "tx"
	case "utah":
		return "ut"
	case "vermont":
		return "vt"
	case "virginia":
		return "va"
	case "virgin islands":
		return "vi"
	case "washington":
		return "wa"
	case "west virginia":
		return "wv"
	case "wisconsin":
		return "wi"
	case "wyoming":
		return "wy"
	default:
		return "not given"

	}

}

func newCounty(newName string, newState string) County {
	county := County{}
	county.name = newName
	county.state = newState
	county.id = fmt.Sprintf("%s.%s", county.state, county.name)
	county.numBeds = 0
	for i := 0; i < len(county.timeline); i++ {
		county.timeline[i] = 0
	}
	return county
}

func findCounty(counties []County, name string, state string) int {
	for i := 0; i < len(counties); i++ {
		if counties[i].name == name && counties[i].state == state {
			return i
		}
	}
	//if you've searched through the entire array and haven't found it
	return -1
}
