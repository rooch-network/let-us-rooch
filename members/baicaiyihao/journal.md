# 学习日志
**task1**  
1、mac安装bitcoin core  
访问 https://bitcoincore.org/en/download/ 下载bitcoin core.app并将其放入应用程序    
![img](./task1/task1-1.jpg)    
然后打开app，配置数据目录    
![img](./task1/task1-2.jpg)  
并且开启rpc节点  
![img](./task1/task1-3.jpg)  
等区块同步完成  
![img](./task1/task1-4.jpg)  

**task2**  
1、修改rpc  
原rpc端口为50051修改为新端口6767  
![img](./task2/task2.jpg)  
![img](./task2/task2-1.jpg)  
2、部署合约  
![img](./task2/task2-2.jpg)  

**task3**  
1、部署dapp  
下载代码  
 `git clone https://github.com/rooch-network/my-first-rooch-dapp.git`  
安装bun  
`curl -fsSL https://bun.sh/install | bash`  
下载依赖  
`bun install`
部署  
`my-first-rooch-dapp % bun dev`  
`counter_contract % rooch move publish --named-addresses quick_start_counter=default`  
然后修改app.tsx  
![img](./task3/task3-1.jpg)   
2、合约交互  
![img](./task3/task3-2.jpg)  