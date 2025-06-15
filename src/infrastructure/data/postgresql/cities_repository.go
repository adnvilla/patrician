package postgresql

import (
	"github.com/adnvilla/patrician/src/domain"
	"gorm.io/gorm"
)

// CityRepository provides CRUD operations for City entities.
type CityRepository struct {
	db *gorm.DB
}

func NewCityRepository(db *gorm.DB) *CityRepository {
	return &CityRepository{db: db}
}

func (r *CityRepository) Create(city *domain.City) error {
	model := toCityModel(city)
	if err := r.db.Create(&model).Error; err != nil {
		return err
	}
	*city = cityModelToDomain(&model)
	return nil
}

func (r *CityRepository) Update(city *domain.City) error {
	model := toCityModel(city)
	return r.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&model).Error
}

func (r *CityRepository) Delete(city *domain.City) error {
	model := toCityModel(city)
	return r.db.Delete(&model).Error
}

func (r *CityRepository) FindAll() ([]domain.City, error) {
	var models []CityModel
	if err := r.db.Preload("MarketHall.Commodities").Find(&models).Error; err != nil {
		return nil, err
	}
	return cityModelsToDomain(models), nil
}

func (r *CityRepository) FindByID(id uint) (*domain.City, error) {
	var model CityModel
	if err := r.db.Preload("MarketHall.Commodities").First(&model, id).Error; err != nil {
		return nil, err
	}
	city := cityModelToDomain(&model)
	return &city, nil
}

func toCityModel(c *domain.City) CityModel {
	m := CityModel{Entity: c.Entity, Name: c.Name}
	if c.MarketHall.Commodities != nil {
		mh := MarketHallModel{Entity: c.MarketHall.Entity}
		for _, com := range c.MarketHall.Commodities {
			mh.Commodities = append(mh.Commodities, *toCommodityModel(com))
		}
		m.MarketHall = mh
	}
	return m
}

func cityModelToDomain(m *CityModel) domain.City {
	city := domain.City{Entity: m.Entity, Name: m.Name}
	commodities := map[string]*domain.Commodity{}
	for _, cm := range m.MarketHall.Commodities {
		commodity := commodityModelToDomain(&cm)
		commodities[commodity.Name] = &commodity
	}
	city.MarketHall = domain.MarketHall{Entity: m.MarketHall.Entity, Commodities: commodities}
	return city
}

func cityModelsToDomain(models []CityModel) []domain.City {
	var cities []domain.City
	for _, m := range models {
		c := cityModelToDomain(&m)
		cities = append(cities, c)
	}
	return cities
}
