package cmd

import (
	"errors"
	"fmt"

	"github.com/lafusew/gokeep/data"
	"github.com/manifoldco/promptui"
)

func PromptGetInput(pc PromptContent) (string, error) {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
		}

		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt: "{{ . }}",
		Valid: "{{ . | green }}",
		Invalid: "{{ . | red }}",
		Success: "{{ . | bold }}",
	}

	prompt := promptui.Prompt{
		Label: pc.label,
		Templates: templates,
		Validate: validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Command cancelled \n%v\n", err)
	}

	fmt.Printf("Input %s\n", result)
	return result, err
}

func PromptGetSelect(options []data.CredID, label string) (data.CredID, error) {
	
	var stringOptions []string
	for i:= 0; i < len(options); i++ {
		option := options[i]
		stringOptions = append(stringOptions, option.Domain)
	}
	
	prompt := promptui.Select{
		Label: label,
		Items: stringOptions,
	}

	atIndx, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Command cancelled \n%v\n", err)
	}

	return options[atIndx], err
}