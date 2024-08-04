package pkg

import (
	"regexp"

	"github.com/spf13/cobra"
)

var (
	outputFormat string
	isRaw        bool
)

func init() {
	RootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "yaml", "Output format (yaml or json)")
	RootCmd.PersistentFlags().BoolVarP(&isRaw, "raw", "r", false, "Print raw output")
}

var RootCmd = &cobra.Command{
	Use:   "cpf [cpf]",
	Short: "Consulte o CPF de uma pessoa.",
	Args:  cobra.ExactArgs(1),
	Run:   fetchCPF,
}

func fetchCPF(cmd *cobra.Command, args []string) {
	var data any
	cnpj := args[0]

	reg := regexp.MustCompile("[^0-9]+")
	cnpj = reg.ReplaceAllString(cnpj, "")

	if cnpj == "" {
		cmd.Println("CPF não pode ser vazio.")
		return
	}

	data, err := fetchCPFData(cmd.Context(), cnpj)
	if err != nil {
		panic(err)
	}

	switch outputFormat {
	case "json":
		PrintJSON(data, isRaw)
	case "yaml":
		PrintYAML(data, isRaw)
	default:
		cmd.Println("Unsupported output format. Please choose 'json' or 'yaml'.")
	}
}
