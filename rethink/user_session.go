package rethink

import (
  r "gopkg.in/dancannon/rethink.v2"
)

type login_tokens struct{
  ID string `gorethink:"id"`
  Token string `gorethink:"token"`
  MobileToken string `gorethink:"mobile_token"`
  WebToken string `gorethink:"web_token"`
}

func CreateUserSession()  {
  store.Info.Println("Creating the User Sessions table")
  _, err := r.Branch(
    r.DB(db_name).TableList().Contains("UserSessions"),
    nil,
    r.DB(db_name).TableCreate("UserSessions"),
  ).Run(store.Session)
  if err != nil{
    store.Error.Println("User Sessions Table creation failed")
  } else {
    store.Success.Println("User Sessions Table creation successful")
  }
}

func GetLoginToken(id string, mobile_device bool) string{
  result, _ := r.DB(db_name).Table("UserSessions").Get(id).Run(store.Session)
  if !result.IsNil() {
    var n login_tokens
    result.One(&n)
    result.Close()
    if (mobile_device){
      return n.MobileToken
    }
    return n.WebToken
  }
}

func CheckTokenExists(id string, mobile_device bool) bool{
  result, _ := r.DB(db_name).Table("UserSessions").Get(id).Run(store.Session)
  if !result.IsNil() {
    var n login_tokens
    result.One(&n)
    result.Close()
    if n.ID == id {
        return true
    }
  }
  return false
}

func AddLoginToken(id, token string, mobile_device bool) bool{
  var mobile_token, web_token string
  if mobile_device {
    mobile_token = token
    web_token = ""
  }else{
    mobile_token = ""
    web_token = token
  }
  if CheckTokenExists(id,mobile_device){
    UpdateToken(id, token, mobile_device)
    return true
  }
  inserr := r.DB(db_name).Table("UserSessions").Insert(login_tokens{
    ID: id,
    MobileToken :  mobile_token,
    WebToken : web_token,
    }).Exec(store.Session)
  return true
}

func UpdateToken(id, token string, mobile_device bool) {
  if mobile_device{
    r.DB(db_name).Table("UserSessions").Get(id).Update(map[string]interface{}{
      "mobile_token" : token,
    }).Exec(store.Session)
  }else{
    r.DB(db_name).Table("UserSessions").Get(id).Update(map[string]interface{}{
      "web_token" : token,
    }).Exec(store.Session)
  }
}
