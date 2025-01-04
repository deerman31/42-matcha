package users

import (
	"database/sql"
	"golang-echo/users/set"

	"github.com/labstack/echo/v4"
)

func UserRoutes(protected *echo.Group, db *sql.DB) {

	user := protected.Group("/users")

	setroute := user.Group("/set")
	setroute.POST("/user_info", set.InitSetUserInfo(db))

	configs := set.NewUpdateFieldConfigs()

	setroute.PATCH("/last-name", set.SetGeneric(db, configs.LastName))
	setroute.PATCH("/first-name", set.SetGeneric(db, configs.FirstName))
	setroute.PATCH("/self-intro", set.SetGeneric(db, configs.SelfIntro))
	setroute.PATCH("/area", set.SetGeneric(db, configs.Area))
	setroute.PATCH("/gender", set.SetGeneric(db, configs.Gender))
	setroute.PATCH("/sexuality", set.SetGeneric(db, configs.Sexuality))
	setroute.PATCH("/is-gps", set.SetGeneric(db, configs.IsGps))
	
	setroute.PATCH("/birthdate", set.SetGeneric(db, configs.BirthDate))//一応作ったが、生年月日は変更する必要がない気がする

	setroute.POST("/image1", set.SetImage(db, set.ImageOne))
	setroute.POST("/image2", set.SetImage(db, set.ImageTwo))
	setroute.POST("/image3", set.SetImage(db, set.ImageThree))
	setroute.POST("/image4", set.SetImage(db, set.ImageFour))
	setroute.POST("/image5", set.SetImage(db, set.ImageFive))

}
