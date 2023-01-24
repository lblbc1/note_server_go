package repository

/*
厦门大学计算机专业 | 前华为工程师
专注《零基础学编程系列》  http://lblbc.cn/blog
包含：Java | 安卓 | 前端 | Flutter | iOS | 小程序 | 鸿蒙
公众号：蓝不蓝编程
*/
import (
	"gorm.io/gorm"
	"lblbc.cn/note/entity"
)

type NoteRepository interface {
	QueryNotes() []entity.Note
	QueryByID(bookID uint64) entity.Note
	AddNote(b entity.Note) entity.Note
	ModifyNote(b entity.Note) entity.Note
	DeleteNote(b entity.Note)
}

type dbConnection struct {
	connection *gorm.DB
}

func NewNoteRepository(connection *gorm.DB) NoteRepository {
	return &dbConnection{connection: connection}
}

func (db *dbConnection) QueryByID(bookID uint64) entity.Note {
	var book entity.Note                              // create variable book
	db.connection.Preload("User").Find(&book, bookID) // get data book from bookID and preload user from book
	return book                                       // return book
}

func (db *dbConnection) QueryNotes() []entity.Note {
	var notes []entity.Note
	db.connection.Preload("User").Find(&notes)
	return notes
}

func (db *dbConnection) AddNote(b entity.Note) entity.Note {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *dbConnection) ModifyNote(b entity.Note) entity.Note {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *dbConnection) DeleteNote(b entity.Note) {
	db.connection.Delete(&b)
}
