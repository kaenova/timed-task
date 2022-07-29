package usecase

import (
	"log"

	"github.com/kaenova/timed-task/backend/entity"
	"github.com/kaenova/timed-task/backend/repository"
)

type Usecase struct {
	r *repository.Repository
}

func NewUseCase(repository *repository.Repository) *Usecase {
	log.Print("Creating use cases")
	return &Usecase{
		r: repository,
	}
}

func (u *Usecase) CreateCheckout(name string) (entity.Checkout, error) {

	obj := entity.CreateCheckout(name)

	if err := u.r.SaveCheckout(&obj); err != nil {
		return entity.Checkout{}, err
	}

	err := u.r.TimedCancelCheckout(obj)
	if err != nil {
		return entity.Checkout{}, err
	}

	return obj, nil
}

func (u *Usecase) DeleteCheckout(id uint) error {

	if err := u.r.DeleteCheckoutByID(id); err != nil {
		return err
	}

	return nil
}

func (u *Usecase) GetAllCheckout() ([]entity.Checkout, error) {

	objs, err := u.r.GetAllCheckout()
	if err != nil {
		return nil, err
	}

	return objs, nil
}

func (u *Usecase) GetCheckoutById(id uint) (entity.Checkout, error) {

	obj, err := u.r.GetCheckoutByID(id)
	if err != nil {
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

	if err := u.r.SaveCheckout(&obj); err != nil {
		return entity.Checkout{}, err
	}

	err = u.r.TimedCancelCheckout(obj)
	if err != nil {
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
