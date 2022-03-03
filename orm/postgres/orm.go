package main

import (
	"d/go/utils"
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

type Field struct {
	Field_type    string
	Field_size    string
	Field_default string
}

type Object struct {
	table_name string
	Fields     map[string]Field
}

func NewObject(name string, fields map[string]Field) *Object {
	return &Object{table_name: name, Fields: fields}
}

// func set_up_table(names ...string, fields ...map[string]Field) {}

func main() {
	// user := NewObject()

	// user := NewObject("users")
	db := utils.Set_db("user=postgres password=1234 dbname=main_db sslmode=disable")

	user_fields := map[string]Field{
		"user_name": {
			Field_type:    "varchar",
			Field_size:    "30",
			Field_default: "",
		},
		"password": {
			Field_type:    "varchar",
			Field_size:    "30",
			Field_default: "",
		},
		"id": {
			Field_type:    "integer",
			Field_size:    "30",
			Field_default: "huy",
		},
		"is_admin": {
			Field_type:    "bool",
			Field_size:    "30",
			Field_default: "huy",
		},
		"is_moderator": {
			Field_type:    "bool",
			Field_size:    "30",
			Field_default: "true",
		},
	}

	user := NewObject("users", user_fields)
	fmt.Println(create_table(db, user))
	// fields := validate_fields(user.Fields)
	// for key, field := range fields {
	// 	fmt.Println(key, field)
	// }
	// db.Exec("create")
}

func validate_fields(fields map[string]Field) map[string]Field {
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
		field = validate_basic(field)
		switch {
		case field_type == "varchar":
			field = validate_varchar(field)
		case field_type == "bool":
			field = validate_bool(field)
		case field_type == "integer":
			field = validate_integer(field)
		case field_type == "date":
			field = validate_date(field)
			// default:
			// 	field = validate_basic(field)
			// }
		}
		fields_validated[key] = field
	}
	return fields_validated
}

func validate_basic(field Field) Field {

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

func validate_integer(field Field) Field {
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

func validate_varchar(field Field) Field {
	return field
}

func validate_bool(field Field) Field {
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

func validate_date(field Field) Field {
	if field.Field_default != "" {
		proper_date := "2006-01-02"
		_, err := time.Parse("01/02/2006", proper_date)
		if err != nil {
			panic(err)
		}
	}
	return field
}

func format_fields(object *Object) string {
	var res string
	var defaults string
	fields := validate_fields(object.Fields)
	for key, field := range fields {
		res += key + " " + field.Field_type + " " + field.Field_size + " " + defaults + ", "
	}
	fmt.Println(res)
	return res[:len(res)-2]
}

func create_table(db *sql.DB, object *Object) sql.Result {
	fields_formatted := format_fields(object)
	result, err := db.Exec("create table " + object.table_name + "(" + fields_formatted + ");")
	if err != nil {
		fmt.Println(err)
	}
	return result
}
