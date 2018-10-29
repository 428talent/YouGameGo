<!DOCTYPE html>
<html>
<head>
    <!-- Standard Meta -->
    <meta charset="utf-8"/>
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1"/>
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0">

    <!-- Site Properties -->
    <title>登录</title>

    <link href="https://cdn.bootcss.com/semantic-ui/2.2.13/semantic.css" rel="stylesheet">
    <script src="https://cdn.bootcss.com/jquery/3.2.1/jquery.js"></script>
    <script src="https://cdn.bootcss.com/semantic-ui/2.2.13/semantic.js"></script>

    <style type="text/css">
        body {

        }

        body > .grid {
            height: 100%;
        }

        .image {
            margin-top: -100px;
        }

        .column {
            max-width: 450px;
        }
    </style>
    <script>
        $(document)
            .ready(function () {
                $('.ui.form')
                    .form({
                        fields: {
                            email: {
                                identifier: 'email',
                                rules: [
                                    {
                                        type: 'empty',
                                        prompt: 'Please enter your e-mail'
                                    },
                                    {
                                        type: 'email',
                                        prompt: 'Please enter a valid e-mail'
                                    }
                                ]
                            },
                            password: {
                                identifier: 'password',
                                rules: [
                                    {
                                        type: 'empty',
                                        prompt: 'Please enter your password'
                                    },
                                    {
                                        type: 'length[6]',
                                        prompt: 'Your password must be at least 6 characters'
                                    }
                                ]
                            }
                        }
                    })
                ;
            })
        ;
    </script>
</head>
<body>
{{template "/components/nav.html" .}}
<div class="ui middle aligned center aligned grid"
     style='background: url("/static/img/login-bg.png") no-repeat;background-size: cover'>
    <div class="column">
        <form class="ui large form" method="post" action="/login/auth">
            <div class="ui stacked segment">
                <h2 class="ui teal image header">
                    {{/*<img src="assets/images/logo.png" class="image">*/}}
                    <div class="content">
                        登录你的账户
                    </div>
                </h2>
                <div class="field">
                    <div class="ui left icon input">
                        <i class="user icon"></i>
                        <input type="text" name="username" placeholder="邮箱地址">
                    </div>
                </div>
                <div class="field">
                    <div class="ui left icon input">
                        <i class="lock icon"></i>
                        <input type="password" name="password" placeholder="密码">
                    </div>
                </div>
                {{if .flash.ErrorTitle}}
                    <div class="ui negative message">
                        <div class="header">
                            {{.flash.ErrorTitle}}
                        </div>
                        <p>{{.flash.ErrorContent}}
                        </p></div>
                {{end}}
                <button class="ui fluid large teal submit button" type="submit">登录</button>
            </div>

            <div class="ui error message"></div>

        </form>

    </div>
</div>
{{template "/components/footer.html" .}}

</body>

</html>