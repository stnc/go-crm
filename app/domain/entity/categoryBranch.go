package entity

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	tr_translations "gopkg.in/go-playground/validator.v9/translations/tr"
)

//CategoriesBranch strcut
type CategoriesBranch struct {
	ID         uint64 `gorm:"primary_key;auto_increment" json:"id"`
	SaveUserID uint64 `gorm:"not null;" json:"saveSserId"`
	BranchID   uint64 `gorm:"not null;" json:"branchId"`
	UserID     uint64 `gorm:"not null;" json:"userId"`

	CreatedAt time.Time  ` json:"created_at"`
	UpdatedAt time.Time  ` json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

//TableName override
func (gk *CategoriesBranch) TableName() string {
	return "categories_branchs"
}

//Validate fluent validation
func (f *CategoriesBranch) Validate() map[string]string {
	var (
		validate *validator.Validate
		uni      *ut.UniversalTranslator
	)
	tr := en.New()
	uni = ut.New(tr, tr)
	trans, _ := uni.GetTranslator("tr")
	validate = validator.New()
	tr_translations.RegisterDefaultTranslations(validate, trans)

	errorLog := make(map[string]string)

	err := validate.Struct(f)
	fmt.Println(err)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		fmt.Println(errs)
		for _, e := range errs {
			// can translate each error one at a time.
			lng := strings.Replace(e.Translate(trans), e.Field(), "Burası", 1)
			errorLog[e.Field()+"_error"] = e.Translate(trans)
			// errorLog[e.Field()] = e.Translate(trans)
			errorLog[e.Field()] = lng
			errorLog[e.Field()+"_valid"] = "is-invalid"
		}
	}
	return errorLog
}
