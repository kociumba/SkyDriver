package styles

import (
	"errors"
	"strings"
)

var (
	errNoDivider    = errors.New("input does not contain the divider")
	errEmptyDivider = errors.New("divider cannot be empty")
)

// EqualSpacingOnDivider equalizes the length before and after the divider.
func EqualSpacingOnDivider(s, divider string) (string, error) {
	split, err := Split(s, divider)
	if err != nil {
		return "", err
	}

	length1 := len(split[0])
	length2 := len(split[1])

	if length1 == length2 {
		return s, nil
	}

	if length1 > length2 {
		return split[0] + divider + strings.Repeat(" ", length1-length2) + split[1], nil
	}

	return split[0] + strings.Repeat(" ", length2-length1) + divider + split[1], nil
}

// EqualSpacingOnDividerFromInput keeps the same length before a divider character.
func EqualSpacingOnDividerFromInput(s, divider string, lengthBeforeDivider int) (string, error) {
	split, err := Split(s, divider)
	if err != nil {
		return "", err
	}

	length1 := len(split[0])
	length2 := len(split[1])

	if length1 == lengthBeforeDivider {
		return s, nil
	}

	if length1 > lengthBeforeDivider {
		return split[0] + divider + strings.Repeat(" ", length1-lengthBeforeDivider) + split[1], nil
	}

	return split[0] + strings.Repeat(" ", lengthBeforeDivider-length1) + divider + split[1] + strings.Repeat(" ", lengthBeforeDivider-length2), nil
}

// Split splits the string s using the divider and returns an error if the divider is not found.
func Split(s, divider string) ([]string, error) {
	if divider == "" {
		return nil, errEmptyDivider
	}

	split := strings.SplitN(s, divider, 2)

	if len(split) < 2 {
		return nil, errNoDivider
	}

	return split, nil
}
