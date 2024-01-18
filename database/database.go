package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/akhil/gql/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	db *mongo.Client
}

func databaseClientInstance() *DB {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientDb, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/?compressors=snappy,zlib,zstd"))
	if err != nil {
		panic(err)
	}
	return &DB{
		db: clientDb,
	}

}

func (d *DB) AddJobListing(newJob model.JobListing) *model.JobListing {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	collection := d.db.Database("jobsListing").Collection("jobs")

	defer cancel()
	res, err := collection.InsertOne(ctx, newJob)
	if err != nil {
		log.Fatal(err)
	}

	return &model.JobListing{
		ID:          res.InsertedID.(primitive.ObjectID).Hex(),
		Description: newJob.Description,
		URL:         newJob.URL,
		Company:     newJob.Company,
		Title:       newJob.Title,
	}
}

func (d *DB) GetAllJobListing() []*model.JobListing {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	collection := d.db.Database("jobsListing").Collection("jobs")
	defer cancel()

	responseDbCursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Sprint(responseDbCursor)

	allJobs := make([]*model.JobListing, 0)

	for responseDbCursor.Next(ctx) {
		var job *model.JobListing
		err := responseDbCursor.Decode(&job)
		if err != nil {
			log.Fatal(err)
		}
		allJobs = append(allJobs, job)
	}

	return allJobs

}

func (d *DB) GetJobById(ID string) *model.JobListing {
	// make the context in calling the database
	ObjectID, err := primitive.ObjectIDFromHex(ID)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	collection := d.db.Database("jobListing").Collection("jobs")
	defer cancel()
	filter := bson.D{{"_id", ObjectID}}
	responseDbCursor := collection.FindOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	var job *model.JobListing
	curErr := responseDbCursor.Decode(&job)
	if curErr != nil {
		log.Fatal(curErr)
	}

	return job

}

func (d *DB) DeleteSingleJob(jobId string) *model.JobListing {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	collection := d.db.Database("jobListing").Collection("jobs")

	defer cancel()
	ObjectID, err := primitive.ObjectIDFromHex(jobId)
	if err != nil {
		log.Fatal(err)
	}

	mongoCursor, mongoErr := collection.DeleteOne(ctx, bson.M{"_id": ObjectID})
	if mongoErr != nil {
		log.Fatal(mongoErr)
	}
	var deletedRecord *model.JobListing

}

func (d *DB) UpdateSingleJob(jobId string) *model.JobListing {

}
