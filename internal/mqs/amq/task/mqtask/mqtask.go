// Copyright 2023 The Ryan SU Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mqtask

import (
	"log"
	"time"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/mqs/amq/handler/payjob"
	"github.com/agui-coder/simple-admin-pay-rpc/internal/mqs/amq/types/pattern"
	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"

	"github.com/hibiken/asynq"
)

type PayMQTask struct {
	svcCtx *svc.ServiceContext
	mux    *asynq.ServeMux
}

func NewPayMQTask(svcCtx *svc.ServiceContext) *PayMQTask {
	return &PayMQTask{
		svcCtx: svcCtx,
	}
}

// Start starts the server.
func (m *PayMQTask) Start() {
	m.Register()
	if err := m.svcCtx.AsynqServer.Run(m.mux); err != nil {
		log.Fatalf("failed to start mqtask server, error: %v", err)
	}
}

// Stop stops the server.
func (m *PayMQTask) Stop() {
	time.Sleep(5 * time.Second)
	m.svcCtx.AsynqServer.Stop()
	m.svcCtx.AsynqServer.Shutdown()
}

// Register adds task to cron. | 在此处定义任务处理逻辑，注册worker. TODO 最好移至mq服务中
func (m *PayMQTask) Register() {
	mux := asynq.NewServeMux()
	// define the handler | 定义处理逻辑
	mux.Handle(pattern.PayNotify, payjob.NewPayNotifyWorldHandler(m.svcCtx))
	m.mux = mux
}
