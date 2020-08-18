package repository

import (
	"net"
)

type Whitelist struct {
	*Base
}

func (repository *Whitelist) Add(ip *net.IPNet) (bool, error) {
	conn := repository.DB.Connection()

	_, err := conn.Exec("INSERT INTO blacklist VALUES ($1) ON CONFLICT DO NOTHING", ip)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (repository *Whitelist) Delete(ip *net.IPNet) (bool, error) {
	conn := repository.DB.Connection()

	_, err := conn.Exec("DELETE FROM blacklist WHERE address = $1", ip)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (repository *Whitelist) IsWhitelisted(ip net.IP) bool {
	conn := repository.DB.Connection()

	err := conn.QueryRow("SELECT address FROM whitelist WHERE address >> ($1)", ip).Scan()

	return err == nil
}
