package handlers

import (
  "sugarchat/rethink"
  "sugarchat/psql"
  "sugarchat/store"
  "github.com/gin-gonic/gin"
)

type login_request struct{
  MobileOrEmail string `json:"mobile_or_email" binding:"required"`
  Password string `json:"password" binding:"required"`
}

func LoginHandler(c *gin.Context){
  var request login_request
  if c.Bind(&request) == nil {
    mobile_device = false
    device_type := c.Request.Header.Get("X-Device-Type")
    if device_type == "mobile"{
      mobile_device = true
    }
    password := psql.FetchPassword(request.MobileOrEmail)
    if password != "" && store.VerifyBcrypt(password, request.Password){
      login_token := store.GenerateToken(request.MobileOrEmail)
      rethink.AddLoginToken(
        request.MobileOrEmail,
        login_token,
        mobile_device,
      )
      c.JSON(200, gin.H{
        "success": true,
        "message": "Login Successful",
        "data" : map[string]string{
          "token" : login_token
        }
      })
    } else {
      c.JSON(200, gin.H{
        "success": false,
        "message": "Invalid Crdentials",
      })
    }
  }
}
