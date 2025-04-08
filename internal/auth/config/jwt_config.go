package config

// GetJWTSecret returns JWT secret key
func (c *Config) GetJWTSecret() string {
	return c.JWT.Secret
}

// GetJWTExpiration returns JWT expiration in seconds
func (c *Config) GetJWTExpiration() int64 {
	return int64(c.JWT.Expiration.Seconds())
}
