package controller

import (
	"fmt"
	"net/http"
	"reflect"
	"stncCms/app/domain/dto"
	"stncCms/app/domain/entity"

	"stncCms/app/domain/helpers/stnccollection"
	"stncCms/app/domain/helpers/stncdatetime"
	"stncCms/app/domain/helpers/stncsession"
	"stncCms/app/services"
	"strconv"

	"github.com/astaxie/beego/utils/pagination"
	"github.com/flosch/pongo2/v4"

	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

//userControl constructor
type UserControl struct {
	UserControlApp       services.UserAppInterface
	Branch               services.BranchAppInterface
	CategoriesBranchJoin services.CategoriesBranchJoinAppInterface
}

const viewPathuserControl = "admin/user/"

//InitUserControl userControl controller constructor
func InitUserControl(KiApp services.UserAppInterface, BranchApp services.BranchAppInterface, categoriesBranchJoinApp services.CategoriesBranchJoinAppInterface) *UserControl {
	return &UserControl{
		UserControlApp:       KiApp,
		Branch:               BranchApp,
		CategoriesBranchJoin: categoriesBranchJoinApp,
	}
}

//Index list
func (access *UserControl) Index(c *gin.Context) {
	// alluserControl, err := access.userControlApp.GetAlluserControl()

	stncsession.IsLoggedInRedirect(c)
	flashMsg := stncsession.GetFlashMessage(c)

	var tarih stncdatetime.Inow
	var total int64
	access.UserControlApp.Count(&total)
	userControlsPerPage := 10
	paginator := pagination.NewPaginator(c.Request, userControlsPerPage, total)
	offset := paginator.Offset()

	userControls, _ := access.UserControlApp.GetAllP(userControlsPerPage, offset)

	// var tarih stncdatetime.Inow
	// fmt.Println(tarih.TarihFullSQL("2020-05-21 05:08"))
	// fmt.Println(tarih.AylarListe("May"))
	// fmt.Println(tarih.Tarih())
	// //	tarih.FormatTarihForMysql("2020-05-17 05:08:40")
	//	tarih.FormatTarihForMysql("2020-05-17 05:08:40")

	viewData := pongo2.Context{
		"paginator": paginator,
		"title":     "İçerik Ekleme",
		"allData":   userControls,
		"tarih":     tarih,
		"flashMsg":  flashMsg,
		"csrf":      csrf.GetToken(c),
	}

	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		viewPathuserControl+"index.html",
		// Pass the data that the page uses
		viewData,
	)
}

//Create all list f
func (access *UserControl) Create(c *gin.Context) {
	//	stncsession.IsLoggedInRedirect(c)
	cats, _ := access.Branch.GetAll()
	viewData := pongo2.Context{
		"title":    "İçerik Ekleme",
		"catsData": cats,
		"csrf":     csrf.GetToken(c),
	}
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		viewPathuserControl+"create.html",
		// Pass the data that the page uses
		viewData,
	)
}

//Store save method
func (access *UserControl) Store(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	var userControl, _, _ = userControlModel(c)
	var saveuserControlError = make(map[string]string)
	// saveuserControlError = userControl.Validate2()

	catsPost := c.PostFormArray("cats")
	//fmt.Println(catsPost)
	catsData, _ := access.Branch.GetAll()
	fmt.Println(reflect.ValueOf(catsData).Kind())
	for key, row := range catsData {
		catsData[key].ID = row.ID
		catsData[key].Title = row.Title
		//a, _ := strconv.Atoi(catsPost[key])
		finding := strconv.FormatInt(int64(row.ID), 10)
		_, found := stnccollection.FindSlice(catsPost, finding)
		if found {
			catsData[key].SelectedID = row.ID
		}
	}

	if len(saveuserControlError) == 0 {

		saveData, saveErr := access.UserControlApp.SaveDto(&userControl)

		if saveErr != nil {
			saveuserControlError = saveErr
		}

		lastID := strconv.FormatUint(uint64(saveData.ID), 10)
		var catPost = entity.CategoriesBranch{}
		for _, row := range catsPost {
			catID, _ := strconv.ParseUint(row, 10, 64)
			catPost.BranchID = catID
			catPost.UserID = saveData.ID
			saveCat, _ := access.CategoriesBranchJoin.Save(&catPost)

			catPost.ID = saveCat.ID + 1
		}
		stncsession.SetFlashMessage("Kayıt başarı ile eklendi", "success", c)

		c.Redirect(http.StatusMovedPermanently, "/"+viewPathuserControl+"edit/"+lastID)
		return
	}

	viewData := pongo2.Context{
		"title":    "içerik ekleme",
		"catsPost": catsPost,
		"catsData": catsData,
		"csrf":     csrf.GetToken(c),
		"err":      saveuserControlError,
		"data":     userControl,
	}
	c.HTML(
		http.StatusOK,
		viewPathuserControl+"create.html",
		viewData,
	)

}

