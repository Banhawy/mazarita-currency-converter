package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
)

var (
	baseCurrency    string
	targetCurrency  string
	amountToConvert string
)

func makeHuhSlection(title string, options map[string]string, value *string) huh.Field {
	selectOptions := make([]huh.Option[string], len(options))
	i := 0
	for k, v := range options {
		selectOptions[i] = huh.NewOption(v, k)
		i++
	}
	return huh.NewSelect[string]().Title(title).Options(selectOptions...).Value(value)
}

func main() {
	var currencies = map[string]string{
		"EGP": "£ EGP",
		"USD": "$ USD",
		"EUR": "€ EUR",
		"GBP": "£ GBP",
		"SAR": "SAR ﷼",
		"KWD": "KWD د.ك",
	}

	form := huh.NewForm(
		huh.NewGroup(
			makeHuhSlection("Convert from base currency", currencies, &baseCurrency),
			makeHuhSlection("to target currency", currencies, &targetCurrency),

			huh.NewInput().
				Title("Amount").
				Placeholder("Enter amount").
				Validate(func(str string) error {
					if str == "" {
						return errors.New("amount is required")
					}

					_, err := strconv.ParseFloat(str, 64)
					if err != nil {
						return errors.New("amount must be a number")
					}

					return nil
				}).
				Value(&amountToConvert),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("You want to convert %s %s to %s\n", amountToConvert, baseCurrency, targetCurrency)

	action := func() { time.Sleep(5 * time.Second) }
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	go action()
	spinErr := spinner.New().
		Type(spinner.Line).
		Title("Making conversion...").
		Context(ctx).
		Run()

	if spinErr != nil {
		log.Fatal(spinErr)
	}

	fmt.Println("Done!")
}
