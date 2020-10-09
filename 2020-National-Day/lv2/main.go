package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"syscall"
)

func md5sum(str string) []byte {
	md5Instance := md5.New()
	io.WriteString(md5Instance, str)
	return md5Instance.Sum(nil)
}

func getInfo(file io.Reader) (map[string][]byte, error) {
	var (
		userinfo map[string][]byte
		username string
		passhash []byte
	)
	userinfo = make(map[string][]byte)
	for {
		_, fErr := fmt.Fscanf(file, "%s %x", &username, &passhash)
		if fErr != nil {
			break
		}
		userinfo[username] = passhash
	}
	return userinfo, nil
}

func main() {
	var (
		username string
		password string
	)
	if len(os.Args) > 1 {
		f, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0666)
		if err == nil {
			userinfo, err := getInfo(f)
			fmt.Fprintf(os.Stderr, "Username: ")
			fmt.Scanln(&username)
			fmt.Fprintf(os.Stderr, "Password: ")
			fmt.Scanln(&password)
			var ok bool
			if err == nil {
				if _, exists := userinfo[username]; exists {
					ok = (bytes.Equal(userinfo[username], md5sum(password)))
				} else {
					ok = false
				}
			} else {
				ok = false
			}
			if ok {
				fmt.Println("Authentication success!")
				syscall.Exit(0)
			} else {
				fmt.Println("Authentication failure.")
				syscall.Exit(1)
			}
		} else {
			panic(err)
		}
	} else {
		panic("No password database given!")
	}

}
