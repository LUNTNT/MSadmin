package Services

import (
	"admin/DB"
	"admin/Model"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetReplyFolderByName(c *fiber.Ctx) error {
	adminColl := DB.MI.DBCol

	paramName := c.Params("name")

	reply := &Model.StandardReply{}
	filter := bson.D{{Key : "name", Value : paramName}}

	err := adminColl.FindOne(c.Context(), filter).Decode(reply)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Log Not found",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    reply,
	})
}

func GetAllReplyFolder(c *fiber.Ctx) error {
	adminColl := DB.MI.DBCol

	cursor, err := adminColl.Find(c.Context(), bson.D{{}})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Replies Not Found",
			"error":   err.Error(),
		})
	}

	var replies []Model.StandardReply = make([]Model.StandardReply, 0)
	err = cursor.All(c.Context(), &replies)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error to interate cursor into result",
			"error":   err.Error(),
		})
	}

	defer cursor.Close(context.TODO())

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    replies,
	})
}

func CreateReply(c *fiber.Ctx) error {
	adminColl := DB.MI.DBCol

	reply := new(Model.StandardReply)

	err := c.BodyParser(&reply)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	sameNameFilter := bson.D{{Key: "name", Value: reply.Name}}
	count, err := adminColl.CountDocuments(c.Context(), sameNameFilter)
	if count > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Name Has Been Used",
		})
	}

	result, err := adminColl.InsertOne(c.Context(), reply)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot insert reply",
			"error":   err,
		})
	}

	//check the inserted data and return
	checkReply := &Model.StandardReply{}
	checkFilter := bson.D{{Key: "_id", Value: result.InsertedID}}

	adminColl.FindOne(c.Context(), checkFilter).Decode(checkReply)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"StandardReply": checkReply,
		},
	})
}

func UpdateReply(c *fiber.Ctx) error {
	adminColl := DB.MI.DBCol

	reply := new(Model.StandardReply)

	err := c.BodyParser(&reply)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	target := bson.D{{Key: "name", Value: reply.Name}}
	update := bson.D{{"$set", reply}}

	result, err := adminColl.UpdateOne(c.Context(), target, update)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot update reply",
			"error":   err,
		})
	}

	checkReply := &Model.StandardReply{}
	checkFilter := bson.D{{Key: "_id", Value: result.UpsertedID}}

	adminColl.FindOne(c.Context(), checkFilter).Decode(checkReply)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"standardReply": checkReply,
		},
	})
}

func DeleteReply(c *fiber.Ctx) error {
	adminColl := DB.MI.DBCol

	paramName := c.Params("name")
	filter := bson.D{{Key : "name", Value : paramName}}

	result, err := adminColl.DeleteOne(c.Context(), filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Folder Not Found",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"number of deletion": result.DeletedCount,
	})
}