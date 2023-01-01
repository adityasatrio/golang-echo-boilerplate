package database

//TODO create interface for db connection  so in the future we can create other implementation other than ent mysql

type (
	DbConnection interface {
		NewSetupConnection() *Client
	}

	Client struct {
		Connection any
	}
)
