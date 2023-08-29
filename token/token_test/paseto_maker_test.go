package token_test

import (
	"testing"
	"time"

	errs "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/errors"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/util"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	userID := util.RandomBigInt(1, 10)
	duration := time.Minute
	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, payload, err := testPasetoMaker.CreateToken(userID, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = testPasetoMaker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	require.NotZero(t, payload.ID)
	require.Equal(t, payload.UserID, userID)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExipredToken(t *testing.T) {
	userID := util.RandomBigInt(1, 10)
	duration := time.Minute

	token, payload, err := testPasetoMaker.CreateToken(userID, -duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = testPasetoMaker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, errs.ErrorExpiredToken.Error())
	require.Nil(t, payload)
}