//Edit edit data
func (access *UserControl) Edit(c *gin.Context) {
	//strconv.Atoi(c.Param("id"))
	//userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	stncsession.IsLoggedInRedirect(c)
	flashMsg := stncsession.GetFlashMessage(c)

	if userID, err := strconv.ParseUint(c.Param("UserID"), 10, 64); err == nil {
		if userControls, err := access.UserControlApp.GetByID(userID); err == nil {
			var catsPost []string
			catsPostData, _ := access.CategoriesBranchJoin.GetAllforUserID(userID)
			for _, row := range catsPostData {
				str := strconv.FormatUint(row.BranchID, 10) //uint64 to stringS
				catsPost = append(catsPost, str)
			}
			catsData, _ := access.Branch.GetAll()
			for key, row := range catsData {
				catsData[key].ID = row.ID
				catsData[key].Title = row.Title
				//a, _ := strconv.Atoi(catsPost[key])
				finding := strconv.FormatInt(int64(row.ID), 10)
				_, found := stnccollection.FindSlice(catsPost, finding)
				if found {
					catsData[key].SelectedID = row.ID
				}
			}

			viewData := pongo2.Context{
				"title":    "kullanıcı düzenleme",
				"catsPost": catsPost,
				"catsData": catsData,
				"data":     userControls,

				"csrf":     csrf.GetToken(c),
				"flashMsg": flashMsg,
			}
			c.HTML(
				http.StatusOK,
				viewPathuserControl+"edit.html",
				viewData,
			)

		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

//Update data
func (access *UserControl) Update(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)

	var userControl, idN, id = userControlModel(c)
	var saveuserControlError = make(map[string]string)

	catsPost := c.PostFormArray("cats")

	catsData, _ := access.Branch.GetAll()
	for key, row := range catsData {
		catsData[key].ID = row.ID
		catsData[key].Title = row.Title
		finding := strconv.FormatInt(int64(row.ID), 10)
		_, found := stnccollection.FindSlice(catsPost, finding)
		if found {
			catsData[key].SelectedID = row.ID
		}
	}
	if len(saveuserControlError) == 0 {

		saveData, saveErr := access.UserControlApp.UpdateDto(&userControl)
		if saveErr != nil {
			saveuserControlError = saveErr
		}

		var catPost = entity.CategoriesBranch{}
		access.CategoriesBranchJoin.DeleteForUserID(idN)
		for _, row := range catsPost {
			catID, _ := strconv.ParseUint(row, 10, 64)
			catPost.BranchID = catID
			catPost.UserID = saveData.ID
			saveCat, _ := access.CategoriesBranchJoin.Save(&catPost)
			catPost.ID = saveCat.ID + 1
		}
		stncsession.SetFlashMessage("Kayıt başarı ile eklendi", "success", c)

		c.Redirect(http.StatusMovedPermanently, "/"+viewPathuserControl+"edit/"+id)
		return
	}

	viewData := pongo2.Context{
		"title": "Kullanıcı düzenleme",
		"err":   saveuserControlError,
		"csrf":  csrf.GetToken(c),

		"data": userControl,
	}
	c.HTML(
		http.StatusOK,
		viewPathuserControl+"edit.html",
		viewData,
	)
}

//TODO: resim silmeyi unutma
//Delete data
func (access *UserControl) Delete(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)

	if userID, err := strconv.ParseUint(c.Param("UserID"), 10, 64); err == nil {

		access.UserControlApp.Delete(userID)
		stncsession.SetFlashMessage("silindi", "success", c)

		c.Redirect(http.StatusMovedPermanently, "/"+viewPathuserControl)
		return

	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

// func (access *UserControl) Status(c *gin.Context) {
// 	stncsession.IsLoggedInRedirect(c)

// 	if userID, err := strconv.ParseUint(c.Param("UserID"), 10, 64); err == nil {
// 		status := stnccollection.StringToint(c.DefaultQuery("status", "false"))
// 		access.UserControlApp.SetUserControlUpdate(userID, status)
// 		stncsession.SetFlashMessage("duurm", "success", c)

// 		c.Redirect(http.StatusMovedPermanently, "/admin/userControl")
// 		return

// 	} else {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	}
// }

//form userControl model
func userControlModel(c *gin.Context) (userControl dto.User, idD uint64, idStr string) {
	id := c.PostForm("ID")
	description := c.PostForm("Title")

	idInt, _ := strconv.Atoi(id)
	var idN uint64
	idN = uint64(idInt)
	// var userControl = dto.User{}
	userControl.ID = idN
	userControl.UserID = 1 //TODO: bunun yok edılmesı lazım s
	userControl.Description = description

	return userControl, idN, id
}
