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
        const [xhr, formData, fileInstance] = [new XMLHttpRequest(), new FormData(), this.state.fileInstance];
        formData.append("source", fileInstance);
        xhr.open("POST", `/gateway/docker-me/v1/updateSource`, true);
        xhr.send(formData);
        xhr.onreadystatechange = () => {
            if (4 === xhr.readyState) {
                if (200 === xhr.status) {
                    try {
                        const result = JSON.parse(xhr.responseText);
                        if (200 !== result.code) {
                           // TODO
                           console.log("load fail ... return code isn't 200 ");
                            return ;
                        }
                    } catch(err) {
                        console.log("load fail ... find some error");
                    }
                }
                console.error("load fail ... xhr status isn't 200");
            }
        };
        // TODO 只有上传成功之后才能进入下一步，这里想忽略掉
        this.props.nextStep({msg: {fileName: fileInstance.name}, step: 0});
    }

    allowNext(event) {
        event.preventDefault();
        const [file] = event.target.files;
        this.setState({allowNextStep: true, fileInstance: file, word: file.name});
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
        const {fileInstance} = this.state;
        return (
           <div className={"fileInstead"}>
               <div class="form-group row">
                   <label class="col-sm-3 col-form-label">文件名</label>
                   <div class="col-sm-9">
                       <input type="text" readonly class="form-control-plaintext" value={fileInstance.name}/>
                   </div>
               </div>
               <div class="form-group row">
                   <label class="col-sm-3 col-form-label">大小</label>
                   <div class="col-sm-9">
                       <input type="text" readonly class="form-control-plaintext"
                              value={ this.formatSize(fileInstance.size) } />
                   </div>
               </div>
               <div class="form-group row">
                   <label class="col-sm-3 col-form-label">类型</label>
                   <div class="col-sm-9">
                       <input type="text" readonly class="form-control-plaintext" value={fileInstance.type} />
                   </div>
               </div>
               <div class="form-group row">
                   <label class="col-sm-3 col-form-label">上次修改时间</label>
                   <div class="col-sm-9">
                       <input type="text" readonly class="form-control-plaintext"
                              value={this.formatTime(fileInstance.lastModifiedDate)} />
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

class CreateImageApp extends React.Component {
    constructor(props) {
        super(props);
        this.nextStep = this.nextStep.bind(this);
        this.state = {
            viewSteps: {
                steps: ["上传镜像压缩包", "构建镜像压缩包", "保存镜像信息"],
                action: 0
            }
        };
    }

    nextStep({msg = {}, step = 0}) {
        const {viewSteps} = this.state;
        const [maxSteps = 0, nextStepAction] = [viewSteps.steps.length, step + 1];
        viewSteps.action = nextStepAction >= maxSteps ? maxSteps: nextStepAction;
        this.setState(viewSteps);
    }

    render() {
        const {viewSteps = {}} = this.state;
        return (
            <div className="createView">
                <ViewStep data={viewSteps}/>
                <span className={"cwWhete"}></span>
                <UploadTarApp nextStep={this.nextStep} />
            </div>
        )
    }
}

ReactDOM.render( <CreateImageApp />, document.getElementById("root"));

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