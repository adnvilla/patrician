package postgresql

import (
	"github.com/adnvilla/patrician/src/domain"
	"gorm.io/gorm"
)

// CommodityRepository handles Commodity persistence.
type CommodityRepository struct {
	db *gorm.DB
}

func NewCommodityRepository(db *gorm.DB) *CommodityRepository {
	return &CommodityRepository{db: db}
}

func (r *CommodityRepository) Create(c *domain.Commodity) error {
	model := toCommodityModel(c)
	if err := r.db.Create(model).Error; err != nil {
		return err
	}
	*c = commodityModelToDomain(model)
	return nil
}

func (r *CommodityRepository) Update(c *domain.Commodity) error {
	model := toCommodityModel(c)
	return r.db.Updates(model).Error
}

func (r *CommodityRepository) Delete(c *domain.Commodity) error {
	return r.db.Delete(toCommodityModel(c)).Error
}

func (r *CommodityRepository) FindByID(id uint) (*domain.Commodity, error) {
	var model CommodityModel
	if err := r.db.First(&model, id).Error; err != nil {
		return nil, err
	}
	commodity := commodityModelToDomain(&model)
	return &commodity, nil
}

func toCommodityModel(c *domain.Commodity) *CommodityModel {
	return &CommodityModel{
		Entity:        c.Entity,
		Name:          c.Name,
		CommodityType: int(c.CommodityType),
		Buy:           c.Buy,
		Sell:          c.Sell,
		Production:    c.Production,
		Consumption:   c.Consumption,
	}
}

func commodityModelToDomain(m *CommodityModel) domain.Commodity {
	return domain.Commodity{
		Entity:        m.Entity,
		Name:          m.Name,
		CommodityType: domain.CommodityType(m.CommodityType),
		Buy:           m.Buy,
		Sell:          m.Sell,
		Production:    m.Production,
		Consumption:   m.Consumption,
	}
}
