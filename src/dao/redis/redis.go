package redis

import (
	context2 "context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/ssh"
	"golang.org/x/net/context"
	"log"
	"net"
	"os"
	"srbbs/src/conf"
	"time"
)

var (
	client *redis.Client
	Nil    = redis.Nil
)

type SliceCmd = redis.SliceCmd

//type StringStringMapCmd = redis.StringStringMapCmd

func init() {
	getClient()
	err := client.Ping(context.TODO()).Err()
	if err != nil {
		panic(err)
	}
	return
}

func Close() {
	_ = client.Close()
}

func getClient() {
	cfg := conf.Cfg.RedisConfig
	switch cfg.ConnectType {
	case "ssh":
		key, err := os.ReadFile(cfg.SSHFile)
		if err != nil {
			log.Fatal(err)
		}
		signer, err := ssh.ParsePrivateKey(key)
		if err != nil {
			log.Fatal(err)
		}
		sshConfig := &ssh.ClientConfig{
			User:            "admin",
			Auth:            []ssh.AuthMethod{ssh.PublicKeys(signer)},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Timeout:         15 * time.Second,
		}
		sshCli, err := ssh.Dial("tcp", cfg.SSHRemoteHost, sshConfig)
		if err != nil {
			log.Fatal(err)
		}
		client = redis.NewClient(&redis.Options{
			Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
			DB:           cfg.DB,
			Password:     cfg.Password,
			PoolSize:     cfg.PoolSize,
			MinIdleConns: cfg.MinIdleConns,
			Dialer: func(ctx context2.Context, network, addr string) (net.Conn, error) {
				return sshCli.Dial(network, addr)
			},
			ReadTimeout:  -2, //点进去看disables SetWriteDeadline的值 和版本有关 网上说-1
			WriteTimeout: -2,
		})
	default:
		client = redis.NewClient(&redis.Options{
			Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
			DB:           cfg.DB,
			Password:     cfg.Password,
			PoolSize:     cfg.PoolSize,
			MinIdleConns: cfg.MinIdleConns,
		})
	}
}
