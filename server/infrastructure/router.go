package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/kindaidensan/UMR/interfaces/controllers"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	ldapHandler := NewLdapHandler()
	redisHandler := NewRedisHandler()
	mailHandler := NewMailHandler()
	accountController := controllers.NewAccountController(ldapHandler, redisHandler, mailHandler)
	settingController := controllers.NewSettingController(redisHandler)
	authenticationController := controllers.NewAuthenticationController(redisHandler)


	router.POST("/admin/create_register_form", func(c *gin.Context) {settingController.CreateRegisterForm(c)})
	router.POST("/admin/get_register_form", func(c *gin.Context) {settingController.GetRegisterForm(c)})
	router.POST("/admin/get_all_accounts", func(c *gin.Context) {accountController.GetAllAccounts(c)})

	router.POST("/register", func(c *gin.Context) {
		err := authenticationController.AuthenticationFormToken(c)
		if err != nil {
			return
		}
		accountController.TemporaryCreate(c)
	})

	router.POST("/authentication", func(c *gin.Context) {accountController.AuthenticationCreate(c)})

	Router = router
}