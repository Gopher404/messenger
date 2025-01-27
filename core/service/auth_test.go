package service

import (
	"github.com/stretchr/testify/require"
	cache "messanger/adapters/cache/local"
	"messanger/config"
	"messanger/core/ports"
	"testing"
)

var s *AuthService
var c ports.Cache

const testId = 123

func init() {
	c = cache.NewCache()

	s = NewAuthService(c, nil, &config.AuthServiceConfig{
		AccessTokenTTLMin:   15,
		RefreshTokenTTLDays: 30,
	})
}

func TestAuthService_AccessToken(t *testing.T) {
	token, _, err := s.newAccessToken(testId)
	require.NoError(t, err, "create access token")

	id, err := s.CheckAccessToken(token)
	require.NoError(t, err, "check access token")
	require.Equal(t, testId, id)
}

func TestAuthService_UpdateTokens(t *testing.T) {
	refreshToken := s.newRefreshToken()
	err := s.cache.Set(refreshToken, testId)
	require.NoError(t, err, "set access token")

	tokens, err := s.UpdateTokens(refreshToken)
	require.NoError(t, err, "update tokens")

	id, err := s.CheckAccessToken(tokens.AccessToken)
	require.NoError(t, err, "check access token")

	require.Equal(t, testId, id)
}
