### 一键部署以下项目

- [x] nodejs                    开发体验    慢
- [x] python
- [x] nginx
- [x] apache/httpd
- [x] nodejs+mongo
- [x] python+redis
- [x] tomcat    
- [x] egg.js                    开发体验    慢
- [x] create-react-app          开发体验    极速
- [x] golang                    开发体验    1.无法监听项目下文件变更，实现动态热更golang项目。2.可以规避的问题
- [x] vue-cli                   开发体验    快
- [x] tensorflow                地址:       [docker pull pengliheng/tf](https://cloud.docker.com/repository/docker/pengliheng/tf/general)


### `docker build --tag=<my project>`  开发环境部署
```bash
docker build --tag=pengliheng/<my project>                                            # 新建docker image
docker run --rm -it -d -v $(pwd):/code -w -p 80:80 /code pengliheng/<my project>      # 在image层下新建一个container
docker ps -a -s                                                                       # 查询container id
docker attach <container id>                                                          # hack 进入container 进行开发
```



### `docker-composer up`  生产环境部署
需要删除image才能修改之前提交的参数,,,,或许重新build一次可以改善问题...
