package main

func main() {
	// init config
	initConfig()

	// db conn
	// mysqlConn := initMysqlDb()

	// repo
	// repo := repository.NewUsersRepository(mysqlConn)

	// grpc
	// log.Fatal().Err(RunGrpcHandler(repo)).Msg("failed to listen")
}

func initConfig() {

}

// func initMysqlDb() *gorm.DB {

// 	return
// }

// func RunGrpcHandler(userRepo repository.UsersRepository) (err error) {
// }
