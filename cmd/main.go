package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hilmanXcode/DiskusiKan/internal/configs"
	"github.com/hilmanXcode/DiskusiKan/internal/handler/memberships"
	membershipRepo "github.com/hilmanXcode/DiskusiKan/internal/repository/memberships"
	membershipSvc "github.com/hilmanXcode/DiskusiKan/internal/service/memberships"
	"github.com/hilmanXcode/DiskusiKan/pkg/internalsql"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)
	err := configs.Init(
		configs.WithConfigFolder(
			[]string{"./internal/configs/"},
		),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatal("Gagal inisiasi config ", err)
	}

	cfg = configs.Get()

	db, err := internalsql.Connect(cfg.Database.DataSourceName)

	if err != nil {
		log.Fatal("Gagal inisiasi database")
	}

	membershipRepo := membershipRepo.NewRepository(db)
	membershipService := membershipSvc.NewService(cfg, membershipRepo)

	membershipHandler := memberships.NewHandler(r, membershipService)
	membershipHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
