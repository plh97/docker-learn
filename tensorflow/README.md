### 机器学习 tensorflow 练手项目


### docker 仓库 
```docker
docker pull pengliheng/tf:latest
docker run -it --name tf --rm -d -v $(pwd):/root/tensorflow pengliheng/tf
docker exec $(docker container --name=tf) # 进入容器内部开发
```

### docker 内容
- Ubuntu 环境 
- 自带zsh终端, 
- apt-get update
- apt-get upgrade
- vim
- tensorflow 做为基础image进行一层封装
- python2
- git

环境是docker的环境,当前目录与docker容器里面共享同一个目录
![image](https://user-images.githubusercontent.com/14355994/51209044-40fb5c00-194a-11e9-866a-3006524a2dc6.png)




### TODOLIST
- [ ] 验证码识别
- 生成100个验证码图片 assets/images




### loss 机器学习-损失函数
