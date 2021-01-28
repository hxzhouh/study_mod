package internal

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"time"
)

type EtcdTools struct {
	Cli     *clientv3.Client
	Timeout time.Duration
}

func (t *EtcdTools) Get(k string) (string, error) {
	timeoutCtx, cancel := context.WithTimeout(context.Background(), t.Timeout)
	defer cancel()
	resp, err := t.Cli.Get(timeoutCtx, k)
	if err != nil {
		return "", err
	}
	kvs := resp.Kvs
	if len(kvs) == 0 {
		return "", nil
	}

	return string(kvs[0].Value), nil
}

func (t *EtcdTools) Put(k, v string) (string, error) {
	timeoutCtx, cancel := context.WithTimeout(context.Background(), t.Timeout)

	defer cancel()
	resp, err := t.Cli.Put(timeoutCtx, k, v)
	if err != nil {
		return "", err
	}
	return string(resp.PrevKv.Key), nil
}
