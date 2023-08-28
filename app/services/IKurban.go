package services

import (
	"stncCms/app/domain/dto"
	"stncCms/app/domain/entity"

	"github.com/gin-gonic/gin"
)

//KurbanAppInterface interface
type KurbanAppInterface interface {
	Save(*entity.SacrificeKurbanlar) (*entity.SacrificeKurbanlar, map[string]string)
	GetByID(uint64) (*entity.SacrificeKurbanlar, error)
	GetAllP(int, int) ([]entity.SacrificeKurbanlar, error)
	ListDataTable(c *gin.Context) (dto.KurbanBilgisiDataTableResult, error)
	GetAll() ([]entity.SacrificeKurbanlar, error)
	GetByGrupID(uint64) ([]entity.SacrificeKurbanlar, error)
	GetByGrupIDCount(uint64, *int)
	Update(*entity.SacrificeKurbanlar) (*entity.SacrificeKurbanlar, map[string]string)
	GetKurbanOpenInfo(slug string) (*dto.KurbanOpenInfoData, error)
	Delete(uint64) error
	Count(*int64)
	OdemelerToplami(uint64) float64
	KalanUcret(uint64) float64
	KurbanFiyati(uint64) float64
	KasaBorcu(uint64) float64
	GetKurbanDurum(kurbanID uint64) (durum int)
	GetKurbanBorcDurum(kurbanID uint64) (durum int)
	GetKurbanGrupID(kurbanID uint64) (GrupID int)
	GetKurbanTuru(kurbanID uint64) (kurbanTuru int)
	GetKurbanHayvanID(grupID uint64) (hayvanID int)
	GetKurbanKisiVarmi(uint64, *int)
	GetGrupLideri(kurbanID uint64) int
	SetKurbanKalanUcretiUpdate(id uint64, price float64)
	SetKurbanBakiyeUpdate(id uint64, price float64)
	SetKurbanFiyatiUpdate(id uint64, price float64)
	SetKurbanAgirligi(id uint64, agirlik int)
	SetKurbanKasaBorcuUpdate(id uint64, price float64)
	SetGrupLideri(kurbanID uint64, value int)
	SetVekaletDurumu(kurbanID uint64, value int)

	SetKurbanBorcDurumUpdate(id uint64, status int)
	SetKurbanDurumUpdate(id uint64, status int)
	UpdatePriceList(post *dto.KurbanUpdateRead) (*dto.KurbanUpdateRead, map[string]string)
	GetAllKurbanAndKisiler(grupId int) ([]dto.KurbanListForGrouplar, error)

	LastID(id uint64, lastid *uint64)
	NextID(id uint64, nextid *uint64)

	GruopLastID(id uint64, gruopID uint64, gnextid *uint64)
	GruopNextID(id uint64, gruopID uint64, glastid *uint64)
}

type GenelKurbanApp struct {
	request KurbanAppInterface
}

var _ KurbanAppInterface = &GenelKurbanApp{}

func (f *GenelKurbanApp) Save(gkurban *entity.SacrificeKurbanlar) (*entity.SacrificeKurbanlar, map[string]string) {
	return f.request.Save(gkurban)
}

func (f *GenelKurbanApp) GetAll() ([]entity.SacrificeKurbanlar, error) {
	return f.request.GetAll()
}

func (f *GenelKurbanApp) GetAllKurbanAndKisiler(grupId int) ([]dto.KurbanListForGrouplar, error) {
	return f.request.GetAllKurbanAndKisiler(grupId)
}

func (f *GenelKurbanApp) GetByGrupID(grupID uint64) ([]entity.SacrificeKurbanlar, error) {
	return f.request.GetByGrupID(grupID)
}

func (f *GenelKurbanApp) GetByID(gkurbanID uint64) (*entity.SacrificeKurbanlar, error) {
	return f.request.GetByID(gkurbanID)
}

func (f *GenelKurbanApp) GetAllP(postsPerPage int, offset int) ([]entity.SacrificeKurbanlar, error) {
	return f.request.GetAllP(postsPerPage, offset)
}

func (f *GenelKurbanApp) Update(gkurban *entity.SacrificeKurbanlar) (*entity.SacrificeKurbanlar, map[string]string) {
	return f.request.Update(gkurban)
}

//UpdatePriceList fiyat günceller
func (f *GenelKurbanApp) UpdatePriceList(post *dto.KurbanUpdateRead) (*dto.KurbanUpdateRead, map[string]string) {
	return f.request.UpdatePriceList(post)
}

func (f *GenelKurbanApp) GetKurbanOpenInfo(slug string) (*dto.KurbanOpenInfoData, error) {
	return f.request.GetKurbanOpenInfo(slug)
}

