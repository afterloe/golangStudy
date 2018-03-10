class ViewStep extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
        const {steps = [], action = 0} = this.props.data;
        return (
            <nav class="nav">
                {
                    steps.map((step, index) =>
                        <a class= {index === action ? "nav-link active": "nav-link disabled"} href="#">
                            <span class="badge badge-danger">{ index + 1 }</span>
                            <span>{ step || "unKnow step" }</span>
                            </a>
                    )
                }
            </nav>
        )
    }
}

class UploadTarApp extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
        return (
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
        )
    }
}

ReactDOM.render(
    <div className="createView">
        <ViewStep data={{steps: ["上传镜像压缩包", "构建镜像压缩包", "保存镜像信息"], action: 0}}/>
        <span className={"cwWhete"}></span>
        <UploadTarApp />
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