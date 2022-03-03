package main

import (
	"fmt"
	"strconv"
	"time"
)

func Validate_fields(fields map[string]Field) map[string]Field {
	// i := 1
	fields_validated := make(map[string]Field)
	for key, field := range fields {
		// field := reflect.ValueOf(&Field).Elem()
		field = Field{
			Field_type:    field.Field_type,
			Field_size:    "(" + field.Field_size + ")",
			Field_default: field.Field_default,
		}
		field_type := field.Field_type
		field = Validate_basic(field)
		switch {
		case field_type == "varchar":
			field = Validate_varchar(field)
		case field_type == "bool":
			field = Validate_bool(field)
		case field_type == "integer":
			field = Validate_integer(field)
		case field_type == "date":
			field = Validate_date(field)
			// default:
			// 	field = validate_basic(field)
			// }
		}
		fields_validated[key] = field
	}
	return fields_validated
}

func Validate_basic(field Field) Field {

	if field.Field_default == "" {
		field = Field{
			Field_type: field.Field_type,
			Field_size: field.Field_size,
		}
	}
	if field.Field_type == "integer" || field.Field_type == "bool" || field.Field_type == "date" {
		field = Field{
			Field_type:    field.Field_type,
			Field_default: field.Field_default,
		}
	}

	return field
}

func Validate_integer(field Field) Field {
	if field.Field_default != "" {
		_, err := strconv.Atoi(field.Field_default)
		if err != nil {
			fmt.Println(err)
			field = Field{
				Field_type: field.Field_type,
			}
		}
	}
	return field
}

func Validate_varchar(field Field) Field {
	return field
}

func Validate_bool(field Field) Field {
	if field.Field_default != "" {
		_, err := strconv.ParseBool(field.Field_default)
		if err != nil {
			fmt.Println(err)
			field = Field{
				Field_type: field.Field_type,
			}
		}
	}
	return field
}

func Validate_date(field Field) Field {
	if field.Field_default != "" {
		proper_date := "2006-01-02"
		_, err := time.Parse("01/02/2006", proper_date)
		if err != nil {
			panic(err)
		}
	}
	return field
}
