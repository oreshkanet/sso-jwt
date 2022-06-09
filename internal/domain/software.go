package domain

type Software struct {
	Id         string `db:"id"`
	Password   string `db:"password"`
	SigningKey string `db:"signing_key"`
}

type ExternalClient struct {
	Id        string `db:"id"`
	Secret    string `db:"secret"`
	GrantType string `db:"grant_type"`
	Scope     string `db:"scope"`
}

type User struct {
	Id       string `db:"id"`
	Password string `db:"password"`
	Scope    string `db:"scope"`
}
