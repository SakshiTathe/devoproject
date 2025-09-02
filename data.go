package main

import "sync"

type Item struct {
    ID   int
    Text string
}

var (
    mu     sync.Mutex
    items  []Item
    nextID = 1
)

