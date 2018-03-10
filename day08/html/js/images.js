class ButtonManager extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
        return (
            <div className={"buttonList"}>
                <a class="btn btn-primary" href={"/manager/create.html"}>构建</a>
                <a class="btn btn-secondary" href={"/manager/update"}>上传</a>
                <a class="btn btn-dark" href={"#"}>管理</a>
            </div>
        )
    }
}

class ImageItem extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
        return (
            <div class="card">
                <img class="card-img-top" src="../images/docker_image.jpg" alt="Card image cap" />
                <div class="card-body">
                    <h5 class="card-title">Card title</h5>
                    <p class="card-text">Some quick example text to build on the card title and make up the bulk of the card's content.</p>
                    <a href="#" class="btn btn-primary">Go somewhere</a>
                    <div class="btn-group">
                        <button type="button" className="btn btn-dark dropdown-toggle">
                            Run With &nbsp;&nbsp;
                        </button>
                        <div class="dropdown-menu">
                            <a class="dropdown-item" href="#">Dropdown link</a>
                            <a class="dropdown-item" href="#">Dropdown link</a>
                        </div>
                    </div>
                </div>
                <div class="card-footer">
                    <small class="text-muted">Last updated 3 mins ago</small>
                </div>
            </div>
        )
    }
}

class ImagesViews extends React.Component {
    constructor(props) {
        super(props);
    }

    buildItemRanks(imageItems = []) {
        return (
            <div className={"card-deck"}>
                {imageItems.map(item => <ImageItem />)}
            </div>
        )
    }

    render() {
        const {imagesList = []} = this.props.data || {};
        const ranks = [];
        for (var i = 0, len = imagesList.length; i < len; i+=3) {
            ranks.push(imagesList.slice(i, i + 3));
        }
        return (
            <div className={"imagesView"}>
                {ranks.map(rank => this.buildItemRanks(rank))}
            </div>
        )
    }
}

class ImagesManager extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
        return (
            <div className={"imagesView"}>
                <ButtonManager />
                <span className={"cwWhete"}></span>
                <ImagesViews data={{imagesList: [1,2,3,4,5,6,7,8,9,10,11,12,13]}}/>
                <span className={"cwWhete"}></span>
                <span className={"cwWhete"}></span>
                <nav aria-label="Page navigation example">
                    <ul class="pagination justify-content-end">
                        <li class="page-item disabled">
                            <a class="page-link" href="#" tabindex="-1">Previous</a>
                        </li>
                        <li class="page-item"><a class="page-link" href="#">1</a></li>
                        <li class="page-item"><a class="page-link" href="#">2</a></li>
                        <li class="page-item"><a class="page-link" href="#">3</a></li>
                        <li class="page-item">
                            <a class="page-link" href="#">Next</a>
                        </li>
                    </ul>
                </nav>
                <span class="cwWhete"></span>
            </div>
        )
    }
}

ReactDOM.render(
    <ImagesManager />,
    document.getElementById('root')
);

const navbarData = {
    linkedHref: "/",
    name: "Cityworks™ 云平台",
    cWNavbarInputForm: {
        word: "搜索获取容器"
    },
    cWNavbarRouters: {
        activeRouter: 0,
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