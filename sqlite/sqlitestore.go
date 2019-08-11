package sqlitecookie

import (
	"fmt"

	"session/sessions"

	gsessions "github.com/gorilla/sessions"

	// "github.com/michaeljs1990/sqlitestore"
	"session/sqlitestore"
)

type Store interface {
	sessions.Store
}

type store struct {
	*sqlitestore.SqliteStore
}

func NewStore(keyPairs ...[]byte) Store {
	// st, err := sqlitestore.NewSqliteStore("./database.db", "sessions", "/", 3600, []byte("<SecretKey>"))
	st, err := sqlitestore.NewSqliteStore("./database.db", "sessions", "/", 3600, keyPairs...)
	if err != nil {
		fmt.Println(err)
	}
	return &store{st}
}

func (c *store) Options(options sessions.Options) {
	c.SqliteStore.Options = &gsessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HttpOnly,
	}
}
