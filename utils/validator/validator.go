package validator

import (
	"github.com/go-playground/locales/zh_Hans_CN"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"log"
	"query/utils/errmsg"
	"reflect"
)

func Validate(data interface{}) (string,int) {
	validate := validator.New()
	uT := unTrans.New(zh_Hans_CN.New())
	translator,_ := uT.GetTranslator("zh_Hans_CN")

	err := zhTrans.RegisterDefaultTranslations(validate,translator)
	if err != nil {
		log.Println("err :",err)
	}

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})
	err = validate.Struct(data)
	for err != nil {
		for _,value := range err.(validator.ValidationErrors) {
			return value.Translate(translator),errmsg.ERROR
		}
	}
	return "",errmsg.SUCCESS
}