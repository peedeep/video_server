package dbops

import "database/sql"

func openConn() *sql.DB {
	return nil
}

func AddUserCredential(loginName string, pwd string) error {
	return nil
}

func GetUserCredentail(loginName string) (string, error) {
	return "", nil
}
