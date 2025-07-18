package usecases

import (
	"context"
	config2 "etcd_cli_pickme/internal/config"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"strconv"
	"time"
)

var ETCDClient *clientv3.Client

func InitEtcdClient() error {
	endpoint := config2.EtcdCltCfg.Server + ":" + strconv.FormatInt(config2.EtcdCltCfg.Port, 10)

	config := clientv3.Config{
		Endpoints:   []string{endpoint},
		Username:    config2.EtcdCltCfg.UserName,
		Password:    config2.EtcdCltCfg.Password,
		DialTimeout: 5 * time.Second,
	}

	client, err := clientv3.New(config)
	if err != nil {
		fmt.Println("Error creating etcd client:", err)
		return err
	}
	ETCDClient = client
	return nil
}

func GetReadResult(searchTerm string) (*clientv3.GetResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := ETCDClient.Get(ctx, searchTerm, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	return res, nil
}

func ViewResponse(response *clientv3.GetResponse) {
	for _, kv := range response.Kvs {
		key := colorText("36", string(kv.Key))
		val := colorText("32", string(kv.Value))

		println(key)
		println(val)
	}

	println("\n\n")
}

func colorText(code, text string) string {
	return fmt.Sprintf("\033[%sm%s\033[0m", code, text)
}
