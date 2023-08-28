package entity

import (
	"fmt"
	"html"
	"strings"
	"time"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	tr_translations "gopkg.in/go-playground/validator.v9/translations/tr"
)

//HayvanSatisYerleriTableName global table name
var HayvanSatisYerleriTableName string = "sacrifice_hayvan_satis_yerleri"

//HayvanSatisYerleri strcut
type SacrificeHayvanSatisYerleri struct {
	ID         uint64 `gorm:"primary_key;auto_increment" json:"id"`
	UserID     uint64 `gorm:"not null;" json:"userId"`
	NotesID    uint64 `gorm:"not null;" json:"notesId"`
	YerAdi     string `gorm:"type:varchar(255); ;null;"  validate:"required" json:"yerAdi"`
	Slug       string `gorm:"type:varchar(255); ;null;" json:"slug"`
	Adresi     string `gorm:"type:text ;" validate:"required" json:"adresi"`
	IlgiliKisi string `gorm:"type:text ;" json:"ilgiliKisi"`
	// Telefon    string     `gorm:"size:255 ;null;" json:"telefon"`//#longText
	Telefon   string     `gorm:"type:varchar(255) ;null;" json:"telefon"`
	Durum     int        `gorm:"type:smallint ;NOT NULL;DEFAULT:'0'" validate:"required" json:"durum"`
	CreatedAt time.Time  ` json:"created_at"`
	UpdatedAt time.Time  ` json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

//BeforeSave init
func (f *SacrificeHayvanSatisYerleri) BeforeSave() {
	f.YerAdi = html.EscapeString(strings.TrimSpace(f.YerAdi))
	f.Adresi = html.EscapeString(strings.TrimSpace(f.Adresi))
	f.IlgiliKisi = html.EscapeString(strings.TrimSpace(f.IlgiliKisi))
	f.Telefon = html.EscapeString(strings.TrimSpace(f.Telefon))
}

//TableName override
func (f *SacrificeHayvanSatisYerleri) TableName() string {
	return HayvanSatisYerleriTableName
}

/*
func (post *Post) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
 }
*/

//Prepare init
func (f *SacrificeHayvanSatisYerleri) Prepare() {
	f.YerAdi = html.EscapeString(strings.TrimSpace(f.YerAdi))
	f.Adresi = html.EscapeString(strings.TrimSpace(f.Adresi))
	f.IlgiliKisi = html.EscapeString(strings.TrimSpace(f.IlgiliKisi))
	f.Telefon = html.EscapeString(strings.TrimSpace(f.Telefon))
	f.CreatedAt = time.Now()
	f.UpdatedAt = time.Now()
}

//Validate fluent validation
func (f *SacrificeHayvanSatisYerleri) Validate() map[string]string {
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
