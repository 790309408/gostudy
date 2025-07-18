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

package engine

import (
	"errors"
	"fmt"
	"go-study/rulego/api/types"
	"go-study/rulego/components/action"
	"go-study/rulego/components/external"
	"go-study/rulego/components/filter"
	"go-study/rulego/components/flow"
	"go-study/rulego/components/transform"
	"go-study/rulego/utils/reflect"
	"plugin"
	"sync"
)

// PluginsSymbol is the symbol used to identify plugins in a Go plugin file.
const PluginsSymbol = "Plugins"

// Registry is the default registry for rule engine components.
// Registry 是规则引擎组件的默认注册表。
var Registry = new(RuleComponentRegistry)

// init registers default components to the default component registry.
// init 将默认组件注册到默认组件注册表。
func init() {
	var components []types.Node
	// Append components from various packages to the components slice.
	// 将来自各个包的组件附加到组件切片。
	components = append(components, action.Registry.Components()...)
	components = append(components, filter.Registry.Components()...)
	components = append(components, transform.Registry.Components()...)
	components = append(components, external.Registry.Components()...)
	components = append(components, flow.Registry.Components()...)

	// Register all components to the default component registry.
	// 将所有组件注册到默认组件注册表。
	for _, node := range components {
		_ = Registry.Register(node)
	}
}

// RuleComponentRegistry is a registry for rule engine components.
type RuleComponentRegistry struct {
	// components is a map of rule engine node components.
	components map[string]types.Node
	// plugins is a map of plugin components.
	plugins map[string][]types.Node
	// endpointComponents is a map of endpoint components.
	endpointComponents map[string]types.Node
	// RWMutex is a read/write mutex lock.
	sync.RWMutex
}

// Register adds a rule engine node component to the registry.
// 注册将规则引擎节点组件添加到注册表
func (r *RuleComponentRegistry) Register(node types.Node) error {
	r.Lock()
	defer r.Unlock()
	if r.components == nil {
		r.components = make(map[string]types.Node)
	}
	if _, ok := r.components[node.Type()]; ok {
		return errors.New("the component already exists. componentType=" + node.Type())
	}
	r.components[node.Type()] = node

	return nil
}

// RegisterPlugin adds a rule engine node component from a plugin file.
func (r *RuleComponentRegistry) RegisterPlugin(name string, file string) error {
	builder := &PluginComponentRegistry{name: name, file: file}
	if err := builder.Init(); err != nil {
		return err
	}
	components := builder.Components()
	for _, node := range components {
		if _, ok := r.components[node.Type()]; ok {
			return errors.New("the component already exists. componentType=" + node.Type())
		}
	}
	for _, node := range components {
		if err := r.Register(node); err != nil {
			return err
		}
	}

	r.Lock()
	defer r.Unlock()
	if r.plugins == nil {
		r.plugins = make(map[string][]types.Node)
	}
	r.plugins[name] = components
	return nil
}

// Unregister removes a component from the registry by its type.
func (r *RuleComponentRegistry) Unregister(componentType string) error {
	r.RLock()
	defer r.RUnlock()
	var removed = false
	// Check if the plugin exists
	if nodes, ok := r.plugins[componentType]; ok {
		for _, node := range nodes {
			// Delete the plugin from the map
			delete(r.components, node.Type())
		}
		delete(r.plugins, componentType)
		removed = true
	}

	// Check if the plugin exists
	if _, ok := r.components[componentType]; ok {
		// Delete the plugin from the map
		delete(r.components, componentType)
		removed = true
	}

	if !removed {
		return fmt.Errorf("component not found. componentType=%s", componentType)
	} else {
		return nil
	}
}

// NewNode creates a new instance of a rule engine node component by its type.
// NewNode 根据其类型创建规则引擎节点组件的新实例。
func (r *RuleComponentRegistry) NewNode(componentType string) (types.Node, error) {
	r.RLock()
	defer r.RUnlock()

	if node, ok := r.components[componentType]; !ok {
		return nil, fmt.Errorf("component not found. componentType=%s", componentType)
	} else {
		return node.New(), nil
	}
}

// GetComponents returns a map of all registered components.
func (r *RuleComponentRegistry) GetComponents() map[string]types.Node {
	r.RLock()
	defer r.RUnlock()
	var components = map[string]types.Node{}
	for k, v := range r.components {
		components[k] = v
	}
	return components
}

// GetComponentForms returns a list of component forms for all registered components.
func (r *RuleComponentRegistry) GetComponentForms() types.ComponentFormList {
	r.RLock()
	defer r.RUnlock()

	var components = make(types.ComponentFormList)
	for _, component := range r.components {
		components[component.Type()] = reflect.GetComponentForm(component.New())
	}
	return components
}

// PluginComponentRegistry is an initializer for Go plugin components.
type PluginComponentRegistry struct {
	name     string
	file     string
	registry types.PluginRegistry
}

// Init initializes the plugin component registry by loading the plugin from a file.
func (p *PluginComponentRegistry) Init() error {
	pluginRegistry, err := loadPlugin(p.file)
	if err != nil {
		return err
	} else {
		p.registry = pluginRegistry
		return nil
	}
}

// Components returns a slice of components provided by the plugin.
func (p *PluginComponentRegistry) Components() []types.Node {
	if p.registry != nil {
		return p.registry.Components()
	}
	return nil
}

// loadPlugin loads a plugin from a file and registers it with a given name
func loadPlugin(file string) (types.PluginRegistry, error) {
	// Use the plugin package to open the file and look up the exported symbol "Plugin"
	p, err := plugin.Open(file)
	if err != nil {
		return nil, err
	}
	sym, err := p.Lookup(PluginsSymbol)
	if err != nil {
		return nil, err
	}
	// Use type assertion to check if the symbol is a Plugin interface implementation
	plugin, ok := sym.(types.PluginRegistry)
	if !ok {
		return nil, errors.New("invalid plugin")
	}
	// Register the plugin with the name
	//pm.plugins[name] = plugin
	return plugin, nil
}
