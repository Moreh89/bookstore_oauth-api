package accesstoken

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	// if expirationTime != 24 {
	// 	t.Error("expiration time modified")
	// }
	assert.EqualValues(t, 24, expirationTime, "expiration time modified")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	if at.IsExpired() {
		t.Error("new access token should not be expired")
	}
	if at.AccessToken != "" {
		t.Error("access token should not have defined token access id")
	}
	if at.UserId != 0 {
		t.Error("access token should not have an associated userId")
	}
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	// if !at.IsExpired() {
	// 	t.Error("empty access token should be expired by default")
	// }
	assert.True(t, at.IsExpired(),"empty access token should be expired by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	// if at.IsExpired() {
	// 	t.Error("access token should not be expired")
	// }
	assert.False(t, at.IsExpired(), "access token should not be expired")
}
