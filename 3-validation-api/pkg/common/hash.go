package common

import (
    "crypto/rand"
    "encoding/base64"
)

const defaultUniqueHashBytes = 12

func UniqueHash() (string, error) {
    b := make([]byte, defaultUniqueHashBytes)
    if _, err := rand.Read(b); err != nil {
        return "", err
    }
    return base64.RawURLEncoding.EncodeToString(b), nil
}
