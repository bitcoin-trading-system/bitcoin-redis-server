package main

import (
	"fmt"
	"os/exec"
	"testing"
	"time"

	"github.com/bitcoin-trading-system/bitcoin-redis-server/config"
	"github.com/bitcoin-trading-system/bitcoin-redis-server/models"
	"github.com/bitcoin-trading-system/bitcoin-redis-server/utils"
)

func TestMainFunction(t *testing.T) {
	go main()

	config, err := config.NewConfig("toml/local.toml", "env/env.local.toml")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}
	redisRepo, err := models.NewRedisRepository(config)
	if err != nil {
		panic(err)
	}
	redisClient := redisRepo.GetRedis(0)

	type args struct {
		curl []string
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "get_redis_handler",
			args: args{
				func() []string {
					key := utils.RandomString(50)
					value := utils.RandomString(50)
					redisClient.Set(key, value, 1*time.Minute)
					return []string{"curl", "-w", "%{http_code}", "-o", "/dev/null", "-s", "-X", "GET", fmt.Sprintf("http://localhost:8003/index/0/key/%s", key)}
				}(),
			},
		},
		{
			name: "set_redis_handler",
			args: args{
				curl: []string{"curl", "-w", "%{http_code}", "-o", "/dev/null", "-s", "-X", "POST", "-H", "Content-Type: application/json", "http://localhost:8003/index/0/key/key1", "-d", `{"value":"value1", "ttl": 30}`},
			},
		},
		{
			name: "del_redis_handler",
			args: args{
				curl: func() []string {
					key := utils.RandomString(50)
					value := utils.RandomString(50)
					redisClient.Set(key, value, 1*time.Minute)
					return []string{"curl", "-w", "%{http_code}", "-o", "/dev/null", "-s", "-X", "DELETE", fmt.Sprintf("http://localhost:8003/index/0/key/%s", key)}
				}(),
			},
		},
	}

	time.Sleep(1 * time.Second)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// curlコマンドを実行
			cmd := exec.Command(tt.args.curl[0], tt.args.curl[1:]...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("Failed to execute curl command: %v", err)
			}

			// ステータスコードを確認
			statusCode := string(output[len(output)-3:]) // 最後の3文字がステータスコード
			if statusCode != "200" && statusCode != "201" {
				t.Fatalf("Server returned non-200/201 status code: %s", statusCode)
			}
		})
	}
}
