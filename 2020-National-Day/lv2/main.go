package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
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

func getInfo(file io.Reader) [2]string {
	var (
		userinfo [2]string
	)
	fmt.Fscanln(file, &userinfo[0])
	fmt.Fscanln(file, &userinfo[1])
	return userinfo
}

func main() {
	var (
		username string
		password string
	)
	fmt.Fprintf(os.Stderr, "Username: ")
	fmt.Scanln(&username)
	fmt.Fprintf(os.Stderr, "Password: ")
	fmt.Scanln(&password)
	if len(os.Args) > 1 {
		f, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0666)
		if err == nil {
			userinfo := getInfo(f)
			passhash, err := hex.DecodeString(userinfo[1])
			var ok bool
			if err == nil {
				ok = (userinfo[0] == username && bytes.Equal(passhash, md5sum(password)))
				// fmt.Fprintf(os.Stderr, "Data : %s %s\nInput: %s %s\n", userinfo[0], userinfo[1], username, hex.EncodeToString(md5sum(password)))
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
