package redirect

import (
	"github.com/pablovgdev/relink/internal/connection"
)

type Redirect struct {
	ID        int    `json:"id"`
	Path      string `json:"path"`
	URL       string `json:"url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

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

func GetRedirects() ([]Redirect, error) {
	db := connection.GetConnection()
	rows, err := db.Query("SELECT * FROM redirects")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	redirects := []Redirect{}
	for rows.Next() {
		r := Redirect{}
		err := rows.Scan(&r.ID, &r.Path, &r.URL, &r.CreatedAt, &r.UpdatedAt)
		if err != nil {
			return nil, err
		}
		redirects = append(redirects, r)
	}

	return redirects, nil
}

func PostRedirect(path, url string) error {
	db := connection.GetConnection()

	_, err := db.Exec("INSERT INTO redirects (path, url) VALUES (?, ?)", path, url)
	if err != nil {
		return err
	}

	return nil
}
