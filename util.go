package main

import (
	"crypto/sha256"
	"encoding/base64"
	"strconv"
	"time"
)

//get timestamp for now
func getNowTimestamp() (timestamp int64) {

	now := time.Now()
	nanos := now.UnixNano()
	timestamp = nanos / 1000000

	return timestamp
}

//get the hash and timestamp for the specified contents
func getWrapping(contents string, key string) (hash string, timestamp int64) {

	timestamp = getNowTimestamp()

	hash = getWrappingHash(contents, timestamp, key)

	return hash, timestamp
}

//get the hash of the contents for the given timestamp
func getWrappingHash(contents string, timestamp int64, key string) (hash string) {

	contentsPlusTime := contents + key + strconv.FormatInt(timestamp, 16)

	hasher := sha256.New()

	bytes := []byte(contentsPlusTime)

	hasher.Write(bytes)

	hash = base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	return hash
}

//checks that the wrapping is valid
func checkWrappingValid(contents string, timestamp int64, hash string, key string) bool {

	checkHash := getWrappingHash(contents, timestamp, key)

	return (checkHash == hash)
}

//get the age of the wrapping
//returns -1 if not valid
func getWrappingAge(contents string, timestamp int64, hash string, key string) (age int64) {

	if checkWrappingValid(contents, timestamp, hash, key) {
		age = getNowTimestamp() - timestamp
	} else {
		age = -1
	}

	return age
}
