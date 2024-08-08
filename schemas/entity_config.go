package schemas

type EntityConfig struct {
	EntityName  string           `json:"entity_name"`
	EntityField []EntityFieldMap `json:"entity_field"`
}

type EntityFieldMap struct {
	DataType   string `json:"data_type"`
	IsOptional bool   `json:"is_optional"`
	FieldName  string `json:"field_name"`
}
