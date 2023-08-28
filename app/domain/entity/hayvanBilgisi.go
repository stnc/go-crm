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

const (
	HayvanBilgisiDurumHayvanBosta   = 1
	HayvanBilgisiDurumHayvanAtanmis = 2
)

//hayvan kartı olacak - ahırı - hayvan kilosu falan olacak hayvanın alış fiyatı satış fiyatı et kilos gibi hesaplamalar olacak

//TODO: buranın sabitlerinde hayvan öldü gibi bir flag olmalı durum için

//HayvanBilgisiTableName global table name
var HayvanBilgisiTableName string = "sacrifice_hayvan_bilgisi"

//HayvanBilgisi yapı
type SacrificeHayvanBilgisi struct {
	Order                uint64                      `gorm:"-" json:"order"`
	ID                   uint64                      `gorm:"primary_key;auto_increment" json:"id"`
	UserID               uint64                      `gorm:"not null;" json:"userId"`
	NotesID              uint64                      `gorm:"null;" json:"notesId"`
	HayvanCinsi          string                      `gorm:"type:varchar(255); null;" validate:"required" json:"hayvanCinsi"`
	Agirlik              int                         `gorm:"type:smallint ; NULL;DEFAULT:'1'"  validate:"required,numeric,min=2,max=1300" json:"agirlik"`
	KupeNo               uint64                      `gorm:"NULL;" validate:"omitempty,numeric,min=1" json:"kupeNo" `
	Resim                string                      `gorm:"type:text; ;null;" json:"resim"`
	Durum                int                         `gorm:"type:smallint ;NOT NULL;DEFAULT:'1'"  validate:"required" json:"durum" `
	AlisFiyatTuru        int                         `gorm:"type:smallint ;NOT NULL;DEFAULT:'1'"  json:"alis_fiyat_turu" `
	AlisFiyati           float64                     `gorm:"type:decimal(10,2); NOT NULL; DEFAULT:'0';" validate:"required,numeric" `
	CreatedAt            time.Time                   ` json:"created_at"`
	UpdatedAt            time.Time                   ` json:"updated_at"`
	DeletedAt            *time.Time                  `json:"deleted_at"`
	HayvanSatisYerleri   SacrificeHayvanSatisYerleri `json:"hayvanSatisYerleri"`
	HayvanSatisYerleriID uint64                      `gorm:"ForeignKey:id;" json:"hayvanSatisYerleriID"`
}

type HayvanBilgisiDataResult struct {
	Total    int                      `json:"recordsTotal"`
	Filtered int                      `json:"recordsFiltered"`
	Data     []SacrificeHayvanBilgisi `json:"data"`
}

//BeforeSave init
func (f *SacrificeHayvanBilgisi) BeforeSave() {
	f.HayvanCinsi = html.EscapeString(strings.TrimSpace(f.HayvanCinsi))

}

//TableName override
func (f *SacrificeHayvanBilgisi) TableName() string {
	return HayvanBilgisiTableName
}

/*
func (post *Post) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
 }
*/

//Prepare init
func (f *SacrificeHayvanBilgisi) Prepare() {
	f.HayvanCinsi = html.EscapeString(strings.TrimSpace(f.HayvanCinsi))

	f.CreatedAt = time.Now()
	f.UpdatedAt = time.Now()
}

//Validate fluent validation
func (f *SacrificeHayvanBilgisi) Validate() map[string]string {
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
