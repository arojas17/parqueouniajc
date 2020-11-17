package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Usuario es el modelo de usuario de la base MongoDB */
type Usuario struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre   string             `bson:"nombre" json:"nombre,omitempty"`
	Apellido string             `bson:"apellido" json:"apellidos,omitempty"`
	Email    string             `bson:"email" json:"email,omitempty"`
	Password string             `bson:"password" json:"password,omitempty"`
}
