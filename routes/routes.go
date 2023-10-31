package routes

import (
	"studentMng/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(inRoutes *gin.Engine) {

	inRoutes.GET("/health", controllers.Health())
	inRoutes.POST("student/add", controllers.AddStudent())
	inRoutes.GET("students/view", controllers.ViewStudents())
	inRoutes.POST("student/update", controllers.UpdateStudentMarks())
	inRoutes.DELETE("student/delete", controllers.DeleteStudent())
	inRoutes.GET("student/getRank", controllers.GetRank())
}
