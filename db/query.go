package mongosh

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SMPPConfig struct {
	Host       string `bson:"host"`
	Port       int    `bson:"port"`
	Username   string `bson:"username"`
	Password   string `bson:"password"`
	SystemType string `bson:"system_type"`
}

type ProductsRepo struct {
	Coll *mongo.Collection
}

func NewProductsRepository(db *mongo.Database) *ProductsRepo {
	return &ProductsRepo{Coll: db.Collection("smpp_config")}
}

func (r *ProductsRepo) GetSMPPConfig(ctx context.Context) (*SMPPConfig, error) {
	var config SMPPConfig
	err := r.Coll.FindOne(ctx, bson.M{}).Decode(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
