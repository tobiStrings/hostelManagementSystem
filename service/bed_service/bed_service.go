package bed_service

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"hostelManagementSystem/models"
	"hostelManagementSystem/repositories"
	"time"
)

var bedCollection *mongo.Collection = repositories.OpenCollection(repositories.Client,"bed")

type BedService interface {
	CreateBed(bedNumber uint64)(models.Bed,error)
	DeleteBed(bedNumber uint64)error
	FindBedById(bedId string) (models.Bed,error)
	FindBedByBedNumber(bedNumber uint64)(models.Bed,error)
}

type BedServiceImpl struct {

}

func (bedService *BedServiceImpl)  CreateBed(bedNumber  uint64)(models.Bed,error){
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	bed := models.Bed{}

	defer cancel()

	err := bedCollection.FindOne(ctx,bson.M{"bednumber":bedNumber}).Decode(&bed)

	if err != nil{
		return models.Bed{},errors.New(err.Error())
	}

	if &bed != nil{
		return models.Bed{},errors.New("Bed with that number has already been created ")
	}

	bed.BedNumber = bedNumber

	_,err = bedCollection.InsertOne(ctx,bed)

	if err != nil{
		return models.Bed{},errors.New(err.Error())
	}

	return bed,nil

}
func (bedService *BedServiceImpl)  DeleteBed(bedNumber uint64)error{
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	_, err := bedCollection.DeleteOne(ctx, bson.M{"bednumber": bedNumber})
	if err != nil {
		return err
	}
	return  nil
}

func (bedService *BedServiceImpl)  FindBedById(bedId string)(models.Bed,error){
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	bed := models.Bed{}

	defer cancel()

	err:=bedCollection.FindOne(ctx,bson.M{"id":bedId}).Decode(&bed)

	if err != nil{
		return models.Bed{}, err
	}

	return bed,nil
}

func(bedService *BedServiceImpl) FindBedByBedNumber(bedNumber uint64)(models.Bed,error){
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	bed := models.Bed{}

	defer cancel()

	err:=bedCollection.FindOne(ctx,bson.M{"bednumber":bedNumber}).Decode(&bed)

	if err != nil{
		return models.Bed{}, err
	}

	return bed,nil
}
