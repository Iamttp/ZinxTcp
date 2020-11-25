package myNet

import (
	"awesomeProject/myZinx/myInterface"
	"sync"
)

type ConnManager struct {
	rwLock sync.RWMutex
	conns  map[uint32]myInterface.IConnect
}

func NewConnManage() *ConnManager {
	return &ConnManager{
		conns: make(map[uint32]myInterface.IConnect),
	}
}

func (cm *ConnManager) Get(index uint32) myInterface.IConnect {
	cm.rwLock.RLock()
	defer cm.rwLock.RUnlock()

	return cm.conns[index]
}

func (cm *ConnManager) Len() int {
	cm.rwLock.RLock() // TODO 是否加读锁
	defer cm.rwLock.RUnlock()

	return len(cm.conns)
}

func (cm *ConnManager) Add(connect myInterface.IConnect) {
	cm.rwLock.Lock()
	defer cm.rwLock.Unlock()

	cm.conns[connect.GetIdConnect()] = connect
}

func (cm *ConnManager) Remove(connect myInterface.IConnect) {
	cm.rwLock.Lock()
	defer cm.rwLock.Unlock()

	delete(cm.conns, connect.GetIdConnect())
}

func (cm *ConnManager) Clear() {
	cm.rwLock.Lock()
	defer cm.rwLock.Unlock()

	for id, conn := range cm.conns {
		conn.Stop()
		delete(cm.conns, id)
	}
}
