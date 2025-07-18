/*
 * Copyright 2023 The RuleGo Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package types

import (
	"math"
	"time"

	"go-study/rulego/api/pool"
)

// OnDebug is a global debug callback function for nodes.
// OnDebug 是节点的全局调试回调函数。
var OnDebug func(ruleChainId string, flowType string, nodeId string, msg RuleMsg, relationType string, err error)

// Config defines the configuration for the rule engine.
// Config 定义规则引擎的配置。
type Config struct {
	// OnDebug is a callback function for node debug information. It is only called if the node's debugMode is set to true.
	// OnDebug 是节点调试信息的回调函数，仅当节点的 debugMode 设置为 true 时才会调用
	// - ruleChainId: The ID of the rule chain.
	// - flowType: The event type, either IN (incoming) or OUT (outgoing) for the component.
	// - nodeId: The ID of the node.
	// - msg: The current message being processed.
	// - relationType: If flowType is IN, it represents the connection relation between the previous node and this node (e.g., True/False).
	//                 If flowType is OUT, it represents the connection relation between this node and the next node (e.g., True/False).
	// - err: Error information, if any.
	OnDebug func(ruleChainId string, flowType string, nodeId string, msg RuleMsg, relationType string, err error)
	// OnEnd is a deprecated callback function that is called when the rule chain execution is complete. If there are multiple endpoints, it will be executed multiple times.
	// Deprecated: Use types.WithEndFunc instead.
	// OnEnd 是一个已弃用的回调函数，在规则链执行完成时调用。如果有多个端点，则会执行多次。
	OnEnd func(msg RuleMsg, err error)
	// ScriptMaxExecutionTime is the maximum execution time for scripts, defaulting to 2000 milliseconds.
	// Script Execution Time 是脚本的最大执行时间，默认为 2000 毫秒
	ScriptMaxExecutionTime time.Duration
	// Pool is the interface for a coroutine pool. If not configured, the go func method is used by default.
	// Pool 是协程池的接口，如果不配置，则默认使用 go func 方法
	// The default implementation is `pool.WorkerPool`. It is compatible with ants coroutine pool and can be implemented using ants.
	// 默认实现是 `pool.WorkerPool`，兼容 Ant 协程池，可以使用 Ant 实现
	// Example:
	//   pool, _ := ants.NewPool(math.MaxInt32)
	//   config := rulego.NewConfig(types.WithPool(pool))
	Pool Pool
	// ComponentsRegistry is the component registry, defaulting to `rulego.Registry`.
	// ComponentsRegistry 是组件注册表，默认为 `rulego.Registry`
	ComponentsRegistry ComponentRegistry
	// Parser is the rule chain parser interface, defaulting to `rulego.JsonParser`.
	// Parser 是规则链解析器接口，默认为 `rulego.JsonParser`
	Parser Parser
	// Logger is the logging interface, defaulting to `DefaultLogger()`.
	// Logger 是日志接口，默认为 `DefaultLogger()`。
	Logger Logger
	// Properties are global properties in key-value format.
	// 属性是键值格式的全局属性。
	// Rule chain node configurations can replace values with ${global.propertyKey}.
	// 规则链节点配置可以用${global.propertyKey}替换值。
	// Replacement occurs during node initialization and only once.
	// 替换发生在节点初始化期间，并且仅发生一次
	Properties Metadata
	// Udf is a map for registering custom Golang functions and native scripts that can be called at runtime by script engines like JavaScript.
	// Udf 是一个用于注册自定义 Golang 函数和原生脚本的映射，可以在运行时被 JavaScript 等脚本引擎调用。
	// Function names can be repeated for different script types.
	// 不同脚本类型的函数名称可以重复。
	Udf map[string]interface{}
	// SecretKey is an AES-256 key of 32 characters in length, used for decrypting the `Secrets` configuration in the rule chain.
	// SecretKey 是一个长度为 32 个字符的 AES-256 密钥，用于解密规则链中的 `Secrets` 配置。
	SecretKey string
	// EndpointEnabled indicates whether the endpoint module in the rule chain DSL is enabled.
	// EndpointEnabled 表示规则链 DSL 中的端点模块是否启用
	EndpointEnabled bool
	// NetPool is the interface for a shared Component Pool.
	// NetPool 是共享组件池的接口
	NetPool NodePool
	// NodeClientInitNow indicates whether to initialize the net client node immediately after creation.
	// NodeClientInitNow 表示是否在创建后立即初始化网络客户端节点
	// True: During the component's Init phase, the client connection is established. If the client initialization fails, the rule chain initialization fails.
	// 组件在Init阶段，会建立客户端连接，如果客户端初始化失败，则规则链初始化失败。
	// False: During the component's OnMsg phase, the client connection is established.
	// 在组件的 OnMsg 阶段，建立客户端连接。
	NodeClientInitNow bool
	// AllowCycle indicates whether nodes in the rule chain are allowed to form cycles.
	// AllowCycle 表示规则链中的节点是否允许形成循环
	AllowCycle bool
}

// RegisterUdf registers a custom function. Function names can be repeated for different script types.
func (c *Config) RegisterUdf(name string, value interface{}) {
	if c.Udf == nil {
		c.Udf = make(map[string]interface{})
	}
	if script, ok := value.(Script); ok {
		// Resolve function name conflicts for different script types.
		name = script.Type + ScriptFuncSeparator + name
	}
	c.Udf[name] = value
}

// NewConfig creates a new Config with default values and applies the provided options.
// NewConfig 使用默认值创建一个新的配置并应用提供的选项。
func NewConfig(opts ...Option) Config {
	c := &Config{
		ScriptMaxExecutionTime: time.Millisecond * 2000,
		Logger:                 DefaultLogger(),
		Properties:             NewMetadata(),
		EndpointEnabled:        true,
	}

	for _, opt := range opts {
		_ = opt(c)
	}
	return *c
}

// DefaultPool provides a default coroutine pool.
// DefaultPool 提供了一个默认的协程池。
func DefaultPool() Pool {
	wp := &pool.WorkerPool{MaxWorkersCount: math.MaxInt32}
	wp.Start()
	return wp
}
