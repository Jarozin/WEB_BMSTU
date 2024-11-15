package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"project/internal/pkg/models"

	"github.com/jackc/pgx/v5/pgxpool"

	authHandler "project/internal/pkg/auth/delivery/http"
	driverHandler "project/internal/pkg/driver/delivery/http"
	featuredHandler "project/internal/pkg/featured/delivery/http"
	grandPrixHandler "project/internal/pkg/grand_prix/delivery/http"
	qualHandler "project/internal/pkg/qual_result/delivery/http"
	raceHandler "project/internal/pkg/race_result/delivery/http"
	teamHandler "project/internal/pkg/team/delivery/http"
	trackHandler "project/internal/pkg/track/delivery/http"
	userHandler "project/internal/pkg/user/delivery/http"

	driverRepository "project/internal/pkg/driver/repository/postgresql"
	featuredRepository "project/internal/pkg/featured/postgresql"
	grandPrixRepository "project/internal/pkg/grand_prix/repository/postgresql"
	qualRepository "project/internal/pkg/qual_result/repository/postgresql"
	raceRepository "project/internal/pkg/race_result/repository/postgresql"
	teamRepository "project/internal/pkg/team/repository/postgresql"
	trackRepository "project/internal/pkg/track/repository/postgresql"
	userRepository "project/internal/pkg/user/repository/postgresql"

	driverUsecase "project/internal/pkg/driver/usecase"
	featuredUsecase "project/internal/pkg/featured/usecase"
	grandPrixUsecase "project/internal/pkg/grand_prix/usecase"
	qualUsecase "project/internal/pkg/qual_result/usecase"
	raceUsecase "project/internal/pkg/race_result/usecase"
	teamUsecase "project/internal/pkg/team/usecase"
	trackUsecase "project/internal/pkg/track/usecase"
	userUsecase "project/internal/pkg/user/usecase"

	"project/internal/app/middleware"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// @title FormulOne Web-Server
// @version 1.0
// @description API Server for F1 Grand-Prix info
// @termsOfService  http://swagger.io/terms/

// @host localhost:5259
// @BasePath /
func main() {
	// params := fmt.Sprintf("user=postgresql dbname=postgresql password=postgresql host=%s port=%s sslmode=disable", os.Getenv("PG_HOST"), os.Getenv("PG_PORT"))
	params := "user=postgres dbname=formula1 password=7303_486 host=localhost port=5432 sslmode=disable"

	db, err := sqlx.Connect("postgres", params)
	if err != nil {
		log.Fatal(err)
	}
	conf, err := pgxpool.ParseConfig("postgres://postgresql:postgresql@" + os.Getenv("PG_HOST") + ":" + os.Getenv("PG_PORT") + "/postgresql?" + "pool_max_conns=100")
	if err != nil {
		log.Fatal(err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), conf)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
	defer db.Close()

	var driverRepo models.DriverRepositoryI
	// if os.Getenv("MODE") == "SQLX" {
	driverRepo = driverRepository.NewPsqlDriverRepository(db)
	// } else {
	// 	driverRepo = pgx.NewPsqlDriverRepositoryPGX(pool)

	// }

	//params := fmt.Sprintf("user=postgresql dbname=postgresql password=postgresql host=%s port=%s sslmode=disable", os.Getenv("PG_HOST"), os.Getenv("PG_PORT"))
	//db, err := sqlx.Connect("postgres", params)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer db.Close()

	//driverRepo := driverRepository.NewPsqlDriverRepository(db)
	teamRepo := teamRepository.NewPsqlTeamRepository(db)
	trackRepo := trackRepository.NewPsqlTrackRepository(db)
	gpRepo := grandPrixRepository.NewPsqlGPRepository(db)
	raceRepo := raceRepository.NewPsqlRaceResultRepository(db)
	qualRepo := qualRepository.NewPsqlQualResultRepository(db)
	userRepo := userRepository.NewPsqlUserRepository(db)
	featuredRepo := featuredRepository.NewPsqlGPRepository(db)

	driverUcase := driverUsecase.NewDriverUsecase(driverRepo)
	teamUcase := teamUsecase.NewTeamUsecase(teamRepo)
	trackUcase := trackUsecase.NewTrackUsecase(trackRepo)
	gpUcase := grandPrixUsecase.NewGrandPrixUsecase(gpRepo)
	raceUcase := raceUsecase.NewRaceResultUsecase(raceRepo)
	qualUcase := qualUsecase.NewQualResultUsecase(qualRepo)
	userUcase := userUsecase.NewUserUsecase(userRepo)
	featuredUcase := featuredUsecase.NewFeaturedUsecase(featuredRepo)

	m := mux.NewRouter()

	driverHandler.NewDriverHandler(m, driverUcase, raceUcase)
	teamHandler.NewTeamHandler(m, teamUcase)
	trackHandler.NewTrackHandler(m, trackUcase)
	grandPrixHandler.NewDriverHandler(m, gpUcase, raceUcase, qualUcase)
	raceHandler.NewRaceResultHandler(m, raceUcase)
	qualHandler.NewQualResultHandler(m, qualUcase)
	authHandler.NewAuthHandler(m, userUcase)
	userHandler.NewUserHandler(m, userUcase)
	featuredHandler.NewFeaturedHandler(m, featuredUcase)

	// mt := metrics.NewPrometheusMetrics("api")
	// err = mt.SetupMetrics()
	if err != nil {
		os.Exit(1)
	}

	// m.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler).Methods(http.MethodGet)
	//m.HandleFunc("/swagger/").Handler(httpSwagger.WrapHandler).Methods("GET")
	mMiddleware := middleware.LogMiddleware(m)
	// pm := middleware.PromMetrics(mMiddleware, mt)

	// go metrics.ServePrometheusHTTP("0.0.0.0:9001")

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", handlers.CORS()(mMiddleware))
}
