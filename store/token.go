package store

import (
  "github.com/dgrijalva/jwt-go"
)

func AuthenticateToken(ID, tokenString string, mobile_device bool) (bool, bool) {
  vertoken, _ := jwt.Parse(tokenString, func (token *jwt.Token) (interface{}, error)  {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
       return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
   }
   return mySigningKey, nil
  })
  if claims, ok := vertoken.Claims.(jwt.MapClaims); ok && vertoken.Valid {
    id := claims["id"].(string)
    valid := CheckTokenExists(id, mobile_device)
    if valid {
      VerifyToken(id, tokenString, mobile_device)
      return false, true
    }
    return false, false
  }
  DeleteToken(ID, mobile_device)
  return true, false
}

func DeleteAuthToken(tokenString string, mobile_device bool) bool{
  vertoken, err := jwt.Parse(tokenString, func (token *jwt.Token) (interface{}, error)  {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
       return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
   }
   return mySigningKey, nil
  })
  if(err != nil){
    return err.Error()
  }
  if claims, ok := vertoken.Claims.(jwt.MapClaims); ok && vertoken.Valid {
    id := claims["id"].(string)
    valid := CheckTokenExists(id, mobile_device)
    if valid {
      VerifyToken(id, tokenString, mobile_device)
      DeleteToken(id, mobile_device)
      return true
    }
  }
  return false
}

func GenerateToken(ID string) login_tokens{
  token := jwt.New(jwt.SigningMethodHS256)
  if expiry {
    token.Claims =  jwt.MapClaims{
      "id"  : ID,
      "exp"   : time.Now().Add(time.Minute * 10).Unix(),
    }
  }else{
    token.Claims =  jwt.MapClaims{
      "id"  : ID,
      "exp"   : time.Now().Add(time.Minute * 100).Unix(),
    }
  }
  tokenString, err := token.SignedString(mySigningKey)
  if(err != nil){
    log.Fatal(err)
  }
  logintoken := login_tokens{
    ID : ID,
    Token: tokenString,
  }
  return logintoken
}
