package main

import (
	"d/go/utils"
	"fmt"
)

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
	fmt.Println(Create_table(db, user))
	// fields := validate_fields(user.Fields)
	// for key, field := range fields {
	// 	fmt.Println(key, field)
	// }
	// db.Exec("create")
}
