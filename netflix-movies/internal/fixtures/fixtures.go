package fixtures

import (
	"database/sql"
	"fmt"
	"netflix-movies/internal/config"
	"netflix-movies/pkg/postgres"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/go-pg/pg/v10"
	"github.com/go-testfixtures/testfixtures/v3"
	_ "github.com/lib/pq"
)

var (
	p           *pg.DB
	db          *sql.DB
	Cfg         *config.Config
	fixtures    *testfixtures.Loader
	err         error
	onceFixture sync.Once
)

func init() {
	postgres.Load(CFG)

	db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		CFG.DB.Host, CFG.DB.Port, CFG.DB.User, CFG.DB.Password, CFG.DB.Name))
	if err != nil {
		panic(err)
	}

	_, filename, _, _ := runtime.Caller(0)
	path := filepath.Dir(filename)

	fixtures, err = testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("postgres"),
		testfixtures.Directory(filepath.Join(path, "data")),
	)
	if err != nil {
		panic(err)
	}

	p = postgres.GetDB()
}

func GetDB() *pg.DB {
	return p
}

func PrepareFixtures() {
	onceFixture.Do(func() {
		if err := fixtures.Load(); err != nil {
			panic(err)
		}
	})
}
