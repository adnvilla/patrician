package postgresql

import (
	"github.com/adnvilla/patrician/src/domain"
	"gorm.io/gorm"
)

// MarketHallRepository persists market halls.
type MarketHallRepository struct {
	db *gorm.DB
}

func NewMarketHallRepository(db *gorm.DB) *MarketHallRepository {
	return &MarketHallRepository{db: db}
}

func (r *MarketHallRepository) Create(m *domain.MarketHall) error {
	model := marketHallToModel(m)
	if err := r.db.Create(&model).Error; err != nil {
		return err
	}
	*m = marketHallModelToDomain(&model)
	return nil
}

func (r *MarketHallRepository) Update(m *domain.MarketHall) error {
	model := marketHallToModel(m)
	return r.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&model).Error
}

func (r *MarketHallRepository) Delete(m *domain.MarketHall) error {
	model := marketHallToModel(m)
	return r.db.Delete(&model).Error
}

func (r *MarketHallRepository) FindByID(id uint) (*domain.MarketHall, error) {
	var model MarketHallModel
	if err := r.db.Preload("Commodities").First(&model, id).Error; err != nil {
		return nil, err
	}
	mh := marketHallModelToDomain(&model)
	return &mh, nil
}

func marketHallToModel(m *domain.MarketHall) MarketHallModel {
	res := MarketHallModel{Entity: m.Entity}
	for _, c := range m.Commodities {
		res.Commodities = append(res.Commodities, *toCommodityModel(c))
	}
	return res
}

func marketHallModelToDomain(m *MarketHallModel) domain.MarketHall {
	mh := domain.MarketHall{Entity: m.Entity, Commodities: map[string]*domain.Commodity{}}
	for _, c := range m.Commodities {
		commodity := commodityModelToDomain(&c)
		mh.Commodities[commodity.Name] = &commodity
	}
	return mh
}
