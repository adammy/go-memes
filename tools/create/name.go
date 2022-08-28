package create

import (
	"errors"
	"fmt"
	"strings"
)

func ValidateName(name string) error {
	switch {
	case name == "":
		return errors.New("name should not be empty")
	case strings.Contains(name, "_"):
		return fmt.Errorf("package name %s should not contain underscores", name)
	case strings.Contains(name, "-"):
		return fmt.Errorf("package name %s should not contain hyphens", name)
	case strings.ToLower(name) != name:
		return fmt.Errorf("package name %s should be all lowercase", name)
	default:
		return nil
	}
}
