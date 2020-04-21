package validator

import (
	"log"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	en_translations "github.com/go-playground/validator/v10/translations/en"

	"github.com/go-playground/validator/v10"
)

// Validator implements json validation methods
type Validator struct {
	validator  *validator.Validate
	translator *ut.Translator
}

// Validate - Validate json body using struct
func (v *Validator) Validate(s interface{}) error {
	return v.validator.Struct(s)
}

// New Validator and Translator with json fields and default en translation
func New() *Validator {
	en := en.New()
	uni := ut.New(en, en)

	trans, _ := uni.GetTranslator("en")

	validate := validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)
	validate.RegisterTranslation("url|fqdn", trans, registrationFunc("url|fqdn", "{0} must be a valid domain", false), translateFunc)

	// Using the names which have been specified for JSON representations of structs, rather than normal Go field names
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return &Validator{
		validator:  validate,
		translator: &trans,
	}
}

// TranslateValidatorErr to english readable messages
func (v *Validator) TranslateValidatorErr(err error) string {
	errs := err.(validator.ValidationErrors)
	translatedErrs := errs.Translate(*v.translator)

	stringErrs := []string{}
	for _, err := range translatedErrs {
		stringErrs = append(stringErrs, err)
	}
	return strings.Join(stringErrs, ", ")
}

func registrationFunc(tag string, translation string, override bool) validator.RegisterTranslationsFunc {
	return func(ut ut.Translator) (err error) {
		if err = ut.Add(tag, translation, override); err != nil {
			return
		}
		return
	}
}

func translateFunc(ut ut.Translator, fe validator.FieldError) string {
	t, err := ut.T(fe.Tag(), fe.Field())
	if err != nil {
		log.Printf("warning: error translating FieldError: %#v", fe)
		return fe.(error).Error()
	}
	return t
}
