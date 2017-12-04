package stockwell

import (
  "io/ioutil"
  "net/http"
)

type Reddit struct {
  BaseURL string
  conn *Connection
  httpClient *http.Client
}

func NewReddit() (r *Reddit) {
  r = &Reddit{
    BaseURL: "https://www.reddit.com",
    conn: NewConnection(),
    httpClient: new(http.Client),
  }

  return r
}

func (r *Reddit) FetchListing(name string) (body string, err error) {
  var respBody []byte

  respBody, err = r.get(r.BaseURL + "/r/" + name + ".json")
  if err != nil {
    return body, err
  }

  body = string(respBody)

  return body, err
}

func (r *Reddit) get(u string) (body []byte, err error) {
  var resp *http.Response

  resp, err = r.httpClient.Get(u)
  if err != nil {
    return body, err
  }
  defer resp.Body.Close()

  return ioutil.ReadAll(resp.Body)
}
