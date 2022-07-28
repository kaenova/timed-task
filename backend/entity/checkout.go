package entity

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Checkout struct {
	gorm.Model
	Name                string
	MaxTimeConfirmation int64
	MaxTimeProcess      *int64

	// Steps
	Confirmation bool
	Processed    bool
	Delivered    bool
	Canceled     bool
}

const (
	CONFIRM_TIME_EXPIRE = time.Minute * 2
	PROCESS_TIME_EXPIRE = time.Minute * 2
)

var (
	ErrAlreadyDeliveredOrCanceled = errors.New("already delivered or canceled")
	ErrAlreadyProcessed           = errors.New("already processed")
	ErrAlreadyDelivered           = errors.New("already delivered")
)

//  ====

func CreateCheckout(name string) Checkout {
	maxTimeToConfirm := time.Now().Add(CONFIRM_TIME_EXPIRE).Unix()
	return Checkout{
		Name:                name,
		Confirmation:        true,
		MaxTimeConfirmation: maxTimeToConfirm,
	}
}

//  ====

func (c *Checkout) GoToProcess() error {
	if c.deliveredOrCanceled() {
		return ErrAlreadyDeliveredOrCanceled
	}
	maxToProcess := time.Now().Add(PROCESS_TIME_EXPIRE).Unix()
	c.MaxTimeProcess = &maxToProcess
	c.Confirmation = true
	c.Processed = true
	return nil
}

func (c *Checkout) GoToDeliver() error {
	if c.deliveredOrCanceled() {
		return ErrAlreadyDeliveredOrCanceled
	}
	c.Confirmation = true
	c.Processed = true
	c.Delivered = true
	return nil
}

func (c *Checkout) CancelFromConfirm() error {
	if c.deliveredOrCanceled() {
		return ErrAlreadyDeliveredOrCanceled
	}
	if c.Processed {
		return ErrAlreadyProcessed
	}
	c.Canceled = true
	return nil
}

func (c *Checkout) CancelFromProcess() error {
	if c.deliveredOrCanceled() {
		return ErrAlreadyDeliveredOrCanceled
	}
	c.Canceled = true
	return nil
}

// ===

func (c *Checkout) deliveredOrCanceled() bool {
	return c.Canceled || c.Delivered
}

//  ====

type TriggerSource int

const (
	TriggerSourceConfirmation TriggerSource = iota
	TriggerSourceProcess
)

type CheckoutTrigger struct {
	ID                uint
	TimeTriggeredLeft int64
	TriggerFrom       TriggerSource
}
