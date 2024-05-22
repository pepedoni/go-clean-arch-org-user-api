package middlewares

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

// ResponseRecorder para capturar a resposta
type ResponseRecorder struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r *ResponseRecorder) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func (r *ResponseRecorder) WriteString(s string) (int, error) {
	r.body.WriteString(s)
	return r.ResponseWriter.WriteString(s)
}

func RedisCacheByRequestUri(rdb *redis.Client, ttl time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		cacheKey := "cache:" + c.Request.RequestURI
		val, err := rdb.Get(ctx, cacheKey).Result()
		if err == redis.Nil {
			recorder := &ResponseRecorder{
				ResponseWriter: c.Writer,
				body:           bytes.NewBufferString(""),
			}
			c.Writer = recorder
			c.Next()

			status := c.Writer.Status()
			if status == 200 {
				rdb.Set(ctx, cacheKey, recorder.body.String(), ttl*time.Second).Err()
			}
		} else if err != nil {
			log.Printf("Erro ao acessar o Redis: %v", err)
			c.Next()
		} else {
			c.Header("Content-Type", "application/json")
			c.String(http.StatusOK, val)
			c.Abort()
		}
	}
}

func InvalidateCacheByRequestUri(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		c.Next()
		status := c.Writer.Status()
		if status >= 200 && status < 300 {
			cacheKey := "cache:" + c.Request.RequestURI
			rdb.Del(ctx, cacheKey)
		}

	}
}

func InvalidateCacheLikeRequestUri(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		c.Next()
		status := c.Writer.Status()
		if status >= 200 && status < 300 {
			cacheKey := "cache:" + c.Request.RequestURI + "*"
			fmt.Println("cacheKey", cacheKey)
			keys, _, err := rdb.Scan(ctx, 0, cacheKey, 0).Result()
			fmt.Println("keys", keys)
			if err != nil {
				log.Printf("Erro ao acessar o Redis: %v", err)
				return
			}
			rdb.Del(ctx, keys...)
		}
	}
}
