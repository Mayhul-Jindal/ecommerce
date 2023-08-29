package token_test

import (
	"testing"
	"time"

	errs "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/errors"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/token"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/util"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/require"
)

func TestJwtMaker(t *testing.T) {
	userID := util.RandomBigInt(1, 10)
	duration := time.Minute
	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, payload, err := testJwtMaker.CreateToken(userID, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = testJwtMaker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	require.NotZero(t, payload.ID)
	require.Equal(t, payload.UserID, userID)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredToken(t *testing.T) {
	userID := util.RandomBigInt(1, 10)
	duration := time.Minute

	token, payload, err := testJwtMaker.CreateToken(userID, -duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = testJwtMaker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, errs.ErrorExpiredToken.Error())
	require.Nil(t, payload)
}

func TestInvalidJWTTokenAlgNone(t *testing.T) {
	// this whole thing if sent by frontend
	payload, err := token.NewPayload(util.RandomBigInt(1, 10), time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	payload, err = testJwtMaker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, errs.ErrorInvalidToken.Error())
	require.Nil(t, payload)
}
