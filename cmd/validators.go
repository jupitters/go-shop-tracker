package main

import (
	"slices"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
	"github.com/jupitters/go-pizza-tracker/internal/models"
)

func RegisterCustomValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.registerValidation("valid_pizza_type", createSliceValidator(models.PizzaTypes))
		v.registerValidation("valid_pizza_size", createSliceValidator(models.PizzaSizes))
	}
}

func createSliceValidator(allowedVallues []string) validator.Func {
	return func(fl validator.FieldValue) bool {
		return slices.Contains(allowedValues, fl.Field().String())
	}
}
