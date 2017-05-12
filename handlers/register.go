package handlers

import (
  "sugarchat/store"
  "sugarchat/psql"
  "github.com/gin-gonic/gin"
)

type register_request struct{
  Firstname string `json:"firstname" binding:"required"`
  Lastname string `json:"lastname" binding:"required"`
  MobileNumber string `json:"mobile_number" binding:"required"`
  EmailID string `json:"email_id" binding:"required"`
  Password string `json:"password" binding:"required"`
  GenderCode int `json:"gender_code" binding:"required"`
  DOB int `json:"dob" binding:"required"`
}

func SendResponse(c *gin.Context, httpstatus int, success bool, message string){
  c.JSON(httpstatus, gin.H{
    "success": success,
    "message": message,
  })
}

func RegisterHandler(c *gin.Context){
  var request register_request
  if c.Bind(&request) == nil {
    if store.ValidateMobile(request.MobileNumber) {
      if store.ValidateEmail(request.EmailID) {
        if !psql.CheckUserConflict(request.MobileNumber) {
          if !psql.CheckUserConflict(request.EmailID){
            if psql.RegisterUser(
              request.MobileNumber,
              request.EmailID,
              request.Firstname,
              request.Lastname,
              store.BcryptPassword(request.Password),
              request.GenderCode,
              request.DOB,
            ){
              SendResponse(c, 200, true, "Registered Successfully")
            } else {
              SendResponse(c, 200, false, "Failed to Register. Please try again later.")
            }
          } else {
            SendResponse(c, 200, false, "Email ID already registered")
          }
        } else {
          SendResponse(c, 200, false, "Mobile Number already registered")
        }
      } else {
        SendResponse(c, 200, false, "Invalid Email ID")
      }
    } else {
      SendResponse(c, 200, false, "Invalid Mobile Number")
    }
  }
}
