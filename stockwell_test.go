package stockwell

import (
  "testing"
)

func TestCreateReddit(t *testing.T) {
  var reddit *Reddit = NewReddit()

  if reddit == nil {
    t.Error("Found nil while creating a Reddit")
  }
}

func TestGetListing(t *testing.T) {
  var err error
  var body string
  var reddit *Reddit = NewReddit()

  body, err = reddit.FetchListing("youtube")
  if err != nil {
    t.Error(err)
  }

  if body == "" {
    t.Error("Empty body")
  }
}
