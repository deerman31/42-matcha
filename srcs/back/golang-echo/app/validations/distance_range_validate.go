package validations

import (
	"reflect"

	"github.com/go-playground/validator"
)

func validateDistanceRange(fl validator.FieldLevel) bool {
	// 構造体のフィールドを取得
	field := fl.Field()

	// 構造体の値を取得
	if field.Kind() != reflect.Struct {
		return false
	}

	// Min値とMax値を取得
	minField := field.FieldByName("Min")
	maxField := field.FieldByName("Max")

	if !minField.IsValid() || !maxField.IsValid() {
		return false
	}

	min := minField.Int()
	max := maxField.Int()

	// バリデーションルール
	if min < 0 || max > 100 || min > max {
		return false
	}

	return true
}
