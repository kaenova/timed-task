package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/kaenova/timed-task/backend/entity"
	"github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

func (r *Repository) GetCheckoutByID(id uint) (entity.Checkout, error) {
	obj := entity.Checkout{
		Model: gorm.Model{
			ID: id,
		},
	}
	if err := r.g.First(&obj).Error; err != nil {
		return entity.Checkout{}, err
	}
	return obj, nil
}

func (r *Repository) GetAllCheckout() ([]entity.Checkout, error) {
	var objs []entity.Checkout
	if err := r.g.Find(&objs).Error; err != nil {
		return nil, err
	}
	return objs, nil
}

func (r *Repository) SaveCheckout(objs ...*entity.Checkout) error {
	tx := r.g.Begin()
	for i := 0; i < len(objs); i++ {
		if err := tx.Save(&objs[i]).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	err := tx.Commit().Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) TimedCancelCheckout(obj entity.Checkout) error {
	now := time.Now().Unix()
	ctx := context.Background()

	if obj.MaxTimeConfirmation > now {
		deltaMiliseconds := (obj.MaxTimeConfirmation - now) * 1000
		trg := entity.CheckoutTrigger{
			ID:                obj.ID,
			TimeTriggeredLeft: deltaMiliseconds,
			TriggerFrom:       entity.TriggerSourceConfirmation,
		}

		body, err := json.Marshal(trg)
		if err != nil {
			return err
		}

		r.rb.PublishWithContext(ctx, r.rbConf.TimedExchangeName, "", false, false,
			amqp091.Publishing{
				ContentType: "application/json",
				Body:        body,
				Headers: amqp091.Table{
					"x-delay": trg.TimeTriggeredLeft,
				},
			})
	}

	if obj.MaxTimeProcess != nil && *obj.MaxTimeProcess > now {
		deltaMiliseconds := (obj.MaxTimeConfirmation - now) * 1000
		trg := entity.CheckoutTrigger{
			ID:                obj.ID,
			TimeTriggeredLeft: deltaMiliseconds,
			TriggerFrom:       entity.TriggerSourceProcess,
		}

		body, err := json.Marshal(trg)
		if err != nil {
			return err
		}

		r.rb.PublishWithContext(ctx, r.rbConf.TimedExchangeName, "", false, false,
			amqp091.Publishing{
				ContentType: "application/json",
				Body:        body,
				Headers: amqp091.Table{
					"x-delay": trg.TimeTriggeredLeft,
				},
			})
	}
	return nil
}
