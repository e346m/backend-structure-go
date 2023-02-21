package di

import (
	"database/sql"

	bInMemory "github.com/e346m/bsg/internal/booking/adapters/memory"
	bPSQL "github.com/e346m/bsg/internal/booking/adapters/psql"
	bHttp "github.com/e346m/bsg/internal/booking/ports/http"
	bUsecase "github.com/e346m/bsg/internal/booking/usecases"
)

func NewbookingHTTP(db *sql.DB) *bHttp.Handler {
	usecase := newBookingSet(db)
	handler := bHttp.NewHandler(usecase)
	return handler
}

func newBookingSet(db *sql.DB) *bUsecase.Usecase {
	repositoryKeeper := bPSQL.NewPSQLClient(db)
	usecase := bUsecase.NewUsecase(repositoryKeeper)
	return usecase
}

// if you want to replace adapter
// you can just generate repo and set to usecase
// because InMemoryClient meet the interface defined inside usecase
func newBookingSetWithInmemory() *bUsecase.Usecase {
	repositoryKeeper := bInMemory.NewInMemoryClient()
	usecase := bUsecase.NewUsecase(repositoryKeeper)
	return usecase
}
