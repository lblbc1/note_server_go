package controllers

/*
厦门大学计算机专业 | 前华为工程师
专注《零基础学编程系列》  http://lblbc.cn/blog
包含：Java | 安卓 | 前端 | Flutter | iOS | 小程序 | 鸿蒙
公众号：蓝不蓝编程
*/
import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"lblbc.cn/note/dto"
	"lblbc.cn/note/entity"
	"lblbc.cn/note/helper"
	"lblbc.cn/note/services"
)

type NoteController interface {
	QueryNotes(c *gin.Context)
	QueryById(c *gin.Context)
	AddNote(c *gin.Context)
	ModifyNote(c *gin.Context)
	DeleteNote(c *gin.Context)
}

type noteController struct {
	noteService services.NoteService
	jwtService  services.JWTService
}

func NewNoteController(noteService services.NoteService, jwtServ services.JWTService) NoteController {
	return &noteController{noteService: noteService, jwtService: jwtServ}
}

func (c *noteController) QueryById(ctx *gin.Context) {
	noteId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		response := helper.ErrorsResponse(-1, err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	var note entity.Note = c.noteService.QueryById(noteId)
	if note == (entity.Note{}) {
		response := helper.ErrorsResponse(-1, "Note Not Found", helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusNotFound, response)
		return
	} else {
		response := helper.SuccessResponse(0, "", note)
		ctx.JSON(http.StatusOK, response)
	}
}

func (c *noteController) QueryNotes(ctx *gin.Context) {
	var note []entity.Note = c.noteService.QueryNotes()
	response := helper.SuccessResponse(0, "", note)
	ctx.JSON(http.StatusOK, response)
}

func (c *noteController) AddNote(ctx *gin.Context) {
	var noteCreateDTO dto.NoteCreateDTORequest
	errDTO := ctx.ShouldBind(&noteCreateDTO)
	if errDTO != nil {
		response := helper.ErrorsResponse(-1, errDTO.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authHeader := ctx.GetHeader("Authorization")
	userID := c.getUserIDByToken(authHeader)
	id, err := strconv.ParseUint(userID, 10, 64)
	if err == nil {
		noteCreateDTO.UserID = id
	}
	result := c.noteService.AddNote(noteCreateDTO)
	response := helper.SuccessResponse(0, "", result)
	ctx.JSON(http.StatusCreated, response)
}

func (c *noteController) ModifyNote(ctx *gin.Context) {
	noteId, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	var noteUpdateDTO dto.NoteUpdateDTORequest
	errDTO := ctx.ShouldBind(&noteUpdateDTO)
	noteUpdateDTO.ID = noteId
	if errDTO != nil {
		result := helper.ErrorsResponse(-1, errDTO.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, result)
		return
	}
	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	id, errID := strconv.ParseUint(userID, 10, 64)
	if errID == nil {
		noteUpdateDTO.UserID = id
	}
	result := c.noteService.ModifyNote(noteUpdateDTO)
	response := helper.SuccessResponse(0, "Update Data Note", result)
	ctx.JSON(http.StatusOK, response)
}

func (c *noteController) DeleteNote(ctx *gin.Context) {
	var note entity.Note
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		response := helper.ErrorsResponse(-1, err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	note.ID = id
	authHeader := ctx.GetHeader("Authorization")
	_, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	c.noteService.DeleteNote(note)
	response := helper.SuccessResponse(0, "", "")
	ctx.JSON(http.StatusOK, response)
}

func (c *noteController) getUserIDByToken(token string) string {
	myToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := myToken.Claims.(jwt.MapClaims)
	return claims["user_id"].(string)
}
