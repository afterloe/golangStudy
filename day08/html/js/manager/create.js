ReactDOM.render(
    <div className="createView">
        <nav class="nav">
            <a class="nav-link active" href="#"><span class="badge badge-danger">1</span> 上传tar包</a>
            <a class="nav-link disabled" href="#"><span class="badge badge-danger">2</span> 构建</a>
            <a class="nav-link disabled" href="#"><span class="badge badge-danger">3</span> 保存镜像</a>
        </nav>
        <span className={"cwWhete"}></span>
        <div className={"uploadTar"}>
            <form>
                <div class="form-group">
                    <label for="selectUploadFile" className={"form-control-lg"}>选择tar包</label>
                    <input type="file" class="form-control-file form-control-lg"
                           aria-describedby="fileHelp" placeholder="请上传镜像包." />
                </div>
                <small class="text-muted">
                    请上传使用tar.gz 格式压缩的tar包，否则将无法解压。tar包中请包含 dockerfile
                </small>
                <span className={"cwWhete"}></span>
                <button type="submit" class="btn btn-primary disabled float-right">下一步</button>
            </form>
        </div>
    </div>
    , document.getElementById("root"));

const navbarData = {
    linkedHref: "/",
    name: "Cityworks™ 云平台",
    cWNavbarInputForm: {
        word: "搜索获取容器"
    },
    cWNavbarRouters: {
        activeRouter: -1,
        routers: [
            {
                name: "镜像",
                href: "/images.html"
            },
            {
                name: "容器",
                href: "/containers.html"
            },
            {
                name: "服务",
                href: "/services.html"
            },
            {
                name: "其他",
                href: "/others.html"
            }
        ]
    },
    barStatus: {
        newMenuItem: [
            {
                name: "创建镜像",
                href: "/manager/create.html"
            },
            {
                name: "创建容器",
                href: "/image/runAsCondition.html"
            },
            {
                name: "创建服务",
                href: "/image/runAsService.html"
            }
        ],
        profileMenuItem: [
            {
                name: "个人资料",
                href: "/personal/info.html"
            },
            {
                name: "报表",
                href: "dashboard/personal.html"
            }
        ]
    }
};

ReactDOM.render(
    <CWNavbar data={navbarData}/>,
    document.getElementById('nav')
);