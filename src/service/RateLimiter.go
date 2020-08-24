package service

import (
	"net"
	"otus_final_project/src/component/config"
	"otus_final_project/src/repository"

	"github.com/kevinms/leakybucket-go"
)

type RateLimiter struct {
	loginBucketsCollector BucketsCollector
	passBucketsCollector  BucketsCollector
	ipBucketsCollector    BucketsCollector

	repWhitelist repository.ListRepository
	repBlacklist repository.ListRepository
}

func NewRateLimiter(whitelist repository.ListRepository, blacklist repository.ListRepository, limits *config.Limits) *RateLimiter {
	return &RateLimiter{
		loginBucketsCollector: NewBucketsCollector(limits.Login),
		passBucketsCollector:  NewBucketsCollector(limits.Password),
		ipBucketsCollector:    NewBucketsCollector(limits.IP),
		repWhitelist:          whitelist,
		repBlacklist:          blacklist,
	}
}

func (cont *RateLimiter) LoginAttempt(ip net.IP, login string, password string) bool {
	// TODO добавть кеширование
	if cont.repBlacklist.IsInList(ip) {
		return false
	}

	if cont.repWhitelist.IsInList(ip) {
		return true
	}

	if !cont.loginBucketsCollector.Add(login) {
		return false
	}

	if !cont.passBucketsCollector.Add(password) {
		return false
	}

	if !cont.ipBucketsCollector.Add(ip.String()) {
		return false
	}

	return true
}

type BucketsCollector struct {
	collector     *leakybucket.Collector
	elementWeight int64
}

func NewBucketsCollector(limitPerMinute int) BucketsCollector {
	clt := leakybucket.NewCollector(1000, 60000, true)

	return BucketsCollector{
		collector:     clt,
		elementWeight: int64(60000 / limitPerMinute),
	}
}

func (bc *BucketsCollector) Add(key string) bool {
	n := bc.collector.Add(key, bc.elementWeight)

	return n == bc.elementWeight
}
