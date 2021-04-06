package datastruct

import (
    "bytes"
    "strconv"
)

type LinkedLister interface {
    PushHead()
    PushBack()
    Head()
    Tail()
    Remove()
    Length()
    Visit()
    Reverse()
}

type Node struct {
    key string
    value int
    next *Node
}

type LinkedList struct {
    head *Node
    tail *Node
    length int
}

// Gonna use pointer receiver here to mutate all the stuff
func (ll *LinkedList) PushHead(key string, value int) int {
    n := Node{key, value, ll.head}

    if ll.head == nil {
        ll.tail = &n
    }
    ll.head = &n
    ll.length += 1

    return ll.length
}

// Gonna use pointer receiver here to mutate all the stuff
func (ll *LinkedList) PushBack(key string, value int) int {
    n := Node{key, value, nil}

    if ll.head == nil {
        ll.head = &n
    } else {
        ll.tail.next = &n
    }
    ll.tail = &n
    ll.length += 1

    return ll.length
}

func (ll *LinkedList) Reverse() {
    curr := ll.head
    ll.head, ll.tail = ll.tail, ll.head
    var prev *Node = nil

    for curr != nil {
        follower := curr.next
        curr.next = prev
        prev = curr
        curr = follower
    }
    ll.head = prev
}

// Remove all the instances O(n)
func (ll *LinkedList) Remove(key string) {
    if ll.head == nil {
        return
    }
    if ll.head.key == key {
        ll.head = ll.head.next
        ll.length -= 1
    }
    var prev *Node = nil
    curr := ll.head

    for curr != nil {
        if curr.key == key {
            prev.next = curr.next
            ll.length -= 1
        }
        prev = curr
        curr = curr.next
    }
}

func (ll LinkedList) Length() int {
    return ll.length
}

func (ll LinkedList) Head() *Node {
    return ll.head
}

func (ll LinkedList) Tail() *Node {
    return ll.tail
}

// Gonna use value receiver here because i'm just reading
func (ll LinkedList) Visit() string {
    var visit bytes.Buffer
    // visit.WriteString("Visit: [ ")
    curr := ll.head
    for curr != nil {
        // visit.WriteString(curr.key)
        // visit.WriteString(": ")
        visit.WriteString(strconv.Itoa(curr.value))
        if curr.next != nil {
            visit.WriteString(",")
        }
        // visit.WriteString(" ")
        curr = curr.next
    }
    // visit.WriteString("]")
    var res string = visit.String()
    return res
}

func (ll LinkedList) Get (key string) int {
    curr := ll.head
    for curr != nil {
        if curr.key == key {
            return curr.value
        }
        curr = curr.next
    }
    return -1
}