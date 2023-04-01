package main

import (
	"fmt"
	"strings"
)

type Chiper interface {
	Encode() string
	Decode() string
}

type Student struct {
	name       string
	nameEncode string
	score      int
}

func (s *Student) Encode() string {
	var sb strings.Builder
	for _, r := range s.name {
		if r >= 'a' && r <= 'z' {
			sb.WriteByte('a' + byte((r-'a'+3)%26))
		} else if r >= 'A' && r <= 'Z' {
			sb.WriteByte('A' + byte((r-'A'+3)%26))
		} else {
			sb.WriteRune(r)
		}
	}
	s.nameEncode = sb.String()
	return s.nameEncode
}

func (s *Student) Decode() string {
	var sb strings.Builder
	for _, r := range s.nameEncode {
		if r >= 'a' && r <= 'z' {
			sb.WriteByte('a' + byte((r-'a'+23)%26))
		} else if r >= 'A' && r <= 'Z' {
			sb.WriteByte('A' + byte((r-'A'+23)%26))
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

func main() {
	var menu int
	var a Student
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of student’s name " + a.name + " is: " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.nameEncode)
		fmt.Print("\nDecode of student’s name " + a.nameEncode + " is: " + c.Decode())
	}
}
