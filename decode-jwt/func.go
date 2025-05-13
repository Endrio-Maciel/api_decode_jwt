package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// aa

type TokenRequest struct {
	Token string `json:"token"`
}

func Handle(ctx context.Context, in io.Reader) (interface{}, error) {
	var req TokenRequest
	if err := json.NewDecoder(in).Decode(&req); err != nil {
		return nil, fmt.Errorf("erro ao decodificar entrada: %w", err)
	}

	parts := strings.Split(req.Token, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("token inválido")
	}

	payloadEncoded := parts[1]
	payloadBytes, err := base64.URLEncoding.DecodeString(payloadEncoded)
	if err != nil {
		return nil, fmt.Errorf("erro ao decodificar payload: %w", err)
	}

	var payload map[string]interface{}
	if err := json.Unmarshal(payloadBytes, &payload); err != nil {
		return nil, fmt.Errorf("payload inválido: %w", err)
	}

	return payload, nil
}
