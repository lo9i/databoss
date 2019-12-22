package main

import (
	"fmt"
	"github.com/lo9i/databoss/server/data"
	"github.com/lo9i/databoss/server/data/repositories"
	"github.com/lo9i/databoss/server/domain"
	"github.com/lo9i/databoss/server/services"
	"go.uber.org/dig"
)

func NewDatabaseConfiguration() data.DatabaseConfiguration {
	return data.DatabaseConfiguration{
		Driver:   "mysql",
		User:     "databoss_user",
		Password: "anibal",
		Port:     "3306",
		Host:     "localhost",
		Name:     "databoss",
	}
}

func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(NewDatabaseConfiguration)
	container.Provide(data.NewDatabase)
	container.Provide(data.NewUnitOfWork)
	container.Provide(repositories.NewCandidateRepository)
	container.Provide(repositories.NewUserRepository)
	container.Provide(services.NewUserService)
	container.Provide(services.NewCandidateService)
	//container.Provide(services.NewStellarKeyPairService)
	//container.Provide(services.NetStellarNetworkService)

	return container
}

func main() {
	fmt.Println("Starting")
	container := BuildContainer()
	err := container.Invoke(func(uService domain.UserService, cService domain.CandidateService) {
		server := NewServer(uService, cService)
		server.Run("0.0.0.0:8080")
		//err := server.Run("")
		//if err != nil {
		//	panic(err)
		//}
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Ending")
}
