package config

import (
	"fmt"
	"go-dummy-project/go-dummy/common"
	"strconv"
	"sync"

	"github.com/jinzhu/gorm"
)

var (
	host               = common.EnvMap["HOST"]
	port               = common.EnvMap["PORT"]
	user               = common.EnvMap["USER"]
	password           = common.EnvMap["PASSWORD"]
	dbname             = common.EnvMap["DB_NAME"]
	maxOpenConnections = common.EnvMap["MAX_OPEN_CONNECTION"]
	maxIdleConnections = common.EnvMap["MAX_IDLE_CONNECTION"]
)

var (
	createDBConnectionPool sync.Once
	db                     *gorm.DB
)

func createDBConnection() {

	mxOpenConnections, err := strconv.Atoi(maxOpenConnections)
	if err != nil {
		common.GetLogger().Fatalf("Error fetching maxOpenConnections for DB %s", err)
		return
	}

	mxIdleConnections, err := strconv.Atoi(maxIdleConnections)
	if err != nil {
		common.GetLogger().Fatalf("Error fetching maxIdleConnections for DB %s", err)
		return
	}

	portNo, err := strconv.Atoi(port)
	if err != nil {
		common.GetLogger().Fatalf("Error fetching port for DB %s", err)
		return
	}

	info := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, portNo, user, password, dbname)

	dbConn, err := gorm.Open("postgres", info)
	if err != nil {
		common.GetLogger().Fatalf("Error while connecting to postgres docker DB %s", err)
		return
	}

	dbConn.DB().SetMaxOpenConns(mxOpenConnections)
	dbConn.DB().SetMaxIdleConns(mxIdleConnections)
	dbConn.SetLogger(common.GetLogger())
	dbConn.BlockGlobalUpdate(true)
	db = dbConn
}

func initDBConnection() {
	createDBConnection()
}

func GetDBConnection() *gorm.DB {
	createDBConnectionPool.Do(initDBConnection)
	if db.Error != nil {
		common.GetLogger().Fatalf("Resetting the DB Error %s", db.Error)
		db.Error = nil
	}
	return db
}
