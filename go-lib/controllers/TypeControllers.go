package controllers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/baguette/go-lib/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

func GetTypes() gin.HandlerFunc{
	return func (c *gin.Context)  {
		ctx,cancel := context.WithTimeout(context.Background(),30*time.Second)
		defer cancel()
		res := []bson.M{}
		opt := options.Find().SetProjection(bson.D{{Key:"_id",Value:0}})
		cursor,_  := TypeCollections.Find(ctx,bson.D{{}},opt)
		cursor.All(ctx,&res)
		c.JSON(http.StatusOK,res)
	}
}
func CreateType() gin.HandlerFunc{
	return func(c *gin.Context) {
		ctx,cancel := context.WithTimeout(context.Background(),30*time.Second)
		defer cancel()
		rec := bson.M{}
		if err := c.ShouldBindJSON(&rec);err != nil {
			c.JSON(http.StatusBadRequest,gin.H{"error":err})
			return
		}
		newType := models.BookTypeModel{}
		newType.Id = primitive.NewObjectID()
		count,_ := TypeCollections.CountDocuments(ctx,bson.D{{}})
		newType.TypeId = strconv.Itoa(int(count))
		newType.TypeName = rec["type_name"].(string)
		_,insertErr := TypeCollections.InsertOne(ctx,newType)
		if insertErr != nil {
			c.JSON(http.StatusBadRequest,gin.H{"error":insertErr})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"type_id":newType.TypeId,
			"type_name":newType.TypeName,
		})
	}
}

func UpdateType() gin.HandlerFunc{
	return func(c *gin.Context) {
		ctx,cancel := context.WithTimeout(context.Background(),30*time.Second)
		defer cancel()
		type_id := c.Param("type_id")
		rec := bson.M{}
		if err := c.ShouldBindJSON(&rec);err != nil {
			c.JSON(http.StatusBadRequest,gin.H{"error":err})
			return
		}
		updateObj := bson.D{}
		if rec["type_name"].(string) != ""{
			updateObj = append(updateObj, bson.E{Key:"typename",Value:""})
		}
		TypeCollections.UpdateOne(ctx,bson.M{"typeid":type_id},bson.D{
			{Key:"$set",Value:updateObj},
		})
		c.JSON(http.StatusOK,gin.H{
			"type_id":type_id,
			"type_name":rec["type_name"].(string),
		})
	}
}

func DeleteType() gin.HandlerFunc{
	return func (c *gin.Context)  {
		ctx,cancel := context.WithTimeout(context.Background(),30*time.Second)
		defer cancel()
		type_id := c.Param("type_id")
		TypeCollections.DeleteOne(ctx,bson.M{"typeid":type_id})
		c.JSON(http.StatusOK,gin.H{
			"delete":"success",
			"type_id":type_id,
		})
	}
}