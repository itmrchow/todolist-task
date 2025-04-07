package main

import (
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"

	"itmrchow/todolist-task/infra"
)

func main() {
	// init config
	initConfig()

	// db conn
	// mysqlConn := initMysqlDb()
	initMysqlDb()

	// repo
	// repo := repository.NewUsersRepository(mysqlConn)

	// grpc
	// log.Fatal().Err(RunGrpcHandler(repo)).Msg("failed to listen")
}

func initConfig() {
	err := infra.InitConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to init config")
	}

	log.Info().Msg("config loaded")
}

func initMysqlDb() *gorm.DB {
	db, err := infra.InitMysqlDb()

	if err != nil {
		log.Fatal().Err(err).Msg("failed to init mysql db")
	}

	log.Info().Msg("mysql db connected")

	return db
}

// func RunGrpcHandler(userRepo repository.UsersRepository) (err error) {
// }
