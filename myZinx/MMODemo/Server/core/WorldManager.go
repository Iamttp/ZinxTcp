package core

import "sync"

type WorldManager struct {
	playerMap map[int32]*Player
	pLock     sync.Mutex
}

func NewWorldManager() *WorldManager {
	return &WorldManager{
		playerMap: make(map[int32]*Player),
	}
}

func (wm *WorldManager) Add(p *Player) {
	wm.pLock.Lock()
	wm.playerMap[p.Pid] = p
	wm.pLock.Unlock()
}

func (wm *WorldManager) Remove(pid int32) {
	wm.pLock.Lock()
	delete(wm.playerMap, pid)
	wm.pLock.Unlock()
}

func (wm *WorldManager) GetPlayerById(pid int32) *Player {
	wm.pLock.Lock()
	defer wm.pLock.Unlock()

	// TODO 判断pid存在
	return wm.playerMap[pid]
}

func (wm *WorldManager) GetAllPlayers() []*Player {
	wm.pLock.Lock()
	defer wm.pLock.Unlock()

	players := make([]*Player, 0, len(wm.playerMap))
	for _, v := range wm.playerMap {
		players = append(players, v)
	}
	return players
}
