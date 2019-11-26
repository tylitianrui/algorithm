package main

func isValid(s string) bool {

	tmp := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			tmp = append(tmp, s[i])
		}
		if len(tmp) == 0 && (s[i] == ')' || s[i] == '}' || s[i] == ']') {
			return false
		}

		if s[i] == ')' {
			if tmp[len(tmp)-1] == '(' {
				tmp = tmp[:len(tmp)-1]
			} else {
				tmp = append(tmp, s[i])
			}

		}

		if s[i] == '}' {
			if tmp[len(tmp)-1] == '{' {
				tmp = tmp[:len(tmp)-1]
			} else {
				tmp = append(tmp, s[i])
			}
		}

		if s[i] == ']' {
			if tmp[len(tmp)-1] == '[' {
				tmp = tmp[:len(tmp)-1]
			} else {
				tmp = append(tmp, s[i])
			}
		}

	}
	if len(tmp) == 0 {
		return true
	}
	return false
}
