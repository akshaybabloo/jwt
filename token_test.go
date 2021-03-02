package jwt

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInTimeSpan(t *testing.T) {
	t1 := time.Date(1984, time.November, 3, 10, 23, 34, 0, time.UTC)
	t2 := time.Date(1984, time.November, 3, 11, 23, 34, 0, time.UTC)
	t3 := time.Date(1984, time.November, 3, 13, 0, 0, 0, time.UTC)

	span := InTimeSpan(t1, t3, t2)
	assert.True(t, span)
}

func TestToken_Valid(t *testing.T) {

	tk := Token{
		AccessToken: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTQ1ODY5ODIsIm5iZiI6MTYxNDU4MzM4MiwidmVyIjoiMS4wIiwiaXNzIjoiaHR0cHM6Ly9nb2xsYWhhbGxpYXV0aC5iMmNsb2dpbi5jb20vNDkyMDY1NmQtMzZlMC00YjM1LTg5NGQtMGEwYTE0YmVjOGIzL3YyLjAvIiwic3ViIjoiNzYyNWY0MDktMWZhNy00ZDc4LTkzYjAtNWMzNDAyNTVjZGYyIiwiYXVkIjoiMTIzNDU2Nzg5Iiwibm9uY2UiOiJkZWZhdWx0Tm9uY2UiLCJpYXQiOjE2MTQ1ODMzODIsImF1dGhfdGltZSI6MTYxNDU4MzM4Miwib2lkIjoic29tZS1vaWQiLCJnaXZlbl9uYW1lIjoiQWtzaGF5IiwiZmFtaWx5X25hbWUiOiJHb2xsYWhhbGxpIiwidGZwIjoiQjJDXzFfU2lnblVwU2lnbkluRmxvdyJ9.CNbe3qUhpUR7U51e9k1miAdQP3ZEbCDip67MlvgTpFsFmB3nbjp7wi8e-66cPoS9z_hmQP3wLc5I8KE-b4_cFCzHZkhuQPA0Mhi9mBAZ_tAPSXNaDeiX1FEalsiH_sWuHWnojxSwlSxeKL9Tlh_0u5vXaABILdDeRWOTBJDHZ5I2BgIk8J_hI-fXDvBb0wfjI4mQe7lQqDZLVos4mlA1Uhpz2wN7Lorc31PZo02STbk5S1j1BMk4eQUS_OHTKZfAlUmfV9giWEMr7qk21eSy_HlybVKQgCB9Cde_rmNWJETyw6X712QGYWaDkDrwocQ99wR6rALWVkkTOhmL6CRT5Q",
	}
	fmt.Println(tk.Valid())

}
