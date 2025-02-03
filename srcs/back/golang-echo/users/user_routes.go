package users

import (
	"database/sql"
	"golang-echo/users/get"
	"golang-echo/users/set"

	"github.com/labstack/echo/v4"
)

func UserRoutes(protected *echo.Group, db *sql.DB) {

	user := protected.Group("/users")

	setter := user.Group("/set")
	setter.POST("/user-info", set.InitSetUserInfo(db))

	setter.PATCH("/username", set.SetUserName(db))
	setter.PATCH("/email", set.SetEmail(db))

	configs := set.NewUpdateFieldConfigs()

	setter.PATCH("/last-name", set.SetGeneric(db, configs.LastName))
	setter.PATCH("/first-name", set.SetGeneric(db, configs.FirstName))
	setter.PATCH("/self-intro", set.SetGeneric(db, configs.SelfIntro))
	setter.PATCH("/area", set.SetGeneric(db, configs.Area))
	setter.PATCH("/gender", set.SetGeneric(db, configs.Gender))
	setter.PATCH("/sexuality", set.SetGeneric(db, configs.Sexuality))

	// setter.PATCH("/is-gps", set.SetGeneric(db, configs.IsGps))

	setter.PATCH("/birthdate", set.SetGeneric(db, configs.BirthDate)) //一応作ったが、生年月日は変更する必要がない気がする

	setter.POST("/image1", set.SetImage(db, set.ImageOne))
	setter.POST("/image2", set.SetImage(db, set.ImageTwo))
	setter.POST("/image3", set.SetImage(db, set.ImageThree))
	setter.POST("/image4", set.SetImage(db, set.ImageFour))
	setter.POST("/image5", set.SetImage(db, set.ImageFive))

	getter := user.Group("/get")
	getter.GET("/image1", get.GetImage(db, set.ImageOne))
	getter.GET("/image2", get.GetImage(db, set.ImageTwo))
	getter.GET("/image3", get.GetImage(db, set.ImageThree))
	getter.GET("/image4", get.GetImage(db, set.ImageFour))
	getter.GET("/image5", get.GetImage(db, set.ImageFive))


	// ルート設定をまとめて定義
	routes := []routeConfig{
		{path: "username", tableName: "users", fieldName: "username", where: "id"},
		{path: "email", tableName: "users", fieldName: "email", where: "id"},
		{path: "lastname", tableName: "user_info", fieldName: "lastname", where: "user_id"},
		{path: "firstname", tableName: "user_info", fieldName: "firstname", where: "user_id"},
		{path: "birthdate", tableName: "user_info", fieldName: "birthdate", where: "user_id"},
		{path: "gender", tableName: "user_info", fieldName: "gender", where: "user_id"},
		{path: "sexuality", tableName: "user_info", fieldName: "sexuality", where: "user_id"},
		{path: "area", tableName: "user_info", fieldName: "area", where: "user_id"},
		{path: "self_intro", tableName: "user_info", fieldName: "self_intro", where: "user_id"},
	}

	// ルートを一括で設定
	for _, route := range routes {
		params := get.QueryParams{
			TableName: route.tableName,
			FieldName: route.fieldName,
			Where:     route.where,
		}
		getter.GET("/"+route.path, get.GetGeneric(db, params))
	}

}

type routeConfig struct {
	path      string
	tableName string
	fieldName string
	where     string
}
