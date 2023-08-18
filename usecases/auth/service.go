package auth

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/fx"
	"golang.org/x/crypto/bcrypt"

	"github.com/marlonlorram/baseapi/config"
	"github.com/marlonlorram/baseapi/entities"
	"github.com/marlonlorram/baseapi/usecases/users"
)

type Service struct {
	repo   users.Repository
	config *config.Config
}

// Login autentica um usuário com base em seu nome de usuário e senha.
// Retorna um token JWT se a autenticação for bem-sucedida.
func (s *Service) Login(ctx context.Context, username, password string) (*entities.AuthResult, error) {
	user, err := s.repo.FindName(ctx, username)
	if err != nil {
		return nil, err
	}

	err = s.checkPassword(user, password)
	if err != nil {
		return nil, entities.NewForbidden("Falha no login")
	}

	token, err := s.generateToken(user.ID.Hex(), user.Name)
	if err != nil {
		return nil, err
	}

	return &entities.AuthResult{
		User:  user,
		Token: *token,
	}, nil
}

// Register registra um novo usuário.
func (s *Service) Register(ctx context.Context, username, password string) (*entities.User, error) {
	if !s.config.Server.AllowRegister {
		return nil, entities.NewBadRequest("Novos registros estão desabilitados")
	}
	if len(password) < 6 {
		return nil, entities.NewBadRequest("Senha muito curta")
	}
	if len(username) == 0 {
		return nil, entities.NewBadRequest("Nome de usuário não fornecido")
	}

	user := &entities.User{
		Name: username,
	}

	err := s.setPassword(user, password)
	if err != nil {
		return nil, err
	}

	user, err = s.repo.Insert(ctx, user)
	if err != nil && strings.Contains(err.Error(), "E11000") {
		return nil, entities.NewConflict("Nome de usuário não está disponível")
	}

	return user, nil
}

// SetPassword criptografa a senha fornecida e a armazena no objeto usuário.
func (s *Service) setPassword(user *entities.User, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = hash
	return nil
}

// CheckPassword compara a senha fornecida com a senha criptografada armazenada no objeto usuário.
func (s *Service) checkPassword(user *entities.User, password string) error {
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		return err
	}

	return nil
}

// generateToken cria um novo JWT para um usuário específico.
// Ele inclui o ID do usuário e o nome de usuário nas reivindicações do token.
func (s *Service) generateToken(userID string, username string) (*string, error) {
	claims := jwt.MapClaims{
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
		"sub":      userID,
		"username": username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(s.config.JWT.Secret))
	if err != nil {
		return nil, err
	}

	return &signedToken, nil
}

// parseToken analisa e valida o token JWT fornecido.
// Retorna informações do usuário do token se ele for válido.
func (s *Service) parseToken(token string) (*entities.User, error) {
	tk, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
		}
		return []byte(s.config.JWT.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := tk.Claims.(jwt.MapClaims); ok && tk.Valid {
		userID, _ := claims["sub"].(string)
		userName, _ := claims["username"].(string)

		objectId, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			return nil, err
		}

		return &entities.User{
			ID:   objectId,
			Name: userName,
		}, nil
	}
	return nil, fmt.Errorf("token inválido")
}

// ValidateToken verifica a validade do token JWT fornecido.
func (s *Service) ValidateToken(token string) (*entities.User, error) {
	return s.parseToken(token)
}

// newAuthService cria uma nova instância do serviço de autenticação.
func newAuthService(repository users.Repository, config *config.Config) *Service {
	return &Service{repository, config}
}

func Build() fx.Option {
	return fx.Provide(newAuthService)
}
