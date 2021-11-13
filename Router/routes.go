package Router

import (
	"admin/Services"

	"github.com/gofiber/fiber/v2"
)

func CustomersRoute(route fiber.Router) {
	//Standard Reply Folder

	route.Get("/name/:name", Services.GetReplyFolderByName)
	route.Get("/all", Services.GetAllReplyFolder)

	route.Post("/create", Services.CreateReply)

	route.Put("/update", Services.UpdateReply)

	route.Delete("/delete/:name", Services.DeleteReply)

}