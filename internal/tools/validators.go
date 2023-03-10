package tools

import (
	"errors"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func Validate(s string, validators ...func(string) error) error {
	for _, v := range validators {
		if err := v(s); err != nil {
			return err
		}
	}
	return nil
}

func IsMinMaxLen(min, max int) func(string) error {
	return func(str string) error {
		l := len([]rune(str))
		if l < min {
			return errors.New("the value is too short length characters (minimum is " + strconv.Itoa(min) + " characters vs your " + strconv.Itoa(l) + " character(s))")
		} else if l > max {
			return errors.New("the value is too long length characters (maximum is " + strconv.Itoa(max) + " characters vs your " + strconv.Itoa(l) + " character(s))")
		}
		return nil
	}
}

func IsLen(length int) func(string) error {
	return func(str string) error {
		l := len([]rune(str))
		if l != length {
			return errors.New("the value is not the required length characters (required is " + strconv.Itoa(length) + " vs your " + strconv.Itoa(l) + " character(s))")
		}
		return nil
	}
}

func IsEmail() func(string) error {
	return func(str string) error {
		regex, _ := regexp.Compile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
		if !regex.MatchString(str) {
			return errors.New("the value is not an email")
		}
		return nil
	}
}

func IsNickname() func(string) error {
	return func(str string) error {
		regex, _ := regexp.Compile(`^[a-zA-Z0-9_-]+$`)
		if !regex.MatchString(str) {
			return errors.New("the value is not a nickname")
		}
		return nil
	}
}

func IsUrl() func(string) error {
	return func(str string) error {
		_, err := url.ParseRequestURI(str)
		if err != nil {
			return errors.New("the value is not a url address")
		}
		return nil
	}
}

func IsInt64() func(string) error {
	return func(str string) error {
		_, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return errors.New("the value is not an integer data type")
		}
		return nil
	}
}

func IsUint64() func(string) error {
	return func(str string) error {
		_, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			return errors.New("the value is not an unsigned integer data type")
		}
		return nil
	}
}

func IsNotBlank() func(string) error {
	return func(str string) error {
		if strings.TrimSpace(str) == "" {
			return errors.New("the value is blank")
		}
		return nil
	}
}

func IsNotContainsSpace() func(string) error {
	return func(str string) error {
		for p, c := range str {
			if unicode.IsSpace(c) {
				return errors.New("the value contains a space (space in " + strconv.Itoa(p+1) + " character)")
			}
		}
		return nil
	}
}
