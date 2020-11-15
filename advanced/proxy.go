package main

import (    
	"bytes"
    "fmt"
    "io"
    "log"
    "net"
    "net/url"
    "strings"
)

func main2() {
    // 设置日志格式
    log.SetFlags(log.LstdFlags|log.Lshortfile)
    // 监听tcp端口
    ln, err := net.Listen("tcp", ":8081")    
    if err != nil {
        // 错误时，记录错误，处理defer退出
        log.Panic(err)
    }    
    for {
        // 等待客户端连接
        client, err := ln.Accept()        
        if err != nil {
            log.Panic(err)
        }        
        // 通过goroutine来处理用户的请求
        go handleClientRequest(client)
    }
}

func handleClientRequest(client net.Conn) {
    if client == nil {        
        return
    } 
    // 关闭连接   
    defer client.Close()    
    var b [1024]byte
    // 读取连接信息
    n, err := client.Read(b[:])    
    if err != nil {
        log.Println(err)        
        return
    }    
    var method, host, address string
    // 返回 s 的的第一个实例的索引
    fmt.Println(bytes.IndexByte(b[:], '\n'))
    http_string := string(b[:bytes.IndexByte(b[:], '\n')])
    // 字符串扫描输入
    fmt.Sscanf(http_string, "%s%s", &method, &host)
    hostPortURL, err := url.Parse(host)    
    if err != nil {
        log.Println(err)        
        return
    }    
    if hostPortURL.Opaque == "443" { //https访问
        address = hostPortURL.Scheme + ":443"
    } else { //http访问
        if strings.Index(hostPortURL.Host, ":") == -1 { //host不带端口，默认80
            address = hostPortURL.Host + ":80"
        } else {
            address = hostPortURL.Host
        }
    }    
    
    //获得了请求的host和port，就开始拨号吧
    server, err := net.Dial("tcp", address)    
    if err != nil {
        log.Println(err)        
        return
    }    

    // ??
    if method == "CONNECT" {
        fmt.Fprint(client, "HTTP/1.1 200 Connection established\r\n")
    } else {
        // ??
        server.Write(b[:n])
    }    
    
    //进行转发
    go io.Copy(server, client)
    io.Copy(client, server)
}