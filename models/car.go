package models

import (
	"github.com/google/uuid"
	"errors"
	"time"
	"strconv"
)

type Car struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	Year string `json:"year"`
	Brand string `json:"brand"`
	FuelType string `json:"fuel-type"`
	Engine Engine `json:"engine"`
	Price float64 `json:"price:`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

}

type CarRequest struct {
	Name string `json:"name"`
	Year string `json:"year"`
	Brand string `json:"brand"`
	FuelType string `json:"fuel-type"`
	Engine Engine `json:"engine"`
	Price float64 `json:"price:`
}


//create validateions

func validateRequest(carReq CarRequest) error {
	if err := validateName(carReq.Name);err != nil {
		return err
	}

	if err := validateYear(carReq.Year);err != nil {
		return err
	}

	if err := validateBrand(carReq.Brand);err != nil {
		return err
	}


	if err := validateFuelType(carReq.FuelType);err != nil {
		return err
	}

	if err := validateEngine(carReq.Engine);err != nil {
		return err
	}

	if err := validatePrice(carReq.Price);err != nil {
		return err
	}

	return nil
}


func validateName(name string) error {
	if name == ""{
		return erros.New("Name is Required")
	}
}

func validateYear(year string) error {
	if year == ""{
		return errors.New("Year is Required")
	}
	//convert year into int
	_ ,err  := strconv.Atoi(year)
	if err != nil {
		return errors.New("year must be a valid number")
	}

	//calculate current year
	currentYear := time.Now().Year()
	intYear ,  _ := strconv.Atoi(year)
	if intYear < 1888 || intYear > currentYear{
		return errors.New("Year must be between 1888 and current year")
	}
	return nil
}

func validateBrand(brand string) error {
	if brand == "" {
	return errors.New("Brand is required")
	}
	return nil
}

func validateFuelType(fuelType string) error {
	validateFuleType := []string{"Petrol", "Diesel", "Electric","Hybrid"}
	for _, validType  := range validateFuelType {
		if fuelTye == validType {
			return nil
		}
	}

	return errors.New("Fuel Type Must be one of : Petrol, Diesel, Elctric, Hybrid")
}

func validateEngine(engine Engine) error {
	if engine.EngineID == uuid.Nil {
		return errors.New("Engine ID is required")
	}

	if engine.Displacement <= 0 {
		return errors.New("Displacement must be a greater tha 0")
	}
	if egiine.NoOfCylinders <= 0{
		return errors.New("no of cylinder must be a greater tha 0")	
	}

	if egiine.CarRange <= 0{
		return errors.New("Car range msut be a greater tha 0")	
	}

	return nil
}

func validatePrice(price float64) error {
	if price <= 0 {
		return errors.New("price must be greater than 0")
	}
}