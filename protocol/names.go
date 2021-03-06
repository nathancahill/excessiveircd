// Copyright (c) 2014 Michael Johnson. All rights reserved.
//
// Use of this source code is governed by the BSD license that can be found in
// the LICENSE file.

package protocol

// CheckFunc checks if a given r at index i is within its specification.
type CheckFunc func(i int, r rune) (valid bool)

// Or returns a CheckFunc that returns true if any of the given funcs return
// true.
// If no funcs are given, it always returns true.
// It will short circuit on the first true func.
func Or(funcs ...CheckFunc) CheckFunc {
	return func(i int, r rune) bool {
		for _, f := range funcs {
			if f(i, r) {
				return true
			}
		}

		return false
	}
}

// And returns a CheckFunc that returns true if all of the given funcs return
// true.
// If no funcs are given, it always returns true.
// It will short circuit on the first false func.
func And(funcs ...CheckFunc) CheckFunc {
	return func(i int, r rune) bool {
		for _, f := range funcs {
			if !f(i, r) {
				return false
			}
		}

		return true
	}
}

// Letter checks if r is a valid letter, as according to RFC 2812.
func Letter(i int, r rune) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') // A-Z / a-z
}

// Special checks if r is a valid special character, as according to RFC 2812.
func Special(i int, r rune) bool {
	return (r >= ']' && r <= '`') || (r >= '{' && r <= '}') // "[", "]", "\", "`", "_", "^", "{", "|", "}"
}

// Digit checks if r is a valid digit, as according to RFC 2812.
func Digit(i int, r rune) bool {
	return r >= '0' && r <= '9' // 0-9
}

// MaxMaximumNickLen specifies the maximum length that a nickname may be.
const MaximumNickLen = 15

// Nickname checks if r at index i is a valid nickname character, as according
// to RFC 2812.
func Nickname(i int, r rune) bool {
	if i >= MaximumNickLen {
		// Nicknames have a maximum length of 9 characters. The runes are
		// 0-indexed.
		return false
	}
	if i == 0 {
		if !Letter(i, r) && !Special(i, r) {
			return false
		}

		return true
	}

	if !Letter(i, r) && !Digit(i, r) && !Special(i, r) && r != '-' {
		return false
	}

	return true
}

// Username checks if r at index i is a valid nickname character, as according
// to RFC 2812.
func Username(i int, r rune) bool {
	// TODO(nightexcessive): Is this correct? It seems to allow overly broad characters such as \t.
	if r == 0x00 || r == '\r' || r == '\n' || r == ' ' || r == '@' || r > 0xFF || // These are characters barred by RFC 2812.
		r == '~' || r == '*' { // These are characters barred by our special uses.
		return false
	}

	return true
}

// IsValid checks s against checkFunc and returns true if all runes are given
// success values by checkFunc.
// It short circuits at the first false value.
func IsValid(s string, checkFunc CheckFunc) bool {
	for i, r := range s {
		if !checkFunc(i, r) {
			return false
		}
	}

	return true
}
