# go logger
构建一个非结构化的通用日志模块，支持输出syslog和terminal上。
注意：由于基于go的标准库 log/syslog 实现syslog，这个包没有在Windows上实现。所以
Windows将不能运行syslog的功能。
