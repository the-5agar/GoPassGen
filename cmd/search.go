package cmd

import (
	"gopassgen/core/cryptor"
	"gopassgen/core/utils"
	"log"
	"github.com/spf13/cobra"
	"fmt"
	"strings"
)

var searchCmd  = &cobra.Command{
	Use:   "search",
    Short: "Search for Stored Password",
    Long:  "Match stored password based on website and regex.",

	Run: func(cmd *cobra.Command, args []string) {
		encodedData := utils.ReadData()
		if encodedData != "" {
			decryptedText, err := cryptor.DecryptWithAES(master, encodedData)
			if err != nil {
				log.Fatal(err)
			}
	
			lines := strings.Split(string(decryptedText), "\n")
	
			for _, line := range lines {
				trimmedLine := strings.TrimSpace(line)
				if strings.Contains(trimmedLine, website)  {
					trimmedLine = strings.TrimRight(trimmedLine, ",") //Outputs in JSON FORMAT
					fmt.Println(trimmedLine)
				}
			}
		}
	},
		
}


func init() {
	searchCmd.Flags().StringVarP(&website, "website", "w", "", "search based on website")
	searchCmd.Flags().StringVarP(&master, "secret-pass", "m", "", "Master password for encryption and decryption")
    rootCmd.AddCommand(searchCmd)
}