package main

import (
    "fmt"
    "os"
    bbolt "go.etcd.io/bbolt"
)

func Buckets(path string) {
    if _, err := os.Stat(path); os.IsNotExist(err) {
        fmt.Println(err)
        return
    }

    db, err := bbolt.Open(path, 0600, nil)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    err = db.View(func(tx *bbolt.Tx) error {
        return tx.ForEach(func(name []byte, _ *bbolt.Bucket) error {
            fmt.Println(string(name))
            return nil
        })
    })
    if err != nil {
        fmt.Println(err)
        return
    }
}

func main() {
    Buckets(os.Args[1])
}
