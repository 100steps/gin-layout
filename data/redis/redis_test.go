package redis

import (
	"testing"
	"time"
)

func TestPing(t *testing.T) {
	conn := GetConn()
	defer conn.Close()
	err := pool.TestOnBorrow(conn, time.Now())
	if err != nil {
		t.Fatal(err)
	}
}
