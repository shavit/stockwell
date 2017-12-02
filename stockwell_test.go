package reddit

import (
  "testing"
)

func TestCreateReddit(t *testing.T) {
  var reddit *Reddit = NewReddit()

  if reddit == nil {
    t.Error("Found nil while creating a Reddit")
  }
}
