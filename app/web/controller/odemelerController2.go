package controller

import (
	"errors"
	"fmt"
	"net/http"
	"stncCms/app/domain/entity"
	"stncCms/app/domain/helpers/stncsession"
	"stncCms/app/domain/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (access *Odemeler) OdemeSil(c *gin.Context) {

	stncsession.IsLoggedInRedirect(c)

	if odeme, _, errorR := odemeSilModel(c); errorR == nil {

		access.OdemelerApp.Save(&odeme)

		stncsession.SetFlashMessage("Kayıt başarı ile silindi", "success", c)
		// return
		c.Redirect(http.StatusMovedPermanently, "/admin/kurban/edit/"+c.Query("kid")+"?action=odemelerList")
		return
	}
}

func odemeSilModel(c *gin.Context) (odemelerData entity.SacrificeOdemeler, idStr string, err error) {
	id, _ := strconv.ParseUint(c.Param("odemeID"), 10, 64)
	idStr = c.Query("kid")
	// odemelerData.ID = id

	var kalanBorcu, yeniKalanBorcu, yeniToplamOdemeler, odemelerToplami float64

	var kurbanID uint64
	kurbanID, _ = strconv.ParseUint(c.Query("kid"), 10, 64)

	db := repository.DB
	appKurban := repository.KurbanRepositoryInit(db)
	appOdeme := repository.OdemelerRepositoryInit(db)

	var silinecekUcret float64

	appOdeme.GetByHirePurchase(id, &silinecekUcret)

	kurbanDurum := appKurban.GetKurbanDurum(kurbanID)
	odemelerToplami = appKurban.OdemelerToplami(kurbanID)
	kurbanFiyati := appKurban.KurbanFiyati(kurbanID)
	kasaBorcu := appKurban.KasaBorcu(kurbanID)

	kalanBorcu = appKurban.KalanUcret(kurbanID) //TODO: hayvasız gruplarda bunu test et sorun olcak mı

	yeniKalanBorcu = kalanBorcu + silinecekUcret

	yeniToplamOdemeler = odemelerToplami - silinecekUcret

	// fmt.Println("odemelerToplami")
	// fmt.Println(odemelerToplami)

	// fmt.Println("silinecekUcret")
	// fmt.Println(silinecekUcret)

	// fmt.Println("kalanBorcu")
	// fmt.Println(kalanBorcu)

	// fmt.Println("yeniKalanBorcu")
	// fmt.Println(yeniKalanBorcu)

	// fmt.Println("yeniToplamOdemeler")
	// fmt.Println(yeniToplamOdemeler)
	// return
	//Grup Olusmus Hayvan Yok ise
	if kurbanDurum == entity.KurbanDurumGrupOlusmusHayvanYok {
		odemelerData.Durum = entity.OdemelerDurumGrupOlusmusHayvanYok
	}

	//TODO: bu ne ıse yaracak????
	if silinecekUcret > odemelerToplami {
		fmt.Println("buyukkkk2")
		return odemelerData, c.Query("kid"), errors.New("Ücret Büyük")
	}

	//TODO: odeme silince odenen tum mıktar ıle kurban fıyatı aynı ıse odeme bıttı flag eklenmesı gerekıyor

	if kurbanDurum == entity.KurbanDurumGrupOlusmusKesimlikHayvaniVar { //	//Grup Olusmus Kesimlik Hayvani Var
		odemelerData.Durum = entity.KurbanDurumGrupOlusmusKesimlikHayvaniVar
		//TODO: burası calışır mı aynı sekılde kalanucret ==0 ıcınde bu duruma bakıalcak
		//	if yeniVerilenUcret == 0 {
		// 	fmt.Println("buyukkkk")
		// 	return odemelerData, c.Query("kid"), errors.New("Ücret Büyük")
		//}

		//TODO: bak bu calısr mı mantıklı mı //TODO: kapattım
		// if yeniVerilenUcret == odemelerToplami {
		// }
	}

	// kurban eklenmiş ama kurban bayramına ait değil yani direk kurban girişi
	if kurbanDurum == entity.KurbanDurumKurbanEklendiKurbanBayraminaAitDegil {
		odemelerData.Durum = entity.KurbanDurumKurbanEklendiKurbanBayraminaAitDegil
	}

	// if aciklama != "" {
	// 	odemelerData.Aciklama = aciklama
	// }

	odemelerData.Alacak = yeniKalanBorcu
	odemelerData.KurbanFiyati = kurbanFiyati
	odemelerData.Borc = kasaBorcu
	odemelerData.Bakiye = yeniKalanBorcu
	odemelerData.VerilenUcret = silinecekUcret
	odemelerData.Aciklama = c.Param("odemeID") + " Numaralı taksit silindi "
	odemelerData.BorcDurum = entity.OdemelerBorcDurumTaksitSilindiBildirim
	appKurban.SetKurbanBakiyeUpdate(kurbanID, yeniToplamOdemeler)
	appKurban.SetKurbanKalanUcretiUpdate(kurbanID, yeniKalanBorcu) //alacak tablo adı
	appKurban.SetKurbanBorcDurumUpdate(kurbanID, entity.KurbanBorcDurumTaksitOdemesi)
	appOdeme.SetOdemelerBorcDurum(id, entity.OdemelerBorcDurumTaksitSilindi)
	//orginal yerleri
	odemelerData.KurbanID = kurbanID
	odemelerData.Makbuz = c.PostForm("Makbuz")
	odemelerData.UserID = stncsession.GetUserID2(c)
	// odemelerData.VerildigiTarih = time.Now()
	odemelerData.KurbanFiyati = kurbanFiyati
	return odemelerData, c.Query("kid"), nil
}
