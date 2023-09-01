package users

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"

	"github.com/marlonlorram/baseapi/entities"
	"github.com/marlonlorram/baseapi/usecases/users"
)

type userRepository struct {
	users *mongo.Collection
}

// FindName busca um usuário pelo nome no banco de dados.
// Retorna um erro específico se o usuário não for encontrado.
func (r *userRepository) FindName(ctx context.Context, name string) (*entities.User, error) {
	user := new(entities.User)
	err := r.users.FindOne(ctx, bson.M{"name": name}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, entities.NewNotFound("Usuário não encontrado!")
		}

		return nil, err
	}

	return user, nil
}

// List lista todos os usuários.
// Retorna uma lista de ponteiros para entidades de usuários
// e um erro se algo der errado.
func (r *userRepository) List(ctx context.Context) ([]*entities.User, error) {
	users := make([]*entities.User, 0, 8)

	cur, err := r.users.Find(ctx, bson.M{}, options.Find())
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var user entities.User
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// Insert insere um novo usuário no banco de dados e retorna o usuário com seu ID gerado.
func (r *userRepository) Insert(ctx context.Context, user *entities.User) (*entities.User, error) {
	result, err := r.users.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	user.ID = result.InsertedID.(primitive.ObjectID)

	return user, nil
}

// newUserRepository inicializa e retorna uma nova instância do repositório de usuários.
// Ele também cria um índice único para o nome do usuário para garantir que não haja usuários duplicados.
func newUserRepository(db *mongo.Database) (users.Repository, error) {
	collection := db.Collection("users")

	indexModel := mongo.IndexModel{
		Keys:    bson.M{"name": 1},
		Options: options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		return nil, err
	}

	return &userRepository{users: collection}, nil
}

func Build() fx.Option {
	return fx.Provide(newUserRepository)
}
