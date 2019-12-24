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

func NewNosisApiConfiguration() domain.NosisApiConfiguration {
	return domain.NosisApiConfiguration{
		BaseURL: "https://ws01.nosis.com/api",
		User:     "62450",
		Group:     "1",
		Password: "179281",
	}
}
func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(NewDatabaseConfiguration)
	container.Provide(NewNosisApiConfiguration)
	container.Provide(data.NewDatabase)
	container.Provide(data.NewUnitOfWork)
	container.Provide(repositories.NewCandidateRepository)
	container.Provide(repositories.NewJobRepository)
	container.Provide(repositories.NewUserRepository)
	container.Provide(repositories.NewNosisUserRepository)
	container.Provide(services.NewUserService)
	container.Provide(services.NewCandidateService)
	container.Provide(services.NewJobService)
	container.Provide(services.NewNosisService)

	return container
}

func main() {
	fmt.Println("Starting")
	container := BuildContainer()
	err := container.Invoke(func(uService domain.UserService,
								cService domain.CandidateService,
								jService domain.JobService) {
		server := NewServer(uService, cService, jService)
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
