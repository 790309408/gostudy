以下是针对鸿蒙ArkTS语言两个核心特性的系统化学习笔记整理，结合官方规范与开发实践，重点关注**运行时对象布局不可变性**和**对象字面量类型强制标注**要求：

---

### ⚠️ 一、**禁止运行时更改对象布局**  
**规则说明**  
ArkTS要求对象在实例化后**不能动态增删属性或方法**，对象结构必须在编译期确定。此设计通过静态内存分配优化性能，减少运行时类型检查开销。  

**违反示例与错误分析**  
```typescript
class Point {
  x: number = 0;
  y: number = 0;
}

const p = new Point();
p.z = 10;                 // 编译错误：Property 'z' does not exist on type 'Point'
delete p.x;               // 编译错误：The operand of a delete operator must be optional
(p as any).z = 10;        // 编译错误：ArkTS禁止通过any绕过类型检查
```  
**根本原因**：动态修改对象布局会导致：  
1. 引擎无法预分配内存，增加GC压力  
2. 破坏类型安全，可能引发运行时异常  
3. 阻碍编译器优化（如内联缓存）  

**正确实践方案**  
- **继承扩展**：通过子类添加新属性  
  ```typescript
  class Point3D extends Point {
    z: number = 0;
  }
  ```
- **组合替代继承**：将新属性封装到独立对象  
  ```typescript
  class PointWithLabel {
    point: Point;
    label: string = "";
  }
  ```
- **不可变数据模式**：使用`Readonly<T>`或`Object.freeze()`  
  ```typescript
  const p: Readonly<Point> = new Point();
  p.x = 1; // 编译错误：Cannot assign to 'x' because it is read-only
  ```

---

### 📌 二、**对象字面量必须显式标注类型**  
**规则说明**  
ArkTS要求所有对象字面量**必须通过接口或类型别名明确定义结构**，禁止依赖类型推断。此约束确保类型系统完整性，消除隐式any风险。  

**违反示例与错误分析**  
```typescript
// 错误写法（TS允许，ArkTS禁止）
const user = { name: "Alice", age: 30 }; 

// 正确写法（显式声明类型）
interface User {
  name: string;
  age: number;
}
const user: User = { name: "Alice", age: 30 }; // 通过编译
```  
**关键差异对比**  
| 特性                | TypeScript       | ArkTS           |  
|---------------------|-----------------|----------------|  
| **类型推断**        | 支持隐式推断     | **禁止**        |  
| **鸭子类型**        | 允许            | **禁止**        |  
| **结构兼容性检查**  | 运行时可能失败  | 编译期强制通过  |  

**复杂类型标注技巧**  
1. **联合类型标注**  
   ```typescript
   type Device = { id: string } & ( 
     { type: "phone"; screenSize: number } | 
     { type: "watch"; waterproof: boolean }
   );
   ```
2. **深度嵌套对象**  
   ```typescript
   interface Company {
     name: string;
     address: {
       city: string;
       zipCode: string;
     };
   }
   ```
3. **泛型约束**  
   ```typescript
   class Response<T extends { status: number }> {
     data: T;
   }
   ```

---

### 🔧 三、规避动态类型的工程实践  
**1. 替代any的方案**  
| 场景                | 替代方案                      | 示例                          |  
|---------------------|------------------------------|-------------------------------|  
| 第三方库数据类型    | 声明.d.ts类型定义文件         | `declare module "lib" { ... }` |  
| 动态API返回         | 使用**类型守卫**              | `if (isApiResponse(res)) { ... }` |  
| 多态数据            | 使用**联合类型+类型断言**     | `(data as VideoContent).url`  |  

**2. 类型守卫实现示例**  
```typescript
function isUser(obj: any): obj is User {
  return typeof obj.name === "string" && 
         typeof obj.age === "number";
}
```

**3. 严格空值检查**  
ArkTS默认启用`strictNullChecks`：  
```typescript
let name: string | null = null;
name.length;  // 编译错误：Object is possibly 'null'
name!.length; // 安全调用（开发者确保非空）
```

---

### 💎 总结：ArkTS静态类型核心优势  
| 特性                | 开发收益                      | 运行时收益                  |  
|---------------------|------------------------------|----------------------------|  
| **对象布局固定**    | 代码可维护性↑，文档化程度↑   | 内存占用↓，执行速度↑       |  
| **强制类型标注**    | 错误在编译期暴露↑            | 消除90%类型检查开销↑       |  
| **空安全**          | 减少NullPointerException↑    | 崩溃率↓，稳定性↑           |  
