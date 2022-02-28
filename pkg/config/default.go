package config

var DefaultConfig = []byte(`
server:
  host: 127.0.0.1
  port: 8080
database:
  host: 127.0.0.1
  port: 5444
  database: wagerdb
  username: wager
  password: wager
`)
