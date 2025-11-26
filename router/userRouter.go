package router

import (
	"net/http"
	"pgsql/handler"
	"pgsql/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	repo *handler.UserRepo
}

func NewUserRouter(r *handler.UserRepo) *UserRouter {
	return &UserRouter{repo: r}
}

func (u *UserRouter) CreateUser(c *gin.Context) {
	var user model.UserModel
	c.BindJSON(&user)

	err := u.repo.Create(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "user Created"})

}
func (u *UserRouter) GetAllUser(c *gin.Context) {
	users, err := u.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
func (u *UserRouter) GetUserById(c *gin.Context) {
	id,err := strconv.Atoi(c.Param("id"))
	if err != nil{
		c.JSON(http.StatusNotFound , gin.H{"err" : err.Error()})
	}
	user, err := u.repo.GetById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (u *UserRouter) UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var user model.UserModel
	c.BindJSON(&user)
	user.ID = id
	err := u.repo.Updata(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(200, gin.H{"user Update": user})

}

func (u *UserRouter) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := u.repo.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}
