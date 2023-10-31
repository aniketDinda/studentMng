package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sort"
	"studentMng/database"
	"studentMng/helpers"
	"studentMng/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var studentCollection *mongo.Collection = database.OpenCollection(database.Client, "students")
var validate = validator.New()

func AddStudent() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var student models.Student

		if err := ctx.BindJSON(&student); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(student)

		if validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		count, err := studentCollection.CountDocuments(c, bson.M{"name": student.Name})
		defer cancel()

		if err != nil {
			log.Panic(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking student"})
		}

		if count > 0 {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "this student already exists"})
			return
		}

		student.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		student.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		student.ID = primitive.NewObjectID()
		student.UserId = student.ID.Hex()
		student.Passed = helpers.IsPassed(student.Marks)

		_, err = studentCollection.InsertOne(ctx, student)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Student not Created!!"})
			return
		}
		defer cancel()

		ctx.JSON(http.StatusOK, student)
	}

}

func ViewStudents() gin.HandlerFunc {
	return func(c *gin.Context) {
		var students []models.Student
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		cursor, err := studentCollection.Find(ctx, bson.D{{}})
		if err != nil {
			c.JSON(http.StatusInternalServerError, "DB Error")
			return
		}
		err = cursor.All(ctx, &students)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		defer cursor.Close(ctx)
		if err := cursor.Err(); err != nil {
			c.JSON(http.StatusBadRequest, "invalid request")
			return
		}
		defer cancel()
		c.IndentedJSON(http.StatusOK, students)
	}

}

func UpdateStudentMarks() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var updatedStudent models.UpdateMarksInput

		if err := ctx.BindJSON(&updatedStudent); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		filter := bson.M{"name": updatedStudent.Name}
		update := bson.M{"$set": bson.M{"marks": updatedStudent.Marks, "passed": helpers.IsPassed(updatedStudent.Marks)}}
		result, err := studentCollection.UpdateOne(c, filter, update)
		if err != nil {
			log.Fatal(err)
		}

		ctx.JSON(http.StatusOK, result.ModifiedCount)
	}
}

func DeleteStudent() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var name models.UserInput
		if err := ctx.BindJSON(&name); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		filter := bson.M{"name": name.Name}

		result, err := studentCollection.DeleteOne(context.Background(), filter)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Deleted Count: ", result.DeletedCount)
	}
}

func GetRank() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		c, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		cursor, err := studentCollection.Find(c, bson.D{{}})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, "DB Error")
			return
		}

		var marksArray []int32

		for cursor.Next(c) {
			var student models.Student
			err := cursor.Decode(&student)
			if err != nil {
				log.Println(err)
				ctx.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			marksArray = append(marksArray, student.Marks)
		}

		sort.Slice(marksArray, func(i, j int) bool {
			return marksArray[i] > marksArray[j]
		})
		var name models.UserInput
		if err := ctx.BindJSON(&name); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var student models.Student
		err = studentCollection.FindOne(c, bson.M{"name": name.Name}).Decode(&student)
		if err != nil {
			ctx.JSON(http.StatusNotFound, "Student not found")
			return
		}

		rank := helpers.GetRank(marksArray, student.Marks)

		ctx.JSON(http.StatusOK, gin.H{"rank": rank})
	}
}

func Health() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		health := database.DbHealth()

		if !health {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "DB Connection Failed"})
		} else {
			ctx.IndentedJSON(http.StatusOK, gin.H{"msg": "DB Connection Successful"})
		}
	}
}
