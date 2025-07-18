package rulego

import (
	"go-study/rulego/api/types"
	"go-study/rulego/endpoint"
	"go-study/rulego/engine"
	"go-study/rulego/node_pool"
)

// Registry 是规则引擎组件的默认注册表。
var Registry = engine.Registry

// Rules 是RuleGo的默认实例，其规则引擎池设置为默认池。
var Rules = &RuleGo{pool: engine.DefaultPool}

var Endpoints = endpoint.DefaultPool

// 确保RuleGo实现RuleEnginePool接口。
var _ types.RuleEnginePool = (*RuleGo)(nil)

// RuleGo 是规则引擎实例的池。
type RuleGo struct {
	pool *engine.Pool
}

// NewRuleGo 创建一个新的RuleGo实例。
func NewRuleGo() *RuleGo {
	return &RuleGo{
		pool: engine.NewPool(),
	}
}

// Pool 返回规则引擎池。
func (g *RuleGo) Pool() *engine.Pool {
	return g.pool
}

// Load 从指定文件夹及其子文件夹加载所有规则链配置到规则引擎实例池中。
// 规则链ID取自规则链文件中指定的ruleChain.id。
func (g *RuleGo) Load(folderPath string, opts ...types.RuleEngineOption) error {
	if g.pool == nil {
		g.pool = engine.NewPool()
	}
	return g.pool.Load(folderPath, opts...)
}

// New 创建一个新的RuleEngine并将其存储在RuleGo规则链池中。
// 如果指定的id为空（""），则使用规则链文件中的ruleChain.id。
func (g *RuleGo) New(id string, rootRuleChainSrc []byte, opts ...types.RuleEngineOption) (types.RuleEngine, error) {
	return g.pool.New(id, rootRuleChainSrc, opts...)
}

// Get 通过ID获取规则引擎实例。
func (g *RuleGo) Get(id string) (types.RuleEngine, bool) {
	return g.pool.Get(id)
}

// Del 通过ID删除规则引擎实例。
func (g *RuleGo) Del(id string) {
	g.pool.Del(id)
}

// Stop 释放所有规则引擎实例。
func (g *RuleGo) Stop() {
	g.pool.Stop()
}

// Range 遍历所有规则引擎实例。
func (g *RuleGo) Range(f func(key, value any) bool) {
	g.pool.Range(f)
}

// Reload 重新加载所有规则引擎实例。
func (g *RuleGo) Reload(opts ...types.RuleEngineOption) {
	g.pool.Reload(opts...)
}

// OnMsg 调用所有规则引擎实例处理消息。
// 规则引擎实例池中的所有规则链都将尝试处理该消息。
func (g *RuleGo) OnMsg(msg types.RuleMsg) {
	g.pool.Range(func(key, value any) bool {
		if item, ok := value.(types.RuleEngine); ok {
			item.OnMsg(msg)
		}
		return true
	})
}

// Load 从指定文件夹及其子文件夹加载所有规则链配置到规则引擎实例池中。
// 规则链ID取自规则链文件中指定的ruleChain.id。
func Load(folderPath string, opts ...types.RuleEngineOption) error {
	return Rules.Load(folderPath, opts...)
}

// New 创建一个新的RuleEngine并将其存储在RuleGo规则链池中。
func New(id string, rootRuleChainSrc []byte, opts ...types.RuleEngineOption) (types.RuleEngine, error) {
	return Rules.New(id, rootRuleChainSrc, opts...)
}

// Get 通过ID获取规则引擎实例。
func Get(id string) (types.RuleEngine, bool) {
	return Rules.Get(id)
}

// Del 通过ID删除规则引擎实例。
func Del(id string) {
	Rules.Del(id)
}

// Stop 释放所有规则引擎实例。
func Stop() {
	Rules.Stop()
}

// OnMsg 调用所有规则引擎实例处理消息。
// 规则引擎实例池中的所有规则链都将尝试处理该消息。
func OnMsg(msg types.RuleMsg) {
	Rules.OnMsg(msg)
}

// Reload 重新加载所有规则引擎实例。
func Reload(opts ...types.RuleEngineOption) {
	Rules.Range(func(key, value any) bool {
		_ = value.(types.RuleEngine).Reload(opts...)
		return true
	})
}

// Range 遍历所有规则引擎实例。
func Range(f func(key, value any) bool) {
	Rules.Range(f)
}

// WithConfig 是一个设置RuleEngine的Config的选项。
func WithConfig(config types.Config) types.RuleEngineOption {
	return engine.WithConfig(config)
}

// NewConfig 创建一个新的Config并应用选项。
// opts ...types.Option 表示可选参数列表。
func NewConfig(opts ...types.Option) types.Config {
	config := engine.NewConfig(opts...)
	if config.NetPool == nil {
		config.NetPool = node_pool.NewNodePool(config) // 创建一个共享组件池
	}
	return config
}
