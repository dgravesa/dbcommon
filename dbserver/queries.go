package dbcommon

const dbFindDatabaseByNameQuery string = `
SELECT EXISTS(
	SELECT datname
	FROM pg_catalog.pg_database
	WHERE lower(datname) = lower('{{.}}')
);`

const dbCreateDatabaseQuery string = `
CREATE DATABASE {{.}};`
