# Offchain Storage

Offchain Storage 是一个围绕去中心化存储、密码学与证明协议的实验性/研究型项目，包含 Go 主体代码、若干协议实现（PoD/PoS/PoST/TMPS）、IPFS 交互示例、以及少量 Node/Python/Solidity 资源。

## 目录结构
- `cmd/`
  - `graph-demo/`: 从根目录迁移的图结构示例入口（Go 可执行）。
  - `ipfs-hello/`: 从 `core/api` 迁移的 IPFS 交互示例入口（Go 可执行）。
- `core/`
  - `database/`: 基于 OrbitDB/IPFS 的数据库集成与交互逻辑。
  - `storage_api/`: 存储相关的 API、CID 计算、上传封装等。
  - 其他 `core` 下的加密/数学/工具等包。
- `blockchain-crypto/`: 曲线、重加密、随机数、数学等独立加密模块（Go）。
- `protocol/`
  - `PoD/`: Proof of Data 相关代码与合约、测试、脚本。
  - `PoS/`: Proof of Space 相关模块（独立 `go.mod`）。
  - `PoST/`: Proof of Space-Time 相关模块与测试。
  - `TMPS/`: 实验性多项式承诺/双线性群相关代码与主程序。
- `core/api/`: Node 侧的 API 示例与 Solidity 合约样例（保留在原位置）。
- `speed.sh`/`speed_data.json`: 网络测速脚本与数据。

说明：为了不破坏现有包路径与导入关系，本次重构仅将两个 Go 可执行入口调整到 `cmd/` 目录，其余包保持不动。后续若需要进一步标准化（如引入 `internal/`、统一 `pkg/` 命名等），可在编译验证后逐步进行。

## 环境要求
- Go 1.19 或更高
- Node.js（如需运行 `core/api` 下的 JS/前端部分）
- Python（如需运行 `protocol` 下的脚本）
- Solidity 0.8+（如需编译/部署合约）
- 本地 IPFS 节点（如需进行 IPFS 交互）

## 快速开始
1. 拉取依赖：
   ```bash
   go mod tidy
   ```
2. 构建示例可执行文件（需已有 `make` 环境）：
   ```bash
   make build
   ```
   构建完成后，二进制位于 `bin/` 目录：
   - `bin/graph-demo`
   - `bin/ipfs-hello`

3. 运行示例：
   - Graph Demo：
     ```bash
     ./bin/graph-demo
     ```
   - IPFS Hello（确保本地 IPFS API 在 `localhost:5001`）
     ```bash
     ./bin/ipfs-hello
     ```

## 重要模块简述
- `core/database`: 与 OrbitDB、libp2p、IPFS 的结合。包含主机配置、事件处理、文档存储、访问控制等。
- `core/storage_api`: 包含 `CalculateFileCID`、本地文件目录组织、上传到远端（结合配置）等。
- `blockchain-crypto`: 椭圆曲线点运算、重加密（`recrypt`）、随机大整数生成、AES-GCM 等。
- `protocol/PoD/PoS/PoST/TMPS`: 各类证明协议的实现、测试与脚本，适合分别进入子目录阅读与运行。

## 开发建议
- 保持包命名清晰，避免在同一目录混放可执行入口与库代码。
- 新的可执行入口统一放置在 `cmd/<name>/main.go`。
- 后续若需要进一步模块化，可考虑：
  - 将通用库移动到 `pkg/`。
  - 将仅供本项目内部使用的代码放到 `internal/`。
  - 为 Node/Python/合约引入 `web/`、`scripts/`、`contracts/` 等清晰的边界目录。

## 运行 Node/合约与脚本
- `core/api`: 保留 Node/JS 与示例合约，可在该目录下按 `package.json` 指南进行。
- `protocol/PoD`: 包含合约与 Python 部署脚本，依说明运行。

## 许可
未明确授权文件，默认遵循仓库现有 LICENSE（若无则需要补充）。
