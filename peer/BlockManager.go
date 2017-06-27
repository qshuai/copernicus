package peer

import (
	
	"github.com/btccom/copernicus/model"
	"sync"
	"container/list"
	"sync/atomic"
	"github.com/btccom/copernicus/utils"
)

type BlockManager struct {
	server          *PeerManager //todo mutual reference
	started         int32
	shutdown        int32
	Chain           *model.Blockchain
	rejectedTxns    map[utils.Hash]struct{}
	requestedTxns   map[utils.Hash]struct{}
	requestedBlocks map[utils.Hash]struct{}
	progressLogger  *BlockProgressLogger //todo do't need?
	syncPeer        *ServerPeer     //todo mutual reference
	
	messageChan      chan interface{}
	waitGroup        sync.WaitGroup
	quit             chan struct{}
	headersFirstMode bool
	headerList       *list.List
	startHeader      *list.Element
	netCheckPoint    *model.Checkpoint
}

func (blockManager *BlockManager) NewPeer(serverPeer *ServerPeer) {
	if atomic.LoadInt32(&blockManager.shutdown) != 0 {
		return
	}
	blockManager.messageChan <- &NewPeerMessage{serverPeer: serverPeer}
	
}
func (blockManager *BlockManager) Start() {
	//todo
}
func (blockmanager *BlockManager)Stop(){

}