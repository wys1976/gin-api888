package middleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/wys1976/gin-api888/utils/crypto"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func CryptoMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 解密请求
		if c.Request.Body != nil {
			bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

			if encrypted := gjson.GetBytes(bodyBytes, "encrypted").String(); encrypted != "" {
				decrypted := crypto.Decrypt(encrypted)
				var data map[string]interface{}
				json.Unmarshal([]byte(decrypted), &data)
				c.Set("decryptedData", data)
			}
		}

		// 加密响应
		c.Writer = &encryptWriter{c.Writer}
		c.Next()
	}
}

type encryptWriter struct {
	gin.ResponseWriter
}

func (w *encryptWriter) Write(data []byte) (int, error) {
	encrypted := crypto.Encrypt(string(data))
	return w.ResponseWriter.Write([]byte(`{"encrypted":"` + encrypted + `"}`))
}
