func addBinary(a string, b string) string {
	var s, t []byte
	i := len(a) - 1
	j := len(b) - 1
	carry := false
	for i > -1 && j > -1 {
		if a[i] == '1' && b[j] == '1' {
			if carry {
				s = append(s, '1')
			} else {
				s = append(s, '0')
			}
			carry = true
		} else if a[i] == '1' || b[j] == '1' {
			if carry {
				s = append(s, '0')
			} else {
				s = append(s, '1')
			}
		} else {
			if carry {
				s = append(s, '1')
				carry = false
			} else {
				s = append(s, '0')
			}
		}
		i--
		j--
	}
	for i > -1 {
		if a[i] == '1' {
			if carry {
				s = append(s, '0')
			} else {
				s = append(s, '1')
			}
		} else {
			if carry {
				s = append(s, '1')
				carry = false
			} else {
				s = append(s, '0')
			}
		}
		i--
	}

	for j > -1 {
		if b[j] == '1' {
			if carry {
				s = append(s, '0')
			} else {
				s = append(s, '1')
			}
		} else {
			if carry {
				s = append(s, '1')
				carry = false
			} else {
				s = append(s, '0')
			}
		}
		j--
	}

	if carry {
		s = append(s, '1')
	}

	for p := len(s) - 1; p > -1; p-- {
		t = append(t, s[p])
	}
	return t.String()
}