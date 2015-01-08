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
