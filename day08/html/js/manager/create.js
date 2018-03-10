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
        this.state = {
            allowNextStep: false,
        };
        this.uploadFile = this.uploadFile.bind(this);
        this.allowNext = this.allowNext.bind(this);
        this.buildFileInstead = this.buildFileInstead.bind(this);
    }

    uploadFile(event) {
        event.preventDefault();
    }

    allowNext(event) {
        event.preventDefault();
        const [file] = event.target.files;
        console.log(file);
        this.setState({allowNextStep: true, fileInstead: file, word: file.name});
    }

    formatTime(d = new Date()) {
        return [ d.getFullYear(), d.getMonth() + 1, d.getDate() ].join('-')
            + ' ' + [ d.getHours(), d.getMinutes(), d.getSeconds() ].join(':');
    }

    formatSize(d = 0) {
        d = d / 1000;
        return d > 1000 ? (d / 1000).toFixed(2) + " Mb": d.toFixed(3) + "Kb";
    }

    buildFileInstead() {
        const {fileInstead} = this.state;
        return (
           <div className={"fileInstead"}>
               <div class="form-group row">
                   <label class="col-sm-3 col-form-label">文件名</label>
                   <div class="col-sm-9">
                       <input type="text" readonly class="form-control-plaintext" value={fileInstead.name}/>
                   </div>
               </div>
               <div class="form-group row">
                   <label class="col-sm-3 col-form-label">大小</label>
                   <div class="col-sm-9">
                       <input type="text" readonly class="form-control-plaintext"
                              value={ this.formatSize(fileInstead.size) } />
                   </div>
               </div>
               <div class="form-group row">
                   <label class="col-sm-3 col-form-label">类型</label>
                   <div class="col-sm-9">
                       <input type="text" readonly class="form-control-plaintext" value={fileInstead.type} />
                   </div>
               </div>
               <div class="form-group row">
                   <label class="col-sm-3 col-form-label">上次修改时间</label>
                   <div class="col-sm-9">
                       <input type="text" readonly class="form-control-plaintext"
                              value={this.formatTime(fileInstead.lastModifiedDate)} />
                   </div>
               </div>
           </div>
       )
    }

    render() {
        const {allowNextStep = false, word = "请选择资源包"} = this.state;
        return (
            <div className={"uploadTar"}>
                <form onSubmit={this.uploadFile}>
                    <div class="custom-file">
                        <input type="file" class="custom-file-input" onChange={this.allowNext} />
                        <label class="custom-file-label" for="customFile">{word}</label>
                    </div>
                    <small class="text-muted">
                        请上传使用tar.gz 格式压缩的tar包，否则将无法解压。tar包中请包含 dockerfile
                    </small>
                    {allowNextStep? this.buildFileInstead(): ""}
                    <span className={"cwWhete"}></span>
                    <button type="submit" class= {
                        allowNextStep ? "btn btn-primary float-right": "btn btn-primary float-right disabled"
                    }>下一步</button>
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