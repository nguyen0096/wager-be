package config

var DefaultConfig = []byte(`
server:
  port: 8080
database:
  host: 127.0.0.1
  port: 5432
  database: wagerdb
  username: wager
  password: wager
`)
