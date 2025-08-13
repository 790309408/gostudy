以下是针对你提供的 ArkTS 学习笔记的系统整理，结合语法规范与开发实践，重点覆盖数据类型、枚举、联合类型和数组的核心特性：

---

### 一、**基础类型**
#### 1. **`number` 数值类型**
- **声明与赋值**  
  ```typescript
  let n1: number = 10;      // 整数
  let n2: number = 43.5;    // 浮点数
  let n3 = n1 + n2;         // 自动类型推断为 number
  ```
- **进制表示**  
  ```typescript
  let b1: number = 0o17;    // 八进制 → 十进制 15
  let b2: number = 0x17;    // 十六进制 → 十进制 25
  let b: number = 0b111;    // 二进制 → 十进制 7
  ```
- **注意**：大整数需用 `BigInt` 避免精度丢失。

#### 2. **`boolean` 布尔类型**
- **逻辑判断基础**  
  ```typescript
  let flag: boolean = false;
  if (flag) {
    console.log("条件成立");
  }
  ```

#### 3. **`string` 字符串类型**
- **三种声明方式**  
  ```typescript
  let s1: string = '单引号字符串';
  let s2: string = "双引号字符串";
  let s3: string = `模板字符串：${s1}`; // 支持表达式嵌入
  ```

---

### 二、**枚举类型（`enum`）**
- **自增数字枚举**  
  ```typescript
  enum Color {
    Red = 1,   // 显式赋值
    Green,     // 自动递增为 2
    Blue       // 自动递增为 3
  }
  ```
- **使用场景**  
  ```typescript
  if (Color.Red === 1) {
    console.log('红色');  // 输出验证
  }
  ```
- **优势**：通过命名常量提升代码可读性，避免魔法数字。

---

### 三、**联合类型（Union Types）**
- **多类型兼容**  
  ```typescript
  type Union = string | number;
  let s: Union = '1';     // 支持字符串
  let n: Union = 2;       // 支持数值
  ```
- **元组扩展**  
  ```typescript
  type MixedTuple = [...string[], string, number];
  let tuple: MixedTuple = ['a', 'b', 'last', 100]; // 灵活结构
  ```
- **作用**：增强类型灵活性，需配合类型守卫确保运行时安全。

---

### 四、**数组（`Array`）**
#### 1. **基础声明**
```typescript
let students: string[] = [];        // 字符串数组
let stus: number[] = new Array(10); // 初始化长度为10的数值数组
```

#### 2. **动态扩容特性**
```typescript
students[100] = 'end'; // 数组自动扩容至长度101，未赋值项为 undefined
```
**注意**：避免跳跃赋值，防止内存浪费和未定义行为。

#### 3. **泛型语法（等价声明）**
```typescript
let altList: Array<string> = ["A", "B"]; // 与 string[] 等效
```

---

### 五、**关键特性总结**
| **类型**       | **语法示例**                     | **核心特性**                                  |
|----------------|----------------------------------|---------------------------------------------|
| `number`       | `let hex = 0x1F;`               | 支持多进制、整数/浮点数                      |
| `boolean`      | `let enabled = true;`           | 仅 `true`/`false`，用于条件逻辑              |
| `string`       | `` `Hello ${name}` ``           | 模板字符串支持表达式插值                     |
| `enum`         | `enum Status { Active, Done }`   | 自增常量，提升可读性                         |
| 联合类型       | `let id: string | number;`       | 兼容多种类型，需类型守卫                     |
| `Array`        | `let nums: number[] = [1,2];`   | 动态扩容，避免越界赋值                       |

> **开发建议**：  
> - 优先使用 `enum` 替代数字常量，增强可维护性；  
> - 联合类型需通过 `typeof` 或 `instanceof` 进行类型收缩；  
> - 数组操作时预分配长度，避免动态扩容的性能开销。