func (f *GenelKurbanApp) Delete(gkurbanID uint64) error {
	return f.request.Delete(gkurbanID)
}

func (f *GenelKurbanApp) Count(postTotalCount *int64) {
	f.request.Count(postTotalCount)
}

func (f *GenelKurbanApp) GetByGrupIDCount(grupID uint64, postTotalCount *int) {
	f.request.GetByGrupIDCount(grupID, postTotalCount)
}

func (f *GenelKurbanApp) GetKurbanHayvanID(grupID uint64) (hayvanID int) {
	return f.request.GetKurbanHayvanID(grupID)
}

func (f *GenelKurbanApp) GetKurbanGrupID(grupID uint64) (GrupID int) {
	return f.request.GetKurbanGrupID(grupID)
}

func (f *GenelKurbanApp) OdemelerToplami(kurbanID uint64) float64 {
	return f.request.OdemelerToplami(kurbanID)
}

func (f *GenelKurbanApp) KalanUcret(kurbanID uint64) float64 {
	return f.request.KalanUcret(kurbanID)
}
func (f *GenelKurbanApp) KurbanFiyati(kurbanID uint64) float64 {
	return f.request.KurbanFiyati(kurbanID)
}

func (f *GenelKurbanApp) KasaBorcu(kurbanID uint64) float64 {
	return f.request.KasaBorcu(kurbanID)
}

func (f *GenelKurbanApp) GetKurbanDurum(kurbanID uint64) int {
	return f.request.GetKurbanDurum(kurbanID)
}

func (f *GenelKurbanApp) GetKurbanBorcDurum(kurbanID uint64) int {
	return f.request.GetKurbanBorcDurum(kurbanID)
}

func (f *GenelKurbanApp) GetKurbanTuru(kurbanID uint64) int {
	return f.request.GetKurbanDurum(kurbanID)
}
func (f *GenelKurbanApp) GetGrupLideri(kurbanID uint64) int {
	return f.request.GetGrupLideri(kurbanID)
}

func (f *GenelKurbanApp) GetKurbanKisiVarmi(kurbanID uint64, kisiID *int) {
	f.request.GetKurbanKisiVarmi(kurbanID, kisiID)
}

func (f *GenelKurbanApp) SetKurbanKalanUcretiUpdate(id uint64, price float64) {
	f.request.SetKurbanKalanUcretiUpdate(id, price)
}

func (f *GenelKurbanApp) SetGrupLideri(kurbanID uint64, value int) {
	f.request.SetGrupLideri(kurbanID, value)
}

func (f *GenelKurbanApp) SetVekaletDurumu(kurbanID uint64, value int) {
	f.request.SetVekaletDurumu(kurbanID, value)
}

func (f *GenelKurbanApp) SetKurbanKasaBorcuUpdate(id uint64, price float64) {
	f.request.SetKurbanKasaBorcuUpdate(id, price)
}

//SetKurbanBakiyeUpdate odeme gncelleme
func (f *GenelKurbanApp) SetKurbanBakiyeUpdate(id uint64, price float64) {
	f.request.SetKurbanBakiyeUpdate(id, price)
}

//SetKurbanFiyatiUpdate
func (f *GenelKurbanApp) SetKurbanFiyatiUpdate(id uint64, price float64) {
	f.request.SetKurbanFiyatiUpdate(id, price)
}
func (f *GenelKurbanApp) SetKurbanBorcDurumUpdate(id uint64, status int) {
	f.request.SetKurbanBorcDurumUpdate(id, status)
}

func (f *GenelKurbanApp) SetKurbanDurumUpdate(id uint64, status int) {
	f.request.SetKurbanDurumUpdate(id, status)
}
func (f *GenelKurbanApp) SetKurbanAgirligi(id uint64, agirlik int) {
	f.request.SetKurbanAgirligi(id, agirlik)
}

func (f *GenelKurbanApp) ListDataTable(c *gin.Context) (dto.KurbanBilgisiDataTableResult, error) {
	return f.request.ListDataTable(c)
}

func (f *GenelKurbanApp) LastID(id uint64, lastid *uint64) {
	f.request.LastID(id, lastid)
}

func (f *GenelKurbanApp) NextID(id uint64, nextid *uint64) {
	f.request.NextID(id, nextid)
}

func (f *GenelKurbanApp) GruopLastID(id uint64, gruopID uint64, glastid *uint64) {
	f.request.GruopLastID(id, gruopID, glastid)
}

func (f *GenelKurbanApp) GruopNextID(id uint64, gruopID uint64, gnextid *uint64) {
	f.request.GruopNextID(id, gruopID, gnextid)
}
