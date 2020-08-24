package service

import (
	"net"
	"otus_final_project/src/component/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRateLimiter_LoginAttempt(t *testing.T) {
	t.Run("Limits", func(t *testing.T) {
		login := "login"
		password := "pass"
		ip := net.ParseIP("127.0.0.1")

		limits := &config.Limits{
			Login:    10,
			Password: 100,
			IP:       2000,
		}

		lists := NewListRepositoryMock(t).IsInListMock.Return(false)

		service := NewRateLimiter(lists, lists, limits)

		for i := 0; i < 10; i++ {
			res := service.LoginAttempt(ip, login, password)
			assert.True(t, res)
		}

		res := service.LoginAttempt(ip, login, password)
		assert.False(t, res)
	})

	t.Run("Black and white lists checks", func(t *testing.T) {
		login := "login"
		password := "pass"
		ip := net.ParseIP("127.0.0.1")

		limits := &config.Limits{
			Login:    1,
			Password: 1,
			IP:       1,
		}

		whiteList := NewListRepositoryMock(t).IsInListMock.Return(true)
		blackList := NewListRepositoryMock(t).IsInListMock.Return(false)

		service := NewRateLimiter(whiteList, blackList, limits)

		for i := 0; i < 10; i++ {
			res := service.LoginAttempt(ip, login, password)
			assert.True(t, res)
		}

		limits = &config.Limits{
			Login:    100,
			Password: 100,
			IP:       100,
		}

		whiteList = NewListRepositoryMock(t).IsInListMock.Return(false)
		blackList = NewListRepositoryMock(t).IsInListMock.Return(true)

		service = NewRateLimiter(whiteList, blackList, limits)

		for i := 0; i < 10; i++ {
			res := service.LoginAttempt(ip, login, password)
			assert.False(t, res)
		}

	})

}
