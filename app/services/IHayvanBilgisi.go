package services

import (
	"stncCms/app/domain/dto"
	"stncCms/app/domain/entity"

	"github.com/gin-gonic/gin"
)

//HayvanBilgisiAppInterface service
type HayvanBilgisiAppInterface interface {
	Save(*entity.SacrificeHayvanBilgisi) (*entity.SacrificeHayvanBilgisi, map[string]string)
	GetByID(uint64) (*entity.SacrificeHayvanBilgisi, error)
	GetByIDRelated(uint64) (*dto.HayvanBilgisi, error)

	GetAll() ([]entity.SacrificeHayvanBilgisi, error)
	GetAllFindDurum(int) ([]entity.SacrificeHayvanBilgisi, error)
	ListDataTable(c *gin.Context) (entity.HayvanBilgisiDataResult, error)
	GetAllP(int, int) ([]entity.SacrificeHayvanBilgisi, error)
	Update(*entity.SacrificeHayvanBilgisi) (*entity.SacrificeHayvanBilgisi, map[string]string)
	Count(*int64)
	KupeNoCount(int, *int)
	Delete(uint64) error
	UpdateSingleStatus(uint64, int)
	LastID(id uint64, lastid *uint64)
	NextID(id uint64, nextid *uint64)
	GetKupeNo(id uint64) (result string)
}

//UpdateSingleStatus update
func (f *hayvanBilgisiApp) GetKupeNo(id uint64) (result string) {
	return f.request.GetKupeNo(id)
}

//UpdateSingleStatus update
func (f *hayvanBilgisiApp) UpdateSingleStatus(id uint64, status int) {
	f.request.UpdateSingleStatus(id, status)
}

//HayvanBilgisiApp struct  init
type hayvanBilgisiApp struct {
	request HayvanBilgisiAppInterface
}

var _ HayvanBilgisiAppInterface = &hayvanBilgisiApp{}

//Save service init
func (f *hayvanBilgisiApp) Save(Cat *entity.SacrificeHayvanBilgisi) (*entity.SacrificeHayvanBilgisi, map[string]string) {
	return f.request.Save(Cat)
}

//GetAll service init
func (f *hayvanBilgisiApp) GetAll() ([]entity.SacrificeHayvanBilgisi, error) {
	return f.request.GetAll()
}

//ListDataTable service init
func (f *hayvanBilgisiApp) ListDataTable(c *gin.Context) (entity.HayvanBilgisiDataResult, error) {
	return f.request.ListDataTable(c)
}

//GetAllStatus1 service init
func (f *hayvanBilgisiApp) GetAllFindDurum(status int) ([]entity.SacrificeHayvanBilgisi, error) {
	return f.request.GetAllFindDurum(status)
}

func (f *hayvanBilgisiApp) Count(totalCount *int64) {
	f.request.Count(totalCount)
}

func (f *hayvanBilgisiApp) KupeNoCount(kupeNo int, totalCount *int) {
	f.request.KupeNoCount(kupeNo, totalCount)
}

func (f *hayvanBilgisiApp) LastID(id uint64, lastid *uint64) {
	f.request.LastID(id, lastid)
}

func (f *hayvanBilgisiApp) NextID(id uint64, nextid *uint64) {
	f.request.NextID(id, nextid)
}

func (f *hayvanBilgisiApp) GetAllP(perPage int, offset int) ([]entity.SacrificeHayvanBilgisi, error) {
	return f.request.GetAllP(perPage, offset)
}

//GetByID service init
func (f *hayvanBilgisiApp) GetByID(CatID uint64) (*entity.SacrificeHayvanBilgisi, error) {
	return f.request.GetByID(CatID)
}

//GetByIDRelated service init
func (f *hayvanBilgisiApp) GetByIDRelated(ID uint64) (*dto.HayvanBilgisi, error) {
	return f.request.GetByIDRelated(ID)
}

//Update service init
func (f *hayvanBilgisiApp) Update(data *entity.SacrificeHayvanBilgisi) (*entity.SacrificeHayvanBilgisi, map[string]string) {
	return f.request.Update(data)
}

//Delete service init
func (f *hayvanBilgisiApp) Delete(id uint64) error {
	return f.request.Delete(id)
}
