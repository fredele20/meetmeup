package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"meetmeup/graph/model"
)

type MeetupsRepo struct{}

//type Repository interface {
//	Save()
//}

func (m *MeetupsRepo) GetMeetups(filter *model.MeetupFilter, limit, offset *int) ([]*model.Meetup, error) {
	var meetups []*model.Meetup
	limitInt64 := int64(*limit)
	offsetInt64 := int64(*offset)

	option := options.Find()

	//option.SetSort(bson.D{{"name", *filter.Name}})
	option.SetSort(map[string]int{"name": 1})
	option.SetLimit(limitInt64)
	option.SetSkip(offsetInt64)
	cursor, err := MeetupsCollection.Find(context.TODO(), bson.M{}, option)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var meetup *model.Meetup
		cursor.Decode(&meetup)
		meetups = append(meetups, meetup)
	}
	return meetups, nil
}

func (m *MeetupsRepo) CreateMeetup(meetup *model.Meetup) (*model.Meetup, error) {
	_, err := MeetupsCollection.InsertOne(context.TODO(), meetup)
	if err != nil {
		log.Fatal(err)
	}
	return meetup, nil
}

func (m *MeetupsRepo) GetByID(id string) (*model.Meetup, error) {
	var meetup model.Meetup
	_, err := MeetupsCollection.Find(context.TODO(), bson.M{"_id": id})
	return &meetup, err
}

func (m *MeetupsRepo) Update(meetup *model.Meetup) (*model.Meetup, error) {
	filter := bson.M{"_id": meetup.ID}
	update := bson.D{
		{"$set", bson.D{
			{"name", meetup.Name},
			{"description", meetup.Description},
		}},
	}
	_, err := MeetupsCollection.UpdateOne(context.TODO(), filter, update)
	return meetup, err
}

func (m *MeetupsRepo) Delete(meetup *model.Meetup) error {
	filter := bson.M{"_id": meetup.ID}
	_, err := MeetupsCollection.DeleteOne(context.TODO(), filter)
	return err
}

func (m *MeetupsRepo) GetMeetupsForUser(user *model.User) ([]*model.Meetup, error) {
	var meetups []*model.Meetup
	cursor, err := MeetupsCollection.Find(context.TODO(), bson.M{"userId": user.ID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var meetup *model.Meetup
		cursor.Decode(&meetup)
		meetups = append(meetups, meetup)
	}
	return meetups, nil
}
