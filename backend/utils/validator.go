package utils

import (
	zhLocale "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zhTrans "gopkg.in/go-playground/validator.v9/translations/zh"
	"reflect"
	"sync"
)

type ValidatorV9 struct {
	once     sync.Once
	validate *validator.Validate
}

func (v *ValidatorV9) ValidateStruct(obj interface{}) error {
	if kindOfData(obj) == reflect.Struct {
		v.lazyinit()
		if err := v.validate.Struct(obj); err != nil {
			return err
		}
	}
	return nil
}

func (v *ValidatorV9) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *ValidatorV9) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("binding")
		// add any custom validations etc. here
		zh := zhLocale.New()
		uni := ut.New(zh, zh)
		trans, _ := uni.GetTranslator("zh")
		_ = zhTrans.RegisterDefaultTranslations(v.validate, trans)
	})
}

func kindOfData(data interface{}) reflect.Kind {
	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}
