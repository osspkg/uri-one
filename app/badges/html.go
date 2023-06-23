/*
 *  Copyright (c) 2020-2023 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a GPL-3.0 license that can be found in the LICENSE file.
 */

package badges

const indexHTML = `
<!DOCTYPE html>
<html>
<head>
    <meta charset='utf-8'>
    <meta http-equiv='X-UA-Compatible' content='IE=edge'>
    <meta name='viewport' content='width=device-width, initial-scale=1'>
    <title>UriOne - Badge</title>
    <style>
        body {
            margin: 30vh auto 0 auto;
            width: 450px;
            text-align: center;
			background-color: darkgrey;
        }
    </style>
</head>
<body>

<h1>Badge Generation</h1>
<code>https://uri.one/badge/[color]/[title]/[data]/image.svg</code>
<p>Color: 
	<span style="color: #0d6efd">primary</span>,
    <span style="color: #6c757d">secondary</span>,
    <span style="color: #198754">success</span>,
    <span style="color: #dc3545">danger</span>,
    <span style="color: #ffc107">warning</span>,
    <span style="color: #0dcaf0">info</span>,
    <span style="color: #f8f9fa">light</span>
</p>
<br><br>
<h1>Example</h1>
<pre>&lt;img src=&quot;https://uri.one/badge/light/User ID/12/image.svg&quot;&gt;</pre>
<p>
	<img src="/badge/primary/User ID/12/image.svg">
	<img src="/badge/secondary/User ID/12/image.svg">
	<img src="/badge/success/User ID/12/image.svg">
	<img src="/badge/danger/User ID/12/image.svg">
	<img src="/badge/warning/User ID/12/image.svg">
	<img src="/badge/info/User ID/12/image.svg">
	<img src="/badge/light/User ID/12/image.svg">
</p>

</body>
</html>
`
