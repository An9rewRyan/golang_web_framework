package main

import (
	"database/sql"
	"fmt"
)

func Format_fields(object *Object) string {
	var res string
	var defaults string
	fields := Validate_fields(object.Fields)
	for key, field := range fields {
		res += key + " " + field.Field_type + " " + field.Field_size + " " + defaults + ", "
	}
	fmt.Println(res)
	return res[:len(res)-2]
}

func Create_table(db *sql.DB, object *Object) sql.Result {
	fields_formatted := Format_fields(object)
	result, err := db.Exec("create table " + object.table_name + "(" + fields_formatted + ");")
	if err != nil {
		fmt.Println(err)
	}
	return result
}

func NewObject(name string, fields map[string]Field) *Object {
	return &Object{table_name: name, Fields: fields}
}
