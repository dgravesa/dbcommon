# dbcommon

[![CircleCI](https://circleci.com/gh/dgravesa/dbcommon.svg?style=svg)](https://circleci.com/gh/dgravesa/dbcommon)

*dbcommon* provides some helper utilities for connecting to a Postgres database.

### StartupDBFromConfigFile

This method takes the config file name `cfgName` and the name of the database `dbName`.
The config file is expected to be YAML format describing the connection to the database server.
If a database with the given name exists on the server, a connection to it is returned.
If no database exists with the given name, it is created, and then a connection to it is returned.
