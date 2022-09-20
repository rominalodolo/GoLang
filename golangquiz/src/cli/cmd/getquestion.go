package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

// getquestionCmd represents the getquestion command
var getquestionCmd = &cobra.Command{
	Use:   "getquestion",
	Short: "Fetch questions from server.",
	Long:  `Fetch question from API server. Use it like that: cli.exe getquestion --id=2`,
	Run: func(cmd *cobra.Command, args []string) {
		questionID, _ := cmd.Flags().GetString("id")

		questionReader, err := FetchQuestion(questionID)
		if err != nil {
			os.Stderr.WriteString("Failed to fetch the question.")
			os.Exit(1)
		}
		defer questionReader.Close()

		respBody, err := ioutil.ReadAll(questionReader)
		if err != nil {
			os.Stderr.WriteString("Error reading response. " + err.Error())
			os.Exit(1)
		}

		fmt.Println("Response:\n", string(respBody))
	},
}

func init() {
	rootCmd.AddCommand(getquestionCmd)

	getquestionCmd.Flags().StringP("id", "i", "", "Question ID")
}
