package config

import (
	"fmt"
	"net"
	"strconv"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/fx"
)

// Prefix é um prefixo opcional usado para variáveis de ambiente.
var Prefix string

// Config representa a configuração geral da aplicação.
type Config struct {
	Server Server `envconfig:"SERVER"`
}

// Server define as propriedades de configuração do servidor.
type Server struct {
	Address string `envconfig:"ADDR" default:"0.0.0.0:7788" desc:"Endereço de escuta do servidor"`
}

// parseConfig lê as variáveis de ambiente, popula a Config e
// valida o endereço do servidor.
func parseConfig() (*Config, error) {
	var c Config

	// Processa as variáveis de ambiente
	if err := envconfig.Process(Prefix, &c); err != nil {
		return nil, fmt.Errorf("falha ao processar variáveis de ambiente: %w", err)
	}

	// Valida o endereço do servidor
	if err := validateListenAddr(c.Server.Address); err != nil {
		return nil, err
	}

	return &c, nil
}

// validateListenAddr verifica se o endereço (IP e porta) fornecido é válido.
// Em caso de falha, retorna um erro específico.
func validateListenAddr(addr string) error {
	host, portStr, err := net.SplitHostPort(addr)
	if err != nil {
		return fmt.Errorf("endereço de escuta inválido: %w", err)
	}

	if net.ParseIP(host) == nil {
		return fmt.Errorf("endereço IP inválido no endereço de escuta: %s", host)
	}

	port, err := strconv.Atoi(portStr)
	if err != nil || port < 1 || port > 65535 {
		return fmt.Errorf("porta inválida no endereço de escuta: %s", portStr)
	}

	return nil
}

func Build() fx.Option {
	return fx.Provide(parseConfig)
}
