package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/lo9i/databoss/server/api/data"
	"github.com/lo9i/databoss/server/api/data/repositories"
	"go.uber.org/dig"
)

func DatabaseConnection() data.Database {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	database := &data.DatabaseImpl{Db: db}
	database.AutoMigrate()
	return database
}

func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(DatabaseConnection)
	//container.Provide(WalletTestConfiguration)
	//container.Provide(StellarTestConfiguration)

	container.Provide(data.NewUnitOfWork)
	container.Provide(repositories.NewCandidateRepository)
	//container.Provide(services.NewWalletService)
	//container.Provide(services.NewStellarKeyPairService)
	//container.Provide(services.NetStellarNetworkService)

	return container
}

func main() {
	fmt.Println("Starting")
	//container := BuildContainer()
	//err := container.Invoke(func(service domain.WalletService) {
	//	server := api.NewServer(service)
	//	err := server.Run()
	//	if err != nil {
	//		panic(err)
	//	}
	//})
	//if err != nil {
	//	panic(err)
	//}
	fmt.Println("Ending")
}
