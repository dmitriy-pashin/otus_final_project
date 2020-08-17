package repository

import (
	"net"
)

type Blacklist struct {
	*Base
}

func (repository *Blacklist) Add(ip *net.IPNet) (bool, error) {
	conn := repository.DB.Connection()

	_, err := conn.Exec("INSERT INTO blacklist VALUES ($1) ON CONFLICT DO NOTHING", ip)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (repository *Blacklist) Delete(ip *net.IPNet) (bool, error) {
	conn := repository.DB.Connection()

	_, err := conn.Exec("DELETE FROM blacklist WHERE address = $1", ip)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (repository *Blacklist) IsBlacklisted(ip net.IP) bool {
	conn := repository.DB.Connection()
	err := conn.QueryRow("SELECT address FROM blacklist WHERE address >> ($1)", ip).Scan()

	return err == nil
}
