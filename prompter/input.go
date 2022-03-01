package prompter

import (
	"errors"
	"fmt"

	"github.com/lafusew/gokeep/data"

	"github.com/manifoldco/promptui"
)

type PromptContent struct {
	ErrorMsg string
	Label string
}

func PromptGetInput(pc PromptContent) (string, error) {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.ErrorMsg)
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
		Label: pc.Label,
		Templates: templates,
		Validate: validate,
	}

	result, err := prompt.Run()
	
	if err != nil {
		fmt.Printf("Command cancelled \n%v\n", err)
	}

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

func TwoStepsSelect(promptContent PromptContent, cred *data.CredID) error{
	domain, err := PromptGetInput(promptContent)

	if err != nil {
		fmt.Println("No credentials found, command cancelled")
	}

	pDomains := data.FindCred(domain)

	if len(pDomains) < 1 {
		err = errors.New("no credentials found, command cancelled")
		return err
	}

	res, err := PromptGetSelect(pDomains, "Confirm selection:")

	if err != nil {
		fmt.Println(err)
	}

	*cred = res

	return err
}