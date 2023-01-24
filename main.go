package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"lblbc.cn/note/config"
	"lblbc.cn/note/controllers"
	"lblbc.cn/note/middleware"
	"lblbc.cn/note/repository"
	"lblbc.cn/note/services"
)

var (
	db              *gorm.DB                   = config.SetupDatabase()
	userRepository  repository.UserRepository  = repository.NewUserRepository(db)
	noteRepository  repository.NoteRepository  = repository.NewNoteRepository(db)
	jwtService      services.JWTService        = services.NewJWTService()
	userService     services.UserService       = services.NewUserService(userRepository)
	noteService     services.NoteService       = services.NewBookService(noteRepository)
	authService     services.LoginService      = services.NewAuthService(userRepository)
	loginController                            = controllers.NewLoginController(authService, jwtService)
	userController  controllers.UserController = controllers.NewUserController(userService, jwtService)
	noteController  controllers.NoteController = controllers.NewNoteController(noteService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("/user")
	{
		authRoutes.POST("/login", loginController.Login)
		authRoutes.POST("/register", loginController.Register)
	}

	noteRoutes := r.Group("note/notes", middleware.AuthorizeJWT(jwtService))
	{
		noteRoutes.GET("/", noteController.QueryNotes)
		noteRoutes.GET("/:id", noteController.QueryById)
		noteRoutes.POST("/", noteController.AddNote)
		noteRoutes.PUT("/:id", noteController.ModifyNote)
		noteRoutes.DELETE("/:id", noteController.DeleteNote)
	}

	r.Run(":8080")

}
