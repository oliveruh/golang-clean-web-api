package dependency

import (
	"golang-clean-web-api/config"
	"golang-clean-web-api/domain/model"
	contractRepository "golang-clean-web-api/domain/repository"
	database "golang-clean-web-api/infra/persistence/database"
	infraRepository "golang-clean-web-api/infra/persistence/repository"
)

func GetCountryRepository(cfg *config.Config) contractRepository.CountryRepository {
	var preloads []database.PreloadEntity = []database.PreloadEntity{{Entity: "Cities"}, {Entity: "Companies"}}
	return infraRepository.NewBaseRepository[model.Country](cfg, preloads)
}

func GetCityRepository(cfg *config.Config) contractRepository.CityRepository {
	var preloads []database.PreloadEntity = []database.PreloadEntity{{Entity: "Country"}}
	return infraRepository.NewBaseRepository[model.City](cfg, preloads)
}

func GetColorRepository(cfg *config.Config) contractRepository.ColorRepository {
	var preloads []database.PreloadEntity = []database.PreloadEntity{}
	return infraRepository.NewBaseRepository[model.Color](cfg, preloads)
}

func GetCompanyRepository(cfg *config.Config) contractRepository.CompanyRepository {
	var preloads []database.PreloadEntity = []database.PreloadEntity{{Entity: "Country"}}
	return infraRepository.NewBaseRepository[model.Company](cfg, preloads)
}
