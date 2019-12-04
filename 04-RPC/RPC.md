客户端在不知道调用细节的情况下，调用存在于远程计算机的某个对象，就像调用本地应用程序中的对象一样

 ![rpc1](http://image.chaindesk.cn/rpc1.png) 

 在golang中实现RPC非常简单，有封装好的官方库和一些第三方库提供支持。Go RPC可以利用tcp或http来传递数据，可以对要传递的数据使用多种类型的编解码方式。  golang官方的`net/rpc`库使用`encoding/gob`进行编解码，支持`tcp`或`http`数据传输方式，由于其他语言不支持`gob`编解码方式，所以使用`net/rpc`库实现的RPC方法没办法进行跨语言调用。 

 golang官方还提供了`net/rpc/jsonrpc`库实现RPC方法，JSON RPC采用JSON进行数据编解码，因而支持跨语言调用。但目前的jsonrpc库是基于tcp协议实现的，暂时不支持使用http进行数据传输。 