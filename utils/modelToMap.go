package utils

import (
	"gorm.io/gorm/utils"
	"reflect"
	"time"
)

func ModelToMap(obj any, formatTime string, ignoreField ...string) map[string]any {
	m := make(map[string]any, 16)
	value := reflect.Indirect(reflect.ValueOf(obj))
	
	for i := 0; i < value.NumField(); i++ {
		types := value.Type()
		field := types.Field(i)
		name := field.Name
		value := value.FieldByName(name)
		tagName := field.Tag.Get("json")
		if utils.Contains(ignoreField, name) {
			continue
		}
		if v, ok := value.Interface().(time.Time); ok {
			timeStr := v.Format(formatTime)
			m[tagName] = timeStr
			continue
		}
		m[tagName] = value.Interface()
	}
	return m
}
