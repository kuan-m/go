- ### acceess by index
![acceess by index](./img/acceess_by_index.png)

When you pass an array to a function in Go, the entire array is copied — which consumes time and memory. To avoid copying, use slices (which are reference types) or pass a pointer to the array