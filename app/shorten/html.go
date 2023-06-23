/*
 *  Copyright (c) 2020-2023 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a GPL-3.0 license that can be found in the LICENSE file.
 */

package shorten

const page404HTML = `
<!DOCTYPE html>
<html>
<head>
    <meta charset='utf-8'>
    <meta http-equiv='X-UA-Compatible' content='IE=edge'>
    <meta name='viewport' content='width=device-width, initial-scale=1'>
    <title>UriOne - 404</title>
    <style>
        body {
            margin: 30vh auto 0 auto;
            width: 400px;
            text-align: center;
        }
        .number {
            font-size: 10vh;
            color: rgb(209, 17, 17);
            text-shadow: -1px 0px 2px rgba(150, 150, 150, 1);
        }
    </style>
</head>
<body>
    <h1 class="number">404</h1>
    <h1>Link not found</h1>
</body>
</html>
`
