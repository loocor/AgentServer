Readme
======

agent_server 是基于 go-zero 微服务框架的试验性项目，提供账户管理等功能，目前正在学习、开发中。

## 版本管理

尝试采用单体仓库（Mono-Repo）进行版本管理，目前托管在 Github 上，后期考虑使用 GitLab 的 CI/CD 流程，可能会迁移、调整。有
关多体仓库与单体仓库的比较，这里有篇文章介绍：http://weibo.ws/TwhhgB

## 目录结构

参考了一些文章和 go-zero 相关的视频，决定采用按模块分组，组内再分 api、docs、model、service 等目录。