package handlers

import (
  "github.com/gin-gonic/gin"
)

type profile_request struct{
  MobileOrEmail string `form:"mobile_or_email" json:"mobile_or_email" binding:"required"`
}

type profile_response struct{
  Firstname string `json:"firstname"`
  Lastname string `json:"lastname"`
  MobileNumber string `json:"mobile_number"`
  EmailID string `json:"email_id"`
  GenderCode int `json:"gender_code"`
  DOB int `json:"dob"`
}

func ProfileHandler(c *gin.Context){
  var request profile_request
  if c.Bind(&request) == nil {
    c.JSON(200, gin.H{
      "success": true,
      "data": "",
    })
  }
}
