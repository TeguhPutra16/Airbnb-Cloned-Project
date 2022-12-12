package service

import (
	"be13/project/features/user"
	"be13/project/mocks"
	"errors"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	repo := new(mocks.UserRepo)
	t.Run("Succes Create User", func(t *testing.T) {
		inputRepo := user.CoreUser{Name: "alta", Email: "alta@mail.id", Password: "qwerty", Address: "Jakarta", Role: "User"}
		inputData := user.CoreUser{Name: "alta", Email: "alta@mail.id", Password: "qwerty", Address: "Jakarta"}
		repo.On("Create", inputRepo).Return(nil).Once()
		srv := New(repo)
		err := srv.Create(inputData)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Create, Duplicate", func(t *testing.T) {
		inputRepo := user.CoreUser{Name: "alta", Email: "alta@mail.id", Password: "qwerty", Address: "Jakarta", Role: "User"}
		inputData := user.CoreUser{Name: "alta", Email: "alta@mail.id", Password: "qwerty", Address: "Jakarta"}
		repo.On("Create", inputRepo).Return(errors.New("Failed Insert")).Once()
		srv := New(repo)
		err := srv.Create(inputData)
		assert.NotNil(t, err)
		assert.Equal(t, "GAGAL MENAMBAH DATA , QUERY ERROR", err.Error()) // samakan dengan di logic
		repo.AssertExpectations(t)
	})

	t.Run("Failed Create, Empty Name", func(t *testing.T) {
		// inputRepo := user.Core{Name: "alta", Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta", Role: "user"} tidak diperulkan kareana tidak dijalankan kareana sudah di return "Lihat Logic"
		inputData := user.CoreUser{Email: "alta@mail.id", Password: "qwerty", Address: "Jakarta"}
		// repo.On("InsertUser", inputRepo).Return(errors.New("Failed Insert,error querry")).Once()
		srv := New(repo)
		err := srv.Create(inputData)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

}
