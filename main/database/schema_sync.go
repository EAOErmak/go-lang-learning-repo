package database

import (
	"go-learn/main/models"

	"gorm.io/gorm"
)

func syncDiarySchema(db *gorm.DB) error {
	if err := dropLegacyDiaryEntrySchema(db); err != nil {
		return err
	}

	if err := dropLegacyDictionarySchema(db); err != nil {
		return err
	}

	if err := dropLegacyTables(db); err != nil {
		return err
	}

	return nil
}

func dropLegacyDiaryEntrySchema(db *gorm.DB) error {
	if db.Migrator().HasIndex(&models.DiaryEntry{}, "idx_diary_user") {
		if err := db.Migrator().DropIndex(&models.DiaryEntry{}, "idx_diary_user"); err != nil {
			return err
		}
	}

	if db.Migrator().HasColumn(&models.DiaryEntry{}, "user_id") {
		if err := db.Migrator().DropColumn(&models.DiaryEntry{}, "user_id"); err != nil {
			return err
		}
	}

	return nil
}

func dropLegacyDictionarySchema(db *gorm.DB) error {
	legacyColumns := []string{
		"parent_id",
		"chart_type",
		"active",
		"allowed_role",
		"entry_field_config_id",
	}

	for _, column := range legacyColumns {
		if !db.Migrator().HasColumn(&models.DictionaryItem{}, column) {
			continue
		}

		if err := db.Migrator().DropColumn(&models.DictionaryItem{}, column); err != nil {
			return err
		}
	}

	return nil
}

func dropLegacyTables(db *gorm.DB) error {
	legacyTables := []string{
		"users",
		"entry_field_config",
		"tag",
		"diary_entry_tag",
	}

	for _, table := range legacyTables {
		if !db.Migrator().HasTable(table) {
			continue
		}

		if err := db.Migrator().DropTable(table); err != nil {
			return err
		}
	}

	return nil
}
