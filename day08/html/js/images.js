ReactDOM.render(
    <div>
        <div className={"buttonList"}>
            <button type="button" class="btn btn-primary">构建</button>
            <button type="button" class="btn btn-secondary">上传</button>
            <button type="button" class="btn btn-success">管理</button>
        </div>
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