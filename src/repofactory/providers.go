package repofactory

import "github.com/smakogon/dbs-samples/src/common"

var RepositoryProviders = map[string]common.RepositoryProvider{
	"postgres": repositories.NewPostgresRepository,
}
