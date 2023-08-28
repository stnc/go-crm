package services

import (
	"stncCms/app/domain/dto"
	"stncCms/app/domain/entity"
)

//GruplarAppInterface interface
type GruplarAppInterface interface {
	Save(*entity.SacrificeGruplar) (*entity.SacrificeGruplar, map[string]string)
	GetByID(uint64) (*entity.SacrificeGruplar, error)
	GetByIDAllRelations(uint64) (*entity.SacrificeGruplar, error)
	GetByIDAllRelationsHayvanOlmayanlar(uint64) (*entity.SacrificeGruplar, error)
	GetByAllRelations() ([]dto.GruplarExcelandIndex, error)
	GetAll() ([]entity.SacrificeGruplar, error)
	GetAllFindDurum(int) ([]entity.SacrificeGruplar, error)
	GetAllFindDurumAndAgirlikTipi(int, int) ([]entity.SacrificeGruplar, error)
	GrupToplamOdemeler(uint64) float64
	GrupKalanBorclar(uint64) float64
	ToplamOdemeler() float64
	KalanBorclar() float64
	KasaBorcu(uint64) float64
	GetAllP(int, int) ([]entity.SacrificeGruplar, error)
	Update(*entity.SacrificeGruplar) (*entity.SacrificeGruplar, map[string]string)
	Delete(uint64) error
	KurbanFiyati(uint64) float64
	SatisFiyatTuru(uint64) int
	Count(*int64)
	SetGrupID(GrupID uint64, kurbanID uint64, grupLideri int)
}

type gruplarApp struct {
	request GruplarAppInterface
}

var _ GruplarAppInterface = &gruplarApp{}

func (f *gruplarApp) Save(Kurban *entity.SacrificeGruplar) (*entity.SacrificeGruplar, map[string]string) {
	return f.request.Save(Kurban)
}

func (f *gruplarApp) GetAll() ([]entity.SacrificeGruplar, error) {
	return f.request.GetAll()
}

func (f *gruplarApp) GetAllP(postsPerPage int, offset int) ([]entity.SacrificeGruplar, error) {
	return f.request.GetAllP(postsPerPage, offset)
}

func (f *gruplarApp) GetByID(kurbanID uint64) (*entity.SacrificeGruplar, error) {
	return f.request.GetByID(kurbanID)
}

func (f *gruplarApp) GetByIDAllRelations(kurbanID uint64) (*entity.SacrificeGruplar, error) {
	return f.request.GetByIDAllRelations(kurbanID)
}

func (f *gruplarApp) GetByIDAllRelationsHayvanOlmayanlar(kurbanID uint64) (*entity.SacrificeGruplar, error) {
	return f.request.GetByIDAllRelationsHayvanOlmayanlar(kurbanID)
}

func (f *gruplarApp) GetAllFindDurum(durum int) ([]entity.SacrificeGruplar, error) {
	return f.request.GetAllFindDurum(durum)
}

func (f *gruplarApp) GetAllFindDurumAndAgirlikTipi(status int, agirilikTipi int) ([]entity.SacrificeGruplar, error) {
	return f.request.GetAllFindDurumAndAgirlikTipi(status, agirilikTipi)
}
func (f *gruplarApp) GetByAllRelations() ([]dto.GruplarExcelandIndex, error) {
	return f.request.GetByAllRelations()
}

func (f *gruplarApp) Update(Kurban *entity.SacrificeGruplar) (*entity.SacrificeGruplar, map[string]string) {
	return f.request.Update(Kurban)
}

func (f *gruplarApp) Delete(kurbanID uint64) error {
	return f.request.Delete(kurbanID)
}
func (f *gruplarApp) KurbanFiyati(kurbanID uint64) float64 {
	return f.request.KurbanFiyati(kurbanID)
}

func (f *gruplarApp) GrupToplamOdemeler(GrupID uint64) float64 {
	return f.request.GrupToplamOdemeler(GrupID)
}

func (f *gruplarApp) GrupKalanBorclar(GrupID uint64) float64 {
	return f.request.GrupKalanBorclar(GrupID)
}

func (f *gruplarApp) ToplamOdemeler() float64 {
	return f.request.ToplamOdemeler()
}

func (f *gruplarApp) KalanBorclar() float64 {
	return f.request.KalanBorclar()
}

func (f *gruplarApp) KasaBorcu(GrupID uint64) float64 {
	return f.request.KasaBorcu(GrupID)
}

func (f *gruplarApp) SatisFiyatTuru(kurbanID uint64) int {
	return f.request.SatisFiyatTuru(kurbanID)
}
func (f *gruplarApp) Count(postTotalCount *int64) {
	f.request.Count(postTotalCount)
}

func (f *gruplarApp) SetGrupID(grupID uint64, kurbanID uint64, grupLideri int) {
	f.request.SetGrupID(grupID, kurbanID, grupLideri)
}
