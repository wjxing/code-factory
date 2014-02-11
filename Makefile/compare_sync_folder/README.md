# Compare and sync folders

------

这个Makefile可以把源目录里面的文件同步到目标目录中。

* 示例

在当前目录下执行：
```
make SYNC_SRC=files1 SYNC_DES=files2
```
可以把files1目录中的文件同步到files2目录中。

* 注意

1. 当源目录和目标目录里面有相同文件时不会拷贝（即使两者的相对路径不同）

2. 当源目录和目标目录里面相同的路径下有同名文件，并且文件名不相同时不会拷贝

3. 可以在**sync_files_prefix**变量中指定过滤的文件后缀
