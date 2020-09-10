package data


//test data
var users = []User{
  {
    Name:     "Kotume kawauso",
    Email:    "kawauso@gmail.com",
    Password: "kawauso_pass",
  },
  {
    Name:     "Maneki neko",
    Email:    "manekineko@gmail.com",
    Password: "neko_pass",
  },
}

func setup() {
  ThreadDeleteAll()
  SessionDeleteAll()
  UserDeleteAll()
}
