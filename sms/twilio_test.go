package sms

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSend(t *testing.T) {
	err := Send("hello")
	require.NoError(t, err)
	fmt.Println(status)
}
