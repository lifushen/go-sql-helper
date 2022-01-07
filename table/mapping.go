package table

import (
	"crypto/md5"
	"encoding/hex"
	"reflect"
)

var tableCache map[string]interface{}

func New(table interface{}, modelInstance interface{}) interface{} {
	refType := reflect.TypeOf(modelInstance)
	key := MD5(refType.PkgPath() + "/" + refType.Name())
	if _, ok := tableCache[key]; !ok {
		setCache(table, key, refType)
	}
	return tableCache[key]
}

func setCache(tb interface{}, key string, modelType reflect.Type) {
	refValue := reflect.ValueOf(tb)
	table := reflect.New(refValue.Type()).Elem()

	fieldCount := modelType.NumField()
	for i := 0; i < fieldCount; i++ {
		fieldName := modelType.Field(i).Name
		fieldTag := modelType.Field(i).Tag

		var tempField Field
		tempField.ColumnName = fieldTag.Get("column")

		refTargetValue := reflect.ValueOf(tempField)
		if table.FieldByName(fieldName).IsValid() {
			table.FieldByName(fieldName).Set(refTargetValue.Convert(table.FieldByName(fieldName).Type()))
		}
	}

	if tableCache == nil {
		tableCache = make(map[string]interface{})
	}
	tableCache[key] = table.Interface()
}

func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}
