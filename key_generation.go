package main

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/tendermint/tendermint/crypto/ed25519"
)

func main() {
	// Generate a new ed25519 private key
	privKey := ed25519.GenPrivKey()

	// Encode the public key in base64 format
	pubKeyBase64 := base64.StdEncoding.EncodeToString(privKey.PubKey().Bytes()[:])

	// Calculate the address and capitalize it
	address := hex.EncodeToString(privKey.PubKey().Address().Bytes())
	capitalizedAddress := strings.ToUpper(address)

	// Encode the private key in base64 format
	privKeyBase64 := base64.StdEncoding.EncodeToString(privKey[:])

	// Create a map for the JSON content
	privValidatorKeys := map[string]interface{}{
		"address": capitalizedAddress,
		"pub_key": map[string]interface{}{
			"type":  "tendermint/PubKeyEd25519",
			"value": pubKeyBase64,
		},
		"priv_key": map[string]interface{}{
			"type":  "tendermint/PrivKeyEd25519",
			"value": privKeyBase64,
		},
	}

	// Marshal the priv_validator_keys map to JSON
	jsonData, err := json.MarshalIndent(privValidatorKeys, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling priv_validator_keys to JSON:", err)
		os.Exit(1)
	}

	// Save the JSON data to priv_validator_keys.json file
	file, err := os.Create("priv_validator_keys.json")
	if err != nil {
		fmt.Println("Error creating priv_validator_keys.json file:", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing to priv_validator_keys.json file:", err)
		os.Exit(1)
	}

	fmt.Println("priv_validator_keys.json generated successfully!")
}
