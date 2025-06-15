package postgresql

import (
	"github.com/adnvilla/patrician/src/domain"
	"gorm.io/gorm"
)

// DistanceRepository persists distance entities.
type DistanceRepository struct {
	db *gorm.DB
}

func NewDistanceRepository(db *gorm.DB) *DistanceRepository {
	return &DistanceRepository{db: db}
}

func (r *DistanceRepository) Create(d *domain.Distance) error {
	model := toDistanceModel(d)
	if err := r.db.Create(model).Error; err != nil {
		return err
	}
	*d = distanceModelToDomain(model)
	return nil
}

func (r *DistanceRepository) Update(d *domain.Distance) error {
	model := toDistanceModel(d)
	return r.db.Updates(model).Error
}

func (r *DistanceRepository) Delete(d *domain.Distance) error {
	return r.db.Delete(toDistanceModel(d)).Error
}

func (r *DistanceRepository) FindAll() ([]domain.Distance, error) {
	var models []DistanceModel
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}
	var res []domain.Distance
	for _, m := range models {
		res = append(res, distanceModelToDomain(&m))
	}
	return res, nil
}

func toDistanceModel(d *domain.Distance) *DistanceModel {
	return &DistanceModel{
		Entity:   d.Entity,
		FromCity: d.FromCity,
		ToCity:   d.ToCity,
		Value:    d.Value,
	}
}

func distanceModelToDomain(m *DistanceModel) domain.Distance {
	return domain.Distance{
		Entity:   m.Entity,
		FromCity: m.FromCity,
		ToCity:   m.ToCity,
		Value:    m.Value,
	}
}
