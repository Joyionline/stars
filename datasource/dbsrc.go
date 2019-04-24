package datasource

import (
	"fmt"
	"stars/conf"
	"sync"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-xorm/xorm"
	"github.com/gpmgo/gopm/modules/log"
)

var (
	masterEngine *xorm.Engine
	slaveEngine  *xorm.Engine
	lock         sync.Mutex
)

func InstanceMaster() *xorm.Engine {
	if masterEngine != nil {
		return masterEngine
	}
	lock.Lock()
	defer lock.Unlock()
	//单例模式中防止同时多个请求被锁住等待，这样之后的请求可以直接返回
	if masterEngine != nil {
		return masterEngine
	}
	c := conf.MasterDbConfig
	driveSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		c.Username, c.Password, c.Host, c.Port, c.DbName)
	engine, err := xorm.NewEngine(conf.DriverName, driveSource)
	if err != nil {
		log.Fatal("Dbsrc.DbInstanceMaster error =", err)
		return nil
	} else {
		masterEngine = engine
		return masterEngine
	}
	return masterEngine
}

func InstanceSlave() *xorm.Engine {
	if slaveEngine != nil {
		return slaveEngine
	}
	lock.Lock()
	defer lock.Unlock()
	//单例模式中防止同时多个请求被锁住等待，这样之后的请求可以直接返回
	if slaveEngine != nil {
		return slaveEngine
	}
	c := conf.MasterDbConfig
	driveSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		c.Username, c.Password, c.Host, c.Port, c.DbName)
	engine, err := xorm.NewEngine(conf.DriverName, driveSource)
	if err != nil {
		log.Fatal("Dbsrc.DbInstanceSlave error =", err)
		return nil
	} else {
		slaveEngine = engine
		return slaveEngine
	}
	return slaveEngine
}
