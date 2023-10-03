package cryptor

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "fmt"
    "golang.org/x/crypto/scrypt"
    "io"
)


func EncryptWithAES(password string, plaintext []byte) (string, error) {
    // Generate a random salt.
    salt := make([]byte, 16)
    if _, err := io.ReadFull(rand.Reader, salt); err != nil {
        return "", err
    }

    // Derive an encryption key from the password and salt.
    key, salt, err := deriveKeyFromPassword(password, salt)
    if err != nil {
        return "", err
    }

    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }

    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }

    // Generate a random nonce (IV).
    nonce := make([]byte, gcm.NonceSize())
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return "", err
    }

    // Encrypt the plaintext.
    ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

    // Prepend salt to ciphertext for storage/transmission.
    saltAndCiphertext := append(salt, ciphertext...)

    // Encode salt and ciphertext combination to a base64 string.
    encodedData := base64.StdEncoding.EncodeToString(saltAndCiphertext)

    return encodedData, nil
}

func DecryptWithAES(password string, encodedData string) ([]byte, error) {
    // Decode the encoded data from base64.
    saltAndCiphertext, err := base64.StdEncoding.DecodeString(encodedData)
    if err != nil {
        return nil, err
    }

    // Extract the salt.
    salt := saltAndCiphertext[:16]

    // Derive an encryption key from the password and the extracted salt.
    key, _, err := deriveKeyFromPassword(password, salt)
    if err != nil {
        return nil, err
    }

    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }

    nonceSize := gcm.NonceSize()
    if len(saltAndCiphertext) < nonceSize {
        return nil, fmt.Errorf("ciphertext is too short")
    }

    nonce, ciphertext := saltAndCiphertext[16:nonceSize+16], saltAndCiphertext[nonceSize+16:]

    // Decrypt the ciphertext.
    plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        return nil, err
    }

    return plaintext, nil
}


func deriveKeyFromPassword(password string, salt []byte) ([]byte, []byte, error) {
    // Use scrypt to derive a key from the password and salt.
    key, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, 32)
    if err != nil {
        return nil, nil, err
    }

    return key, salt, nil
}
