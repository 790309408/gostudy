


### 🔑 **1. 声明与初始化**  
- **声明格式**：`var mapName map[keyType]valueType`  
  声明后为 `nil` 未初始化，直接赋值会 panic（运行时错误）。  
- **正确初始化**：  
  ```go
  // 方式1：make 初始化（推荐预分配容量）
  m1 := make(map[string]int, 10) // 容量 hint 减少扩容
  
  // 方式2：字面量初始化
  m2 := map[string]int{"张三": 90, "李四": 85}
  ```  
  预分配容量提升性能，避免频繁扩容 。

---

### ⚠️ **2. 键存在性检查**  
使用双返回值模式：  
```go
value, ok := scoreMap["张三"]  
if ok { 
    fmt.Println("存在，值:", value) 
} else { 
    fmt.Println("键不存在") // value 为值类型零值（int 返回 0）
}  
```  
- `ok` 为 `true` 表示键存在，`value` 为对应值；  
- `ok` 为 `false` 时 `value` 无意义 。

---

### 🔄 **3. 遍历（无序性）**  
```go
for k, v := range scoreMap { 
    fmt.Println(k, v) 
}  
```  
- **重要特性**：遍历顺序随机！Go 设计如此，避免开发者依赖顺序 。  
- 仅遍历键：`for k := range scoreMap`  
  仅遍历值：`for _, v := range scoreMap`。

---

### 🗑️ **4. 删除键值对**  
```go
delete(scoreMap, "王五") // 删除键"王五"对应的条目  
```  
- 安全操作：键不存在时 `delete` 不会报错 。  
- 清空 map：需遍历逐个删除（Go 无清空函数）。

---

### ⚙️ **5. 高级特性与技巧**  
1. **并发安全**：  
   原生 map 非并发安全！并发读写需同步：  
   ```go
   var mu sync.Mutex
   mu.Lock()
   m["key"] = value // 写操作
   mu.Unlock()
   ```  
   或使用 `sync.Map`（适合读多写少场景）。  

2. **值类型限制**：  
   - **Key 必须可比较**（支持 `==`）：可用基础类型、结构体（字段均为可比较类型）、指针；  
     禁用切片、函数、包含不可比较字段的结构体 。  
   - **Value 无限制**：可为任意类型（包括嵌套 map）。  

3. **修改结构体 Value**：  
   需整体替换，不可直接修改字段：  
   ```go
   type User struct{ Name string }
   m := make(map[int]User)
   u := m[1]
   u.Name = "new"  // 无效！未修改 map 中的数据
   m[1] = u       // 正确：整体赋值
   ```

---

### 💎 **最佳实践建议**  
- **预分配容量**：已知大小时用 `make(map[k]v, cap)` 减少扩容开销 。  
- **避免大对象 Key**：复杂结构体作 Key 时复制开销大，改用指针（需确保指针指向内容不变）。  
- **零值陷阱**：未初始化 map 赋值会 panic，始终用 `make` 或字面量初始化 。
