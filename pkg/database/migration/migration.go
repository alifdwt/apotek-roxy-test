package migration

import (
	"apotek-roxy/models"

	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.TransaksiDetail{},
		&models.TransaksiHider{},
		&models.MasterBarang{},
	)

	if err != nil {
		panic(err)
	}

	return err
}

func DropTable(db *gorm.DB) error {
	err := db.Migrator().DropTable(
		&models.TransaksiHider{},
		&models.TransaksiDetail{},
		&models.MasterBarang{},
	)

	if err != nil {
		panic(err)
	}

	return err
}
