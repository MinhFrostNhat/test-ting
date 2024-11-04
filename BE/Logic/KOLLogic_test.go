package Logic

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
	"wan-api-kol-event/DTO"
	"wan-api-kol-event/Initializers"
	"wan-api-kol-event/Models"
)

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&Models.Kol{})
	return db
}

func TestGetKolLogic(t *testing.T) {
	db := setupTestDB()
	Initializers.DB = db

	// Prepare test data
	kols := []Models.Kol{
		{KolID: 1, UserProfileID: 1, Language: "EN"},
		{KolID: 2, UserProfileID: 2, Language: "EN"},
		{KolID: 3, UserProfileID: 3, Language: "FR"},
	}

	for _, kol := range kols {
		if err := db.Create(&kol).Error; err != nil {
			t.Fatalf("Failed to seed data: %v", err)
		}
	}

	tests := []struct {
		name         string
		pageIndex    int64
		pageSize     int64
		expectedKols []*DTO.KolDTO
		expectedErr  error
	}{
		{
			name:      "First page with two results",
			pageIndex: 1,
			pageSize:  2,
			expectedKols: []*DTO.KolDTO{
				{KolID: 1, UserProfileID: 1, Language: "EN"},
				{KolID: 2, UserProfileID: 2, Language: "EN"},
			},
			expectedErr: nil,
		},
		{
			name:      "Second page with one result",
			pageIndex: 2,
			pageSize:  2,
			expectedKols: []*DTO.KolDTO{
				{KolID: 3, UserProfileID: 3, Language: "FR"},
			},
			expectedErr: nil,
		},
		{
			name:         "Page out of range",
			pageIndex:    3,
			pageSize:     2,
			expectedKols: []*DTO.KolDTO{},
			expectedErr:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kolsDTO, err := GetKolLogic(tt.pageIndex, tt.pageSize)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
				return
			}
			assert.NoError(t, err)

			assert.Equal(t, len(tt.expectedKols), len(kolsDTO))
			for i, expectedKol := range tt.expectedKols {
				assert.Equal(t, expectedKol.KolID, kolsDTO[i].KolID)
				assert.Equal(t, expectedKol.UserProfileID, kolsDTO[i].UserProfileID)
				assert.Equal(t, expectedKol.Language, kolsDTO[i].Language)
			}
		})
	}
}
