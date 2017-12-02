package stockwell

type Reddit struct {
  conn *Connection
}

func NewReddit() (r *Reddit) {
  r = &Reddit{
    conn: NewConnection(),
  }

  return r
}
