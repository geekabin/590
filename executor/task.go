package main

import (
	"github.com/zeromicro/go-zero/core/executors"
	"time"
)

type NodeMsgBatchInsertIntoDb struct {
}

func (receiver NodeMsgBatchInsertIntoDb) init() {
	executors.NewBulkExecutor(
		dts.insertIntoCk,
		executors.WithBulkInterval(time.Second*3), // 3s会自动刷一次container中task去执行
		executors.WithBulkTasks(10240),            // container最大task数。一般设为2的幂次
	)
}
