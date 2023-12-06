package ui

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
)

type PromptContent struct {
	ErrorMsg string
	Label    string
}

// func PromptGetInput интерактивный ввод
func PromptGetInput(pc PromptContent, text string) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.ErrorMsg)
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     pc.Label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s: %s\n", text, result)

	return result
}

// func PromptGetInputPsw интерактивный ввод пароля
func PromptGetInputPsw(pc PromptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.ErrorMsg)
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     pc.Label,
		Templates: templates,
		Validate:  validate,
		Mask:      '*',
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}

// func PromptChose3 интерактивное меню
func PromptChose3() string {
	items := []string{"Изменить", "Запросить данные", "Выход"}
	prompt := promptui.Select{
		Label: "Выберите один из вариантов",
		Items: items,
	}
	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
}

// func PromptConfirm подтверждение для сохранения
func PromptConfirm() string {

	prompt := promptui.Prompt{
		Label:     "Сохранить данные?",
		IsConfirm: true,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result

}
