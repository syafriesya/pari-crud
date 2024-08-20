package crud

import (
	"context"
	"pari/internal/domain"
)

func (cr CrudRepository) GetAll(ctx context.Context) ([]domain.Items, error) {
	var result []domain.Items

	if err := cr.db.WithContext(ctx).Debug().Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (cr CrudRepository) GetById(ctx context.Context, id string) (domain.Items, error) {
	var result domain.Items

	if err := cr.db.WithContext(ctx).Debug().
		Where("id = ?", id).
		Find(&result).Error; err != nil {
		return domain.Items{}, err
	}

	return result, nil
}
