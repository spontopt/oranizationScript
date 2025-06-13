package data

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDb struct {
	Db *mongo.Database
}

func DBConnect() (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(DBUrl))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	db := client.Database(DBName)

	return db, nil
}

func (m MongoDb) BulkSaveUserOrganizations(users []UserOrganization) error {

	coll := m.Db.Collection("UserOrganizations")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	models := make([]mongo.WriteModel, len(users))

	for i, o := range users {
		m := mongo.
			NewReplaceOneModel().
			SetFilter(bson.M{"organizationId": o.OrganizationID, "subscriptionId": o.SubscriptionID, "userId": o.UserID}).
			SetReplacement(o).SetUpsert(true)
		models[i] = m
	}

	opt := options.BulkWrite()
	_, err := coll.BulkWrite(ctx, models, opt)

	return err

}

func (m MongoDb) SaveOrganization(organization Organization) Organization {

	coll := m.Db.Collection("Organizations")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	filter := bson.M{"subscriptionId": organization.SubscriptionID}
	options := options.FindOneAndReplace().SetUpsert(true).SetReturnDocument(options.After)
	result := coll.FindOneAndReplace(ctx, filter, organization, options)

	o := Organization{}
	err := result.Decode(&o)
	if err != nil {
		panic(err)
	}

	return o
}

func (m MongoDb) GetUsers(subscriptionId string) []IUser {

	coll := m.Db.Collection("Users")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	cursor, err := coll.Find(ctx, bson.M{"subscriptionId": subscriptionId})
	if err != nil {
		panic(err)
	}
	var users []IUser
	cursor.All(ctx, &users)
	return users
}

func (m MongoDb) SaveRoles(roles []Profile, organizationId string, subscriptionId string) {

	coll := m.Db.Collection("Roles")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	for i := range roles {
		roles[i].OrganizationID = organizationId
		roles[i].SubscriptionID = subscriptionId
		roles[i].ID = primitive.NilObjectID
	}

	models := make([]mongo.WriteModel, len(roles))

	for i, o := range roles {
		m := mongo.
			NewReplaceOneModel().
			SetFilter(bson.M{"id": o.LegacyID, "subscriptionId": o.SubscriptionID}).
			SetReplacement(o).SetUpsert(true)
		models[i] = m
	}

	opt := options.BulkWrite()
	_, err := coll.BulkWrite(ctx, models, opt)

	if err != nil {
		panic(err)
	}
	return
}

func (m MongoDb) GetAllUserOrganizations() []UserOrganization {
	coll := m.Db.Collection("UserOrganizations")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		panic(err)
	}
	var users []UserOrganization
	cursor.All(ctx, &users)
	return users

}
