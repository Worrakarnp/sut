package main

import (
	"github.com/Worrakarnp/suta/controller"
	"github.com/Worrakarnp/suta/entity"
	"github.com/Worrakarnp/suta/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())
	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{
			//Officer Routes
			protected.GET("/officers", controller.ListOfficers)
			protected.GET("/officer/:id", controller.GetOfficer)
			protected.PATCH("/officers", controller.UpdateOfficer)
			protected.DELETE("/officers/:id", controller.DeleteOfficer)

			// Category Routes
			protected.GET("/categories", controller.ListCategories)
			protected.GET("/category/:id", controller.GetCategory)
			protected.POST("/categories", controller.CreateCategory)
			protected.PATCH("/categories", controller.UpdateCategory)
			protected.DELETE("/categories/:id", controller.DeleteCategory)

			// Branch Routes
			protected.GET("/branches", controller.ListBranch)
			protected.GET("/branch/:id", controller.GetBranch)
			protected.POST("/branches", controller.CreateBranch)
			protected.PATCH("/branches", controller.UpdateBranch)
			protected.DELETE("/branches/:id", controller.DeleteBranch)

			// SubDistrict Routes
			protected.GET("/sub_districts", controller.ListSubDistrict)
			protected.GET("/sub_district/:id", controller.GetSubDistrict)
			protected.POST("/sub_districts", controller.CreateSubDistrict)
			protected.PATCH("/sub_districts", controller.UpdateSubDistrict)
			protected.DELETE("/sub_districts/:id", controller.DeleteSubDistrict)

			// District Routes
			protected.GET("/districts", controller.ListDistrict)
			protected.GET("/district/:id", controller.GetDistrict)
			protected.POST("/districts", controller.CreateDistrict)
			protected.PATCH("/districts", controller.UpdateDistrict)
			protected.DELETE("/districts/:id", controller.DeleteDistrict)

			// Province Routes
			protected.GET("/provinces", controller.ListProvince)
			protected.GET("/province/:id", controller.GetProvince)
			protected.POST("/provinces", controller.CreateProvince)
			protected.PATCH("/provinces", controller.UpdateProvince)
			protected.DELETE("/provinces/:id", controller.DeleteProvince)

			// Size Routes
			protected.GET("/sizes", controller.ListSize)
			protected.GET("/size/:id", controller.GetSize)
			protected.POST("/sizes", controller.CreateSize)
			protected.PATCH("/sizes", controller.UpdateSize)
			protected.DELETE("/sizes/:id", controller.DeleteSize)

			// Blazer Routes
			protected.GET("/blazers", controller.ListBlazer)
			protected.GET("/blazer/:id", controller.GetBlazer)
			protected.POST("/blazers", controller.CreateBlazer)
			protected.PATCH("/blazers", controller.UpdateBlazer)
			protected.DELETE("/blazers/:id", controller.DeleteBlazer)

			// Degree Routes
			protected.GET("/degrees", controller.ListDegree)
			protected.GET("/degree/:id", controller.GetDegree)
			protected.POST("/degrees", controller.CreateDegree)
			protected.PATCH("/degrees", controller.UpdateDegree)
			protected.DELETE("/degrees/:id", controller.DeleteDegree)

			// Syndicate Routes
			protected.GET("/syndicate", controller.ListSyndicate)
			protected.GET("/syndicate/:id", controller.GetSyndicate)
			protected.POST("/syndicate", controller.CreateSyndicate)
			protected.PATCH("/syndicate", controller.UpdateSyndicate)
			protected.DELETE("/syndicate/:id", controller.DeleteSyndicate)

			// Member Routes
			protected.GET("/member", controller.ListMember)
			protected.GET("/member/:id", controller.GetMember)
			protected.POST("/member", controller.CreateMember)
			protected.PATCH("/member", controller.UpdateMember)
			protected.DELETE("/member/:id", controller.DeleteMember)

			// Trader Routes
			protected.GET("/trader", controller.ListTrader)
			protected.GET("/trader/:id", controller.GetTrader)
			protected.POST("/trader", controller.CreateTrader)
			protected.PATCH("/trader", controller.UpdateTrader)
			protected.DELETE("/trader/:id", controller.DeleteTrader)

			// Gown Routes
			protected.GET("/gown", controller.ListGown)
			protected.GET("/gown/:id", controller.GetGown)
			protected.POST("/gown", controller.CreateGown)
			protected.PATCH("/gown", controller.UpdateGown)
			protected.DELETE("/gown/:id", controller.DeleteGown)

		}
	}
	// User Routes
	r.POST("/officers", controller.CreateOfficer)

	// Authentication Routes
	r.POST("/login", controller.Login)

	// Run the server
	r.Run()

}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
