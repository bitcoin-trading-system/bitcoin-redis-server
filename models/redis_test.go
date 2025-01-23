package models

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"testing"
	"time"

	"github.com/bitcoin-trading-system/bitcoin-redis-server/config"
	"github.com/bitcoin-trading-system/bitcoin-redis-server/utils"
	"github.com/redis/go-redis/v9"
)

func TestNewRedis(t *testing.T) {
	type args struct {
		cfg   config.Config
		index int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestNewRedis",
			args: args{
				cfg:   modelsTestConfig,
				index: redisTestIndex,
			},
			wantErr: false,
		},
		{
			name: "index is out of range",
			args: args{
				cfg:   modelsTestConfig,
				index: 1000,
			},
			wantErr: true,
		},
		{
			name: "index is minus",
			args: args{
				cfg:   modelsTestConfig,
				index: -1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewRedis(tt.args.cfg, tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRedis() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestConnectRedis(t *testing.T) {
	type args struct {
		cfg config.Config
		db  int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestConnectRedis",
			args: args{
				cfg: modelsTestConfig,
				db:  redisTestIndex,
			},
		},
		{
			name: "index is out of range",
			args: args{
				cfg: modelsTestConfig,
				db:  1000,
			},
			wantErr: true,
		},
		{
			name: "index is minus",
			args: args{
				cfg: modelsTestConfig,
				db:  -1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ConnectRedis(tt.args.cfg, tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConnectRedis() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestRedis_Set(t *testing.T) {
	type fields struct {
		Client *redis.Client
		Config config.Config
	}
	type args struct {
		key   string
		value interface{}
		ttl   time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "TestRedis_Set",
			fields: fields{
				Client: func() *redis.Client {
					client, err := ConnectRedis(modelsTestConfig, redisTestIndex)
					if err != nil {
						panic(err)
					}
					return client
				}(),
				Config: modelsTestConfig,
			},
			args: args{
				key: func() string {
					return fmt.Sprintf("time_%s_%s", "TestRedis_Set", strconv.FormatInt(time.Now().Unix(), 10))
				}(),
				value: "test",
				ttl:   100 * time.Second,
			},
			wantErr: false,
		},
		{
			name: "TestRedis_Set",
			fields: fields{
				Client: func() *redis.Client {
					client, err := ConnectRedis(modelsTestConfig, redisTestIndex)
					if err != nil {
						panic(err)
					}
					return client
				}(),
				Config: modelsTestConfig,
			},
			args: args{
				key: func() string {
					return fmt.Sprintf("time_%s_%s", "TestRedis_Set", strconv.FormatInt(time.Now().Unix(), 10))
				}(),
				value: rand.Int32(),
				ttl:   100 * time.Second,
			},
			wantErr: false,
		},
		{
			name: "TestRedis_Set",
			fields: fields{
				Client: func() *redis.Client {
					client, err := ConnectRedis(modelsTestConfig, redisTestIndex)
					if err != nil {
						panic(err)
					}
					return client
				}(),
				Config: modelsTestConfig,
			},
			args: args{
				key: func() string {
					return fmt.Sprintf("time_%s_%s", "TestRedis_Set", strconv.FormatInt(time.Now().Unix(), 10))
				}(),
				value: rand.Float64(),
				ttl:   1 * time.Second,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				Client: tt.fields.Client,
				Config: tt.fields.Config,
			}
			if err := r.Set(tt.args.key, tt.args.value, tt.args.ttl); (err != nil) != tt.wantErr {
				t.Errorf("Redis.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedis_Get(t *testing.T) {
	redisClient := func() *redis.Client {
		client, err := ConnectRedis(modelsTestConfig, redisTestIndex)
		if err != nil {
			panic(err)
		}
		return client
	}()

	type fields struct {
		Client *redis.Client
		Config config.Config
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "TestRedis_Get_String",
			fields: fields{
				Client: redisClient,
				Config: modelsTestConfig,
			},
			args: args{
				"test_Redis_Get_key_string",
			},
			want: func() string {
				val := "test"
				if err := redisClient.Set(ctx, "test_Redis_Get_key_string", val, 100*time.Second).Err(); err != nil {
					panic(err)
				}
				return val
			}(),
			wantErr: false,
		},
		{
			name: "TestRedis_Get_Int",
			fields: fields{
				Client: redisClient,
				Config: modelsTestConfig,
			},
			args: args{
				"test_Redis_Get_key_int",
			},
			want: func() string {
				val := int(rand.Int64())
				if err := redisClient.Set(ctx, "test_Redis_Get_key_int", val, 100*time.Second).Err(); err != nil {
					panic(err)
				}
				return strconv.Itoa(val)
			}(),
			wantErr: false,
		},
		{
			name: "TestRedis_Get_Float64",
			fields: fields{
				Client: redisClient,
				Config: modelsTestConfig,
			},
			args: args{
				"test_Redis_Get_key_float64",
			},
			want: func() string {
				val := 0.528407
				if err := redisClient.Set(ctx, "test_Redis_Get_key_float64", val, 100*time.Second).Err(); err != nil {
					panic(err)
				}
				return fmt.Sprintf("%f", val)
			}(),
			wantErr: false,
		},
		{
			name: "TestRedis_Get_Json_not_found",
			fields: fields{
				Client: redisClient,
				Config: modelsTestConfig,
			},
			args: args{
				"test_Redis_Get_key_json_not_found",
			},
			want: func() string {
				if err := redisClient.Del(ctx, "test_Redis_Get_key_json_not_found").Err(); err != nil {
					panic(err)
				}
				return ""
			}(),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			time.Sleep(1 * time.Second)
			r := &Redis{
				Client: tt.fields.Client,
				Config: tt.fields.Config,
			}
			got, err := r.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Redis.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Redis.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Del(t *testing.T) {
	redisClient := func() *redis.Client {
		client, err := ConnectRedis(modelsTestConfig, redisTestIndex)
		if err != nil {
			panic(err)
		}
		return client
	}()

	type fields struct {
		Client *redis.Client
		Config config.Config
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "TestRedis_Del_not_found",
			fields: fields{
				Client: redisClient,
				Config: modelsTestConfig,
			},
			args: args{
				utils.RandomString(100), // ほぼ存在しないkey
			},
			wantErr: false,
		},
		{
			name: "TestRedis_Del_found",
			fields: fields{
				Client: redisClient,
				Config: modelsTestConfig,
			},
			args: args{
				func() string {
					key := utils.RandomString(100)
					if err := redisClient.Set(ctx, key, "test", 100*time.Second).Err(); err != nil {
						panic(err)
					}
					return key
				}(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				Client: tt.fields.Client,
				Config: tt.fields.Config,
			}
			if err := r.Del(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Redis.Del() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
