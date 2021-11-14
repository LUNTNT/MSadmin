package Router

import (
	"admin/Services"

	"github.com/gofiber/fiber/v2"
)

func CustomersRoute(route fiber.Router) {

	//Standard Reply
	route.Get("/getReplyByID/:id", Services.GetReplyFolderByID)
	route.Get("/getAllReply", Services.GetAllReplyFolder)

	route.Post("/createReply", Services.CreateReply)

	route.Put("/updateReply", Services.UpdateReply)

	route.Delete("/deleteReply/:id", Services.DeleteReply)

}