package services

/*
厦门大学计算机专业 | 前华为工程师
专注《零基础学编程系列》  http://lblbc.cn/blog
包含：Java | 安卓 | 前端 | Flutter | iOS | 小程序 | 鸿蒙
公众号：蓝不蓝编程
*/
import (
	"log"

	"github.com/mashingan/smapping"
	"lblbc.cn/note/dto"
	"lblbc.cn/note/entity"
	"lblbc.cn/note/repository"
)

type NoteService interface {
	AddNote(b dto.NoteCreateDTORequest) entity.Note
	ModifyNote(b dto.NoteUpdateDTORequest) entity.Note
	DeleteNote(b entity.Note)
	QueryById(noteId uint64) entity.Note
	QueryNotes() []entity.Note
}

type noteService struct {
	noteRepository repository.NoteRepository
}

func NewBookService(bookRepo repository.NoteRepository) NoteService {
	return &noteService{noteRepository: bookRepo}
}

func (s *noteService) QueryById(noteId uint64) entity.Note {
	return s.noteRepository.QueryByID(noteId)
}

func (s *noteService) QueryNotes() []entity.Note {
	return s.noteRepository.QueryNotes()
}

func (s *noteService) AddNote(b dto.NoteCreateDTORequest) entity.Note {
	note := entity.Note{}
	err := smapping.FillStruct(&note, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed to map fields %v: ", err)
	}
	result := s.noteRepository.AddNote(note)
	return result
}

func (s *noteService) ModifyNote(b dto.NoteUpdateDTORequest) entity.Note {
	note := entity.Note{}
	err := smapping.FillStruct(&note, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed to map fields %v: ", err)
	}
	result := s.noteRepository.ModifyNote(note)
	return result
}

func (s *noteService) DeleteNote(b entity.Note) {
	s.noteRepository.DeleteNote(b) // delete book
}
