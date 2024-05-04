package endpoint

import "minang-kos-service/repository"

func getUserRepository() repository.UserRepository {
	return repository.NewUserRepository()
}
