package config

type Bot struct {
	Token string 
}

type DB struct {
    Name string
    User string
    Password string
    Host string
    Port string
}

func NewBotConfig() *Bot {
    return &Bot{}
}

func NewDataBases() *DB {
    return &DB{}
}