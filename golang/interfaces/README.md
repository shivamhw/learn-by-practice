# Go Interfaces: Practice & Notes


---

## Go Interface Questions & Answers

### 1. How does passing an interface to a function work?  
When you pass an interface to a function, the interface value contains two things:  
- The **type** of the concrete value it holds  
- The **pointer** (address) to the concrete value

So, when you pass an interface, a copy of the interface value is made, but both the original and the copy point to the same underlying concrete value.

**Diagram:**

```
+-------------------+         +-------------------+
|   Interface (A)   |         |   Interface (B)   |
|-------------------|         |-------------------|
| Type: *MyStruct   |         | Type: *MyStruct   |
| Ptr: 0x1234abcd   |         | Ptr: 0x1234abcd   |
+-------------------+         +-------------------+
           |                           |
           +-----------+---------------+
                       |
              +-------------------+
              |   MyStruct value  |
              |   (at 0x1234abcd) |
              +-------------------+
```

Both interface variables (`A` and `B`) contain the same type and pointer, so they reference the same underlying value. Ref ex1.go

### 2. How to marshal/unmarshal interfaces?  
Interfaces cant have custom marhal/unmashal function as functions can only be implemented on concrete classes.

### 3. How to check the type of an interface?  
easy one would be switch and other one would be reflect.TypeOf

### 4. Why does the error "pointer does not implement the interface" occur?  
Never use pointers with interfaces, there is little to no use for them, they are allowed but increase complexity and is almost never useful

---

## References

-
