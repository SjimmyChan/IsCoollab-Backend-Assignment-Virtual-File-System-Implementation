package cmd_test

import (
	"testing"

	"github.com/SjimmyChan/IsCoollab-Backend-Assignment-Virtual-File-System-Implementation/cmd"
	"github.com/spf13/cobra"
)

func TestCreateFolderCmd(t *testing.T) {
	
	// store exist user inforamtion and initial json file
	exist_users_inforamtion := cmd.GetUsersInformation()
	cmd.InitialUsersInformation()

	testCreateFileCmd := &cobra.Command{
		Use: "create-folder",
		Run: CreateFolderCmdRunE,
	}
	CreateFolderCmdFlags(testCreateFileCmd)

	cmd.SaveUsersInformation(cmd.CreateFakeData("user1", "", "", "", ""))

	createFileTests := []struct {
		input 	[]string
		output 	string
	}{
		{
			input: []string{"-u", "user1", "-f", "folder1"},
			output: "Create folder1 in user1 successfully.\n",
		},
		{
			input: []string{"-u", "abc", "-f", "folder1"},
			output: "Error: The username:abc doesn't exist.\n",
		},
		{
			input: []string{"-u", "user1", "-f", "folder1"},
			output: "Error: The foldername:folder1 has already existed.\n",
		},
		{
			input: []string{"-u", "user?", "-f", "folder1"},
			output: "Error: username contains invalid chars.\n",
		},
		{
			input: []string{"-u", "user1", "-f", "folder?"},
			output: "Error: foldername contains invalid chars.\n",
		},
		{
			input: []string{"-u", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "-f", "folder1"},
			output: "Error: The username must be less than 30 chars and greater than 1 char.\n",
		},
		{
			input: []string{"-u", "user1", "-f", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
			output: "Error: The foldername must be less than 30 chars and greater than 1 char.\n",
		},
	}

	for _, test := range createFileTests {
		actual_output := execute(t, testCreateFileCmd, test.input ... )

		expected_output := test.output
		if expected_output != actual_output {
			t.Errorf("Expected output '%s', but got '%s'", expected_output, actual_output)
		}
	}

	// store back originial user inforamtion
	cmd.SaveUsersInformation(exist_users_inforamtion)
}

func CreateFolderCmdRunE(c *cobra.Command, args []string) {
	username, _ := c.Flags().GetString("username")
		if err := cmd.CheckValidation(0, username, 30); err != nil {
			c.Print(err.Error())
			return
		}

		foldername, _ := c.Flags().GetString("foldername")
		if err := cmd.CheckValidation(1, foldername, 30); err != nil {
			c.Print(err.Error())
			return
		}

		description, _ := c.Flags().GetString("description")

		if succeed := cmd.CreateFolder(c, username, foldername, description); succeed {
			c.Println("Create " + foldername + " in " + username + " successfully.")
		}
}

func CreateFolderCmdFlags(c *cobra.Command) {
	c.Flags().StringP("username", "u", "", "username")
	c.MarkFlagRequired("username")

	c.Flags().StringP("foldername", "f", "", "foldername")
	c.MarkFlagRequired("foldername")

	c.Flags().StringP("description", "d", "", "description")
}