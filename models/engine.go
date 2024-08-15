package models


import (
	"github.com/google/uuid"
	"errors"
)

type Engine struct {
	EngineID  uuid.UUID `json:"enigne_id"`
	Displacement int64 `json:"displacement"`
	NoOfCylinders int64 `json:"noOfCylinders"`
	CarRange int64 `json:"carRange"`
}

type EngineRequest struct {
	Displacement int64 `json:"displacement"`
	NoOfCylinders int64 `json:"noOfCylinders"`
	CarRange int64 `json:"carRange"`
}


//validation
func validateEngineRequest(engineReq EngineRequest) error {
	if err := validateDisplacement(engineReq.Displacement) ; err != nil {
		return err 
	}

	if err := validateNoOfCylinders(engineReq.NoOfCylinders) ; err != nil {
		return err 
	}

	if err := validateCarRange(engineReq.CarRange) ; err != nil {
		return err 
	}
}


func validateDisplacement(displacement int64) error {
	if displacement <= 0 {
		return errors.New("displacement greater than 0")	
	}
	return nil
}

func validateNoOfCylinders(noOfCylinders int64) error {
	if noOfCylinders <= 0 {
		return errors.New("no of cylinders must be greater than 0")
	}
	return nil
}

func validateCarRange(carRange int64) error {
	if carRange <= 0 {
		return errors.New("carRange must be greater than 0" )
	}
	return nil
}