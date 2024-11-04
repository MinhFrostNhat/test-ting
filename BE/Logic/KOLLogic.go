package Logic

import (
	"fmt"
	"log/slog"
	"wan-api-kol-event/DTO"
	"wan-api-kol-event/Initializers"
	"wan-api-kol-event/Models"
	"wan-api-kol-event/Utils"
)

// * Get Kols from the database based on the range of pageIndex and pageSize
// ! USE GORM TO QUERY THE DATABASE
// ? There are some support function that can be access in Utils folder (/BE/Utils)
// --------------------------------------------------------------------------------
// @params: pageIndex
// @params: pageSize
// @return: List of KOLs and error message
func GetKolLogic(pageIndex, pageSize int64) ([]*DTO.KolDTO, error) {
	slog.Info("GetKolLogic started", "pageIndex", pageIndex, "pageSize", pageSize)
	var kols []Models.Kol
	var totalCount int64

	offset := (pageIndex - 1) * pageSize
	slog.Info("Calculated offset", "offset", offset)

	if err := Initializers.DB.Model(&Models.Kol{}).Count(&totalCount).Error; err != nil {
		slog.Error("Error when counting KOLs", "error", err)
		return nil, fmt.Errorf("failed to count KOLs: %v", err)
	}

	if err := Initializers.DB.Offset(int(offset)).Limit(int(pageSize)).Find(&kols).Error; err != nil {
		slog.Error("Error when retrieving KOLs", "error", err)
		return nil, fmt.Errorf("failed to retrieve KOLs: %v", err)
	}

	slog.Info("KOLs retrieved", "count", len(kols))
	kolsDTO := make([]*DTO.KolDTO, len(kols))
	for i, kol := range kols {
		kolsDTO[i] = &DTO.KolDTO{
			KolID:                kol.KolID,
			UserProfileID:        kol.UserProfileID,
			Language:             kol.Language,
			Education:            kol.Education,
			ExpectedSalary:       kol.ExpectedSalary,
			ExpectedSalaryEnable: kol.ExpectedSalaryEnable,
			ChannelSettingTypeID: kol.ChannelSettingTypeID,
			IDFrontURL:           kol.IDFrontURL,
			IDBackURL:            kol.IDBackURL,
			PortraitURL:          kol.PortraitURL,
			RewardID:             kol.RewardID,
			PaymentMethodID:      kol.PaymentMethodID,
			TestimonialsID:       kol.TestimonialsID,
			VerificationStatus:   kol.VerificationStatus,
			Enabled:              kol.Enabled,
			ActiveDate:           Utils.TimeToString(kol.ActiveDate),
			Active:               kol.Active,
			CreatedBy:            kol.CreatedBy,
			CreatedDate:          Utils.TimeToString(kol.CreatedDate),
			ModifiedBy:           kol.ModifiedBy,
			ModifiedDate:         Utils.TimeToString(kol.ModifiedDate),
			IsRemove:             kol.IsRemove,
			IsOnBoarding:         kol.IsOnBoarding,
			Code:                 kol.Code,
			PortraitRightURL:     kol.PortraitRightURL,
			PortraitLeftURL:      kol.PortraitLeftURL,
			LivenessStatus:       kol.LivenessStatus,
		}
	}

	slog.Info("GetKolLogic completed successfully", "resultCount", len(kolsDTO))
	return kolsDTO, nil
}
