package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

var Revision string

type User struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min:12"`
}

func main() {

	user := User{
		Email:    "info@freecoder.dev",
		Password: "password",
	}

	fmt.Println(ValidateData(user))
}

func ValidateData(data interface{}) error {
	v := reflect.ValueOf(data)
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("validate")

		if tag != "" {
			value := v.Field(i)
			tags := strings.Split(tag, ",")
			for _, tag := range tags {
				switch {
				case tag == "required":
					if value.Interface() == reflect.Zero(value.Type()).Interface() {
						return fmt.Errorf("validation failed: field '%s' is required", field.Name)
					}
				case strings.HasPrefix(tag, "min:"):
					var min int
					fmt.Sscanf(tag, "min:%d", &min)
					if len(value.String()) < min {
						return fmt.Errorf("validation failed: field '%s' should have a minimum lenght of %d",
							field.Name,
							min)
					}
				case tag == "email":
					regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
					valid, _ := regexp.MatchString(regex, value.String())
					if !valid {
						return fmt.Errorf("validation failed: field '%s' should be a valid email",
							field.Name)
					}
				}
			}
		}
	}
	return nil
}
