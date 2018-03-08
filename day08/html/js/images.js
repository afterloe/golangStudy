ReactDOM.render(
    <div>
        <div className={"buttonList"}>
            <button type="button" class="btn btn-primary">构建</button>
            <button type="button" class="btn btn-secondary">上传</button>
            <button type="button" class="btn btn-dark">管理</button>
        </div>
        <span class="cwWhete"></span>
        <div>
            <div class="card" style={{width: "200px"}}>
                <img class="card-img-top" src="..." alt="Card image cap" />
                <div class="card-body">
                    <h5 class="card-title">Card title</h5>
                    <p class="card-text">Some quick example text to build on the card title and make up the bulk of the card's content.</p>
                    <a href="#" class="btn btn-primary">Go somewhere</a>
                </div>
            </div>
        </div>
        <span class="cwWhete"></span>
    </div>,
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
    }
};

ReactDOM.render(
    <CWNavbar data={navbarData}/>,
    document.getElementById('nav')
);