package reader

const (
	max19 = 1844674407370955161
	c1e19 = 10000000000000000000
)

var eTab = []uint64{
	1e0, 1e1, 1e2, 1e3, 1e4, 1e5, 1e6, 1e7, 1e8, 1e9, 1e10,
	1e11, 1e12, 1e13, 1e14, 1e15, 1e16, 1e17, 1e18, 1e19,
}

func (r *Reader) skipNumber() error {
	dat, i := r.dat[r.pos:], 0
	dp, ln := 0, len(dat)

	// sign
	if i < ln && dat[0] == '-' {
		i++
	}

	// digits
	if i >= ln {
		return NewEndOfFileError()

	} else if dat[i] == '0' {
		if i++; i < ln && dat[i]-'0' <= 9 {
			return NewInvalidCharacterError(dat[i], r.pos+i)
		}

	} else if dat[i]-'0' <= 9 {
		for ; i < ln && dat[i]-'0' <= 9; i++ {
			/* do nothing */
		}
	} else {
		return NewInvalidCharacterError(dat[i], r.pos+i)
	}

	if i < ln && dat[i] == '.' {
		dp = i
		for i++; i < ln && dat[i]-'0' <= 9; i++ {
			/* do nothing */
		}
		if dp+1 == i {
			if i == ln {
				return NewEndOfFileError()
			} else {
				return NewInvalidCharacterError(dat[i], r.pos+i)
			}
		}
	}

	// exponent
	if i < ln && dat[i]|0x20 == 'e' {
		if i++; i >= ln {
			return NewEndOfFileError()
		} else if dat[i] != '-' && dat[i] != '+' {
			return NewInvalidCharacterError(dat[i], r.pos+i)
		}

		if i++; i >= ln {
			return NewEndOfFileError()
		} else if dat[i]-'0' > 9 {
			return NewInvalidCharacterError(dat[i], r.pos+i)
		}

		for ; i < ln && dat[i]-'0' <= 9; i++ {
			/* do nothing */
		}
	}

	r.pos += i
	return nil
}

func readInteger(dat []byte) (integer uint64, negative bool, i int, ok bool) {
	var intg uint64
	var dig, exp, dp, sp int
	var neg, trc, en bool
	var d20 byte
	ln := len(dat)

	// sign
	if ln != 0 && dat[0] == '-' {
		neg = true
		i++
	}

	// digits
	if i >= ln {
		return
	} else if dat[i]-'0' == 0 {
		if i++; i == ln {
			/* integer zero */
		} else if dat[i]-'0' <= 9 {
			return
		} else if dat[i] == '.' {
			dp = i
			for i++; i < ln && dat[i] == '0'; i++ {
				// do nothing
			}
			sp = i
		}
	} else if dat[i]-'0' <= 9 {
		sp = i
		for ; i < ln && dat[i]-'0' <= 9; i++ {
			if dig < 19 {
				intg = intg*10 + uint64(dat[i]-'0')
				dig++
			}
		}
		if i < ln && dat[i] == '.' {
			dp = i
			i++
		}
	} else {
		return
	}
	for ; i < ln && dat[i]-'0' <= 9; i++ {
		if dig < 19 {
			intg = intg*10 + uint64(dat[i]-'0')
			dig++
		}
	}

	// no digit after dot
	if dp > 0 && dp+1 == i {
		return
	}
	// digits were truncated
	if dig < 19 {
		/* do nothing */
	} else if dp <= sp {
		if trc = sp+dig < i; trc {
			d20 = dat[sp+19] - '0'
		}
	} else if dp > 19 {
		if trc = sp+dig+1 < i; trc {
			d20 = dat[sp+19] - '0'
		}
	} else {
		if trc = sp+dig+1 < i; trc {
			d20 = dat[sp+20] - '0'
		}
	}
	// no dot found
	if dp == 0 {
		dp = i
	}

	// exponent
	if i < ln && (dat[i]|0x20) == 'e' {
		if i++; i >= ln {
			return
		} else if dat[i] == '-' {
			en = true
		} else if dat[i] != '+' {
			return
		}
		if i++; i >= ln || dat[i]-'0' > 9 {
			return
		}

		for ; i < ln && dat[i]-'0' <= 9; i++ {
			if exp < 10000 {
				exp *= 10
				exp += int(dat[i] - '0')
			}
		}
		if intg == 0 {
			exp = 0
		} else if en {
			exp = -exp
		}
	}

	// finalizing
	exp += dotExp(dig, dp, sp, trc)
	if intg == 0 || dig+exp <= 0 {
		return 0, false, i, true
	} else if dig+exp > 20 {
		return 0, true, i, false
	} else if exp == 0 {
		return intg, neg, i, true
	} else if exp < 0 {
		return intg / eTab[-exp], neg, i, true
	} else if dig+exp < 20 {
		return intg * eTab[exp], neg, i, true
	} else if intg < max19 || d20 <= 5 {
		return intg*eTab[exp] + uint64(d20), neg, i, true
	} else {
		return 0, true, i, false
	}
}

func readFloat(dat []byte) (
	man uint64, dig int, exp int, neg bool, trc bool, dp int, sp int, i int, ok bool,
) {
	ln, en := len(dat), false

	// sign
	if ln != 0 && dat[0] == '-' {
		neg = true
		i++
	}

	// digits
	if i == ln {
		return
	} else if dat[i]-'0' == 0 {
		if i++; i == ln {
			/* integer zero */
		} else if dat[i]-'0' <= 9 {
			return
		} else if dat[i] == '.' {
			dp = i
			for i++; i < ln && dat[i] == '0'; i++ {
				/* do nothing */
			}
			sp = i
		}
	} else if dat[i]-'0' <= 9 {
		sp = i
		for ; i < ln && dat[i]-'0' <= 9; i++ {
			if dig < 19 {
				man = man*10 + uint64(dat[i]-'0')
				dig++
			}
		}
		if i < ln && dat[i] == '.' {
			dp = i
			i++
		}
	} else {
		return
	}
	for ; i < ln && dat[i]-'0' <= 9; i++ {
		if dig < 19 {
			man = man*10 + uint64(dat[i]-'0')
			dig++
		}
	}

	// no digit after dot
	if dp > 0 && dp+1 == i {
		return
	}
	// mantissa was truncated
	if dig < 19 {
		/* do nothing */
	} else if dp <= sp {
		trc = sp+19 < i
	} else {
		trc = sp+20 < i
	}
	// no dot found
	if dp == 0 {
		dp = i
	}

	// exponent
	if i < ln && (dat[i]|0x20) == 'e' {
		if i++; i >= ln {
			return
		} else if dat[i] == '-' {
			en = true
		} else if dat[i] != '+' {
			return
		}
		if i++; i >= ln || dat[i]-'0' > 9 {
			return
		}

		for ; i < ln && dat[i]-'0' <= 9; i++ {
			if exp < 10000 {
				exp *= 10
				exp += int(dat[i] - '0')
			}
		}
		if man == 0 {
			exp = 0
		} else if en {
			exp = -exp
		}
	}

	ok = true
	return
}

func dotExp(d int, dp int, sp int, trc bool) int {
	if trc && sp+d < dp {
		// dp - (sp+d)
		return dp - sp - d
	} else if sp < dp {
		// (dp+1)-(sp+d+1)
		return dp - sp - d
	} else {
		// (dp+1)-(sp+d)
		return dp + 1 - sp - d
	}
}
