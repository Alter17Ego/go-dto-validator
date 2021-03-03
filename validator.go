package validator

import (
	"reflect"
	"strconv"
)

func Validate(mystruct interface{}) (err error) {
	field := reflect.TypeOf(mystruct)
	value := reflect.ValueOf(mystruct)
	if field.Kind() != reflect.Struct && field.Kind() != reflect.Slice {
		return myerr(STRUCT_ERROR)
	}
	switch field.Kind() {
	case reflect.Slice:
		err = fetchSlice(value)
		if err != nil {
			return
		}
	default:
		for i := 0; i < field.NumField(); i++ {
			type_kind := field.Field(i).Type.Kind()
			val := value.Field(i).Interface()
			value_kind := reflect.TypeOf(val).Kind()
			if type_kind == reflect.Struct || type_kind == reflect.Slice && value_kind != reflect.Slice {
				err = Validate(val)
			} else if type_kind != reflect.Map || type_kind != reflect.Array {
				err = validateTags(field.Field(i), val)
			}
			if err != nil {
				return
			}
		}
	}
	return
}
func validateTags(field reflect.StructField, value interface{}) (err error) {
	for i := 0; i < len(tags); i++ {
		tag, stat := field.Tag.Lookup(tags[i])
		field_tag, stat_required := field.Tag.Lookup("field")
		required := false
		if stat == true {
			if stat_required == true && field_tag == "required" {
				required = true
			}
			structInfo := structDetail{
				name:      field.Name,
				val:       value,
				tag_name:  tags[i],
				tag_value: tag,
				required:  required,
			}
			err := getType(&structInfo).validate()
			if err != nil {
				return err
			}
		}
	}
	return err
}
func getType(s *structDetail) myvalidator {
	tagsValue := map[string]myvalidator{
		"email":        email{s},
		"text":         text{s},
		"number":       number{s},
		"required":     required{s},
		"default":      def{s},
		"alphabet":     alphabet{s},
		"alphanumeric": alphanumeric{s},
	}
	res := tagsValue["default"]
	switch s.tag_name {
	case "min":
		lenChar, err := strconv.Atoi(s.tag_value)
		if err != nil {
			res = tagsValue["default"]
		}
		tagsValue[s.tag_name] = min{min: lenChar, detail: s}
		res = tagsValue["min"]
	case "max":
		lenChar, err := strconv.Atoi(s.tag_value)
		if err != nil {
			res = tagsValue["default"]
		}
		tagsValue[s.tag_name] = max{max: lenChar, detail: s}
		res = tagsValue["max"]
	case "startswith":
		tagsValue[s.tag_name] = startswith{s}
		res = tagsValue["startswith"]
	case "endswith":
		tagsValue[s.tag_name] = endswith{s}
		res = tagsValue["endswith"]
	case "value_of":
		tagsValue[s.tag_name] = valueOf{s}
		res = tagsValue["value_of"]
	default:
		if _, ok := tagsValue[s.tag_value]; ok {
			res = tagsValue[s.tag_value]
		}
	}
	return res
}

func fetchSlice(t reflect.Value) (err error) {
	for i := 0; i < t.Len(); i++ {
		val := t.Index(i).Interface()
		err = Validate(val)
		if err != nil {
			return
		}
	}
	return err
}
