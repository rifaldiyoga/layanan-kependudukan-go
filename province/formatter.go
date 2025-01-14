package province

import (
	"fmt"
	"layanan-kependudukan-api/helper"
)

type ProvinceFormatter struct {
	ID        int    `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

func FormatProvince(province Province) ProvinceFormatter {
	fmt.Print(helper.CapitalizeEachWord(province.Name))
	formatter := ProvinceFormatter{
		ID:        province.ID,
		Code:      province.Code,
		Name:      helper.CapitalizeEachWord(province.Name),
		CreatedAt: helper.FormatDateToString(province.CreatedAt),
	}

	return formatter
}

func FormatProvinces(provinces []Province) []ProvinceFormatter {
	var provincesFormatter []ProvinceFormatter

	for _, province := range provinces {
		provinceFormatter := FormatProvince(province)
		provincesFormatter = append(provincesFormatter, provinceFormatter)
	}

	return provincesFormatter
}
