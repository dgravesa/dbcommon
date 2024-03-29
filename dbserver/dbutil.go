package dbserver

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"strings"
	"text/template"
)

// StartupDBFromConfigFile connects to a database named dbName or creates it if it doesn't
// exist on the server specified by the configuration in file; the resulting database
// connection is returned on success.
func StartupDBFromConfigFile(cfgName, dbName string) (*sql.DB, error) {
	var config Config
	var err error

	if config, err = ReadConfigFile(cfgName); err != nil {
		return nil, err
	}

	return StartupDB(config, dbName)
}

// StartupDB connects to a database named dbName or creates it if it doesn't exist on the
// server specified by the configuration in file; the resulting database connection is
// returned on success.
func StartupDB(cfg Config, dbName string) (*sql.DB, error) {
	var srv *sql.DB
	var err error

	if srv, err = Connect(cfg); err != nil {
		return nil, err
	}

	if !DBExists(srv, dbName) {
		if err = CreateDB(srv, dbName); err != nil {
			return nil, err
		}
	}

	var db *sql.DB
	if db, err = ConnectToDB(cfg, dbName); err != nil {
		return nil, err
	}

	return db, err
}

// Connect creates a connection to a database server.
func Connect(config Config) (*sql.DB, error) {
	return ConnectToDB(config, "")
}

// ConnectToDB creates a connection to a database named dbName on a server pointed to by config.
func ConnectToDB(config Config, dbName string) (*sql.DB, error) {
	var passphrase string
	if content, err := ioutil.ReadFile(config.Passkey); err == nil {
		passphrase = ":" + string(content)[:len(content)]
	}

	var connectStr string
	if len(dbName) > 0 {
		connectStr = fmt.Sprintf("postgres://%s%s@%s:%d/%s?sslmode=disable",
			config.Username, passphrase, config.Host, config.Port, dbName)
	} else {
		connectStr = fmt.Sprintf("postgres://%s%s@%s:%d?sslmode=disable",
			config.Username, passphrase, config.Host, config.Port)
	}

	db, err := sql.Open("postgres", connectStr)

	if err == nil {
		// ensure connection
		err = db.Ping()
	}

	return db, err
}

// CreateDB creates a database named dbName on server srv.
func CreateDB(srv *sql.DB, dbName string) error {
	var sb strings.Builder
	tmpl, _ := template.New("").Parse(dbCreateDatabaseQuery)
	tmpl.Execute(&sb, dbName)
	query := sb.String()

	_, err := srv.Exec(query)
	return err
}

// DBExists returns true if a database named dbName exists on server srv, otherwise false.
func DBExists(srv *sql.DB, dbName string) bool {
	var sb strings.Builder
	tmpl, _ := template.New("").Parse(dbFindDatabaseByNameQuery)
	tmpl.Execute(&sb, dbName)
	query := sb.String()

	var dbFound bool
	srv.QueryRow(query).Scan(&dbFound)
	return dbFound
}
