package usecase

import (
	"github.com/kaenova/timed-task/backend/entity"
	"github.com/kaenova/timed-task/backend/repository"
)

type Usecase struct {
	r repository.Repository
}

func NewUseCase(repository repository.Repository) *Usecase {
	return &Usecase{
		r: repository,
	}
}

func (u *Usecase) CreateCheckout(name string) (entity.Checkout, error) {

	obj := entity.CreateCheckout(name)

	err := u.r.TimedCancelCheckout(obj)
	if err != nil {
		return entity.Checkout{}, err
	}

	if err := u.r.SaveCheckout(&obj); err != nil {
		return entity.Checkout{}, err
	}

	return obj, nil
}

func (u *Usecase) CheckoutGoToProcess(id uint) (entity.Checkout, error) {
	obj, err := u.r.GetCheckoutByID(id)
	if err != nil {
		return entity.Checkout{}, err
	}

	err = obj.GoToProcess()
	if err != nil {
		return entity.Checkout{}, err
	}

	err = u.r.TimedCancelCheckout(obj)
	if err != nil {
		return entity.Checkout{}, err
	}

	if err := u.r.SaveCheckout(&obj); err != nil {
		return entity.Checkout{}, err
	}

	return obj, nil
}

func (u *Usecase) CheckoutGoToDeliver(id uint) (entity.Checkout, error) {
	obj, err := u.r.GetCheckoutByID(id)
	if err != nil {
		return entity.Checkout{}, err
	}

	err = obj.GoToDeliver()
	if err != nil {
		return entity.Checkout{}, err
	}

	if err := u.r.SaveCheckout(&obj); err != nil {
		return entity.Checkout{}, err
	}

	return obj, nil
}

func (u *Usecase) CheckoutCancelFromConfirm(id uint) (entity.Checkout, error) {
	obj, err := u.r.GetCheckoutByID(id)
	if err != nil {
		return entity.Checkout{}, err
	}

	err = obj.CancelFromConfirm()
	if err != nil {
		return entity.Checkout{}, err
	}

	if err := u.r.SaveCheckout(&obj); err != nil {
		return entity.Checkout{}, err
	}

	return obj, nil
}

func (u *Usecase) CheckoutCancelFromProcess(id uint) (entity.Checkout, error) {
	obj, err := u.r.GetCheckoutByID(id)
	if err != nil {
		return entity.Checkout{}, err
	}

	err = obj.CancelFromProcess()
	if err != nil {
		return entity.Checkout{}, err
	}

	if err := u.r.SaveCheckout(&obj); err != nil {
		return entity.Checkout{}, err
	}

	return obj, nil
}
