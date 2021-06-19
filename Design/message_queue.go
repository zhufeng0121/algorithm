package Design

import "sync"

type MessageQueue struct {

	msgdata []interface{} // 缓冲区
	len     int32

	readPos int32
	readMutex sync.Mutex

	writePos int32
	writeMutex sync.Mutex

	emptyCond *sync.Cond
	fullCond  *sync.Cond

}

func NewMQ(len int32) *MessageQueue {
	if len < 1 {
		panic("new meg queue fail: len <１")
		return nil
	}
	l := &sync.Mutex{}
	return &MessageQueue{
		msgdata: make([]interface{}, len+1),
		len:     len + 1,
		readPos: 0,
		writePos: 0,

		emptyCond: sync.NewCond(1),
		fullCond:  sync.NewCond(1),

	}

}

func (mq *MessageQueue) Put(in interface{}) {
	mq.writeMutex.Lock()
	defer mq.writeMutex.Unlock()

	mq.fullCond.L.Lock()
	defer mq.fullCond.L.Unlock()
	for (mq.writePos + 1) % mq.len == mq.readPos {
		//缓冲区为满
		mq.fullCond.Wait()
	}
	//写入一个数据
	mq.msgdata[mq.writePos] = in
	mq.writePos = (mq.writePos + 1) % mq.len

	//通知消费者缓冲区存在数据
	mq.emptyCond.Signal()

}
func (mq *MessageQueue) Get(out interface{}) {
	//获取读锁，读取的优先级也是一样的
	mq.readMutex.Lock()
	defer mq.readMutex.Unlock()

	//判断缓冲区是否为空
	mq.emptyCond.L.Lock()
	defer mq.emptyCond.L.Unlock()

	for mq.writePos == mq.readPos {
		//缓冲区为空，等待生产者通知缓冲区有数据写入
		mq.emptyCond.Wait()
	}
	//读取
	out = mq.msgdata[(mq.readPos) % mq.len]
	mq.readPos = (mq.readPos + 1) % mq.len
	mq.fullCond.Signal()
	return

}

func (mq *MessageQueue) Len() int32 {
	if mq.writePos < mq.readPos {
		return mq.writePos + mq.len - mq.readPos
	}
	return mq.writePos - mq.readPos
}
