package service

import (
	"github.com/kevinms/leakybucket-go"
	"net"
	"otus_final_project/src/component/config"
	"otus_final_project/src/repository"
)

type Login struct {
	loginBucketsCollector BucketsCollector
	passBucketsCollector  BucketsCollector
	ipBucketsCollector    BucketsCollector
	repWhitelist          *repository.Whitelist
	repBlacklist          *repository.Blacklist
}

func NewLogin(whitelist *repository.Whitelist, blacklist *repository.Blacklist, limits *config.Limits) *Login {
	return &Login{
		loginBucketsCollector: NewBucketsCollector(limits.Login),
		passBucketsCollector:  NewBucketsCollector(limits.Password),
		ipBucketsCollector:    NewBucketsCollector(limits.IP),
		repWhitelist:          whitelist,
		repBlacklist:          blacklist,
	}
}

func (cont *Login) Login(ip net.IP, login string, password string) bool {
	//TODO добавть кеширование
	if cont.repBlacklist.IsBlacklisted(ip) {
		return false
	}

	if cont.repWhitelist.IsWhitelisted(ip) {
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

	//fmt.Print("N = ", n, " Weight = ", bc.elementWeight, " CHECK: = ", n == bc.elementWeight, " COUNT: = ",bc.collector.Count(key), "\n")
	return n == bc.elementWeight
}
