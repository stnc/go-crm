package services

import (
	"stncCms/app/domain/dto"
	"stncCms/app/domain/entity"
)

//OdemelerAppInterface interface
type OdemelerAppInterface interface {
	Save(*entity.SacrificeOdemeler) (*entity.SacrificeOdemeler, map[string]string)
	SaveBatch([]entity.SacrificeOdemeler) ([]entity.SacrificeOdemeler, map[string]string)
	GetByID(uint64) (*entity.SacrificeOdemeler, error)
	GetByHirePurchase(kurbanID uint64, hirePurchase *float64)
	GetAll() ([]entity.SacrificeOdemeler, error)
	Update(*entity.SacrificeOdemeler) (*entity.SacrificeOdemeler, map[string]string)
	SetUpdateHesaplar(*dto.SacrificeOdemeler) (*dto.SacrificeOdemeler, map[string]string)
	Delete(uint64) error
	SonOdemeTablosu(uint64) (*dto.OdemelerSonFiyat, error)
	OdemelerToplami(kurbanID uint64) float64
	KurbanSonKalanUcret(kurbanID uint64) float64
	KasaBorcu(kurbanID uint64) float64
	GetOdemeRelation(odemeID uint64) (*dto.OdemeMakbuzu, error)
	SetOdemelerBorcDurum(id uint64, status int)
}
type odemelerApp struct {
	request OdemelerAppInterface
}

var _ OdemelerAppInterface = &odemelerApp{}

func (f *odemelerApp) SetOdemelerBorcDurum(id uint64, status int) {
	f.request.SetOdemelerBorcDurum(id, status)
}

func (f *odemelerApp) Save(odemeler *entity.SacrificeOdemeler) (*entity.SacrificeOdemeler, map[string]string) {
	return f.request.Save(odemeler)
}

func (f *odemelerApp) SaveBatch(odemeler []entity.SacrificeOdemeler) ([]entity.SacrificeOdemeler, map[string]string) {
	return f.request.SaveBatch(odemeler)
}

func (f *odemelerApp) GetAll() ([]entity.SacrificeOdemeler, error) {
	return f.request.GetAll()
}

func (f *odemelerApp) GetByID(odemelerID uint64) (*entity.SacrificeOdemeler, error) {
	return f.request.GetByID(odemelerID)
}

func (f *odemelerApp) Update(odemeler *entity.SacrificeOdemeler) (*entity.SacrificeOdemeler, map[string]string) {
	return f.request.Update(odemeler)
}

func (f *odemelerApp) SetUpdateHesaplar(odemeler *dto.SacrificeOdemeler) (*dto.SacrificeOdemeler, map[string]string) {
	return f.request.SetUpdateHesaplar(odemeler)
}

func (f *odemelerApp) Delete(odemelerID uint64) error {
	return f.request.Delete(odemelerID)
}

func (f *odemelerApp) SonOdemeTablosu(kurbanID uint64) (*dto.OdemelerSonFiyat, error) {
	return f.request.SonOdemeTablosu(kurbanID)
}

func (f *odemelerApp) OdemelerToplami(kurbanID uint64) float64 {
	return f.request.OdemelerToplami(kurbanID)
}

func (f *odemelerApp) KurbanSonKalanUcret(kurbanID uint64) float64 {
	return f.request.KurbanSonKalanUcret(kurbanID)
}

func (f *odemelerApp) KasaBorcu(kurbanID uint64) float64 {
	return f.request.KasaBorcu(kurbanID)
}

func (f *odemelerApp) GetOdemeRelation(odemeID uint64) (*dto.OdemeMakbuzu, error) {
	return f.request.GetOdemeRelation(odemeID)
}

func (f *odemelerApp) GetByHirePurchase(kurbanID uint64, hirePurchase *float64) {
	f.request.GetByHirePurchase(kurbanID, hirePurchase)
}
