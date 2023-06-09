package cmd


import (
	"time"
	"regexp"
	"errors"
	"strconv"
	"io/ioutil"
	"encoding/json"
	"path/filepath"
)

type File struct {
	Filename 	string `json:"filename"`
	Description string `json:"description"`
	Created_at 	time.Time `json:"created_at"`
}

type Folder struct {
	Foldername 	string `json:"foldername"`
	Description string `json:"description"`
	Created_at 	time.Time `json:"created_at"`
	Files 		[]File `json:"files"`
}

type User struct {
	Username 	string `json:"username"`
	Folders 	[]Folder `json:"folders"`
}

func getFilePath() (path string) {
	absPath, err := filepath.Abs("../cmd/data/users_information.json")
	if err != nil {
		panic(err)
		return
	}
	return absPath
}

func GetUsersInformation()(users []User) {
	fileBytes, err := ioutil.ReadFile(getFilePath())

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileBytes, &users)

	if err != nil {
		panic(err)
	}

	return users
}

func SaveUsersInformation(users []User)(err error) {

	userBytes, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(getFilePath(), userBytes, 0644)
	if err != nil {
		return err
	}
	return nil
}

func InitialUsersInformation()(err error) {
	userBytes, err := json.MarshalIndent([]User{}, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(getFilePath(), userBytes, 0644)
	if err != nil {
		return err
	}
	return nil
}

func CheckValidation(options int, input string, length int) (err error) {
	
	var option string
	switch options {
		case 0:
			option = "username"
		case 1:
			option = "foldername"
		case 2:
			option = "filename"
		case 3:
			option = "new_foldername"
	}

	if len(input) > length || len(input) == 0 {
		return errors.New("Error: The " + option + " must be less than " + strconv.Itoa(length) + " chars and greater than 1 char.\n")
	}

	var alphanumeric = regexp.MustCompile("^[a-zA-Z0-9_]*$")

	if !alphanumeric.MatchString(input) {
		return errors.New("Error: " + option + " contains invalid chars.\n")
	}

	return nil
}

func CreateFakeData(username string, foldername string, filename string, folder_description string, file_description string) (users []User) {
	var files []File
	if filename != "" {
		file := File{
			Filename: filename,
			Description: file_description,
			Created_at: time.Now(),
		}
		files = append(files, file)
	}
	var folders []Folder
	if foldername != "" {
		folder := Folder{
			Foldername: foldername,
			Description: folder_description,
			Created_at: time.Now(),
			Files: files,
		}
		folders = append(folders, folder)
	}
	user := User{
		Username: username,
		Folders: folders,
	}
	return append(users, user)
}

func CreateFakeListData(username string, foldernames []string, filenames []string, option int, t time.Time) (users []User) {
	var files []File
	var folders []Folder

	switch option {
		case 0:
			for index, foldername := range foldernames {
				folder := Folder{
					Foldername: foldername,
					Description: "",
					Created_at: t.Add(time.Hour * time.Duration(index)),
					Files: files,
				}
				folders = append(folders, folder)
			}
		case 1:
			for index, filename := range filenames {
				file := File{
					Filename: filename,
					Description: "",
					Created_at: t.Add(time.Hour * time.Duration(index)),
				}
				files = append(files, file)
			}
			folder := Folder{
				Foldername: foldernames[0],
				Description: "",
				Created_at: t,
				Files: files,
			}
			folders = append(folders, folder)
	}

	user := User{
		Username: username,
		Folders: folders,
	}
	return append(users, user)
}

func checkUserExist(users []User, username string)(exist bool, index int) {
	if len(users) > 0 {
		for index, exist_user := range users {
			if exist_user.Username == username {
				return true, index
			}
		}
	}
	return false, -1
}

func checkFolderExist(folders *[]Folder, foldername string)(exist bool, index int) {
	if len(*folders) > 0 {
		for index := 0; index < len(*folders); index++ {
			if (*folders)[index].Foldername == foldername {
				return true, index
			}
		}
	}
	return false, -1
}

func checkFileExist(files *[]File, filename string)(exist bool, index int) {
	if len(*files) > 0 {
		for index := 0; index < len(*files); index++ {
			if (*files)[index].Filename == filename {
				return true, index
			}
		}
	}
	return false, -1
}