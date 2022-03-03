package main

type Field struct {
	Field_type    string
	Field_size    string
	Field_default string
}

type Object struct {
	table_name string
	Fields     map[string]Field
}
