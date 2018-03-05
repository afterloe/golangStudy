class CWNavbarStatus extends React.Component {
    constructor(props) {
        super(props)
    }

    render() {
        return (
            <div class="collapse navbar-collapse">
                <div class="d-flex">
                    <ul class="navbar-nav">
                        <li class="dropdown js-menu-container">
                    <span class="d-inline-block  px-2">
                        <a href="/notifications" aria-label="You have no unread notifications"
                           class="notification-indicator tooltipped tooltipped-s  js-socket-channel js-notification-indicator"
                           data-channel="notification-changed:11550317" data-ga-click="Header, go to notifications, icon:read"
                           data-hotkey="g n">
                            <span class="mail-status "></span>
                            <svg aria-hidden="true" class="octicon octicon-bell" height="16" version="1.1" viewBox="0 0 14 16"
                                 width="14">
                                <path fill-rule="evenodd"
                                      d="M14 12v1H0v-1l.73-.58c.77-.77.81-2.55 1.19-4.42C2.69 3.23 6 2 6 2c0-.55.45-1 1-1s1 .45 1 1c0 0 3.39 1.23 4.16 5 .38 1.88.42 3.66 1.19 4.42l.66.58H14zm-7 4c1.11 0 2-.89 2-2H5c0 1.11.89 2 2 2z"></path>
                            </svg>
                        </a>
                    </span>
                        </li>
                        <li class="dropdown js-menu-container">
                            <details class="dropdown-details details-reset js-dropdown-details d-flex px-2 flex-items-center">
                                <summary class="HeaderNavlink" aria-label="Create new…" data-ga-click="Header, create new, icon:add">
                                    <svg aria-hidden="true" class="octicon octicon-plus float-left mr-1 mt-1" height="16" version="1.1" viewBox="0 0 12 16" width="12"><path fill-rule="evenodd" d="M12 9H7v5H5V9H0V7h5V2h2v5h5z"></path></svg>
                                    <span class="dropdown-caret mt-1"></span>
                                </summary>
                                <ul class="dropdown-menu dropdown-menu-sw">
                                    <a class="dropdown-item" href="/new" data-ga-click="Header, create new repository">
                                        New repository
                                    </a>
                                    <a class="dropdown-item" href="/new/import" data-ga-click="Header, import a repository">
                                        Import repository
                                    </a>
                                    <a class="dropdown-item" href="https://gist.github.com/" data-ga-click="Header, create new gist">
                                        New gist
                                    </a>
                                    <a class="dropdown-item" href="/organizations/new" data-ga-click="Header, create new organization">
                                        New organization
                                    </a>
                                </ul>
                            </details>
                        </li>
                        <li class="dropdown js-menu-container">
                            <details class="dropdown-details details-reset js-dropdown-details d-flex pl-2 flex-items-center">
                                <summary class="HeaderNavlink name mt-1" aria-label="View profile and more" data-ga-click="Header, show menu, icon:avatar">
                                    <img alt="@afterloe" class="avatar float-left mr-1" src="https://avatars0.githubusercontent.com/u/11550317?s=40&amp;v=4" height="20" width="20" />
                                    <span class="dropdown-caret"></span>
                                </summary>
                                <ul class="dropdown-menu dropdown-menu-sw">
                                    <li class="dropdown-header header-nav-current-user css-truncate">
                                        Signed in as <strong class="css-truncate-target">afterloe</strong>
                                    </li>
                                    <li class="dropdown-divider"></li>
                                    <li><a class="dropdown-item" href="/afterloe" data-ga-click="Header, go to profile, text:your profile">
                                        Your profile
                                    </a></li>
                                    <li><a class="dropdown-item" href="/afterloe?tab=stars" data-ga-click="Header, go to starred repos, text:your stars">
                                        Your stars
                                    </a></li>
                                    <li><a class="dropdown-item" href="https://gist.github.com/" data-ga-click="Header, your gists, text:your gists">Your gists</a></li>
                                    <li class="dropdown-divider"></li>
                                    <li><a class="dropdown-item" href="https://help.github.com" data-ga-click="Header, go to help, text:help">
                                        Help
                                    </a></li>
                                    <li><a class="dropdown-item" href="/settings/profile" data-ga-click="Header, go to settings, icon:settings">
                                        Settings
                                    </a></li>
                                    <li>
                                        <form accept-charset="UTF-8" action="/logout" class="logout-form" method="post">
                                            <div style={{ margin:0, padding:0, display:"inline" }}>
                                                <input name="utf8" type="hidden" value="✓"/>
                                                <input name="authenticity_token" type="hidden" value="qyfbCm6oUdIa6sX53DEyUQdNwFDZb0vQl9LgvL25GO6KlbPLbaQUr2LFCZMxAs3dTYVA/jvmAwcsHbcUGISJlQ=="/>
                                            </div>
                                            <button type="submit" class="dropdown-item dropdown-signout" data-ga-click="Header, sign out, icon:logout">
                                                Sign out
                                            </button>
                                        </form>
                                    </li>
                                </ul>
                            </details>
                        </li>
                    </ul>
                </div>
            </div>
        )
    }
}

class CWNavbarRouters extends React.Component {
    constructor(props) {
        super(props);
    }

    renderRouters(index = 0, routeSlice = []) {
        return routeSlice.map((route, i) => index === i ? (
            <li className={"nav-item active"}>
                <a className={"nav-link"} href={route["href"]}>{route["name"]}</a>
            </li>
        ) : (
            <li className={"nav-item"}>
                <a className={"nav-link"} href={route["href"]}>{route["name"]}</a>
            </li>
        ));
    }

    render() {
        const {activeRouter, routers} = this.props.data || {};
        return (
            <div className={"collapse navbar-collapse"}>
                <ul className={"navbar-nav mr-auto"}>
                    {this.renderRouters(activeRouter, routers)}
                </ul>
            </div>
        )
    }
}

class CWNavbarInputForm extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
        const {word = "Search"} = this.props.data || {};
        return (
            <form className={"form-inline my-2 my-lg-0"}>
                <input className={"form-control mr-sm-2"} type={"search"} placeholder={word}
                       aria-label={word}/>
            </form>
        )
    }
}

class CWNavbar extends React.Component {

    constructor(props) {
        super(props);
    }

    render() {
        const {linkedHref = "/", name = "Cityworks™", cWNavbarInputForm = {},
            cWNavbarRouters = {}} = this.props.data || {};
        return (
            <nav className={"navbar navbar-expand-lg navbar-light"}>
                <a className={"navbar-brand"} href={linkedHref}>{name}</a>
                <CWNavbarInputForm data={cWNavbarInputForm}/>
                <CWNavbarRouters data={cWNavbarRouters}/>
                <CWNavbarStatus />
            </nav>
        )
    }
}

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