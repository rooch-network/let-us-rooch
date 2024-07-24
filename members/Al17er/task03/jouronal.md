##安装bun：
curl -fsSL https://bun.sh/install | bash

![alt text](image.png)

##配置bun环境变量：
PATH=/root/.bun/bin/:$PATH
![
](image-1.png)

##下载dapp源码：
git clone https://github.com/rooch-network/my-first-rooch-dapp.git
![alt text](image-2.png)

##安装依赖:
bun install
![alt text](image-3.png)

##启动dapp:bun dev --host 192.168.221.147
![alt text](image-4.png)

##访问192.168.221.147:5173

![alt text](image-5.png)

##发布合约程序：
rooch move publish --named-addresses quick_start_counter=default
![alt text](image-6.png)

##将发布程序地址添加到：src/App.tsx
![alt text](image-7.png)

##浏览器连接钱包：
![
](image-8.png)

##创建session
![alt text](image-10.png)

##调用increase Counter Value
![alt text](image-11.png)