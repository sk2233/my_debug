### Go 语言实现的命令行 debugger
> 参考：https://www.hitzhangjie.pro/debugger101.io/

TODO：参考 https://github.com/hitzhangjie/tinydbg 实现语句 debug
#### exec
![img.png](img.png)<br>
自动新开进程作为调试进程，可以再代码开始处断点
#### attach
![img_1.png](img_1.png)<br>
附件调试进程到已有进程
#### disass
![img_2.png](img_2.png)<br>
反汇编当前 pc 指针后指定字节数的代码
#### break
![img_3.png](img_3.png)<br>
在指定 pc 寄存器中设置断点
#### breaks
![img_4.png](img_4.png)<br>
查看当前所有断点
#### clear
![img_5.png](img_5.png)<br>
清楚指定 id 的传断点，若是不指定，全部清楚
#### step
![img_6.png](img_6.png)<br>
单步执行
#### continue
![img_7.png](img_7.png)<br>
直接执行到断点处
#### mem
![img_8.png](img_8.png)<br>
查看内存（写入也是类似的）
#### regs
![img_9.png](img_9.png)<br>
查看寄存器（写入也是类似的）
#### exit
![img_10.png](img_10.png)<br>
退出调试器
