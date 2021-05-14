package storage

type MapStorage map[string]string

func (m MapStorage) GetOriginUrl(shortURL string) (originURL string, found bool) {
	originURL, found = m[shortURL]
	return
}

func (m MapStorage) Save(originURL, shortURL string) error {
	m[shortURL] = originURL
	return nil
}
