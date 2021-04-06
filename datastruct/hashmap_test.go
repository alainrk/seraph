package datastruct

import (
	// "fmt"
    "testing"
    "strconv"
)

func TestHashMap(t *testing.T) {
    hm := HashMap{}
    hm.Add("key1", 32)
    hm.Add("key1", 34)
    hm.Add("key1", 31)

	given := hm.Get("key1")
    expected := 32
    if given != expected {
        t.Errorf("Explicit key update. Wrong value [key1] = %d; want %d", given, expected)
    }

    // Massive add
    for i := 0; i < 1024; i++ {
        hm.Add(strconv.Itoa(i), i)
    }

    given = hm.Get("112")
    expected = 112
    if given != expected {
        t.Errorf("Explicit key update. Wrong value [key1] = %d; want %d", given, expected)
    }

    given = hm.Get("476")
    expected = 476
    if given != expected {
        t.Errorf("Explicit key update. Wrong value [key1] = %d; want %d", given, expected)
    }

    // fmt.Println(hm.Print())
}

