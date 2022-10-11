package redisx

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

// doCmd server命令
type doCmd struct {
	cmd      []interface{}
	ret      string
	retError error
	callback func(string, error)
}

func (c *doCmd) ExecCmd(cli IClient) error {
	ret, err := cli.Do(c.cmd...)
	c.ret = ret
	c.retError = err
	return err
}

func (c *doCmd) Key() string {
	if len(c.cmd) > 0 {
		return fmt.Sprintf("%v", c.cmd[0])
	}
	return ""
}

func (c *doCmd) Callback() {
	if c.callback != nil {
		c.callback(c.ret, c.retError)
	}
}

type hmsetCmd struct {
	key   string
	value ArrayReq
	err   error
}

func (c *hmsetCmd) ExecCmd(cli IClient) error {
	c.err = cli.HMSet(c.key, c.value)
	return c.err
}

func (c *hmsetCmd) Key() string {
	return c.key
}

type hmgetCmd struct {
	key      string
	field    []string
	ret      []string
	retError error
	callback func([]string, error)
}

func (c *hmgetCmd) ExecCmd(cli IClient) error {
	c.ret, c.retError = cli.HMGet(c.key, c.field...)
	return c.retError
}

func (c *hmgetCmd) Key() string {
	return c.key
}

func (c *hmgetCmd) Callback() {
	if c.callback != nil {
		c.callback(c.ret, c.retError)
	}
}

// delCmd delete命令
type delCmd struct {
	key      string
	err      error
	callback func(error)
}

func (c *delCmd) ExecCmd(cli IClient) error {
	c.err = cli.Del(c.key)
	return c.err
}

func (c *delCmd) Key() string {
	return c.key
}

type getCmd struct {
	key      string
	ret      string
	retErr   error
	callback func(string, error)
}

func (c *getCmd) ExecCmd(cli IClient) error {
	c.ret, c.retErr = cli.Get(c.key)
	return c.retErr
}

func (c *getCmd) Key() string {
	return c.key
}

func (c *getCmd) Callback() {
	if c.callback != nil {
		c.callback(c.ret, c.retErr)
	}
}

type setCmd struct {
	key   string
	value interface{}
	err   error
}

func (c *setCmd) ExecCmd(cli IClient) error {
	c.err = cli.Set(c.key, c.value, 0)
	return c.err
}

func (c *setCmd) Key() string {
	return c.key
}

type zIncrbyCmd struct {
	key   string
	value map[string]int
	err   error
}

func (c *zIncrbyCmd) ExecCmd(cli IClient) error {
	c.err = cli.ZIncrby(c.key, c.value)
	return c.err
}

func (c *zIncrbyCmd) Key() string {
	return c.key
}

type zAddCmd struct {
	key   string
	value map[string]float64
	err   error
}

func (c *zAddCmd) ExecCmd(cli IClient) error {
	arr := NewZAddReq()
	for mem, score := range c.value {
		arr.Add(score, mem)
	}
	_, c.err = cli.ZAdd(c.key, arr)
	return c.err
}

func (c *zAddCmd) Key() string {
	return c.key
}

type zRemCmd struct {
	key   string
	value []string
	err   error
}

func (c *zRemCmd) ExecCmd(cli IClient) error {
	c.err = cli.ZRem(c.key, c.value...)
	return c.err
}

func (c *zRemCmd) Key() string {
	return c.key
}

type zrevrangeCmd struct {
	key      string
	start    int64
	stop     int64
	ret      []redis.Z
	retError error
	callback func([]redis.Z, error)
}

func (c *zrevrangeCmd) ExecCmd(cli IClient) error {
	c.ret, c.retError = cli.ZRevRangeWithScores(c.key, c.start, c.stop)
	return c.retError
}

func (c *zrevrangeCmd) Key() string {
	return c.key
}

func (c *zrevrangeCmd) Callback() {
	if c.callback != nil {
		c.callback(c.ret, c.retError)
	}
}

type zrevrankCmd struct {
	key      string
	mem      string
	ret      int64
	retError error
	callback func(int64, error)
}

func (c *zrevrankCmd) ExecCmd(cli IClient) error {
	c.ret, c.retError = cli.ZRevRank(c.key, c.mem)
	return c.retError
}

func (c *zrevrankCmd) Key() string {
	return c.key
}

func (c *zrevrankCmd) Callback() {
	if c.callback != nil {
		c.callback(c.ret, c.retError)
	}
}

type zscore struct {
	key      string
	mem      string
	ret      float64
	retError error
	callBack func(float64, error)
}

func (c *zscore) ExecCmd(cli IClient) error {
	c.ret, c.retError = cli.ZScore(c.key, c.mem)
	return c.retError
}

func (c *zscore) Key() string {
	return c.key
}

func (c *zscore) Callback() {
	if c.callBack != nil {
		c.callBack(c.ret, c.retError)
	}
}
