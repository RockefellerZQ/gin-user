package controller

import (
	"gin-user/common"
	"gin-user/dto"
	"gin-user/model"
	"gin-user/vo"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IUserController interface {
	RestController
	Login(*gin.Context)
}

type UserController struct {
	db *gorm.DB
}

func NewUserController() IUserController {
	db := model.GetDB()
	db.AutoMigrate(&model.User{})
	return &UserController{
		db: db,
	}
}

func (u *UserController) Create(ctx *gin.Context) {
	var userRequest vo.UserRequest
	if err := ctx.ShouldBind(&userRequest); err != nil {
		dto.Failed(ctx, 500, err.Error())
		return
	}
	userModel := userRequest.ToUserModel()
	if u.isUserExist(userModel) {
		dto.Failed(ctx, 500, "用户已存在")
		return
	}
	u.db.Create(userModel)
	if userModel.ID == 0 {
		dto.Failed(ctx, 500, "注册失败")
		return
	}
	token, err := common.GenerateToken(userModel.ID)
	if err != nil {
		dto.Failed(ctx, 500, "系统异常")
		return
	}
	dto.Success(ctx, gin.H{
		"token": token,
	})
}

func (u *UserController) Login(ctx *gin.Context) {
	var userRequest vo.UserRequest
	if err := ctx.ShouldBind(&userRequest); err != nil {
		dto.Failed(ctx, 500, err.Error())
		return
	}
	var user model.User
	if err := u.db.Where("email = ?", userRequest.Email).First(&user).Error; err != nil {
		dto.Failed(ctx, 500, err.Error())
		return
	}
	if user.Password != userRequest.MD5Password() {
		dto.Failed(ctx, 500, "登录失败")
		return
	}
	token, err := common.GenerateToken(user.ID)
	if err != nil {
		dto.Failed(ctx, 500, "登录失败")
		return
	}
	dto.Success(ctx, gin.H{
		"token": token,
	})
}

func (u *UserController) Show(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		dto.Failed(ctx, 500, "获取用户信息失败")
		return
	}
	dto.Success(ctx, dto.UserDTOFromUserModel(user.(model.User)))
}

func (u *UserController) Update(ctx *gin.Context) {
}

func (u *UserController) Delete(ctx *gin.Context) {
}

func (u *UserController) isUserExist(user *model.User) bool {
	var userModel model.User
	u.db.Where("email = ?", user.Email).First(&userModel)
	if userModel.ID != 0 {
		return true
	}
	return false
}




