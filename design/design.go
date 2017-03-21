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
	DefaultMedia(RecommendationsMediaType)

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
	DefaultMedia(UserMediaType)

	Action("list", func() {
		Routing(GET(""))
		Description("List up to N user resources. N can be adjusted using the 'limit' and 'offset' parameters.")
		Params(func() {
			Param("limit", Integer, func() {
				Default(20)
			})
			Param("offset", Integer, func() {
				Default(0)
			})
		})

		Response(OK, CollectionOf(UserMediaType))
	})

	Action("show", func() {
		Routing(GET("/:id"))
		Description("Get a user resource with the given ID.")
		Params(func() {
			Param("id")
		})

		Response(OK, func() {
			Media(UserMediaType, "full")
		})
		Response(NotFound, ErrorMedia)
	})

	Action("create", func() {
		Routing(POST(""))
		Payload(User)

		Response(Created, func() {
			Media(UserMediaType, "link")
		})
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
	DefaultMedia(MusicMediaType)

	Action("list", func() {
		Routing(GET(""))
		Description("List up to N music resources. N can be adjusted using the 'limit' and 'offset' parameters.")
		Params(func() {
			Param("limit", Integer, func() {
				Default(20)
			})
			Param("offset", Integer, func() {
				Default(0)
			})
		})

		Response(OK, CollectionOf(MusicMediaType))
	})

	Action("show", func() {
		Routing(GET("/:id"))
		Description("Get a music resource with the given ID")
		Params(func() {
			Param("id")
		})

		Response(OK, func() {
			Media(MusicMediaType, "full")
		})
		Response(NotFound, ErrorMedia)
	})

	Action("create", func() {
		Routing(POST(""))
		Payload(Music)

		Response(Created, func() {
			Media(MusicMediaType, "link")
		})
	})
})

var RecommendationsMediaType = MediaType("application/vnd.bluelens.recommendations+json", func() {
	Description("A list of recommendations for the specified user")
	ContentType("application/json")

	Attributes(func() {
		Attribute("musicID", ArrayOf(String))
		Attribute("list", CollectionOf(MusicMediaType))
		Attribute("user", UserMediaType)

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

var User = Type("user", func() {
	Description("A user resource")

	Attribute("id", String)
	Attribute("followees", ArrayOf("user")) // avoid initialization loop
	Attribute("history", ArrayOf(Music))

	Required("id")
})

var UserMediaType = MediaType("application/vnd.bluelens.user+json", func() {
	Description("Media type of a user resource")
	Reference(User)
	ContentType("application/json")

	Attributes(func() {
		Attribute("id")
		Attribute("followees", CollectionOf("application/vnd.bluelens.user+json")) // avoid initialization loop
		Attribute("history", CollectionOf(MusicMediaType))
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

	View("full", func() {
		Attribute("id")
		Attribute("followees")
		Attribute("history")
		Attribute("href")
	})
})

var Music = Type("music", func() {
	Description("A music resource")

	Attribute("id", String)
	Attribute("tags", ArrayOf(String))

	Required("id")
})

var MusicMediaType = MediaType("application/vnd.bluelens.music+json", func() {
	Description("Media type of a music resource")
	Reference(Music)
	ContentType("application/json")

	Attributes(func() {
		Attribute("id")
		Attribute("tags")
		Attribute("href")

		Required("id", "href")
	})

	View("default", func() {
		Attribute("id")
		Attribute("tags")
		Attribute("href")
	})

	View("full", func() {
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
