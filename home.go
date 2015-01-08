package shorty

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
