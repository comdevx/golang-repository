package helper

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const RefreshTokenCollection = "refresh-tokens"

type SignedDetails struct {
	Username string
	ID       int
	jwt.StandardClaims
}

func GenerateToken(username string, uid int) (string, error) {

	secretKey := os.Getenv("JWT_SECRET")
	jwtExpire, err := strconv.Atoi(os.Getenv("JWT_EXPIRED_MINUTES"))

	if err != nil {
		jwtExpire = 15
	}

	claims := SignedDetails{
		Username: username,
		ID:       uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(jwtExpire)).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return token, err
}

// func UpdateAllTokens(accessToken string, refreshToken string, userId primitive.ObjectID, email string) {
// 	db, ctx := dbs.StartDB()
// 	var object model.RefreshToken

// 	refreshTokenExire, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRED_HOURS"))
// 	if err != nil {
// 		refreshTokenExire = 168
// 	}

// 	object.ID = primitive.NewObjectID()
// 	object.RefreshToken = refreshToken
// 	object.UserID = userId
// 	object.Email = email
// 	object.ExpiredAt = time.Now().Add(time.Hour * time.Duration(refreshTokenExire))
// 	object.CreatedAt = time.Now()

// 	refreshTokenCollection := db.Database(os.Getenv("MONGO_DB_NAME")).Collection(RefreshTokenCollection)

// 	_, insertErr := refreshTokenCollection.InsertOne(ctx, object)
// 	if insertErr != nil {
// 		log.Panic(insertErr)
// 		return
// 	}
// }

func ParseToken(token string) (int, string, error) {

	parse, err := jwt.ParseWithClaims(
		token,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		},
	)
	if err != nil {
		return 0, "", err
	}

	claims, ok := parse.Claims.(*SignedDetails)
	if !ok {
		return 0, "", errors.New("Invalid token")
	}

	return claims.ID, claims.Username, nil
}
