package Logic

import (
	"log"
	"math/rand"
	"strconv"
	"time"
	"wan-api-kol-event/Initializers"
	"wan-api-kol-event/Models"
)

func GenerateDummyKOLs(count int) []Models.Kol {
	dummies := make([]Models.Kol, count)
	languages := []string{"en", "vn", "es", "fr"}
	educations := []string{"Bachelor's in Computer Science", "Bachelor's in Marketing", "Master's in Business Administration", "PhD in Economics"}

	for i := 0; i < count; i++ {
		id := int64(100 + i)
		dummies[i] = Models.Kol{
			KolID:                id,
			UserProfileID:        2000 + id,
			Language:             languages[rand.Intn(len(languages))],
			Education:            educations[rand.Intn(len(educations))],
			ExpectedSalary:       int64(rand.Intn(80000) + 30000),
			ExpectedSalaryEnable: rand.Intn(2) == 1,
			ChannelSettingTypeID: int64(rand.Intn(5) + 1),
			IDFrontURL:           "https://example.com/id-front-" + strconv.FormatInt(id, 10) + ".jpg",
			IDBackURL:            "https://example.com/id-back-" + strconv.FormatInt(id, 10) + ".jpg",
			PortraitURL:          "https://example.com/portrait-" + strconv.FormatInt(id, 10) + ".jpg",
			RewardID:             300 + id,
			PaymentMethodID:      400 + id,
			TestimonialsID:       500 + id,
			VerificationStatus:   rand.Intn(2) == 1,
			Enabled:              true,
			ActiveDate:           time.Now().AddDate(0, -rand.Intn(12), 0),
			Active:               rand.Intn(2) == 1,
			CreatedBy:            "admin",
			CreatedDate:          time.Now().AddDate(0, -rand.Intn(12), 0),
			ModifiedBy:           "admin",
			ModifiedDate:         time.Now(),
			IsRemove:             false,
			IsOnBoarding:         rand.Intn(2) == 1,
			Code:                 "KOL2024" + strconv.FormatInt(id, 10),
			PortraitRightURL:     "https://example.com/portrait-right-" + strconv.FormatInt(id, 10) + ".jpg",
			PortraitLeftURL:      "https://example.com/portrait-left-" + strconv.FormatInt(id, 10) + ".jpg",
			LivenessStatus:       rand.Intn(2) == 1,
		}
	}

	err := Initializers.DB.AutoMigrate(&Models.Kol{})
	if err != nil {
		log.Printf("failed to migrate database: %v", err)
		return nil
	}

	for _, kol := range dummies {
		Initializers.DB.Create(&kol)
	}

	return dummies
}
