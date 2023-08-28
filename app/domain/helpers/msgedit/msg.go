package msgedit

import (
	"fmt"
	"net/url"
	"stncCms/app/domain/repository"
	"strings"
)

//oku
//https://faq.whatsapp.com/general/chats/how-to-use-click-to-chat/?lang=tr

type (
	// Inow : Custom Renderer for templates
	Msg struct{ Debug bool }
)

func (msgModul Msg) Wp(slug string) string {
	return SendMsgWp1(slug)
}

func (msgModul Msg) Wp2(slug string) string {
	return SendMsgWp2(slug)
}

func (msgModul Msg) Wp3(slug string) string {
	return SendMsgWp3(slug)
}

func telReplace(tel string) string {
	return strings.Replace(tel, " ", "", -1)
}

func msgReplace(msg string, slug string, url string) string {
	return strings.Replace(msg, "[link]", url+"kurbanBilgi/"+slug+"/", -1)
}

func SendMsgWp1(slug string) string {
	db := repository.DB
	access := repository.KurbanRepositoryInit(db)
	// kurbanData, _ := access.GetKurbanOpenInfo(slug)

	if kurbanData, err := access.GetKurbanOpenInfo(slug); err == nil {

		tel := telReplace(kurbanData.Telefon)

		appOptions := repository.OptionRepositoryInit(db)
		msgg := msgReplace(appOptions.GetOption("whatsAppMsg"), slug, appOptions.GetOption("siteUrl"))
		fmt.Println(appOptions.GetOption("siteUrl"))
		return "https://wa.me/09" + tel + "?text=" + url.QueryEscape(msgg)
		// return "https://wa.me/09" + tel + "?text=" + url.QueryEscape(appOptions.GetOption("whatsAppMsg")) + " " + "https://kysmerkez.kurbandefteri.com/login/"
	} else {
		return ""
	}
}

func SendMsgWp2(slug string) string {
	db := repository.DB
	access := repository.KurbanRepositoryInit(db)
	// kurbanData, _ := access.GetKurbanOpenInfo(slug)

	if kurbanData, err := access.GetKurbanOpenInfo(slug); err == nil {

		tel := telReplace(kurbanData.Telefon)

		appOptions := repository.OptionRepositoryInit(db)
		msgg := appOptions.GetOption("whatsAppMsgMap")
		return "https://wa.me/09" + tel + "?text=" + url.QueryEscape(msgg)
	} else {
		return ""
	}
}

func SendMsgWp3(slug string) string {
	db := repository.DB
	access := repository.KurbanRepositoryInit(db)
	// kurbanData, _ := access.GetKurbanOpenInfo(slug)

	if kurbanData, err := access.GetKurbanOpenInfo(slug); err == nil {

		tel := telReplace(kurbanData.Telefon)

		appOptions := repository.OptionRepositoryInit(db)
		msgg := appOptions.GetOption("whatsAppMsgEk1")
		return "https://wa.me/09" + tel + "?text=" + url.QueryEscape(msgg)
	} else {
		return ""
	}
}
