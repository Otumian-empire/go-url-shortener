package repository

import (
	"database/sql"
	"fmt"

	"github.com/otumian-empire/go-url-shortener/entity"
	"github.com/otumian-empire/go-url-shortener/util"
)

// Repository for the url entity, if there is another entity, create its store (repository)
type UrlStore struct {
	*sql.DB
}

func (urlStore *UrlStore) CreateUrl(shortUrl, originalUrl string) (int, error) {
	result, err := urlStore.Exec(INSERT_QUERY, shortUrl, originalUrl)

	if util.IsNotNil(err) {
		return DEFAULT_INT, err
	}

	rowsAffected, err := result.RowsAffected()

	if util.IsNotNil(err) {
		return DEFAULT_INT, err
	}

	if rowsAffected < DEFAULT_INT {
		return DEFAULT_INT, fmt.Errorf(INSERT_ERROR)
	}

	lastInsertId, err := result.LastInsertId()

	if util.IsNotNil(err) {
		return DEFAULT_INT, err
	}

	return int(lastInsertId), nil
}

func (urlStore *UrlStore) Url(id int) (entity.Url, error) {
	row := urlStore.QueryRow(SELECT_URL_BY_ID_QUERY, id)

	var url entity.Url

	err := row.Scan(&url.Id, &url.Short, &url.Original, &url.CreatedAt)

	if util.IsNotNil(err) {
		return entity.Url{}, err
	}

	return url, nil
}

func (urlStore *UrlStore) Urls() ([]entity.Url, error) {
	var urls []entity.Url

	rows, err := urlStore.Query(SELECT_URL_QUERY)

	if util.IsNotNil(err) || util.IsNotNil(rows.Err()) {
		return []entity.Url{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var url entity.Url
		err = rows.Scan(&url.Id, &url.Short, &url.Original, &url.CreatedAt)

		if util.IsNotNil(err) && err == sql.ErrNoRows {
			return []entity.Url{}, fmt.Errorf("%s: %w", READING_ERROR, err)
		}

		urls = append(urls, url)
	}

	return urls, nil
}

func (urlStore *UrlStore) DeleteUrl(id int) error {
	result, err := urlStore.Exec(DELETE_QUERY, id)

	if util.IsNotNil(err) {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if util.IsNotNil(err) {
		return err
	}

	if rowsAffected < DEFAULT_INT {
		return fmt.Errorf(DELETE_ERROR)
	}

	return nil
}

func (urlStore *UrlStore) OriginalUrl(shortUrl string) (entity.Url, error) {
	row := urlStore.QueryRow(SELECT_URL_BY_SHORT_URL_QUERY, shortUrl)

	var url entity.Url

	err := row.Scan(&url.Id, &url.Short, &url.Original, &url.CreatedAt)

	if util.IsNotNil(err) {
		return entity.Url{}, err
	}

	return url, nil
}
