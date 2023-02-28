package out

import "github.com/hugovallada/golang-hexagonal-architecture/application/core/entity"

type PersistUserInDatabaseOutputPort interface {
	Execute(user entity.User) error
}
