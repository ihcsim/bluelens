package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var apiHost = "localhost:8080"

var _ = API("bluelens", func() {
	Title("The bluelens API")
	Description("This API provides a set of endpoints to manage users' followees, music history and recommendations.")
	Host(apiHost)
	Scheme("http")
	BasePath("/bluelens")

	License(func() {
		Name("MIT")
		URL("https://github.com/ihcsim/bluelens/blob/master/LICENSE")
	})

	Docs(func() {
		Description("Swagger docs")
		URL("http://localhost:8080/swagger.json")
	})

	Consumes("application/json")
	Produces("application/json")
})

var _ = Resource("recommendations", func() {
	BasePath("/recommendations")
	DefaultMedia(Recommendations)

	Action("recommend", func() {
		Routing(GET("/:userID/:maxCount"))
		Description("Make music recommendations for a user.")
		Params(func() {
			Param("userID", String, "ID of the user these recommendations are meant for.")
			Param("maxCount", Integer, "Maximum number of recommendations to be returned to the user. Set to zero to use server's default.", func() {
				Minimum(0)
			})
		})
		Response(OK)
		Response(NotFound, ErrorMedia)
	})
})

var _ = Resource("user", func() {
	BasePath("/user")
	DefaultMedia(User)
	Params(func() {
		Param("id")
	})

	Action("get", func() {
		Routing(GET("/:id"))
		Description("Get a user resource with the given ID")
		Response(OK)
		Response(NotFound, ErrorMedia)
	})

	Action("listen", func() {
		Routing(POST("/:id/listen/:musicID"))
		Description("Add a music to a user's history.")
		Params(func() {
			Param("musicID", String, "ID of the music.")
		})
		Response(OK)
		Response(NotFound, ErrorMedia)
		Response(BadRequest, ErrorMedia)
	})

	Action("follow", func() {
		Routing(POST("/:id/follows/:followeeID"))
		Description("Update a user's followees list with a new followee.")
		Params(func() {
			Param("followeeID", String, "ID of the followee.")
		})
		Response(OK)
		Response(NotFound, ErrorMedia)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("music", func() {
	BasePath("/music")
	DefaultMedia(Music)
	Params(func() {
		Param("id")
	})

	Action("get", func() {
		Routing(GET("/:id"))
		Description("Get a music resource with the given ID")
		Response(OK)
		Response(NotFound, ErrorMedia)
	})
})

var Recommendations = MediaType("application/vnd.bluelens.recommendations+json", func() {
	Description("A list of recommendations for the specified user")
	ContentType("application/json")

	Attributes(func() {
		Attribute("musicID", ArrayOf(String))
		Attribute("list", CollectionOf("application/vnd.bluelens.music+json"))
		Attribute("user", User)

		Links(func() {
			Link("list")
			Link("user")
		})

		Required("list", "user")
	})

	View("default", func() {
		Attribute("musicID")
		Attribute("links")
	})

	View("all", func() {
		Attribute("list")
		Attribute("user")
		Attribute("links")
	})
})

var User = MediaType("application/vnd.bluelens.user+json", func() {
	Description("A user resource")
	ContentType("application/json")

	Attributes(func() {
		Attribute("id", String)
		Attribute("followees", CollectionOf("application/vnd.bluelens.user+json"))
		Attribute("history", CollectionOf("application/vnd.bluelens.music+json"))
		Attribute("href", String)

		Links(func() {
			Link("followees")
			Link("history")
		})

		Required("id", "href")
	})

	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("links")
	})

	View("link", func() {
		Attribute("href")
	})

	View("all", func() {
		Attribute("id")
		Attribute("followees")
		Attribute("history")
		Attribute("href")
	})
})

var Music = MediaType("application/vnd.bluelens.music+json", func() {
	Description("A music resource")
	ContentType("application/json")

	Attributes(func() {
		Attribute("id", String)
		Attribute("tags", ArrayOf(String))
		Attribute("href", String)

		Required("id", "href")
	})

	View("default", func() {
		Attribute("id")
		Attribute("tags")
		Attribute("href")
	})

	View("link", func() {
		Attribute("href")
	})
})

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/swagger.json", "server/swagger/swagger.json")
})
