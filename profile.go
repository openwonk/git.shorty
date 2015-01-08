package shorty

import (
    "fmt"
    "os"
)

func Profile() []byte {
    filepath := shorty.HomeDir() + "/.bash_profile"
    file,err := os.Open(filepath)
    if err != nil {
        fmt.Println(err)
    }

    f := file.Read()

    return f
}
 
func Source(filepath string) []byte {
    file,err := os.Open(filepath)
    if err != nil {
        fmt.Println(err)
    }

    f := file.Read()

    return f
}   


func ModProfile() {
	originFile, err := os.Open("shorty")
	check(err)

	var data []byte
	_, err = originFile.Read(data)
	check(err)

	targetFile, err := os.Open("~/.bash_profile_test")
	check(err)

	targetFile.Write(data)

}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
