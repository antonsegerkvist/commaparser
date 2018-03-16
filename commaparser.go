package commaparser

// InvalidCharacterError is the error returned when the string contains a
// invalid character.
type InvalidCharacterError struct{}

func (*InvalidCharacterError) Error() string {
	return "Invalid character found in string"
}

// InvalidFormatError is the error returned when the string is malformated.
type InvalidFormatError struct{}

func (*InvalidFormatError) Error() string {
	return "Invalid format of query"
}

// alphabet is the alphabet of valid field characters.
var alphabet = map[rune]bool{
	'a': true,
	'b': true,
	'c': true,
	'd': true,
	'e': true,
	'f': true,
	'g': true,
	'h': true,
	'i': true,
	'j': true,
	'k': true,
	'l': true,
	'm': true,
	'n': true,
	'o': true,
	'p': true,
	'q': true,
	'r': true,
	's': true,
	't': true,
	'u': true,
	'v': true,
	'w': true,
	'x': true,
	'y': true,
	'z': true,
}

const (
	// delimiter is the character separating the field characters.
	delimiter = ','

	// all is the character indicating all fields should be used.
	all = '*'
)

// ParseString parses a comma separated string and returns a hash mapping to
// true for fields that the string contained.
func ParseString(str string, fields []string) (map[string]bool, error) {
	runes := []rune(str)
	state := 0
	sIndex := 0
	ret := make(map[string]bool)

	for index, v := range str {
		switch state {
		case 0:
			if v == all {
				state = 1
			} else if v == delimiter {
				return make(map[string]bool), &InvalidFormatError{}
			} else if alphabet[v] == true {
				sIndex = index
				state = 2
			} else {
				return make(map[string]bool), &InvalidCharacterError{}
			}
			break

		case 1:
			return make(map[string]bool), &InvalidFormatError{}

		case 2:
			if v == all {
				return make(map[string]bool), &InvalidFormatError{}
			} else if v == delimiter {
				ret[string(runes[sIndex:index])] = true
				state = 3
			} else if alphabet[v] != true {
				return make(map[string]bool), &InvalidCharacterError{}
			}
			break

		case 3:
			if v == all || v == delimiter {
				return make(map[string]bool), &InvalidFormatError{}
			} else if alphabet[v] == true {
				sIndex = index
				state = 2
			} else {
				return make(map[string]bool), &InvalidCharacterError{}
			}
			break
		}
	}

	if state == 1 {
		for _, v := range fields {
			ret[v] = true
		}
		return ret, nil
	} else if state == 2 {
		ret[string(runes[sIndex:len(runes)])] = true
		return ret, nil
	}

	return make(map[string]bool), &InvalidFormatError{}
}
