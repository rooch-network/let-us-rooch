# 学习日志

## 搭建节点 & 节点同步
Reference: https://github.com/rooch-network/rooch/tree/main/scripts/bitcoin

遇到的问题：
- 权限问题:
    ```
    Error: Settings file could not be written:
    - Error: Unable to open settings file /data/.bitcoin/settings.json.tmp for writing
    ```

    solution：

    `touch ~/.bitcoin/ `

- 连接主网:
    更新 https://github.com/rooch-network/rooch/blob/main/scripts/bitcoin/node/run_local_node_docker.sh， 去掉 `-chain=regtest \`




