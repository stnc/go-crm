package controller

import (
	"net/http"
	"stncCms/app/domain/helpers/stnccollection"
	"stncCms/app/domain/helpers/stncsession"
	"stncCms/app/services"

	"github.com/flosch/pongo2/v4"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

//Options constructor
type Options struct {
	OptionsApp services.OptionsAppInterface
}

const viewPathOptions = "admin/options/"

//InitOptions post controller constructor
func InitOptions(OptionsApp services.OptionsAppInterface) *Options {
	return &Options{
		OptionsApp: OptionsApp,
	}
}

//Index list
func (access *Options) Index(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	flashMsg := stncsession.GetFlashMessage(c)

	hisseAdeti := access.OptionsApp.GetOption("hisse_adeti")
	satisBirimFiyati1 := access.OptionsApp.GetOption("satis_birim_fiyati_1")
	satisBirimFiyati2 := access.OptionsApp.GetOption("satis_birim_fiyati_2")
	satisBirimFiyati3 := access.OptionsApp.GetOption("satis_birim_fiyati_3")
	alisBirimFiyati1 := access.OptionsApp.GetOption("alis_birim_fiyati_1")
	alisBirimFiyati2 := access.OptionsApp.GetOption("alis_birim_fiyati_2")
	alisBirimFiyati3 := access.OptionsApp.GetOption("alis_birim_fiyati_3")
	whatsAppMsg := access.OptionsApp.GetOption("whatsAppMsg")
	whatsAppMsgMap := access.OptionsApp.GetOption("whatsAppMsgMap")
	whatsAppMsgEk1 := access.OptionsApp.GetOption("whatsAppMsgEk1")
	otomatik_sira_buyukbas_2021 := access.OptionsApp.GetOption("otomatik_sira_buyukbas_2021")
	makbuz_otomatik_sira_no := access.OptionsApp.GetOption("makbuz_otomatik_sira_no")
	// dusukagirlik := access.OptionsApp.GetOption("hayvan_dusuk_agirligi")
	// ortaagirlik := access.OptionsApp.GetOption("hayvan_orta_agirligi")
	// yuksekagirlik := access.OptionsApp.GetOption("hayvan_yuksek_agirligi")
	kurbanYili := access.OptionsApp.GetOption("kurban_yili")

	viewData := pongo2.Context{
		"title":                       "Ayarlar",
		"csrf":                        csrf.GetToken(c),
		"hisse_adeti":                 hisseAdeti,
		"satis_birim_fiyati_1":        satisBirimFiyati1,
		"satis_birim_fiyati_2":        satisBirimFiyati2,
		"satis_birim_fiyati_3":        satisBirimFiyati3,
		"alis_birim_fiyati_1":         alisBirimFiyati1,
		"alis_birim_fiyati_2":         alisBirimFiyati2,
		"alis_birim_fiyati_3":         alisBirimFiyati3,
		"whatsAppMsg":                 whatsAppMsg,
		"whatsAppMsgEk1":              whatsAppMsgEk1,
		"otomatik_sira_buyukbas_2021": otomatik_sira_buyukbas_2021,
		"makbuz_otomatik_sira_no":     makbuz_otomatik_sira_no,
		// "hayvan_dusuk_agirligi":  dusukagirlik,
		// "hayvan_orta_agirligi":   ortaagirlik,
		// "hayvan_yuksek_agirligi": yuksekagirlik,
		"kurban_yili":    kurbanYili,
		"flashMsg":       flashMsg,
		"whatsAppMsgMap": whatsAppMsgMap,
	}

	c.HTML(
		http.StatusOK,
		viewPathOptions+"index.html",
		viewData,
	)
}

//Update list
func (access *Options) Update(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)

	access.OptionsApp.SetOption("hisse_adeti", c.PostForm("hisse_adeti"))
	access.OptionsApp.SetOption("satis_birim_fiyati_1", c.PostForm("satis_birim_fiyati_1"))
	access.OptionsApp.SetOption("satis_birim_fiyati_2", c.PostForm("satis_birim_fiyati_2"))
	access.OptionsApp.SetOption("satis_birim_fiyati_3", c.PostForm("satis_birim_fiyati_3"))
	access.OptionsApp.SetOption("alis_birim_fiyati_1", c.PostForm("alis_birim_fiyati_1"))
	access.OptionsApp.SetOption("alis_birim_fiyati_2", c.PostForm("alis_birim_fiyati_2"))
	access.OptionsApp.SetOption("alis_birim_fiyati_3", c.PostForm("alis_birim_fiyati_3"))
	access.OptionsApp.SetOption("whatsAppMsg", c.PostForm("whatsAppMsg"))
	access.OptionsApp.SetOption("whatsAppMsgMap", c.PostForm("whatsAppMsgMap"))
	access.OptionsApp.SetOption("whatsAppMsgEk1", c.PostForm("whatsAppMsgEk1"))
	// access.OptionsApp.SetOption("hayvan_dusuk_agirligi", c.PostForm("hayvan_dusuk_agirligi"))
	// access.OptionsApp.SetOption("hayvan_orta_agirligi", c.PostForm("hayvan_orta_agirligi"))
	// access.OptionsApp.SetOption("hayvan_yuksek_agirligi", c.PostForm("hayvan_yuksek_agirligi"))
	access.OptionsApp.SetOption("kurban_yili", c.PostForm("kurban_yili"))
	access.OptionsApp.SetOption("otomatik_sira_buyukbas_2021", c.PostForm("otomatik_sira_buyukbas_2021"))
	access.OptionsApp.SetOption("makbuz_otomatik_sira_no", c.PostForm("makbuz_otomatik_sira_no"))
	stncsession.SetFlashMessage("Kayıt başarı ile eklendi", "success", c)
	c.Redirect(http.StatusMovedPermanently, "/admin/options")
	return

}

func (access *Options) MakbuzNo(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)

	makbuzOtomatikSiraNo := access.OptionsApp.GetOption("makbuz_otomatik_sira_no")
	makbuzOtomatikSiraNoint := stnccollection.StringToint(makbuzOtomatikSiraNo) + 1
	makbuzOtomatikSiraNostr := stnccollection.IntToString(makbuzOtomatikSiraNoint)
	access.OptionsApp.SetOption("makbuz_otomatik_sira_no", makbuzOtomatikSiraNostr)
	viewData := pongo2.Context{
		"title":                "Otomatik Sıra No",
		"status":               "ok",
		"makbuzOtomatikSiraNo": makbuzOtomatikSiraNo,
		"errMsg":               "",
	}
	c.JSON(http.StatusOK, viewData)
	return

}
