package shorty
// http://stackoverflow.com/questions/7922270/obtain-users-home-directory

import (
    "os/user"
    "fmt"
    "log"
)

func HomeDir() string {
    usr, err := user.Current()
    if err != nil {
        log.Fatal( err )
    }
    // fmt.Println( usr.HomeDir )
    return usr.HomeDir
}
