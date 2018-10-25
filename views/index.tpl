<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>YouGame</title>
    <link href="https://cdn.bootcss.com/semantic-ui/2.2.13/semantic.css" rel="stylesheet">
    <script src="https://cdn.bootcss.com/jquery/3.2.1/jquery.js"></script>
    <script src="https://cdn.bootcss.com/semantic-ui/2.2.13/semantic.js"></script>
</head>
<body>

<div>
    {{template "/components/nav.html" .}}
    <div style="background-image: url('/static/img/game_cover1.jpg');width: 100% ; height: 500px ;background-size: cover;padding-top: 100px;padding-left: 70px">
    </div>

    <div class="container" style="padding-top: 60px;min-height: 1024px">
        <section>
            <h2 class="ui header container" style="margin-left: 10%">热门</h2>
            <div class="ui grid container">
                {{range .GameList}}
                    <div class="four wide column">
                        <div class="ui card">
                            <div class="image">
                                <img src="{{.Band.Path}}">
                            </div>
                            <div class="content">
                                <a class="header" href="/game/{{.Id}}">{{.Name}}</a>
                            </div>
                            <div class="extra content">
                                <a>
                                    <i class="user icon"></i>
                                    22 Friends
                                </a>
                            </div>
                        </div>

                    </div>
                {{end}}
            </div>
        </section>
        <section style="margin-top: 60px">
            <h2 class="ui header container">游戏</h2>
            <div class="container">
                <div class="ui pointing secondary menu container">
                    <a class="item" data-tab="first">新发售</a>
                    <a class="item active" data-tab="second">排行榜</a>
                    <a class="item" data-tab="third">预购</a>
                </div>
                <div class="ui tab segment container" data-tab="first">
                    <div class="ui divided items container">
                        <div class="item">
                            <div class="image">
                                <img src="/static/img/game_band1.jpg">
                            </div>
                            <div class="content">
                                <a class="header">12 Years a Slave</a>
                                <div class="meta">
                                    <span class="cinema">Union Square 14</span>
                                </div>
                                <div class="description">
                                    <p></p>
                                </div>
                                <div class="extra">
                                    <div class="ui label">IMAX</div>
                                    <div class="ui label"><i class="globe icon"></i> Additional Languages</div>
                                </div>
                            </div>
                        </div>
                        <div class="item">
                            <div class="image">
                                <img src="/static/img/game_band1.jpg">
                            </div>
                            <div class="content">
                                <a class="header">My Neighbor Totoro</a>
                                <div class="meta">
                                    <span class="cinema">IFC Cinema</span>
                                </div>
                                <div class="description">
                                    <p></p>
                                </div>
                                <div class="extra">
                                    <div class="ui right floated primary button">
                                        Buy tickets
                                        <i class="right chevron icon"></i>
                                    </div>
                                    <div class="ui label">Limited</div>
                                </div>
                            </div>
                        </div>
                        <div class="item">
                            <div class="image">
                                <img src="/static/img/game_band1.jpg">
                            </div>
                            <div class="content">
                                <a class="header">Watchmen</a>
                                <div class="meta">
                                    <span class="cinema">IFC</span>
                                </div>
                                <div class="description">
                                    <p></p>
                                </div>
                                <div class="extra">
                                    <div class="ui right floated primary button">
                                        Buy tickets
                                        <i class="right chevron icon"></i>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="ui tab segment active container" data-tab="second">
                    <div class="ui divided items container">
                        <div class="item">
                            <div class="image">
                                <img src="/static/img/game_band1.jpg">
                            </div>
                            <div class="content">
                                <a class="header">12 Years a Slave</a>
                                <div class="meta">
                                    <span class="cinema">Union Square 14</span>
                                </div>
                                <div class="description">
                                    <p></p>
                                </div>
                                <div class="extra">
                                    <div class="ui label">IMAX</div>
                                    <div class="ui label"><i class="globe icon"></i> Additional Languages</div>
                                </div>
                            </div>
                        </div>
                        <div class="item">
                            <div class="image">
                                <img src="/static/img/game_band1.jpg">
                            </div>
                            <div class="content">
                                <a class="header">My Neighbor Totoro</a>
                                <div class="meta">
                                    <span class="cinema">IFC Cinema</span>
                                </div>
                                <div class="description">
                                    <p></p>
                                </div>
                                <div class="extra">
                                    <div class="ui right floated primary button">
                                        Buy tickets
                                        <i class="right chevron icon"></i>
                                    </div>
                                    <div class="ui label">Limited</div>
                                </div>
                            </div>
                        </div>
                        <div class="item">
                            <div class="image">
                                <img src="/static/img/game_band1.jpg">
                            </div>
                            <div class="content">
                                <a class="header">Watchmen</a>
                                <div class="meta">
                                    <span class="cinema">IFC</span>
                                </div>
                                <div class="description">
                                    <p></p>
                                </div>
                                <div class="extra">
                                    <div class="ui right floated primary button">
                                        Buy tickets
                                        <i class="right chevron icon"></i>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="ui tab segment container" data-tab="third">
                    <div class="ui divided items container">
                        <div class="item">
                            <div class="image">
                                <img src="/static/img/game_band1.jpg">
                            </div>
                            <div class="content">
                                <a class="header">12 Years a Slave</a>
                                <div class="meta">
                                    <span class="cinema">Union Square 14</span>
                                </div>
                                <div class="description">
                                    <p></p>
                                </div>
                                <div class="extra">
                                    <div class="ui label">IMAX</div>
                                    <div class="ui label"><i class="globe icon"></i> Additional Languages</div>
                                </div>
                            </div>
                        </div>
                        <div class="item">
                            <div class="image">
                                <img src="/static/img/game_band1.jpg">
                            </div>
                            <div class="content">
                                <a class="header">My Neighbor Totoro</a>
                                <div class="meta">
                                    <span class="cinema">IFC Cinema</span>
                                </div>
                                <div class="description">
                                    <p></p>
                                </div>
                                <div class="extra">
                                    <div class="ui right floated primary button">
                                        Buy tickets
                                        <i class="right chevron icon"></i>
                                    </div>
                                    <div class="ui label">Limited</div>
                                </div>
                            </div>
                        </div>
                        <div class="item">
                            <div class="image">
                                <img src="/static/img/game_band1.jpg">
                            </div>
                            <div class="content">
                                <a class="header">Watchmen</a>
                                <div class="meta">
                                    <span class="cinema">IFC</span>
                                </div>
                                <div class="description">
                                    <p></p>
                                </div>
                                <div class="extra">
                                    <div class="ui right floated primary button">
                                        Buy tickets
                                        <i class="right chevron icon"></i>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </section>
    </div>
    <div class="ui inverted vertical footer segment">
        <div class="ui container">
            <div class="ui stackable inverted divided equal height stackable grid">
                <div class="three wide column">
                    <h4 class="ui inverted header">About</h4>
                    <div class="ui inverted link list">
                        <a href="#" class="item">Sitemap</a>
                        <a href="#" class="item">Contact Us</a>
                        <a href="#" class="item">Religious Ceremonies</a>
                        <a href="#" class="item">Gazebo Plans</a>
                    </div>
                </div>
                <div class="three wide column">
                    <h4 class="ui inverted header">Services</h4>
                    <div class="ui inverted link list">
                        <a href="#" class="item">Banana Pre-Order</a>
                        <a href="#" class="item">DNA FAQ</a>
                        <a href="#" class="item">How To Access</a>
                        <a href="#" class="item">Favorite X-Men</a>
                    </div>
                </div>
                <div class="seven wide column">
                    <h4 class="ui inverted header">Footer Header</h4>
                    <p>Extra space for a call to action inside the footer that could help re-engage users.</p>
                </div>
            </div>
        </div>
    </div>
</div>
<script>
    $(document)
        .ready(function () {
            $('.ui.menu .ui.dropdown').dropdown({
                on: 'hover'
            });
            $('.ui.menu a.item')
                .on('click', function () {
                    $(this)
                        .addClass('active')
                        .siblings()
                        .removeClass('active')
                    ;
                })
            ;
        })
    ;
    $('.menu .item')
        .tab()
    ;
</script>
</body>
</html>