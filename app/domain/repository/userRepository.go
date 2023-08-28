package repository

import (
	"errors"
	"fmt"
	"stncCms/app/domain/dto"
	"stncCms/app/domain/entity"
	"stncCms/app/infrastructure/security"
	"stncCms/app/services"
	"strings"

	"github.com/jinzhu/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func UserRepositoryInit(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

//UserRepo implements the repository.UserRepository interface
var _ services.UserAppInterface = &UserRepo{}

func (r *UserRepo) SaveUser(user *entity.User) (*entity.User, map[string]string) {
	dbErr := map[string]string{}
	err := r.db.Debug().Create(&user).Error

	if err != nil {
		//If the email is already taken
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["email_taken"] = "email already taken"
			return nil, dbErr
		}
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return user, nil
}

func (r *UserRepo) GetUser(id uint64) (*entity.User, error) {
	var user entity.User
	var err error
	err = r.db.Debug().Where("id = ?", id).Take(&user).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (r *UserRepo) GetUsers() ([]entity.User, error) {
	var users []entity.User
	var err error
	err = r.db.Debug().Find(&users).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("user not found")
	}
	return users, nil
}

//api kullanacak
func (r *UserRepo) GetUserByEmailAndPassword(u *entity.User) (*entity.User, map[string]string) {
	var user entity.User
	dbErr := map[string]string{}
	var err error
	err = r.db.Debug().Where("email = ?", u.Email).Take(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		dbErr["no_user"] = "user not found"
		return nil, dbErr
	}
	if err != nil {
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	//Verify the password

	ret := security.VerifyPasswordApi(user.Password, u.Password)
	if ret == false {
		dbErr["incorrect_password"] = "incorrect password"
		return nil, dbErr
	}
	return &user, nil
}
func (r *UserRepo) GetUserByEmailAndPassword2(email string, InputPassword string) (*entity.User, map[string]string) {
	var user entity.User

	dbErr := map[string]string{}
	var err error
	err = r.db.Debug().Where("email = ?", email).Take(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		dbErr["no_user"] = "user not found"
		return nil, dbErr
	}
	if err != nil {
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	//Verify the password

	ret := security.VerifyPassword(user.Password, InputPassword)
	if ret == false {
		dbErr["incorrect_password"] = "incorrect password"
		return nil, dbErr
	}
	fmt.Println(dbErr)
	return &user, nil
}

///*******************************  sondaradan eklendı

//Save data
func (r *UserRepo) Save(data *entity.User) (*entity.User, map[string]string) {
	dbErr := map[string]string{}
	//The images are uploaded to digital ocean spaces. So we need to prepend the url. This might not be your use case, if you are not uploading image to Digital Ocean.
	var err error
	err = r.db.Debug().Create(&data).Error
	if err != nil {
		//since our title is unique
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["unique_title"] = "data title already taken"
			return nil, dbErr
		}
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return data, nil
}

func (r *UserRepo) SaveDto(data *dto.User) (*dto.User, map[string]string) {
	dbErr := map[string]string{}
	//The images are uploaded to digital ocean spaces. So we need to prepend the url. This might not be your use case, if you are not uploading image to Digital Ocean.
	var err error
	err = r.db.Debug().Table("user").Create(&data).Error
	if err != nil {
		//since our title is unique
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["unique_title"] = "data title already taken"
			return nil, dbErr
		}
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return data, nil
}

//Update upate data
func (r *UserRepo) Update(data *entity.User) (*entity.User, map[string]string) {
	dbErr := map[string]string{}
	err := r.db.Debug().Save(&data).Error
	//db.Table("libraries").Where("id = ?", id).Update(dataData)

	if err != nil {
		//since our title is unique
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["unique_title"] = "title already taken"
			return nil, dbErr
		}
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return data, nil
}

func (r *UserRepo) UpdateDto(data *dto.User) (*dto.User, map[string]string) {
	dbErr := map[string]string{}
	err := r.db.Debug().Table("user").Save(&data).Error
	//db.Table("libraries").Where("id = ?", id).Update(dataData)

	if err != nil {
		//since our title is unique
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["unique_title"] = "title already taken"
			return nil, dbErr
		}
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return data, nil
}

//Count fat
func (r *UserRepo) Count(dataTotalCount *int64) {
	var data entity.User
	var count int64
	r.db.Debug().Model(data).Count(&count)
	*dataTotalCount = count
}

//Delete data
func (r *UserRepo) Delete(id uint64) error {
	var data entity.User
	var err error
	err = r.db.Debug().Where("id = ?", id).Delete(&data).Error
	if err != nil {
		return errors.New("database error, please try again")
	}
	return nil
}

//GetByID get data
func (r *UserRepo) GetByID(id uint64) (*entity.User, error) {
	var data entity.User
	var err error
	err = r.db.Debug().Where("id = ?", id).Take(&data).Error
	if err != nil {
		return nil, errors.New("database error, please try again")
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("data not found")
	}
	return &data, nil
}

//GetAll all data
func (r *UserRepo) GetAll() ([]entity.User, error) {
	var datas []entity.User
	var err error
	err = r.db.Debug().Order("created_at desc").Find(&datas).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("data not found")
	}
	return datas, nil
}

//GetAllP pagination all data
func (r *UserRepo) GetAllP(datasPerPage int, offset int) ([]entity.User, error) {
	var datas []entity.User
	var err error
	err = r.db.Debug().Limit(datasPerPage).Offset(offset).Order("created_at desc").Find(&datas).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("data not found")
	}
	return datas, nil
}

//SetKioskSliderUpdate update data
func (r *UserRepo) SetKioskSliderUpdate(id uint64, status int) {
	r.db.Debug().Table("user").Where("id = ?", id).Update("status", status)
}
