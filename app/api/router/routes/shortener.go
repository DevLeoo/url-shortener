package routes

import (
	"net/http"
	"url-shortener/app/api/controllers"
)

var shortnerRoutes = Routes{
	Uri:                "/shorten",
	Method:             http.MethodPost,
	Fn:                 controllers.Shorten,
	NeedAuthentication: false,
}
