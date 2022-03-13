package redispkg

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-redis/redis/v8"
)

func TestCustomCmd(t *testing.T) {
	type args struct {
		ctx context.Context
		rdb *redis.Client
		key string
	}
	var tests []struct {
		name string
		args args
		want *redis.StringCmd
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CustomCmd(tt.args.ctx, tt.args.rdb, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExampleClient(t *testing.T) {
	type args struct {
		ctx context.Context
		rdb *redis.Client
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ExampleClient(tt.args.ctx, tt.args.rdb)
		})
	}
}

func TestTrigger(t *testing.T) {
	var tests []struct {
		name string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Trigger()
		})
	}
}
