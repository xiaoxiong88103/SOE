# 关于本前端项目

**不喜欢写英文所以写了中文文档方便查阅**
- [开发技术栈](#技术栈)
- [SOE介绍](#SOE 简单运维平台面板)
- 作者：xiaoxiong
- 邮箱：xiaoxiong88103@hotmail.com

## 技术栈

### Web框架和中间件

- `github.com/gin-gonic/gin` v1.9.1: 高性能的HTTP web框架。
- `github.com/swaggo/files` v1.0.1: 用于Swaggo文件服务的库。
- `github.com/swaggo/gin-swagger` v1.6.0: Gin框架的Swagger集成库。

### 配置文件处理

- `github.com/go-ini/ini` v1.67.0: 用于INI文件的Go语言库。

### 身份验证与授权

- `github.com/golang-jwt/jwt/v4` v4.4.3: JWT(JSON Web Tokens)的Go实现。

### 数据库和数据处理

- `github.com/influxdata/influxdb-client-go/v2` v2.13.0: InfluxDB的官方Go客户端。

### 系统监控与性能

- `github.com/shirou/gopsutil` v3.21.11+incompatible: 用于检索系统、进程和网络信息的库。

### RPC和协议

- `google.golang.org/grpc` v1.39.0: gRPC的Go实现，用于RPC通信。
- `google.golang.org/protobuf` v1.32.0: Google的协议缓冲区库的Go实现。

## 间接依赖（部分）

- `github.com/KyleBanks/depth` v1.2.1
- `github.com/bytedance/sonic` v1.10.2
- `github.com/gabriel-vasile/mimetype` v1.4.3
- `github.com/go-playground/validator/v10` v10.17.0
- `github.com/google/uuid` v1.6.0
- `github.com/golang/protobuf` v1.5.3
- `github.com/json-iterator/go` v1.1.12
- `golang.org/x/net` v0.20.0
- `google.golang.org/genproto` v0.0.0-20240102182953-50ed04b92917



# SOE 简单运维平台面板

SOE简单运维面板是一个为国产化系统设计的全面监控解决方案，旨在提供对嵌入式系统、单片机和服务器的高效、可靠的监控支持。此面板特别适用于需要集群部署的环境，无论是小规模的应用场景还是大型的企业级部署，SOE都能提供强大的监控和管理功能。

## 核心功能

- **多平台监控**：支持对嵌入式系统、单片机和服务器的全方位监控，确保系统的高可用性和稳定性。
- **集群部署方案**：提供灵活的集群部署选项，支持多节点部署，增强系统的负载均衡和故障转移能力。
- **节点管理**：自动记录节点（Node）的基础信息，并将信息反馈给主控节点（Master），进而同步至数据库中，实现节点信息的集中管理。
- **节点状态检查**：定期检查节点状态，确保每个节点的运行健康。如果节点发生掉线，系统会自动启用离线记录模式，待节点恢复后自动同步数据至主节点，保证数据的完整性和一致性。

## 技术特点

- **国产化支持**：面向国产化系统设计，兼容国内主流的操作系统和硬件平台，助力国产化生态的发展。
- **高可靠性设计**：通过集群部署和节点状态自动检查机制，确保系统的高可靠性，即使在部分节点出现故障的情况下也能保持系统的整体稳定运行。
- **灵活的部署方案**：支持多种部署模式，包括云部署、本地部署等，满足不同规模和需求的应用场景。
- **易于扩展**：系统设计考虑了未来的扩展性，支持轻松添加新的监控节点和服务，适应业务规模的增长。
- **完全用Go编写**：采用Go语言编写，支持所有架构的Linux操作系统，确保了高性能和跨平台的兼容性。

## 应用场景

- **国产化数据中心**：为国产化数据中心提供全面的监控和管理，确保数据中心的稳定运行。
- **智能制造**：监控生产线上的嵌入式设备和控制系统，提高生产效率和设备可靠性。
- **智慧城市**：对城市基础设施中的各类监控设备进行管理和监控，保障城市运行的智能化和高效率。
- **云计算和大数据中心**：通过集群部署方案，为云计算和大数据中心提供强大的运维支持。
