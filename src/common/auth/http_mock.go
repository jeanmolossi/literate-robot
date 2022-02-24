package auth

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	commonerrors "github.com/jeanmolossi/literate-robot/src/common/errors"
	"github.com/jeanmolossi/literate-robot/src/common/server/httperr"
)

type (
	User struct {
		Id    int
		Email string
		Role  string

		Username string
	}

	ctxKey int
)

const (
	userContextKey ctxKey = iota
)

var (
	// if we expect that the user of the function may be interested with concrete error,
	// it's a good idea to provide variable with this error
	NoUserInContextError = commonerrors.NewAuthorizationError("no user in context", "no-user-found")
)

// HttpMockMiddleware is used in the local environment (which doesn't depend on Firebase)
func HttpMockMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var claims jwt.MapClaims
		token, err := request.ParseFromRequest(
			r,
			request.AuthorizationHeaderExtractor,
			func(token *jwt.Token) (i interface{}, e error) {
				return []byte("mock_secret"), nil
			},
			request.WithClaims(&claims),
		)
		if err != nil {
			httperr.BadRequest("unable-to-get-jwt", err, w, r)
			return
		}

		if !token.Valid {
			httperr.BadRequest("invalid-jwt", nil, w, r)
			return
		}

		ctx := context.WithValue(r.Context(), userContextKey, User{
			Id:       int(claims["user_id"].(float64)),
			Email:    claims["email"].(string),
			Role:     claims["role"].(string),
			Username: claims["name"].(string),
		})
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func UserFromCtx(ctx context.Context) (User, error) {
	u, ok := ctx.Value(userContextKey).(User)
	if ok {
		return u, nil
	}

	return User{}, NoUserInContextError
}
