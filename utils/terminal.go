package utils

import (
	"os"
	"os/exec"
	"text/tabwriter"

	"github.com/AlecAivazis/survey/v2"
	"github.com/olekukonko/tablewriter"
)

type Termianl struct {
	homeBasePath string
}

func NewTerminal(basePath string) Termianl {

	t := Termianl{
		homeBasePath: basePath,
	}

	return t
}

func (t Termianl) Clear() Termianl {
	cexec("clear")
	return t
}

func (t Termianl) SelectOne(title string, items []string) string {

	pp := survey.Select{
		Message: title,
		Options: items,
	}

	var pick string
	MustCheckError(survey.AskOne(&pp, &pick))
	return pick
}

func (t Termianl) SelectMultiple(title string, items []string) []string {
	prompt := &survey.MultiSelect{
		Message: title,
		Options: items,
	}

	var result []string
	MustCheckError(survey.AskOne(prompt, &result))
	return result
}

func (t Termianl) TableWriter(title []string, values [][]string) {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(title)
	table.SetAlignment(int(tabwriter.AlignRight))

	for _, value := range values {
		table.Append(value)
	}

	table.Render()
}

func cexec(cmds ...string) {

	cmd := exec.Command(cmds[0], cmds[1:]...)
	cmd.Stdout = os.Stdout
	MustCheckError(cmd.Run())
}
