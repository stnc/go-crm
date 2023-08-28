//  golang gin framework mvc and clean code project
//  Licensed under the Apache License 2.0
//  @author Selman TUNÇ <selmantunc@gmail.com>
//  @link: https://github.com/stnc/go-mvc-blog-clean-code
//  @license: Apache License 2.0
package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"stncCms/app/domain/repository"
	"stncCms/app/infrastructure/auth"

	apiController "stncCms/app/web.api/controller"
	"stncCms/app/web.api/controller/middleware"
	"stncCms/app/web/controller"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stnc/pongo4gin"

	csrf "github.com/utrack/gin-csrf"
)

//https://github.com/stnc-go/gobyexample/blob/master/pongo2render/render.go
func init() {
	//To load our environmental variables.

	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}

	/* //bu sunucuda çalışıyor
		    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	        if err != nil {
	            log.Fatal(err)
	        }
	        environmentPath := filepath.Join(dir, ".env")
	        err = godotenv.Load(environmentPath)
	        fatal(err)
	*/

}

func main() {

	// appEnv := os.Getenv("APP_ENV")
	// if appEnv == "local" {
	// 	err := beeep.Alert("Uygulama çalıştı", "Web Server Çalışmaya Başladı localhost:"+port, "assets/warning.png")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	//redis details

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	debugMode := os.Getenv("MODE")

	//
	db := repository.DbConnect()
	services, err := repository.RepositoriesInit(db)
	if err != nil {
		panic(err)
	}
	//defer services.Close()
	services.Automigrate()

	redisService, err := auth.RedisDBInit(redisHost, redisPort, redisPassword)
	if err != nil {
		log.Fatal(err)
	}

	token := auth.NewToken()

	// upload := stncupload.NewFileUpload()

	usersAPI := apiController.InitUsers(services.User, redisService.Auth, token)

	postsAPI := apiController.InitPost(services.Post, services.User, redisService.Auth, token)

	//posts := controller.InitPost(services.Post, services.CatPost, services.Cat, services.Lang, services.User)

	kurbanHandle := controller.InitGkurban(services.Kurban, services.Kisiler, services.Media)

	odemelerHandle := controller.InitOdemeler(services.Kodemeler, services.Kurban)

	kisilerHandle := controller.InitGKisiler(services.Kisiler)

	GruplarHandle := controller.InitGruplar(services.HayvanBilgisi, services.Kurban, services.Gruplar, services.Options, services.Kodemeler, services.Kisiler)

	hayvanSatisYerleriHandle := controller.InitHayvanSatisYerleri(services.HayvanSatisYerleri)

	hayvanBilgisiHandle := controller.InitHayvanBilgisi(services.HayvanBilgisi, services.Media)

	// OptionHandle := controller.InitAyarlar(services.Ayarlar)

	optionsHandle := controller.InitOptions(services.Options)
	kioskSlider := controller.InitkioskSlider(services.KioskSlider, services.User, services.Media, services.CategoriesKiosk, services.CategoriesKioskJoin)

	userHandle := controller.InitUserControl(services.User, services.Branch, services.CategoriesBranchJoin)

	login := controller.InitLogin(services.User)

	//cat := controller.InitPost(services.Post, services.User)
	authenticate := apiController.NewAuthenticate(services.User, redisService.Auth, token)

	webArchive := controller.InitWebArchive(services.WebArchive, services.WebArchiveLink, services.User)

	switch debugMode {
	case "RELEASE":
		gin.SetMode(gin.ReleaseMode)

	case "DEBUG":
		gin.SetMode(gin.DebugMode)

	case "TEST":
		gin.SetMode(gin.TestMode)

	default:
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	//TODO: https://github.com/denisbakhtin/ginblog/blob/master/main.go burada memstore kullanımı var ona bakılablir

	store := cookie.NewStore([]byte("SpeedyGonzales"))
	////60 dakika olan 1 saat tam olarak ( 60x60) 3600 saniyedir.
	//60 saniye * 60 = 1 saat //60*60
	//3600 (1 saat ) * 5 = 5 saat
	store.Options(sessions.Options{Path: "/", HttpOnly: true, MaxAge: 3600 * 8}) //Also set Secure: true if using SSL, you should though
	// store.Options(sessions.Options{Path: "/", HttpOnly: true, MaxAge: -1}) //Also set Secure: true if using SSL, you should though

	r.Use(sessions.Sessions("myCRM", store))

	r.Use(middleware.CORSMiddleware()) //For CORS

	//TODO: csrf kontrolu nasıl olacak
	r.Use(csrf.Middleware(csrf.Options{
		Secret: "SpeedyGonzales",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	}))

	r.HTMLRender = pongo4gin.TemplatePath("public/view")

	r.MaxMultipartMemory = 1 >> 20 // 8 MiB

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	r.Static("/assets", "./public/static")

	r.StaticFS("/upload", http.Dir("./public/upl"))
	//r.StaticFile("/favicon.ico", "./resources/favicon.ico")

	// Use the following code if you need to write the logs to file and console at the same time.
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r.GET("/", controller.Index)
	r.GET("login", login.Login)
	r.GET("sifre", login.SifreVer) 

	r.POST("login", login.LoginPost)
	r.GET("logout", login.Logout)

	r.GET("optionsDefault", controller.OptionsDefault)

	//
	r.GET("/kurbanBilgi/:slug", kurbanHandle.KurbanBilgi)

	//api routes
	v1 := r.Group("/api/v1")
	{
		v1.POST("users", usersAPI.SaveUser)
		v1.GET("users", usersAPI.GetUsers)

		v1.GET("users/:user_id", usersAPI.GetUser)
		v1.GET("postall", postsAPI.GetAllPost)
		v1.POST("post", postsAPI.SavePost)
		v1.PUT("post/:post_id", middleware.AuthMiddleware(), postsAPI.UpdatePost)
		v1.GET("post/:post_id", postsAPI.GetPostAndCreator)
		v1.DELETE("post/:post_id", middleware.AuthMiddleware(), postsAPI.DeletePost)
		// cs.GET("/allcoins", controller.AllCoins())
		// cs.GET("/mycoins/:id", controller.MyCoins())
		// cs.GET("/create", controller.CreateCoin())
		// cs.POST("/store", controller.StoreCoin())
		// cs.GET("/edit/:id", controller.EditCoin())
		// cs.PUT("/update/:id", controller.UpdateCoin())
		// e.GET("/cpr/:slug", controller.CoinPreview())
		// cs.DELETE("/:id", controller.DeleteCoin())
		//authentication routes
		v1.POST("login", authenticate.Login)
		v1.POST("logout", authenticate.Logout)
		v1.POST("refresh", authenticate.Refresh)
	}

	//GENEL KURBAN
	kurbanGroup := r.Group("/admin/kurban")
	{
		kurbanGroup.GET("/", kurbanHandle.Index)

		kurbanGroup.GET("create", kurbanHandle.Create)

		kurbanGroup.POST("store", kurbanHandle.Store)
		kurbanGroup.GET("edit/:kurbanID", kurbanHandle.Edit)
		kurbanGroup.POST("update", kurbanHandle.Update)
		//api
		kurbanGroup.GET("referansCreateModalBox/:viewID", kurbanHandle.ReferansCreateModalBox)
		kurbanGroup.GET("odemeEkleCreateModalBox", kurbanHandle.OdemeEkleCreateModalBox)
		kurbanGroup.GET("grupLideri/:kurbanID", kurbanHandle.GrupLideriAta)
		kurbanGroup.GET("vekaletDurumu/:kurbanID", kurbanHandle.VekaletDurumu)
		kurbanGroup.POST("upload/:ID", kurbanHandle.Upload)
		kurbanGroup.GET("media-delete/:ID", kurbanHandle.MediaDelete)
		kurbanGroup.GET("listDataTable", kurbanHandle.ListDataTable)
	}

	kisilerGroup := r.Group("/admin/kisiler")
	{
		kisilerGroup.GET("/", kisilerHandle.Index)
		kisilerGroup.GET("index", kisilerHandle.Index)
		kisilerGroup.GET("IndexV1", kisilerHandle.IndexV1)
		kisilerGroup.GET("create", kisilerHandle.Create)
		kisilerGroup.POST("store", kisilerHandle.Store)
		kisilerGroup.GET("edit/:ID", kisilerHandle.Edit)
		kisilerGroup.GET("delete/:ID", kisilerHandle.Delete)
		kisilerGroup.POST("update", kisilerHandle.Update)
		kisilerGroup.POST("kisiAra", kisilerHandle.KisiAraAjax)
		kisilerGroup.POST("kisiEkleAjax", kisilerHandle.KisiEkleAjax)
		kisilerGroup.GET("listDataTable", kisilerHandle.ListDataTable)
		kisilerGroup.GET("kisiAciklama/:ID", kisilerHandle.KisiAciklama)
		kisilerGroup.GET("kisiAciklamaEdit/:ID", kisilerHandle.KisiAciklamaEdit)
	}

	odemelerGroup := r.Group("/admin/odemeler")
	{
		odemelerGroup.GET("create", odemelerHandle.Create)
		odemelerGroup.POST("store", odemelerHandle.Store)
		odemelerGroup.POST("odemeEkleAjax", odemelerHandle.OdemeEkleAjax)
		odemelerGroup.GET("edit/:gID", odemelerHandle.Edit)
		odemelerGroup.POST("update", odemelerHandle.Update)
		odemelerGroup.GET("makbuz/:gID", odemelerHandle.Makbuz)
		odemelerGroup.GET("delete/:odemeID", odemelerHandle.OdemeSil)
	}

	gruplar := r.Group("/admin/gruplar")
	{
		gruplar.GET("/", GruplarHandle.Index)
		gruplar.GET("index", GruplarHandle.Index)
		gruplar.GET("create", GruplarHandle.Create)
		gruplar.POST("store", GruplarHandle.Store)
		gruplar.GET("createEmpty", GruplarHandle.CreateEmpty)
		gruplar.POST("storeEmpty", GruplarHandle.StoreEmpty)
		gruplar.GET("edit/:kID", GruplarHandle.Edit)
		gruplar.GET("editEmpty/:kID", GruplarHandle.EditEmpty)
		gruplar.POST("update", GruplarHandle.Update)
		// gruplar.POST("grupDegisim", GruplarHandle.grupDegisim)
		gruplar.GET("degistir", GruplarHandle.Degistir)
		gruplar.POST("degistir", GruplarHandle.DegistirStore)
		gruplar.GET("grupBilgisiAPI/:kID", GruplarHandle.GrupBilgisiAPI)
		gruplar.GET("GrupBilgisiHayvanBosOlanlarAPI/:kID", GruplarHandle.GrupBilgisiHayvanBosOlanlarAPI)
		gruplar.GET("hayvanata", GruplarHandle.HayvanAtamasiYap)
		gruplar.POST("hayvanata", GruplarHandle.HayvanAtamasiStore)
		gruplar.GET("gruplarListeAjaxAgirlikTuru/:agirlikID", GruplarHandle.GruplarListeAjaxAgirlikTuru)
		gruplar.GET("excel", GruplarHandle.Excel)
		gruplar.GET("IndexAta", GruplarHandle.IndexAta)
		gruplar.POST("yerDegistir", GruplarHandle.YerDegistir)

	}

	hayvanSatisYerleriGroup := r.Group("/admin/hayvanSatisYerleri")
	{
		hayvanSatisYerleriGroup.GET("/", hayvanSatisYerleriHandle.Index)
		hayvanSatisYerleriGroup.GET("index", hayvanSatisYerleriHandle.Index)
		hayvanSatisYerleriGroup.GET("create", hayvanSatisYerleriHandle.Create)
		hayvanSatisYerleriGroup.POST("store", hayvanSatisYerleriHandle.Store)
		hayvanSatisYerleriGroup.GET("edit/:kID", hayvanSatisYerleriHandle.Edit)
		hayvanSatisYerleriGroup.POST("update", hayvanSatisYerleriHandle.Update)
		hayvanSatisYerleriGroup.GET("listDataTable", hayvanSatisYerleriHandle.ListDataTable)
	}

	hayvanBilgisiGroup := r.Group("/admin/hayvanBilgisi")
	{
		hayvanBilgisiGroup.GET("/", hayvanBilgisiHandle.Index)
		hayvanBilgisiGroup.GET("index", hayvanBilgisiHandle.Index)
		hayvanBilgisiGroup.GET("listDataTable", hayvanBilgisiHandle.ListDataTable)
		hayvanBilgisiGroup.GET("create", hayvanBilgisiHandle.Create)
		hayvanBilgisiGroup.POST("store", hayvanBilgisiHandle.Store)
		hayvanBilgisiGroup.GET("edit/:kID", hayvanBilgisiHandle.Edit)
		hayvanBilgisiGroup.POST("update", hayvanBilgisiHandle.Update)
		hayvanBilgisiGroup.GET("hayvanListeAjax/:hID", hayvanBilgisiHandle.HayvanListeAjax)
		hayvanBilgisiGroup.POST("upload/:ID", hayvanBilgisiHandle.Upload)
		hayvanBilgisiGroup.GET("media-delete/:ID", hayvanBilgisiHandle.MediaDelete)
	}

	optionsGroup := r.Group("/admin/options")
	{
		optionsGroup.GET("/", optionsHandle.Index)
		optionsGroup.POST("update", optionsHandle.Update)
		optionsGroup.GET("makbuzNo", optionsHandle.MakbuzNo)
	}

	adminWebarchive := r.Group("/admin/webarchive")
	{
		adminWebarchive.GET("/", webArchive.Index)
		adminWebarchive.GET("/index", webArchive.Index)
		adminWebarchive.GET("create", webArchive.Create)
		adminWebarchive.POST("store", webArchive.Store)
		adminWebarchive.GET("edit/:ID", webArchive.Edit)
		adminWebarchive.POST("update", webArchive.Update)
		adminWebarchive.GET("delete/:ID", webArchive.Delete)
		adminWebarchive.GET("run/:ID", webArchive.LinkPdfRun)
	}

	//kiosk başladı
	kioskGroup := r.Group("/admin/kioskSlider")
	{
		kioskGroup.GET("/", kioskSlider.Index)
		kioskGroup.GET("index", kioskSlider.Index)
		kioskGroup.GET("create", kioskSlider.Create)
		kioskGroup.POST("store", kioskSlider.Store)
		kioskGroup.GET("edit/:KioskSliderID", kioskSlider.Edit)
		kioskGroup.GET("delete/:ID", kioskSlider.Delete)
		kioskGroup.GET("media-delete/:ID", kioskSlider.MediaDelete)
		kioskGroup.GET("status/:KioskSliderID", kioskSlider.Status)
		kioskGroup.POST("update", kioskSlider.Update)
		kioskGroup.POST("upload/:KioskSliderID", kioskSlider.Upload)
	}

	userGroup := r.Group("/admin/user")
	{
		userGroup.GET("/", userHandle.Index)
		userGroup.GET("index", userHandle.Index)
		userGroup.GET("create", userHandle.Create)
		userGroup.POST("store", userHandle.Store)
		userGroup.GET("edit/:UserID", userHandle.Edit)
		userGroup.GET("delete/:ID", userHandle.Delete)
		userGroup.POST("update", userHandle.Update)
	}

	// adminPost := r.Group("/admin/post")
	// {
	// 	adminPost.GET("/", posts.Index)
	// 	adminPost.GET("index", posts.Index)
	// 	adminPost.GET("create", posts.Create)
	// 	adminPost.POST("store", posts.Store)
	// 	adminPost.GET("edit/:postID", posts.Edit)
	// 	adminPost.POST("update", posts.Update)
	// 	adminPost.POST("upload", posts.Upload)
	// }

	// defaultOptionsGroup := r.Group("/admin/defaultOptions")
	// {
	// 	defaultOptionsGroup.GET("/", optionsDefaultHandle.Index)
	// 	defaultOptionsGroup.GET("index", optionsDefaultHandle.Index)
	// 	defaultOptionsGroup.GET("create", optionsDefaultHandle.Create)
	// 	defaultOptionsGroup.POST("store", optionsDefaultHandle.Store)
	// 	defaultOptionsGroup.GET("edit/:kID", optionsDefaultHandle.Edit)
	// 	defaultOptionsGroup.POST("update", optionsDefaultHandle.Update)
	// }

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	//Starting the application
	appPort := os.Getenv("PORT")
	if appPort == "" {
		appPort = "8080" //localhost
	}
	log.Fatal(r.Run(":" + appPort))

}
