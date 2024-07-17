package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"go_gateway/public"
	"gopkg.in/go-playground/validator.v9"
	enTranslations "gopkg.in/go-playground/validator.v9/translations/en"
	zhTranslations "gopkg.in/go-playground/validator.v9/translations/zh"
	"reflect"
)

// TranslationMiddleware 设置Translation
func TranslationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//参照：https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go

		//设置支持语言
		en := en.New()
		zh := zh.New()

		//设置国际化翻译器
		uni := ut.New(zh, zh, en)
		val := validator.New()

		//根据参数取翻译器实例
		locale := c.DefaultQuery("locale", "zh")
		trans, _ := uni.GetTranslator(locale)

		//翻译器注册到validator
		switch locale {
		case "en":
			err := enTranslations.RegisterDefaultTranslations(val, trans)
			if err != nil {
				println(err)
			}
			val.RegisterTagNameFunc(func(fld reflect.StructField) string {
				return fld.Tag.Get("en_comment")
			})
			break
		default:
			err := zhTranslations.RegisterDefaultTranslations(val, trans)
			if err != nil {
				println(err)
			}
			val.RegisterTagNameFunc(func(fld reflect.StructField) string {
				return fld.Tag.Get("comment")
			})

			//自定义验证方法
			//https://github.com/go-playground/validator/blob/v9/_examples/custom-validation/main.go
			err = val.RegisterValidation("is_valid_username", func(fl validator.FieldLevel) bool {
				return fl.Field().String() == "admin"
			})
			if err != nil {
				println(err)
			}

			//自定义翻译器
			//https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
			err = val.RegisterTranslation("is_valid_username", trans, func(ut ut.Translator) error {
				return ut.Add("is_valid_username", "{0} 填写不正确哦", true)
			}, func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("is_valid_username", fe.Field())
				return t
			})
			if err != nil {
				println(err)
			}
			break
		}
		c.Set(public.TranslatorKey, trans)
		c.Set(public.ValidatorKey, val)
		c.Next()
	}
}
