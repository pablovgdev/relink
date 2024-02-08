package redirect

import (
	"github.com/pablovgdev/relink/internal/connection"
)

func GetRedirectUrlByPath(path string) (string, error) {
	db := connection.GetConnection()
	row := db.QueryRow("SELECT url FROM redirects WHERE path = ?", path)

	var url string
	err := row.Scan(&url)
	if err != nil {
		return "", err
	}

	return url, nil
}

func PostRedirect(path, url string) error {
	db := connection.GetConnection()

	_, err := db.Exec("INSERT INTO redirects (path, url) VALUES (?, ?)", path, url)
	if err != nil {
		return err
	}

	return nil
}
