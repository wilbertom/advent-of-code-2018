package main

import "testing"

type destroy_entry struct {
  x, y string
  expected bool
}

func TestDestroy(t *testing.T) {
  table := []destroy_entry {
    destroy_entry {"a", "a", false},
    destroy_entry {"a", "A", true},
    destroy_entry {"A", "A", false},
    destroy_entry {"A", "a", true},
    destroy_entry {"B", "A", false},
    destroy_entry {"B", "a", false},
    destroy_entry {"b", "a", false},
  }


  for _, entry := range table {
    if destroy(entry.x, entry.y) != entry.expected {
      t.Error("Failed with ", entry.x, " and ", entry.y)
    }
  }
}

func TestTrigger(t *testing.T) {
  s := "dabAcCaCBAcCcaDA"
  expected := "dabCBAcaDA"
  actual := *trigger(&s)

  if actual != expected {
    t.Error("Failed with ", actual)
  }

}
