/*
 * air_condition
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package router

import (
	"net/http"

	"github.com/DnullP/se_work/controller"
	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	router := gin.Default()
	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

var routes = Routes{
	{
		"Index",
		http.MethodGet,
		"/",
		Index,
	},

	{
		"AdminControlPost",
		http.MethodPost,
		"/admin_control",
		controller.AdminControlPost,
	},

	{
		"GetAllDeviceStatusGet",
		http.MethodGet,
		"/get_all_device_status",
		controller.GetAllDeviceStatusGet,
	},

	{
		"GetDeviceStatusGet",
		http.MethodGet,
		"/get_device_status",
		controller.GetDeviceStatusGet,
	},

	{
		"LoginGet",
		http.MethodGet,
		"/login",
		controller.LoginGet,
	},

	{
		"LogoutGet",
		http.MethodGet,
		"/logout",
		controller.LogoutGet,
	},

	{
		"BillCostGet",
		http.MethodGet,
		"/bill_cost",
		controller.BillCostGet,
	},

	{
		"BillDetailGet",
		http.MethodGet,
		"/bill_detail",
		controller.BillDetailGet,
	},

	{
		"CheckInPost",
		http.MethodPost,
		"/check_in",
		controller.CheckInPost,
	},

	{
		"RemoteControlPost",
		http.MethodPost,
		"/remote_control",
		controller.RemoteControlPost,
	},

	{
		"GetDailyReportPost",
		http.MethodPost,
		"/get_daily_report",
		controller.GetDailyReportPost,
	},

	{
		"GetWeeklyReportPost",
		http.MethodPost,
		"/get_weekly_report",
		controller.GetWeeklyReportPost,
	},
	{
		"GetPanelStatus",
		http.MethodGet,
		"/get_panel_status",
		controller.GetPanelStatusGet,
	},
}
