package services

import (
	"stncCms/app/domain/entity"
)

//CCategoriesBranchJoinAppInterface service
type CategoriesBranchJoinAppInterface interface {
	Save(*entity.CategoriesBranch) (*entity.CategoriesBranch, map[string]string)
	GetAllforUserID(uint64) ([]entity.CategoriesBranch, error)
	GetAllforBranchID(uint64) ([]entity.CategoriesBranch, error)
	GetAll() ([]entity.CategoriesBranch, error)
	Update(*entity.CategoriesBranch) (*entity.CategoriesBranch, map[string]string)
	Delete(uint64) error
	DeleteForUserID(uint64) error
	DeleteForBranchID(uint64) error
}

//CCategoriesBranchJoinApp struct  init
type CCategoriesBranchJoinApp struct {
	fr CategoriesBranchJoinAppInterface
}

var _ CategoriesBranchJoinAppInterface = &CCategoriesBranchJoinApp{}

//Save service init
func (f *CCategoriesBranchJoinApp) Save(Cat *entity.CategoriesBranch) (*entity.CategoriesBranch, map[string]string) {
	return f.fr.Save(Cat)
}

//GetAll service init
func (f *CCategoriesBranchJoinApp) GetAll() ([]entity.CategoriesBranch, error) {
	return f.fr.GetAll()
}

//GetAllforUserID service init
func (f *CCategoriesBranchJoinApp) GetAllforUserID(UserID uint64) ([]entity.CategoriesBranch, error) {
	return f.fr.GetAllforUserID(UserID)
}

//GetAllforCatID service init
func (f *CCategoriesBranchJoinApp) GetAllforBranchID(branchID uint64) ([]entity.CategoriesBranch, error) {
	return f.fr.GetAllforBranchID(branchID)
}

//Update service init
func (f *CCategoriesBranchJoinApp) Update(Cat *entity.CategoriesBranch) (*entity.CategoriesBranch, map[string]string) {
	return f.fr.Update(Cat)
}

//Delete service init
func (f *CCategoriesBranchJoinApp) Delete(ID uint64) error {
	return f.fr.Delete(ID)
}

//DeleteForKioskID service init
func (f *CCategoriesBranchJoinApp) DeleteForUserID(userId uint64) error {
	return f.fr.DeleteForUserID(userId)
}

//DeleteForCatID service init
func (f *CCategoriesBranchJoinApp) DeleteForBranchID(branchID uint64) error {
	return f.fr.DeleteForBranchID(branchID)
}
