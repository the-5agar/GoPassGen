package cmd

import (
	"encoding/json"
	"fmt"
	"gopassgen/core/cryptor"
	"gopassgen/core/passwordutils"
	"gopassgen/core/utils"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
    strength int
    email    string
    username string
    website  string
    master   string
)

var jsonLinesToProcess []string


var newCmd = &cobra.Command{
    Use:   "new",
    Short: "Generate a new password",
    Long:  "Generate a new random password and associate it with a username, email, or website.",
	PreRunE: func (cmd *cobra.Command, args []string) error{
		if username == "" || email == "" {
			return fmt.Errorf("either --username or --email must be specified")
		}
        
        if master == "" {
            return fmt.Errorf("--secret-pass master password  required")
        } else {
            previousEncodedData := utils.ReadData()
            if previousEncodedData != "" {
                decryptedText, err := cryptor.DecryptWithAES(master, previousEncodedData)
                if err != nil {
                    log.Fatal(err)
                    os.Exit(1) // Exits if decryption Key does not match.
                }
                jsonLines := strings.Split(string(decryptedText), "\n,")
                for _, line := range jsonLines {
                    line = strings.TrimSpace(line)

                    if line != "" {
                        jsonLinesToProcess = append(jsonLinesToProcess, line)
                        fmt.Println(jsonLinesToProcess)
                    }
                }               
        }
        } 

		
        return nil
	},


    Run: func(cmd *cobra.Command, args []string) {
        password := password.GeneratePassword(strength)
        fmt.Printf("username: %s | email: %s | password: %s | website: %s", username, email, password, website)

        data := map[string] string {
            "website": website,
            "username" : username,
            "email" : email,
            "password": password,
        }

        jsonData, err := json.Marshal(data)
        if err != nil {
            fmt.Printf("could not marshal json: %s\n", err)
            return
        }
        jsonLinesToProcess = append(jsonLinesToProcess, string(jsonData))

        newData := strings.Join(jsonLinesToProcess, ",\n")
        encodedData, err := cryptor.EncryptWithAES(master, []byte(newData))
        if err != nil {
            log.Fatal(err)
        }

        utils.AddData(encodedData)


    },
}

func init() {
    // Add flags to the new command
    newCmd.Flags().IntVarP(&strength, "strength", "s", 8, "Password strength (length)")
    newCmd.Flags().StringVarP(&username, "username", "u", "", "Username associated with the password")
    newCmd.Flags().StringVarP(&email, "email", "e", "", "Email associated with the password")
    newCmd.Flags().StringVarP(&website, "website", "w", "", "Website associated with the password")
    newCmd.Flags().StringVarP(&master, "secret-pass", "m", "", "Master password for encryption and decryption")
    // Add the newCmd to the rootCmd in root.go
    rootCmd.AddCommand(newCmd)
}
