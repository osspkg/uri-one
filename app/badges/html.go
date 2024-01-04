/*
 *  Copyright (c) 2020-2024 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
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
    <title>Badge</title>
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
<code>%s/badge/[color]/[title]/[data]/image.svg</code>
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
<pre>&lt;img src=&quot;%s/badge/light/User ID/12/image.svg&quot;&gt;</pre>
<p>
	<img src="/badge/primary/primary/color/image.svg">
	<img src="/badge/secondary/secondary/color/image.svg">
	<img src="/badge/success/success/color/image.svg">
	<img src="/badge/danger/danger/color/image.svg">
	<img src="/badge/warning/warning/color/image.svg">
	<img src="/badge/info/info/color/image.svg">
	<img src="/badge/light/light/color/image.svg">
</p>

</body>
</html>
`
