// this acts as a simple in memory db to handle users and refresh tokens.
// replace these functions with actual calls to your db

package db

import (
	"context"
	"csrf/db/models"
	"csrf/randomstrings"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)


var client *mongo.Client
var userCollection *mongo.Collection
var tokenCollection *mongo.Collection
var yourDBName = "csrf"

func InitDB() error {
    mongoURL := "replace with your connection string"
    clientOptions := options.Client().ApplyURI(mongoURL) // Update the MongoDB connection string
    var err error
    client, err = mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        return err
    }

    // Initialize user and token collections
    userCollection = client.Database(yourDBName).Collection("users")
    tokenCollection = client.Database(yourDBName).Collection("refresh_tokens")

    return nil
}

// password is hashed before getting here
func StoreUser(username string, password string, role string) (string, error) {
    uuid, err := randomstrings.GenerateRandomString(32)
    if err != nil {
        return "", err
    }

	passwordHash, hashErr := generateBcryptHash(password)
	if hashErr != nil {
		return "", hashErr
	}

    user := models.User{
        UUID:         uuid,
        Username:     username,
        PasswordHash: passwordHash,
        Role:         role,
    }

    user.UUID = uuid

    _, err = userCollection.InsertOne(context.TODO(), user)
    if err != nil {
        return "", err
    }

    return uuid, nil
}


func DeleteUser(uuid string) error {
    _, err := userCollection.DeleteOne(context.TODO(), bson.M{"UUID": uuid})
    if err != nil {
        return err
    }
    return nil
}


func FetchUserById(uuid string) (models.User, error) {
    filter := bson.M{"UUID": uuid}
    var user models.User
    err := userCollection.FindOne(context.TODO(), filter).Decode(&user)
    if err != nil {
        return models.User{}, err
    }
    return user, nil
}

// returns the user and the userId or an error if not found
func FetchUserByUsername(username string) (models.User, string, error) {
    filter := bson.M{"Username": username}
    var user models.User
    err := userCollection.FindOne(context.TODO(), filter).Decode(&user)
    if err != nil {
        return models.User{},"", err
    }
    return user, "", nil
}


func StoreRefreshToken() (string, error) {
    jti, err := randomstrings.GenerateRandomString(32)
    if err != nil {
        return jti, err
    }

    _, err = tokenCollection.InsertOne(context.TODO(), bson.M{"JTI": jti, "Status": "valid"})
    if err != nil {
        return jti, err
    }

    return jti, nil
}


func DeleteRefreshToken(jti string) error {
    _, err := tokenCollection.DeleteOne(context.TODO(), bson.M{"JTI": jti})
    if err != nil {
        return err
    }
    return nil
}


func CheckRefreshToken(jti string) (bool) {
    filter := bson.M{"JTI": jti}
    var result struct {
        JTI string `bson:"JTI"`
    }

    err := tokenCollection.FindOne(context.TODO(), filter).Decode(&result)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return false // Token not found, return false
        }
        return false // An error occurred while checking the token
    }

    return result.JTI == jti // Token found and matches the provided JTI
}


func LogUserIn(username string, password string) (models.User, string, error) {
    user, uuid, userErr := FetchUserByUsername(username)
    if userErr != nil {
        return models.User{}, "", userErr
    }

    // Compare the provided password with the stored hash
    if err := checkPasswordAgainstHash(user.PasswordHash, password); err != nil {
        return models.User{}, "", errors.New("Invalid password")
    }

    return user, uuid, nil
}


func generateBcryptHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash[:]), err
}

func checkPasswordAgainstHash(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
