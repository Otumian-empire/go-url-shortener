package util

import (
	"fmt"
	"hash/fnv"
	"log"
	"strconv"
)

func IsEmpty(obj string) bool {
	return len(obj) < 1
}

func IsNil(obj interface{}) bool {
	return obj == nil
}

func IsNotNil(obj interface{}) bool {
	return !IsNil(obj)
}

// log println
func Log(obj ...interface{}) {
	log.Println(obj...)
}

// log fatal
func FLog(obj ...interface{}) {
	log.Fatal(obj...)
}

func Recover() {
	if err := recover(); IsNotNil(err) {
		Log(SERVER_RECOVER_FROM_ERROR)
		Log(err)
	}
}

// pass a string to a int64 using strconv.ParseInt and propagating the error
func ToInt(str string) (int, error) {
	return strconv.Atoi(str)
}

func StringValidation(str string) error {
	if IsNil(str) || IsEmpty(str) {
		return fmt.Errorf(VALUE_INVALID)
	}

	return nil
}

func CreateHash(url string) string {
	fnvHash := fnv.New32a()
	fnvHash.Write([]byte(url))
	hash := fnvHash.Sum(nil)

	return fmt.Sprintf("%x", hash)
}
