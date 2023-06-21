package gocache

type Driver interface {
	Set(string, []byte) error
	Get(string) ([]byte, error)
}

type Cache struct {
	driver Driver
}

func New(driver Driver) *Cache {
	return &Cache{
		driver: driver,
	}
}

func (c *Cache) Set(key string, value []byte) error {
	if err := c.driver.Set(key, value); err != nil {
		return err
	}
	return nil
}

func (c *Cache) Get(key string) ([]byte, error) {
	d, err := c.driver.Get(key)
	if err != nil {
		return []byte{}, err
	}
	return d, nil
}
