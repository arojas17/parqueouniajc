package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de conexion a la base de datos*/
var MongoCN = ConnectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://offcaos17:CIuJrCER13vcIZvQ@parking.4bffa.mongodb.net/dbparking?retryWrites=true&w=majority")

/*ConnectarBD es la funcion de conectar a la conexion la la base de datos*/
func ConnectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión Exitosa con la BD")
	return client
}

/*CheckConnection es el ping a la base de datos*/
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
