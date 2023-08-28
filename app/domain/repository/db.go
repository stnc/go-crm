package repository

import (
	"fmt"
	"os"
	"stncCms/app/domain/entity"
	"stncCms/app/services"

	"github.com/hypnoglow/gormzap"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"

	_ "github.com/lib/pq" // here
	_ "gorm.io/driver/mysql"
	_ "gorm.io/driver/postgres"
)

var DB *gorm.DB

//Repositories strcut
type Repositories struct {
	User services.UserAppInterface
	Post services.PostAppInterface

	Lang               services.LanguageAppInterface
	HayvanSatisYerleri services.HayvanSatisYerleriAppInterface
	HayvanBilgisi      services.HayvanBilgisiAppInterface
	Kurban             services.KurbanAppInterface
	Kodemeler          services.OdemelerAppInterface
	Gruplar            services.GruplarAppInterface
	Kisiler            services.KisilerAppInterface
	WebArchive         services.WebArchiveAppInterface
	WebArchiveLink     services.WebArchiveLinksAppInterface
	Options            services.OptionsAppInterface
	Media              services.MediaAppInterface

	KioskSlider         services.KioskSliderAppInterface
	CategoriesKiosk     services.CategoriesKioskAppInterface
	CategoriesKioskJoin services.CategoriesKioskJoinAppInterface

	Branch               services.BranchAppInterface
	CategoriesBranchJoin services.CategoriesBranchJoinAppInterface

	DB *gorm.DB
}

//DbConnect initial
/*TODO: burada db verisi pointer olarak işaretlenecek oyle gidecek veri*/
func DbConnect() *gorm.DB {
	dbdriver := os.Getenv("DB_DRIVER")
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	gormAdvancedLogger := os.Getenv("GORM_ZAP_LOGGER")
	debug := os.Getenv("MODE")
	//	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPassword) //bu postresql

	//DBURL := "root:sel123C#@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local" //mysql
	var DBURL string

	if dbdriver == "mysql" {
		DBURL = dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local"
	} else if dbdriver == "postgres" {
		DBURL = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable ", dbHost, dbPort, dbUser, dbPassword, dbName) //Build connection string
	}

	// dsn := fmt.Sprintf("host=%s  user=%s password=%s dbname=%s sslmode=disable",
	// HOST, PORT, username, password, database)

	db, err := gorm.Open(dbdriver, DBURL)
	db.Set("gorm:table_options", "charset=utf8")
	// }

	// db, err := gorm.Open(dbdriver, DBURL)
	//nunlar gorm 2 versionunda prfexi falan var
	// db, err := gorm.Open(mysql.Open(DBURL), &gorm.Config{
	// 	NamingStrategy: schema.NamingStrategy{
	// 		TablePrefix:   "krbn_", // table name prefix, table for `User` would be `t_users`
	// 		SingularTable: true,    // use singular table name, table for `User` would be `user` with this option enabled
	// 	},
	// 	// Logger: gorm_logrus.New(),
	// })

	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	if debug == "DEBUG" && gormAdvancedLogger == "ENABLE" {
		db.LogMode(true)
		log := zap.NewExample()
		db.SetLogger(gormzap.New(log, gormzap.WithLevel(zap.DebugLevel)))
	} else if debug == "DEBUG" || debug == "TEST" && gormAdvancedLogger == "ENABLE" {
		db.LogMode(true)
	} else if debug == "RELEASE" {
		fmt.Println(debug)
		db.LogMode(false)
	}
	DB = db

	db.SingularTable(true)

	return db
}

//https://techinscribed.com/different-approaches-to-pass-database-connection-into-controllers-in-golang/

//RepositoriesInit initial
func RepositoriesInit(db *gorm.DB) (*Repositories, error) {

	return &Repositories{
		User: UserRepositoryInit(db),
		Post: PostRepositoryInit(db),

		Lang:                LanguageRepositoryInit(db),
		Kurban:              KurbanRepositoryInit(db),
		Kodemeler:           OdemelerRepositoryInit(db),
		Gruplar:             GruplarRepositoryInit(db),
		Kisiler:             KisilerRepositoryInit(db),
		HayvanSatisYerleri:  HayvanSatisYerleriRepositoryInit(db),
		HayvanBilgisi:       HayvanBilgisiRepositoryInit(db),
		WebArchive:          WebArchiveRepositoryInit(db),
		WebArchiveLink:      WebArchiveLinksRepositoryInit(db),
		Options:             OptionRepositoryInit(db),
		Media:               MediaRepositoryInit(db),
		KioskSlider:         KioskSliderRepositoryInit(db),
		CategoriesKiosk:     CategoriesKioskRepositoryInit(db),
		CategoriesKioskJoin: CatKioskJoinRepositoryInit(db),

		Branch:               BranchRepositoryInit(db),
		CategoriesBranchJoin: CategoriesBranchJoinRepositoryInit(db),

		DB: db,
	}, nil
}

//Close closes the  database connection
// func (s *Repositories) Close() error {
// 	return s.db.Close()
// }

//Automigrate This migrate all tables
func (s *Repositories) Automigrate() error {
	s.DB.AutoMigrate(&entity.SacrificeKurbanlar{}, &entity.SacrificeOdemeler{}, &entity.SacrificeGruplar{}, &entity.Kisiler{}, &entity.User{},
		&entity.SacrificeHayvanSatisYerleri{}, &entity.Languages{}, &entity.Modules{}, &entity.Notes{},

		&entity.CategoriesKiosk{}, &entity.CategoriesKioskJoin{}, &entity.KioskSlider{},

		&entity.SacrificeHayvanBilgisi{}, &entity.Options{}, &entity.Media{}, &entity.RootBranch{}, &entity.Branches{}, &entity.CategoriesBranch{})

	// &entity.Post{}, &entity.RootBranch{}, &entity.CategoryRootBranch{}, &entity.WebArchive{}, &entity.WebArchiveLinks{}
	s.DB.Model(&entity.SacrificeHayvanBilgisi{}).AddForeignKey("hayvan_satis_yerleri_id", "sacrifice_hayvan_satis_yerleri(id)", "CASCADE", "CASCADE") // one to one (one=hayvan_satis_yerleri) (one=HayvanBilgisi)

	s.DB.Model(&entity.SacrificeOdemeler{}).AddForeignKey("kurban_id", "sacrifice_kurbanlar(id)", "CASCADE", "CASCADE")           // one to many (one=kurbanlar) (many=odemeler)
	s.DB.Model(&entity.SacrificeKurbanlar{}).AddForeignKey("grup_id", "sacrifice_gruplar(id)", "CASCADE", "CASCADE")              // one to many (one=gruplar) (many=kurbanlar)
	return s.DB.Model(&entity.SacrificeKurbanlar{}).AddForeignKey("kisi_id", "sacrifice_kisiler(id)", "CASCADE", "CASCADE").Error // one to many (one=kisiler) (many=kurbanlar)
	// s.DB.Model(&entity.WebArchiveLinks{}).AddForeignKey("web_archive_id", "web_archive(id)", "CASCADE", "CASCADE").Error // one to many (one=web_archives) (many=WebArchiveLinks)

}
