package db

import "time"

//MakeTimestamp is a little in64 generator
func MakeTimestamp() int64 {
	return time.Now().Unix()
}
