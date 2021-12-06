- [ 简介](#head1)
- [ 特性](#head2)
- [ 可以考虑集成进去的东西](#head3)
- [ 待实现功能](#head4)
- [ 快速启动](#head5)
- [ 项目结构](#head6)
	- [ 根目录](#head7)
	- [ util/](#head8)
	- [ data/](#head9)
	- [ router/](#head10)
	- [ middleware/](#head11)
	- [ controller/](#head12)
	- [ service/](#head13)
	- [ dao/](#head14)
	- [ dao/migration/](#head15)

# <span id="head1"> 简介</span>

`gin-layout`是一个单点web后端项目开发模板，致力于打造一个开箱即用的脚手架，使开发人员更专注于业务逻辑。如果有任何问题或者想法欢迎提issue。

# <span id="head2"> 特性</span>

* 通过`gboot`一键生成代码，开箱即用（有关二进制文件在release里面可以找到）
* 优雅重启（linux系统下）
* 内含若干内容的使用示例
* 集成`gorm`作为MySQL驱动和ORM，以及简单的连接池配置
* 内置了简单的分页
* 内置redis连接池
* 方便集成微服务调用
* 定时任务
* 在项目根目录下编写.env来配置应用
* 内置email发送服务
* 通过依赖注入管理应用容器，使代码逻辑清晰
* 集成了用于一键生成TOC for gitea的脚本
* 有Dockerfile，支持基于docker的流水线发布

# <span id="head3"> 可以考虑集成进去的东西</span>

* `"github.com/wonderivan/logger"` 提供了功能更加强大的logger

# <span id="head4"> 待实现功能</span>

* 更完善的日志打印功能
* 便捷的表单验证
* jwt
* 接入消息队列
* email认证功能
* 插件系统
* ……

# <span id="head5"> 快速启动</span>

1. 下载`gboot`并将其放在项目目录或者path中
2. 在命令行中输入`gboot myproject dev`并执行（这里最后的dev是工具用来指定开发分支的参数，省略掉的情况下是拉取main版本）
3. 配置本地环境以及`.env`文件
4. 执行`go build`
5. 运行生成的二进制文件

# <span id="head6"> 项目结构</span>

```
├─controller // 控制器层，用于解析请求、传参到service层并封装返回数据
├─service // service层，处理具体的业务逻辑
├─dao  // dao层，数据库对象及其操作都应该定义在这里
│  └─migration // 数据库迁移，数据库表更变的操作都定义在这里
├─data // data层，提供数据库驱动
│  ├─mysql
│  └─redis
├─middleware // 中间件层，这个应该不需要解释了
├─app.go // 定义了应用容器及其依赖
├─router.go // router，定义路由，建立请求和控制器方法之间的映射
├─server_xx.go // 定义了容器启动http服务的方法，一般不修改
├─wire.go // 定义了wire注入的方法，一般也不修改
├─wire_gen.go // 自动生成的代码，不修改
└─util // 可复用的非业务逻辑代码集合（其实更推荐直接单独开一个util仓库然后push上去）
```

## <span id="head7"> 根目录</span>

根目录下有程序的入口，还有封装了不同操作系统下的server启动函数。在linux中，我们会使用`endless`包着`gin`来实现优雅重启；而在windows中，由于`endless`实现依赖的系统信号只在unix操作系统中才有，所以去掉了这个组件，直接启动gin服务。

## <span id="head8"> util/</span>

业务无关的工具代码库。这个文件夹实际上不应该存在……工具类的代码最好还是封装成库然后push到仓库再进行引用，这样一来方便管理而来方便引用。有时候一些代码急着要用图个方便直接放在这里也是ok的（本目录存在的唯一意义）

## <span id="head9"> data/</span>

这个文件夹下会实例化项目使用的数据库的驱动。目前只实现了MySQL的驱动并且没有对连接池之类的配置进行设置。后续会对其进行完善并且加入其他的数据库驱动（例如redis）。

## <span id="head10"> router/</span>

这个文件夹中定义了所有`path`到控制器层的映射。

> 其实关于路由要不要统一管理这个事一直都是有争议的，像是php和golang这边一般倡导都写在一个文件里面，这样方便开发人员拿着接口查到对应的函数；而像python和java这种则是把路由直接写在控制器（蓝图）的做法比较常见，因为这样开发会相对简单一点点。考虑到我们部门之前的技术栈一直都是php，所以还是直接沿用了以前的习惯，将路由和控制器分开管理。

## <span id="head11"> middleware/</span>

中间件层。中间件提供了一种方便的机制来过滤进入应用程序的 HTTP 请求。一般用于鉴权、处理跨域、打日志等场景，用于在请求到达控制器层之前拦截之或进行其他的一些操作。

## <span id="head12"> controller/</span>

控制器层。这一业务层的主要职责是解析请求、对得到的入参做一个validation、调用服务层的服务、封装请求的响应。也就是说，这里不应该出现具体的业务逻辑。

## <span id="head13"> service/</span>

服务层。这里需要实现绝大部分具体业务逻辑，也是业务逻辑能到达的最底层。`*gin.Context`不应该在服务层中出现。服务层中的出参和入参应该都是基本类型或者业务相关的一些结构体。大部分业务相关的单元测试都应该在这里编写。

## <span id="head14"> dao/</span>

数据库操作对象层。这里定义了所有数据库对象的实体，以及这些实体的基本crud操作。同样地，这一层更加不应该出现`*gin.Context`以及其他和具体业务相关的内容（包括一些乱七八糟的结构体：因为要传进来肯定要引入包，但是dao层几乎是最底层，不能引用上层的任何东西，不然会发生循环引用的问题）

## <span id="head15"> dao/migration/</span>

数据库迁移。只有一个文件是需要修改的：那就是`schema.go`。里面有示例如何创建示例，只要跟着示例继续平铺代码即可。一般使用`gorm`的`AutoMigrate()`或者`Migrator`来修改数据库表。里面内置了一个简单的版本管理，所以不用担心迁移是否会被重复执行的问题。注意：迁移里面的内容只能增加，不能删除或者修改！也就是说，如果你发现前面的实现有问题，你只能发布一个新的迁移来进行修改，而不应该修改过往版本的迁移（即使是可以删库重开也是这样，因为别人的电脑上也有开发环境，为了节省成本还是应该发布新的迁移实现修改。这方面的话可以灵活处理，但最好还是遵守约定。）
