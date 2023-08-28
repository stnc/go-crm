package services

import (
	"stncCms/app/domain/dto"
	"stncCms/app/domain/entity"
)

//UserAppInterface interface
type UserAppInterface interface {
	SaveUser(*entity.User) (*entity.User, map[string]string)
	GetUser(uint64) (*entity.User, error)
	GetUsers() ([]entity.User, error)
	GetUserByEmailAndPassword(*entity.User) (*entity.User, map[string]string)
	GetUserByEmailAndPassword2(email string, password string) (*entity.User, map[string]string)

	Save(*entity.User) (*entity.User, map[string]string)
	SaveDto(*dto.User) (*dto.User, map[string]string)
	GetByID(uint64) (*entity.User, error)
	GetAll() ([]entity.User, error)
	GetAllP(int, int) ([]entity.User, error)
	Update(*entity.User) (*entity.User, map[string]string)
	UpdateDto(*dto.User) (*dto.User, map[string]string)
	Count(*int64)
	Delete(uint64) error
	// SetUserUpdate(uint64, int)

}

type userApp struct {
	request UserAppInterface
}

//UserApp implements the UserAppInterface
var _ UserAppInterface = &userApp{}

func (u *userApp) SaveUser(user *entity.User) (*entity.User, map[string]string) {
	return u.request.SaveUser(user)
}

func (u *userApp) GetUser(userID uint64) (*entity.User, error) {
	return u.request.GetUser(userID)
}

func (u *userApp) GetUsers() ([]entity.User, error) {
	return u.request.GetUsers()
}

func (u *userApp) GetUserByEmailAndPassword(user *entity.User) (*entity.User, map[string]string) {
	return u.request.GetUserByEmailAndPassword(user)
}
func (u *userApp) GetUserByEmailAndPassword2(email string, password string) (*entity.User, map[string]string) {
	return u.request.GetUserByEmailAndPassword2(email, password)
}

///

func (f *userApp) Count(UserTotalCount *int64) {
	f.request.Count(UserTotalCount)
}

func (f *userApp) Save(User *entity.User) (*entity.User, map[string]string) {
	return f.request.Save(User)
}

func (f *userApp) SaveDto(User *dto.User) (*dto.User, map[string]string) {
	return f.request.SaveDto(User)
}

func (f *userApp) GetAll() ([]entity.User, error) {
	return f.request.GetAll()
}

func (f *userApp) GetAllP(UsersPerPage int, offset int) ([]entity.User, error) {
	return f.request.GetAllP(UsersPerPage, offset)
}

func (f *userApp) GetByID(UserID uint64) (*entity.User, error) {
	return f.request.GetByID(UserID)
}

func (f *userApp) Update(User *entity.User) (*entity.User, map[string]string) {
	return f.request.Update(User)
}

func (f *userApp) UpdateDto(User *dto.User) (*dto.User, map[string]string) {
	return f.request.UpdateDto(User)
}

func (f *userApp) Delete(UserID uint64) error {
	return f.request.Delete(UserID)
}

// func (f *userApp) SetUserUpdate(id uint64, status int) {
// 	f.request.SetUserUpdate(id, status)
// }
