package main

import (
	"database/sql"
	"log"
	"pgsql/handler"
	"pgsql/router"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	db := ConnectedDB()
	UserRepo := handler.NewUserRepo(db)
	userRouter := router.NewUserRouter(UserRepo)

	r1 := gin.Default()
	r1.Use(gin.Logger(),gin.Recovery())
	r1.POST("/addUser", userRouter.CreateUser)
	r1.GET("/getAll", userRouter.GetAllUser)
	r1.GET("/getUserById/:id", userRouter.GetUserById)
	r1.PUT("/updateUser/:id", userRouter.UpdateUser)
	r1.DELETE("/deleteUser/:id", userRouter.DeleteUser)

	r1.Run(":5010")

}
func ConnectedDB() *sql.DB {
	dns := "postgres://postgres:123456@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("pgx", dns)
	if err != nil {
		log.Fatal("error while opening sql")
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("err while pinging database")
	}
	return db
}
