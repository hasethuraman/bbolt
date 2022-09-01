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

func dumpkeys(path string, bucket string) {
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
    db.View(func(tx *bbolt.Tx) error {
        // Assume bucket exists and has keys
        b := tx.Bucket([]byte(bucket))

        c := b.Cursor()

        for k, _ := c.First(); k != nil; k, _ = c.Next() {
            fmt.Printf("%s\n", k)
        }
        return nil
    })
    if err != nil {
        fmt.Println(err)
        return
    }
}
// note: take a copy of db (or snapshot also) and then do otherwise it wont work
// ./boltcli listbuckets /tmp/db
// ./boltcli dumpkeys /tmp/db <bucketname>
func main() {
    if os.Args[1] == "listbuckets" {
        Buckets(os.Args[2])
    } else if os.Args[1] == "dumpkeys" {
        dumpkeys(os.Args[2], os.Args[3])
    }
}
