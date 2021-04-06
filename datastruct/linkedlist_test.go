package datastruct

import (
    "testing"
)

func TestLinkedList(t *testing.T) {
    ll := LinkedList{}
    ll.PushHead("Head1", 3)
    ll.PushHead("Head2", 2)
    ll.PushHead("Head3", 1)

    tail := ll.Tail().value
    if tail != 3 {
        t.Errorf("Wrong tail = %d; want 3", tail)
    }

    ll.PushBack("Back1", 4)
    ll.PushBack("Back2", 5)
    ll.PushBack("Back3", 6)

    visit := ll.Visit()
    correct := "1,2,3,4,5,6"
    if visit != correct {
        t.Errorf("Wrong visit = %s; want %s", visit, correct)
    }

    return

    ll.Reverse()
    visit = ll.Visit()
    correct = "6,5,4,3,2,1"
    if visit != correct {
        t.Errorf("Wrong visit = %s; want %s", visit, correct)
    }

    ll.Reverse()
    visit = ll.Visit()
    correct = "1,2,3,4,5,6"
    if visit != correct {
        t.Errorf("Wrong visit = %s; want %s", visit, correct)
    }

    ll.Remove("Head2")
    ll.Remove("Back2")
    ll.Remove("Head3")

    visit = ll.Visit()
    correct = "3,4,6"
    if visit != correct {
        t.Errorf("Wrong visit = %s; want %s", visit, correct)
    }

    ll.Remove("Back3")
    ll.Remove("Back1")
    ll.Remove("Head1")

    visit = ll.Visit()
    correct = ""
    if visit != correct {
        t.Errorf("Wrong visit = %s; want %s", visit, correct)
    }

    l := ll.Length()
    if l != 0 {
        t.Errorf("Length is %d; want 0", l)
    }
}

