package multipass

import (
    "fmt"
	"encoding/base64"
	"encoding/json"
    "crypto/sha256"
    "crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"io"
)


// Struct for keys used to encrypt and sign the token


type MultipassKeys struct {
	EncryptionKey []byte
	SignatureKey  []byte
}

// Generate a pair of keys for encryption and signing

func GenerateKeys(secret string) *MultipassKeys {
    hash := sha256.New()
	hash.Write([]byte(secret))
	key := hash.Sum(nil)
	return &MultipassKeys{
		EncryptionKey: key[0:16] ,
		SignatureKey:  key[16:],
	}
}

// Encrypt the payload with the encryption key

func Encrypt(key, text []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    ciphertext := make([]byte, aes.BlockSize + len(text))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return nil, err
    }
    cfb := cipher.NewCFBEncrypter(block, iv)
    cfb.XORKeyStream(ciphertext[aes.BlockSize:], text)
    return ciphertext, nil
}

// Sign the token with the signature key

func Sign(key, text []byte) ([]byte, error) {
    h := hmac.New(sha256.New, key)
    h.Write(text)
    return h.Sum(nil), nil
}


// Main logic for generating the token

func GenerateToken(secret string, payload interface{}) (string, error) {

    keys := GenerateKeys(secret)
    payloadBytes, err := json.Marshal(payload)
    if err != nil {     
        return "", err
    }

    ciphertext, err := Encrypt([]byte(keys.EncryptionKey), payloadBytes)
	if err != nil {
        return "", err
	}
    
	signedtext, err := Sign([]byte(keys.SignatureKey), ciphertext)
	if err != nil {
        return "", err
	}

    return base64.URLEncoding.EncodeToString(signedtext), nil

}

// helper func to generate url 

func GenerateURL(token string, store string) (string) {
    return fmt.Sprintf("http://%s.myshopify.com/account/login/multipass/%s", store, token)
}



