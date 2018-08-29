<%@ page language="java" contentType="text/html; charset=utf-8"
	pageEncoding="utf-8"%>
<!DOCTYPE html>
<html>
<head>
<title>chattingAPI</title>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<style type="text/css">
html, body, div, span, applet, object, iframe, h1, h2, h3, h4, h5, h6, p,
	blockquote, pre, a, abbr, acronym, address, big, cite, code, del, dfn,
	em, img, ins, kbd, q, s, samp, small, strike, strong, sub, sup, tt, var,
	b, u, i, center, dl, dt, dd, ol, ul, li, fieldset, form, label, legend,
	table, caption, tbody, tfoot, thead, tr, th, td, article, aside, canvas,
	details, embed, figure, figcaption, footer, header, hgroup, menu, nav,
	output, ruby, section, summary, time, mark, audio, video {
	margin: 0;
	padding: 0;
	border: 0;
}

/* BODY

=============================================================================*/
body {
	font-family: Helvetica, arial, freesans, clean, sans-serif;
	font-size: 14px;
	line-height: 1.6;
	color: #333;
	background-color: #fff;
	padding: 20px;
	max-width: 960px;
	margin: 0 auto;
}

body>*:first-child {
	margin-top: 0 !important;
}

body>*:last-child {
	margin-bottom: 0 !important;
}

/* BLOCKS

=============================================================================*/
p, blockquote, ul, ol, dl, table, pre {
	margin: 15px 0;
}

/* HEADERS

=============================================================================*/
h1, h2, h3, h4, h5, h6 {
	margin: 20px 0 10px;
	padding: 0;
	font-weight: bold;
	-webkit-font-smoothing: antialiased;
}

h1 tt, h1 code, h2 tt, h2 code, h3 tt, h3 code, h4 tt, h4 code, h5 tt,
	h5 code, h6 tt, h6 code {
	font-size: inherit;
}

h1 {
	font-size: 28px;
	color: #000;
}

h2 {
	font-size: 24px;
	border-bottom: 1px solid #ccc;
	color: #000;
}

h3 {
	font-size: 18px;
}

h4 {
	font-size: 16px;
}

h5 {
	font-size: 14px;
}

h6 {
	color: #777;
	font-size: 14px;
}

body>h2:first-child, body>h1:first-child, body>h1:first-child+h2, body>h3:first-child,
	body>h4:first-child, body>h5:first-child, body>h6:first-child {
	margin-top: 0;
	padding-top: 0;
}

a:first-child h1, a:first-child h2, a:first-child h3, a:first-child h4,
	a:first-child h5, a:first-child h6 {
	margin-top: 0;
	padding-top: 0;
}

h1+p, h2+p, h3+p, h4+p, h5+p, h6+p {
	margin-top: 10px;
}

/* LINKS

=============================================================================*/
a {
	color: #4183C4;
	text-decoration: none;
}

a:hover {
	text-decoration: underline;
}

/* LISTS

=============================================================================*/
ul, ol {
	padding-left: 30px;
}

ul li>:first-child, ol li>:first-child, ul li ul:first-of-type, ol li ol:first-of-type,
	ul li ol:first-of-type, ol li ul:first-of-type {
	margin-top: 0px;
}

ul ul, ul ol, ol ol, ol ul {
	margin-bottom: 0;
}

dl {
	padding: 0;
}

dl dt {
	font-size: 14px;
	font-weight: bold;
	font-style: italic;
	padding: 0;
	margin: 15px 0 5px;
}

dl dt:first-child {
	padding: 0;
}

dl dt>:first-child {
	margin-top: 0px;
}

dl dt>:last-child {
	margin-bottom: 0px;
}

dl dd {
	margin: 0 0 15px;
	padding: 0 15px;
}

dl dd>:first-child {
	margin-top: 0px;
}

dl dd>:last-child {
	margin-bottom: 0px;
}

/* CODE

=============================================================================*/
pre, code, tt {
	font-size: 12px;
	font-family: Consolas, "Liberation Mono", Courier, monospace;
}

code, tt {
	margin: 0 0px;
	padding: 0px 0px;
	white-space: nowrap;
	border: 1px solid #eaeaea;
	background-color: #f8f8f8;
	border-radius: 3px;
}

pre>code {
	margin: 0;
	padding: 0;
	white-space: pre;
	border: none;
	background: transparent;
}

pre {
	background-color: #f8f8f8;
	border: 1px solid #ccc;
	font-size: 13px;
	line-height: 19px;
	overflow: auto;
	padding: 6px 10px;
	border-radius: 3px;
}

pre code, pre tt {
	background-color: transparent;
	border: none;
}

kbd {
	-moz-border-bottom-colors: none;
	-moz-border-left-colors: none;
	-moz-border-right-colors: none;
	-moz-border-top-colors: none;
	background-color: #DDDDDD;
	background-image: linear-gradient(#F1F1F1, #DDDDDD);
	background-repeat: repeat-x;
	border-color: #DDDDDD #CCCCCC #CCCCCC #DDDDDD;
	border-image: none;
	border-radius: 2px 2px 2px 2px;
	border-style: solid;
	border-width: 1px;
	font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
	line-height: 10px;
	padding: 1px 4px;
}

/* QUOTES

=============================================================================*/
blockquote {
	border-left: 4px solid #DDD;
	padding: 0 15px;
	color: #777;
}

blockquote>:first-child {
	margin-top: 0px;
}

blockquote>:last-child {
	margin-bottom: 0px;
}

/* HORIZONTAL RULES

=============================================================================*/
hr {
	clear: both;
	margin: 15px 0;
	height: 0px;
	overflow: hidden;
	border: none;
	background: transparent;
	border-bottom: 4px solid #ddd;
	padding: 0;
}

/* TABLES

=============================================================================*/
table {
	margin: 10px 0 15px 0;
	border-collapse: collapse;
}

td, th {
	border: 1px solid #000;
	padding: 3px 10px;
}

th {
	padding: 5px 10px;
}

/* IMAGES

=================================================================*/
img {
	max-width: 100%
}
</style>
</head>
<body>
	<h1>
		<span id="0"> chatting项目API接口文档 </span> <br />
	</h1>
	<h4>文档结构：</h4>
	<p>
		1.<a href="#1"> 文档规范 </a> <br /> 2.<a href="#2"> 文档变更记录 </a> <br /> 3.<a
			href="#3"> API的Signature </a> <br /> &ensp;&ensp; 3.1.<a href="#3.1">
			测试版Header的parameter </a> <br /> &ensp;&ensp; 3.2.<a href="#3.2">
			正式版Header的parameter </a> <br /> &ensp;&ensp; 3.3.<a href="#3.3">
			Signature有效期 </a> <br /> &ensp;&ensp; 3.4.<a href="#3.4">
			计算并获取Signature </a> <br /> &ensp;&ensp; 3.5.<a href="#3.5">
			进行Base64加密 </a> <br /> 4.<a href="#4"> 模块序列号与返回状态码说明 </a> <br />
		&ensp;&ensp; 4.1.<a href="#4.1"> 返回状态码说明 </a> <br /> &ensp;&ensp; 4.2.<a
			href="#4.2"> 客户端错误状态码 </a> <br /> &ensp;&ensp; 4.3.<a href="#4.3">
			服务端及第三方错误状态码 </a> <br /> 5.<a href="#5">URL接口说明</a> <br /> &ensp;&ensp;
		5.1.<a href="#5.1"> 发送验证码 </a> <br /> &ensp;&ensp; 5.2.<a href="#5.2">
			校验验证码 </a> <br /> &ensp;&ensp; 5.3.<a href="#5.3"> 注册 </a> <br />
		&ensp;&ensp; 5.4.<a href="#5.4"> 修改个人信息 </a> <br /> &ensp;&ensp; 5.5.<a
			href="#5.5"> 修改密码 </a> <br /> &ensp;&ensp; 5.6.<a href="#5.6">
			实名认证 </a> <br /> &ensp;&ensp; 5.7.<a href="#5.7"> 找回密码 </a> <br />
		&ensp;&ensp; 5.8.<a href="#5.8"> 登录 </a> <br /> &ensp;&ensp; 5.9.<a
			href="#5.9"> 退出 </a> <br /> &ensp;&ensp; 5.10.<a href="#5.10">
			获取联系人 </a> <br /> &ensp;&ensp; 5.11.<a href="#5.11"> 备份联系人 </a> <br />
		&ensp;&ensp; 5.12.<a href="#5.12"> 发送私信 </a> <br /> &ensp;&ensp; 5.13.<a
			href="#5.13"> 接收指定联系人私信 </a> <br /> &ensp;&ensp; 5.14.<a href="#5.14">
			升级或更新群 </a> <br /> &ensp;&ensp; 5.15.<a href="#5.15"> 踢人出群 </a> <br />
		&ensp;&ensp; 5.16.<a href="#5.16"> 解散群 </a> <br /> &ensp;&ensp; 5.17.<a
			href="#5.17"> 编辑群资料 </a> <br /> &ensp;&ensp; 5.18.<a href="#5.18">
			主动退群 </a> <br /> &ensp;&ensp; 5.19.<a href="#5.19"> 获取全部群资料 </a> <br />
		&ensp;&ensp; 5.20.<a href="#5.20"> 上传群聊记录 </a> <br /> &ensp;&ensp;
		5.21.<a href="#5.21"> 获取群聊记录列表 </a> <br /> &ensp;&ensp; 5.22.<a
			href="#5.22"> 获取群聊记录内容 </a> <br /> &ensp;&ensp; 5.23.<a href="#5.23">
			消息通知 </a> <br /> &ensp;&ensp; 5.24.<a href="#5.24"> 上传文件 </a> <br />
		&ensp;&ensp; 5.25.<a href="#5.25"> 下载文件 </a> <br /> &ensp;&ensp; 5.26.<a
			href="#5.26"> 版本更新 </a> <br /> 6.<a href="#6"> 正则文档 </a>
	</p>
	<h2>
		一、<span id="1">文档规范</span>
	</h2>
	<p>
		1.本文所有请求：POST + key/value + UTF-8编码; <br />
		2.Content-Type类型：application/x-www-form-urlencoded;charset=utf-8; <br />
		3.返回值格式：JSON 、UTF-8编码； <br /> 4.以下各个模块所用ip或域名地址由001模块返回； <br />
		5.文档所有时间全部取自从1970年1月1日0点0 分0 秒开始到现在的毫秒值； <br /> 6.参数验证详情参考<strong><a
			href="#6"> 正则文档 </a></strong>； <br /> 7.所有变更都需要填写变更记录； <br />
	</p>
	<h2>
		二、<span id="2">文档变更记录</span>
	</h2>
	<table>
		<thead>
			<tr>
				<th align="center">版本</th>
				<th align="center">服务</th>
				<th align="center">说明</th>
				<th align="center">姓名</th>
				<th align="center">时间</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="center">v1.0.0</td>
				<td align="center">全部</td>
				<td align="center">第一版，初始化</td>
				<td align="center">王东阳</td>
				<td align="center">2017.03.30</td>
			</tr>
			<tr>
				<td align="center">v1.0.1</td>
				<td align="center">联系人、私信、群</td>
				<td align="center">将url符合restful</td>
				<td align="center">王东阳</td>
				<td align="center">2017.04.10</td>
			</tr>
			<tr>
				<td align="center">v1.0.1</td>
				<td align="center">联系人、私信、群</td>
				<td align="center">重新整理各个接口、状态码与返回集</td>
				<td align="center">王东阳</td>
				<td align="center">2017.04.13</td>
			</tr>
			<tr>
				<td align="center">v1.0.1</td>
				<td align="center">联系人、私信、群</td>
				<td align="center">ccat更名为chatting</td>
				<td align="center">王东阳</td>
				<td align="center">2017.04.14</td>
			</tr>
			<tr>
				<td align="center">v1.0.1</td>
				<td align="center">全部</td>
				<td align="center">兼容用户服务</td>
				<td align="center">王东阳</td>
				<td align="center">2017.04.21</td>
			</tr>
			<tr>
				<td align="center">v1.0.2</td>
				<td align="center">修改密码</td>
				<td align="center">添加修改密码功能</td>
				<td align="center">王东阳</td>
				<td align="center">2017.05.03</td>
			</tr>
			<tr>
				<td align="center">v1.0.3</td>
				<td align="center">全部</td>
				<td align="center">整理测试版API接口文档</td>
				<td align="center">王东阳</td>
				<td align="center">2017.05.11</td>
			</tr>
		</tbody>
	</table>
	<h2>
		三、<span id="3">API的Signature</span>
	</h2>
	<h4>
		<span id="3.1">测试版：APP签名，放在Http Request Header中：</span>
	</h4>
	<table>
		<thead>
			<tr>
				<th>参数</th>
				<th align="center">类型</th>
				<th align="center">必须</th>
				<th align="center">参数说明</th>
				<th align="center">参数值</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td>appSignature</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">app端的签名</td>
				<td align="center">ios或android</td>
			</tr>
			<tr>
				<td>version</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">版本号</td>
				<td align="center">v1.0.3</td>
			</tr>
		</tbody>
	</table>
	<h4>
		<span id="3.2">正式版：以下参数需要放在Http Request Header中</span>
	</h4>
	<table>
		<thead>
			<tr>
				<th>参数</th>
				<th align="center">类型</th>
				<th align="center">必须</th>
				<th align="center">参数说明</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td>AppSecret</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">开发者平台分配的AppSecret；AppSecret由用户的id(永久id或临时id)经Base64加密的结果</td>
			</tr>
			<tr>
				<td>CurTime</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">当前UTC时间戳，从1970年1月1日0点0 分0 秒开始到现在的分钟数(String)</td>
			</tr>
			<tr>
				<td>Nonce</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">(用户的IMEI)+(当前UTC时间戳的毫秒值对1000取余）进行MD5加密</td>
			</tr>
			<tr>
				<td>Signature</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">SHA1(AppSecret + Nonce +
					CurTime),三个参数拼接的字符串，进行SHA1哈希计算，转化成16进制字符(String，小写)</td>
			</tr>
			<tr>
				<td>IMEI</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">手机的唯一标识，后由Base64加密（android为串号，IOS为UUID）</td>
			</tr>
		</tbody>
	</table>
	<p>
		<span id="3.3"><strong>Signature有效期:</strong></span>
		出于安全性考虑，每个Signature的有效期为10分钟(用CurTime计算)，建议每次请求都生成新的Signature，同时请确认发起请求的服务器是与标准时间
		同步的，比如有NTP服务。Signature检验失败时会返回HHTP协议错误码(-1)，具体参看code状态表。
	</p>
	<p>
		<span id="3.4"><strong>计算并获取Signature:</strong></span>
	</p>
	<pre>
		<code>public static String getSignature(String appSecret, String nonce, 
    String curTime) throws Exception {
    byte[] appKey = Base64.decryptBASE64(appSecret);
    return encode(&quot;sha1&quot;, new String(appKey) + nonce + curTime);
}

private static String encode(String algorithm, String value) {
    try {
        if (value == null) {
            return null;
        }
        MessageDigest messageDigest = MessageDigest.getInstance(algorithm);
        messageDigest.update(value.getBytes());
        return getFormattedText(messageDigest.digest());
    } catch (Exception e) {
        throw new RuntimeException(e);
    }
}

/**
 * 获取十六进制字符串形式的MD5摘要
 * @throws NoSuchAlgorithmException
 */
public synchronized static String getMD5Hex(String src) throws  

    NoSuchAlgorithmException {
    MessageDigest md5 = MessageDigest.getInstance(&quot;MD5&quot;);
    byte[] bs = md5.digest(src.getBytes());
    return new String(new Hex().encode(bs));
}

private static String getFormattedText(byte[] bytes) {
    int len = bytes.length;
    StringBuilder buf = new StringBuilder(len * 2);
    for (int j = 0; j &lt; len; j++) {
        buf.append(HEX_DIGITS[(bytes[j] &gt;&gt; 4) &amp; 0x0f]);
        buf.append(HEX_DIGITS[bytes[j] &amp; 0x0f]);
    }
    return buf.toString();
}

private static final char[] HEX_DIGITS = { '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd',
        'e', 'f' };
</code>
	</pre>

	<p>
		<span id="3.5"><strong>进行Base64加密：</strong></span>
	</p>
	<pre>
		<code>/**
 * BASE64解密
 * @param key
 * @return byte[]
 * @throws Exception
 */
public static byte[] decryptBASE64(String key) throws Exception {
    return (new BASE64Decoder()).decodeBuffer(key);
}

/**
 * BASE64加密
 * @param key
 * @return String
 * @throws Exception
 */
public static String encryptBASE64(byte[] key) throws Exception {
    return (new BASE64Encoder()).encodeBuffer(key);
}
</code>
	</pre>

	<h2>
		四、<span id="4">状态码与返回集说明</span>
	</h2>
	<h3>
		1.<span id="4.1">返回结果说明</span>
	</h3>
	<p>HTTP请求返回结果示例:</p>
	<table>
		<thead>
			<tr>
				<th align="center">结果</th>
				<th align="center">成功</th>
				<th align="center">失败</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="center">需返回数据</td>
				<td align="center">{&quot;code&quot;:10000,&quot;message&quot;:&quot;成功(OK)&quot;，&quot;result&quot;:{<strong><em>xxx</em></strong>}}
				</td>
				<td align="center">--</td>
			</tr>
			<tr>
				<td align="center">不需返回数据</td>
				<td align="center">{&quot;code&quot;:10000&quot;,&quot;message&quot;:&quot;成功(OK)&quot;}</td>
				<td align="center">{&quot;code&quot;:21001,&quot;message&quot;:&quot;参数不能为空&quot;}</td>
			</tr>
		</tbody>
	</table>
	<h6>
		&ensp;&ensp; message为友情错误提示； <br />
	</h6>
	<h3>
		2.<span id="4.2">客户端错误状态码</span>
	</h3>
	<table>
		<thead>
			<tr>
				<th align="center">code</th>
				<th align="center">说明</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="center">20xxx</td>
				<td align="center">由用户操作出现的错误</td>
			</tr>
			<tr>
				<td align="center">20001</td>
				<td align="center">请填写用户名</td>
			</tr>
			<tr>
				<td align="center">20002</td>
				<td align="center">手机号/邮箱格式错误</td>
			</tr>
			<tr>
				<td align="center">20003</td>
				<td align="center">请填写密码</td>
			</tr>
			<tr>
				<td align="center">20004</td>
				<td align="center">密码格式错误（6-16位的数字或字母）</td>
			</tr>
			<tr>
				<td align="center">20005</td>
				<td align="center">用户名或密码错误</td>
			</tr>
			<tr>
				<td align="center">20006</td>
				<td align="center">该用户信用太低，禁止登陆</td>
			</tr>
			<tr>
				<td align="center">20007</td>
				<td align="center">该用户状态异常，禁止登陆</td>
			</tr>
			<tr>
				<td align="center">20008</td>
				<td align="center">该用户名已被注册</td>
			</tr>
			<tr>
				<td align="center">20009</td>
				<td align="center">该用户名不存在</td>
			</tr>
			<tr>
				<td align="center">20010</td>
				<td align="center">短信发送次数达到当日上限</td>
			</tr>
			<tr>
				<td align="center">20011</td>
				<td align="center">请填写验证码</td>
			</tr>
			<tr>
				<td align="center">20012</td>
				<td align="center">验证码格式错误</td>
			</tr>
			<tr>
				<td align="center">20013</td>
				<td align="center">验证码错误</td>
			</tr>
			<tr>
				<td align="center">20014</td>
				<td align="center">验证码失效</td>
			</tr>
			<tr>
				<td align="center">20015</td>
				<td align="center">请填写昵称</td>
			</tr>
			<tr>
				<td align="center">20016</td>
				<td align="center">请填写正确的昵称</td>
			</tr>
			<tr>
				<td align="center">20017</td>
				<td align="center">姓名不能为空</td>
			</tr>
			<tr>
				<td align="center">20018</td>
				<td align="center">请填写正确的姓名</td>
			</tr>
			<tr>
				<td align="center">20019</td>
				<td align="center">身份证号码不能为空</td>
			</tr>
			<tr>
				<td align="center">20020</td>
				<td align="center">身份证号码格式错误</td>
			</tr>
			<tr>
				<td align="center">20021</td>
				<td align="center">姓名与身份证号码不符</td>
			</tr>
			<tr>
				<td align="center">20022</td>
				<td align="center">该账户已被实名认证</td>
			</tr>
			<tr>
				<td align="center">20023</td>
				<td align="center">无数据，需要先备份</td>
			</tr>
			<tr>
				<td align="center">20024</td>
				<td align="center">密码错误</td>
			</tr>
			<tr>
				<td align="center">21xxx</td>
				<td align="center">app自动获取数据而出现的错误</td>
			</tr>
			<tr>
				<td align="center">21001</td>
				<td align="center">参数不能为空</td>
			</tr>
			<tr>
				<td align="center">21002</td>
				<td align="center">参数格式不合法</td>
			</tr>
			<tr>
				<td align="center">21003</td>
				<td align="center">记录被封禁</td>
			</tr>
			<tr>
				<td align="center">21004</td>
				<td align="center">用户被警告</td>
			</tr>
			<tr>
				<td align="center">21005</td>
				<td align="center">非法操作或没有权限</td>
			</tr>
			<tr>
				<td align="center">21006</td>
				<td align="center">文件不能为空</td>
			</tr>
			<tr>
				<td align="center">21007</td>
				<td align="center">文件ID错误</td>
			</tr>
		</tbody>
	</table>
	<h3>
		3.<span id="4.3">服务端或第三方错误状态码</span>
	</h3>
	<table>
		<thead>
			<tr>
				<th align="center">code</th>
				<th align="center">说明</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="center">10001</td>
				<td align="center">系统繁忙（不在服务控制内发生的异常）</td>
			</tr>
			<tr>
				<td align="center">10002</td>
				<td align="center">服务繁忙，稍后再试（服务器错误或数据库错误）</td>
			</tr>
			<tr>
				<td align="center">10003</td>
				<td align="center">服务接口维护（接口停用或维护）</td>
			</tr>
			<tr>
				<td align="center">10004</td>
				<td align="center">短信运营商繁忙（发送验证码）</td>
			</tr>
			<tr>
				<td align="center">10005</td>
				<td align="center">邮件服务器繁忙（发送验证码）</td>
			</tr>
		</tbody>
	</table>
	<h2>
		五、<span id="5">URL接口说明</span>
	</h2>
	<h4>
		1：<span id="5.1">发送验证码</span>
	</h4>
	<p>
		功能描述：给自己的手机号码或邮箱发送验证码。<br /> 接口地址：http://<strong><em>ip</em></strong>/chatting/visitors/send_code/(interfaceId)
		<br /> 请求参数说明：
	</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">interfaceId</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">接口序号</td>
				<td align="center">11/18</td>
			</tr>
			<tr>
				<td align="left">username</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">手机号码或者邮箱</td>
				<td align="center">17600117962/wangdy@radacat.com</td>
			</tr>
		</tbody>
	</table>
	<p>返回成功示例：</p>
	<p>
		<img src="http://i.imgur.com/QQNHPpY.png" alt="aas" />
	</p>
	<p>
		返回失败示例： <br />
	</p>
	<p>
		<img src="http://i.imgur.com/aJhHAyv.png" alt="bbb" />
	</p>
	<p>返回参数说明：</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">code</td>
				<td align="center">int</td>
				<td align="center">是</td>
				<td align="center">状态码</td>
				<td align="center">200</td>
			</tr>
		</tbody>
	</table>
	<h4>
		2：<span id="5.2">校验验证码</span>
	</h4>
	<p>
		功能描述：处于安全考虑，需要验证码证明自己的身份。 <br /> 接口地址：http://<strong><em>ip</em></strong>/chatting/visitors/verify_code
		<br /> 请求参数说明：
	</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">username</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">手机号码或者邮箱</td>
				<td align="center">17600117962/wangdy@radacat.com</td>
			</tr>
			<tr>
				<td align="left">checkCode</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">对应发送给用户的验证码或者是发送给邮箱的验证码</td>
				<td align="center">454545</td>
			</tr>
		</tbody>
	</table>
	<p>返回成功示例：</p>
	<p>
		<img src="http://i.imgur.com/QQNHPpY.png" alt="aas" />
	</p>
	<p>
		返回失败示例： <br />
	</p>
	<p>
		<img src="http://i.imgur.com/aJhHAyv.png" alt="bbb" />
	</p>
	<p>返回参数说明：</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">code</td>
				<td align="center">int</td>
				<td align="center">是</td>
				<td align="center">状态码</td>
				<td align="center">200</td>
			</tr>
		</tbody>
	</table>
	<h4>
		3：<span id="5.3">注册</span>
	</h4>
	<p>
		功能描述：用户拥有自己的账户。 <br /> 接口地址：http://<strong><em>ip</em></strong>/chatting/visitors/register
		<br /> 请求参数说明：
	</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">username</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">手机号码或者邮箱</td>
				<td align="center">17600117962/wangdy@radacat.com</td>
			</tr>
			<tr>
				<td align="left">password</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">账户密码</td>
				<td align="center">123456789</td>
			</tr>
			<tr>
				<td align="left">nickname</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">用户昵称</td>
				<td align="center">王东阳</td>
			</tr>
			<tr>
				<td align="left">checkCode</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">验证码</td>
				<td align="center">123456</td>
			</tr>
		</tbody>
	</table>
	<p>返回成功示例：</p>
	<p>
		<img src="http://i.imgur.com/QQNHPpY.png" alt="aas" />
	</p>
	<p>
		返回失败示例： <br />
	</p>
	<p>
		<img src="http://i.imgur.com/aJhHAyv.png" alt="bbb" />
	</p>
	<p>返回参数说明：</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">code</td>
				<td align="center">int</td>
				<td align="center">是</td>
				<td align="center">状态码</td>
				<td align="center">200</td>
			</tr>
		</tbody>
	</table>
	<h4>
		4：<span id="5.4">编辑个人资料</span>
	</h4>
	<p>
		功能描述：编辑与更改自己的资料。 <br /> 接口地址：http://<strong><em>ip</em></strong>/chatting/user/basic_info
		<br /> 请求参数说明：
	</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">pid</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">用户的嗒嗒ID</td>
				<td align="center">1000001</td>
			</tr>
			<tr>
				<td align="left">nickname</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">用户昵称</td>
				<td align="center">王东阳</td>
			</tr>
			<tr>
				<td align="left">gender</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">性别，0：不变；1：男； 2：女；3：无</td>
				<td align="center">1</td>
			</tr>
			<tr>
				<td align="left">birthday</td>
				<td align="center">String</td>
				<td align="center">否</td>
				<td align="center">生日（毫秒值）</td>
				<td align="center">1480294800000</td>
			</tr>
			<tr>
				<td align="left">avatar</td>
				<td align="center">String</td>
				<td align="center">否</td>
				<td align="center">头像ID</td>
				<td align="center">1</td>
			</tr>
			<tr>
				<td align="left">address</td>
				<td align="center">String</td>
				<td align="center">否</td>
				<td align="center">具体地址</td>
				<td align="center">北京市朝阳区</td>
			</tr>
			<tr>
				<td align="left">job</td>
				<td align="center">String</td>
				<td align="center">否</td>
				<td align="center">工作ID</td>
				<td align="center">1</td>
			</tr>
			<tr>
				<td align="left">signature</td>
				<td align="center">String</td>
				<td align="center">否</td>
				<td align="center">个人签名</td>
				<td align="center">有对象的程序员</td>
			</tr>
			<tr>
				<td align="left">commonEmail</td>
				<td align="center">String</td>
				<td align="center">否</td>
				<td align="center">常用邮箱</td>
				<td align="center">wangdy@radacat.com</td>
			</tr>
		</tbody>
	</table>
	<p>返回成功示例：</p>
	<p>
		<img src="http://i.imgur.com/QQNHPpY.png" alt="aas" />
	</p>
	<p>
		返回失败示例： <br />
	</p>
	<p>
		<img src="http://i.imgur.com/aJhHAyv.png" alt="bbb" />
	</p>
	<p>返回参数说明：</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">code</td>
				<td align="center">int</td>
				<td align="center">是</td>
				<td align="center">状态码</td>
				<td align="center">200</td>
			</tr>
		</tbody>
	</table>
	<h4>
		5：<span id="5.5">修改密码</span>
	</h4>
	<p>
		功能描述：用户登录之后，可以将旧密码修改为新密码。<br /> 接口地址：http://<strong><em>ip</em></strong>/chatting/user/pwd/(action)
		<br />
	</p>
	<h6>action:(verify或change)</h6>
	<p>请求参数说明：</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">pid</td>
				<td align="center">long</td>
				<td align="center">是</td>
				<td align="center">用户id</td>
				<td align="center">1000001</td>
			</tr>
			<tr>
				<td align="left">pwd</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">账号旧密码</td>
				<td align="center">123456</td>
			</tr>
			<tr>
				<td align="left">newPwd</td>
				<td align="center">String</td>
				<td align="center">否</td>
				<td align="center">修改账号新密码(校验旧密码时newPwd不需要)</td>
				<td align="center">111111</td>
			</tr>
		</tbody>
	</table>
	<p>返回成功示例：</p>
	<p>
		<img src="http://i.imgur.com/QQNHPpY.png" alt="aas" />
	</p>
	<p>
		返回失败示例： <br />
	</p>
	<p>
		<img src="http://i.imgur.com/aJhHAyv.png" alt="bbb" />
	</p>
	<p>返回参数说明：</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">code</td>
				<td align="center">int</td>
				<td align="center">是</td>
				<td align="center">状态码</td>
				<td align="center">200</td>
			</tr>
		</tbody>
	</table>
	<h4>
		6：<span id="5.6">实名认证</span>
	</h4>
	<p>
		功能描述：将个人账户进行实名认证。 <br /> 接口地址：http://<strong><em>ip</em></strong>/chatting/user/real_name
		<br /> 请求参数说明：
	</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">pid</td>
				<td align="center">long</td>
				<td align="center">是</td>
				<td align="center">用户的嗒嗒ID</td>
				<td align="center">1000001</td>
			</tr>
			<tr>
				<td align="left">realName</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">用户的真实姓名</td>
				<td align="center">王东阳</td>
			</tr>
			<tr>
				<td align="left">idCard</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">身份证号码</td>
				<td align="center">130524********1019</td>
			</tr>
		</tbody>
	</table>
	<p>返回成功示例：</p>
	<p>
		<img src="http://i.imgur.com/QQNHPpY.png" alt="aas" />
	</p>
	<p>
		返回失败示例： <br />
	</p>
	<p>
		<img src="http://i.imgur.com/aJhHAyv.png" alt="bbb" />
	</p>
	<p>返回参数说明：</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">code</td>
				<td align="center">int</td>
				<td align="center">是</td>
				<td align="center">状态码</td>
				<td align="center">200</td>
			</tr>
		</tbody>
	</table>
	<h4>
		7：<span id="5.7">找回密码</span>
	</h4>
	<p>
		功能描述：实现忘记密码之后的重置密码。 接口地址：http://<strong><em>ip</em></strong>/chatting/visitors/reset_pwd
		请求参数说明：
	</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">username</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">用户名</td>
				<td align="center">17600117962</td>
			</tr>
			<tr>
				<td align="left">checkCode</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">验证码</td>
				<td align="center">123456</td>
			</tr>
			<tr>
				<td align="left">password</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">密码</td>
				<td align="center">1234567890</td>
			</tr>
		</tbody>
	</table>
	<p>返回成功示例：</p>
	<p>
		<img src="http://i.imgur.com/QQNHPpY.png" alt="aas" />
	</p>
	<p>
		返回失败示例： <br />
	</p>
	<p>
		<img src="http://i.imgur.com/aJhHAyv.png" alt="bbb" />
	</p>
	<p>返回参数说明：</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">code</td>
				<td align="center">int</td>
				<td align="center">是</td>
				<td align="center">状态码</td>
				<td align="center">200</td>
			</tr>
		</tbody>
	</table>
	<h4>
		8：<span id="5.8">登录</span>
	</h4>
	<p>
		功能描述：游客未登录状态转为用户登录状态。 <br /> 接口地址：http://<strong><em>ip</em></strong>/chatting/user/login
		<br /> 请求参数说明：
	</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">username</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">用户名</td>
				<td align="center">17600117962</td>
			</tr>
			<tr>
				<td align="left">password</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">密码</td>
				<td align="center">1234567890</td>
			</tr>
			<tr>
				<td align="left">device</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">设备型号</td>
				<td align="center">meizu MX5</td>
			</tr>
			<tr>
				<td align="left">imei</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">手机串号或唯一识别码</td>
				<td align="center">123456789123456</td>
			</tr>
		</tbody>
	</table>
	<p>返回成功示例：</p>
	<p>
		<img src="http://i.imgur.com/QQNHPpY.png" alt="aas" />
	</p>
	<p>
		返回失败示例： <br />
	</p>
	<p>
		<img src="http://i.imgur.com/aJhHAyv.png" alt="bbb" />
	</p>
	<p>返回参数说明：</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">code</td>
				<td align="center">int</td>
				<td align="center">是</td>
				<td align="center">状态码</td>
				<td align="center">200</td>
			</tr>
			<tr>
				<td align="left">result</td>
				<td align="center">T</td>
				<td align="center">是</td>
				<td align="center">以下参数为用户信息对象</td>
				<td align="center">--</td>
			</tr>
			<tr>
				<td align="left">pid</td>
				<td align="center">long</td>
				<td align="center">是</td>
				<td align="center">嗒嗒ID</td>
				<td align="center">1000001</td>
			</tr>
			<tr>
				<td align="left">nickname</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">昵称</td>
				<td align="center">王东阳</td>
			</tr>
			<tr>
				<td align="left">username</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">用户名</td>
				<td align="center">17600117962</td>
			</tr>
			<tr>
				<td align="left">gender</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">性别</td>
				<td align="center">1</td>
			</tr>
			<tr>
				<td align="left">birthday</td>
				<td align="center">long</td>
				<td align="center">是</td>
				<td align="center">生日</td>
				<td align="center">1489909978734</td>
			</tr>
			<tr>
				<td align="left">avatar</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">头像</td>
				<td align="center">1</td>
			</tr>
			<tr>
				<td align="left">province</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">省或直辖市</td>
				<td align="center">1</td>
			</tr>
			<tr>
				<td align="left">city</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">市</td>
				<td align="center">1</td>
			</tr>
			<tr>
				<td align="left">address</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">具体地址</td>
				<td align="center">北京市朝阳区</td>
			</tr>
			<tr>
				<td align="left">jobId</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">工作id</td>
				<td align="center">1</td>
			</tr>
			<tr>
				<td align="left">signature</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">签名</td>
				<td align="center">有对象的程序员</td>
			</tr>
			<tr>
				<td align="left">commonEmail</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">常用邮箱</td>
				<td align="center">wangdy@radacat.com</td>
			</tr>
			<tr>
				<td align="left">name</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">姓名</td>
				<td align="center">王东阳</td>
			</tr>
			<tr>
				<td align="left">idCard</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">身份证号码</td>
				<td align="center">130524********1019</td>
			</tr>
			<tr>
				<td align="left">credits</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">信誉点</td>
				<td align="center">666</td>
			</tr>
			<tr>
				<td align="left">point</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">积分</td>
				<td align="center">66</td>
			</tr>
			<tr>
				<td align="left">grade</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">等级</td>
				<td align="center">6</td>
			</tr>
		</tbody>
	</table>
	<h4>
		9：<span id="5.9">退出</span>
	</h4>
	<p>
		功能描述：从登录状态退出。 <br /> 接口地址：http://<strong><em>ip</em></strong>/chatting/user/exit
		<br /> 请求参数说明：
	</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">pid</td>
				<td align="center">long</td>
				<td align="center">是</td>
				<td align="center">用户的嗒嗒ID</td>
				<td align="center">1000001</td>
			</tr>
			<tr>
				<td align="left">imei</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">设备唯一标识</td>
				<td align="center">11111111111111</td>
			</tr>
		</tbody>
	</table>
	<p>返回成功示例：</p>
	<p>
		<img src="http://i.imgur.com/QQNHPpY.png" alt="aaa" />
	</p>
	<p>
		返回失败示例： <br />
	</p>
	<p>
		<img src="http://i.imgur.com/aJhHAyv.png" alt="bbb" />
	</p>
	<p>返回参数说明：</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">code</td>
				<td align="center">int</td>
				<td align="center">是</td>
				<td align="center">状态码</td>
				<td align="center">200</td>
			</tr>
		</tbody>
	</table>
	<h4>
		10：<span id="5.10">获取联系人</span>
	</h4>
	<p>
		功能描述：获取自己存储在云端的通讯录。 <br /> 请求URL：http://<strong><em>ip</em></strong>/chatting/friend/get_friend
		<br /> 请求参数说明：
	</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">pid</td>
				<td align="center">long</td>
				<td align="center">是</td>
				<td align="center">用户名</td>
				<td align="center">1000001</td>
			</tr>
		</tbody>
	</table>
	<p>
		返回成功示例： <br />
	</p>
	<p>
		<img src="http://i.imgur.com/EJ9mdZa.png" alt="aaa" />
	</p>
	<p>
		返回失败示例： <br />
	</p>
	<p>
		<img src="http://i.imgur.com/1gbg9ui.png" alt="bbb" />
	</p>
	<p>返回参数说明：</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">code</td>
				<td align="center">int</td>
				<td align="center">是</td>
				<td align="center">状态码</td>
				<td align="center">200</td>
			</tr>
			<tr>
				<td align="left">message</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">友情错误提示</td>
				<td align="center">--</td>
			</tr>
			<tr>
				<td align="left">result</td>
				<td align="center">T</td>
				<td align="center">是</td>
				<td align="center">以下参数为联系人信息对象List集合,以下参数都为对象内容</td>
				<td align="center">--</td>
			</tr>
			<tr>
				<td align="left">pid</td>
				<td align="center">long</td>
				<td align="center">是</td>
				<td align="center">嗒嗒ID</td>
				<td align="center">1000001</td>
			</tr>
			<tr>
				<td align="left">nickname</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">昵称</td>
				<td align="center">王东阳</td>
			</tr>
			<tr>
				<td align="left">note</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">备注</td>
				<td align="center">17600117962</td>
			</tr>
			<tr>
				<td align="left">gender</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">性别</td>
				<td align="center">1</td>
			</tr>
			<tr>
				<td align="left">birthday</td>
				<td align="center">long</td>
				<td align="center">是</td>
				<td align="center">生日</td>
				<td align="center">1489909978734</td>
			</tr>
			<tr>
				<td align="left">avatar</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">头像</td>
				<td align="center">1</td>
			</tr>
			<tr>
				<td align="left">province</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">省或直辖市</td>
				<td align="center">1</td>
			</tr>
			<tr>
				<td align="left">city</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">市</td>
				<td align="center">1</td>
			</tr>
			<tr>
				<td align="left">address</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">具体地址</td>
				<td align="center">北京市朝阳区</td>
			</tr>
			<tr>
				<td align="left">jobId</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">工作</td>
				<td align="center">1</td>
			</tr>
			<tr>
				<td align="left">signature</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">签名</td>
				<td align="center">有对象的程序员</td>
			</tr>
			<tr>
				<td align="left">commonEmail</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">常用邮箱</td>
				<td align="center">wangdy@radacat.com</td>
			</tr>
			<tr>
				<td align="left">credits</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">信誉点</td>
				<td align="center">666</td>
			</tr>
			<tr>
				<td align="left">grade</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">等级</td>
				<td align="center">6</td>
			</tr>
			<tr>
				<td align="left">status</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">联系人状态</td>
				<td align="center">1</td>
			</tr>
		</tbody>
	</table>
	<h4>
		11：<span id="5.11">备份联系人</span>
	</h4>
	<p>
		功能描述：在无运营商网络状态下添加的好友，当处于运营商网络下时，可以选择上传好友。 <br /> 请求URL：http://<strong><em>ip</em></strong>/chatting/friend/backup_friend
		<br /> 请求参数说明：
	</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">pid</td>
				<td align="center">long</td>
				<td align="center">是</td>
				<td align="center">用户名</td>
				<td align="center">1000001</td>
			</tr>
			<tr>
				<td align="left">longitude</td>
				<td align="center">double</td>
				<td align="center">是</td>
				<td align="center">经度</td>
				<td align="center">12.123456</td>
			</tr>
			<tr>
				<td align="left">latitude</td>
				<td align="center">double</td>
				<td align="center">是</td>
				<td align="center">纬度</td>
				<td align="center">12.123456</td>
			</tr>
			<tr>
				<td align="left">infos</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">由map转为json之后的字符串：opid与note</td>
				<td align="center"></td>
			</tr>
			<tr>
				<td align="left">opid</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">联系人的pid</td>
				<td align="center">1000002</td>
			</tr>
			<tr>
				<td align="left">note</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">对联系人的备注</td>
				<td align="center">老滑</td>
			</tr>
		</tbody>
	</table>
	<p>
		返回成功示例： <br />
	</p>
	<p>
		<img src="http://i.imgur.com/QQNHPpY.png" alt="aaa" />
	</p>
	<p>
		返回失败示例： <br />
	</p>
	<p>
		<img src="http://i.imgur.com/7DEMRiJ.png" alt="bbb" />
	</p>
	<p>返回参数说明：</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">code</td>
				<td align="center">int</td>
				<td align="center">是</td>
				<td align="center">状态码</td>
				<td align="center">200</td>
			</tr>
			<tr>
				<td align="left">message</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">友情错误提示</td>
				<td align="center">参数格式不合法</td>
			</tr>
		</tbody>
	</table>
	<h4>
		12：<span id="5.12">发送私信</span>
	</h4>
	<p>
		功能描述：实现用户对联系人发送私信的功能。 <br /> 请求URL：http://<strong><em>ip</em></strong>/chatting/letter/send_letter
		<br /> 请求参数说明：
	</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">pid</td>
				<td align="center">long</td>
				<td align="center">是</td>
				<td align="center">用户的嗒嗒ID</td>
				<td align="center">1000001</td>
			</tr>
			<tr>
				<td align="left">opid</td>
				<td align="center">long</td>
				<td align="center">是</td>
				<td align="center">联系人的嗒嗒ID</td>
				<td align="center">1000002</td>
			</tr>
			<tr>
				<td align="left">type</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">信息的类型{1：字符串文本，2：文件}</td>
				<td align="center">1</td>
			</tr>
			<tr>
				<td align="left">letter</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">字符串文本内容或文件名称</td>
				<td align="center">hello world！</td>
			</tr>
		</tbody>
	</table>
	<p>
		返回成功示例： <br />
	</p>
	<p>
		<img src="http://i.imgur.com/LtievgK.png" alt="aaa" />
	</p>
	<p>
		返回失败示例： <br />
	</p>
	<p>
		<img src="http://i.imgur.com/T5uH6CN.png" alt="bbb" />
	</p>
	<p>返回参数说明：</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">code</td>
				<td align="center">int</td>
				<td align="center">是</td>
				<td align="center">状态码</td>
				<td align="center">200</td>
			</tr>
			<tr>
				<td align="left">message</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">友情错误提示</td>
				<td align="center">参数格式不合法</td>
			</tr>
		</tbody>
	</table>
	<h4>
		13：<span id="5.13">接收指定联系人私信</span>
	</h4>
	<p>
		功能描述：实现用户接收指定联系人的私信内容。 <br />
		请求URL：http://ip/chatting/letter/get_letter <br /> 请求参数说明：
	</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">pid</td>
				<td align="center">long</td>
				<td align="center">是</td>
				<td align="center">用户的嗒嗒ID</td>
				<td align="center">1000001</td>
			</tr>
			<tr>
				<td align="left">opid</td>
				<td align="center">long</td>
				<td align="center">是</td>
				<td align="center">联系人的嗒嗒ID</td>
				<td align="center">1000002</td>
			</tr>
		</tbody>
	</table>
	<p>
		返回成功示例： <br />
	</p>
	<p>
		<img src="http://i.imgur.com/xjUQLpc.png" alt="aaa" />
	</p>
	<p>
		返回失败示例： <br />
	</p>
	<p>
		<img src="http://i.imgur.com/T5uH6CN.png" alt="bbb" />
	</p>
	<p>返回参数说明：</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">code</td>
				<td align="center">int</td>
				<td align="center">是</td>
				<td align="center">状态码</td>
				<td align="center">200</td>
			</tr>
			<tr>
				<td align="left">message</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">友情错误提示</td>
				<td align="center">参数格式不合法</td>
			</tr>
			<tr>
				<td align="left">result</td>
				<td align="center">T</td>
				<td align="center">是</td>
				<td align="center">以下参数为私信信息对象List集合,以下参数都为对象内容</td>
				<td align="center">--</td>
			</tr>
			<tr>
				<td align="left">sendTime</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">私信发送时间</td>
				<td align="center">1489909978734</td>
			</tr>
			<tr>
				<td align="left">type</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">私信信息类型</td>
				<td align="center">1</td>
			</tr>
			<tr>
				<td align="left">messageInfo</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">私信内容或文件名称</td>
				<td align="center">hello world！</td>
			</tr>
		</tbody>
	</table>
	<h4>
		14：<span id="5.14">升级或更新群</span>
	</h4>
	<p>
		接口描述：将临时群升级为永久群或者从群主处同步群信息和成员列表到云端。 <br /> 接口地址：http://<strong><em>ip</em></strong>/chatting/crowd/sync
		<br /> 请求参数说明：
	</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">tgid</td>
				<td align="center">long</td>
				<td align="center">是</td>
				<td align="center">永久群ID或临时群ID</td>
				<td align="center">100011330</td>
			</tr>
			<tr>
				<td align="left">tgname</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">群名称</td>
				<td align="center">qunliao</td>
			</tr>
			<tr>
				<td align="left">pid</td>
				<td align="center">long</td>
				<td align="center">是</td>
				<td align="center">群主嗒嗒ID</td>
				<td align="center">1000001</td>
			</tr>
			<tr>
				<td align="left">members</td>
				<td align="center">String[]</td>
				<td align="center">是</td>
				<td align="center">群成员的嗒嗒ID</td>
				<td align="center">1000002,1000003</td>
			</tr>
			<tr>
				<td align="left">notice</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">群公告</td>
				<td align="center">welcome</td>
			</tr>
			<tr>
				<td align="left">intro</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">群描述</td>
				<td align="center">望京聊天群</td>
			</tr>
			<tr>
				<td align="left">longitude</td>
				<td align="center">double</td>
				<td align="center">是</td>
				<td align="center">经度，用于知道用户聊天的地点</td>
				<td align="center">12.345678</td>
			</tr>
			<tr>
				<td align="left">latitude</td>
				<td align="center">double</td>
				<td align="center">是</td>
				<td align="center">纬度，用于知道用户聊天的地点</td>
				<td align="center">12.345678</td>
			</tr>
			<tr>
				<td align="left">createTime</td>
				<td align="center">long</td>
				<td align="center">是</td>
				<td align="center">群创建时的时间</td>
				<td align="center">1489909978734</td>
			</tr>
			<tr>
				<td align="left">icon</td>
				<td align="center">int</td>
				<td align="center">是</td>
				<td align="center">群头像</td>
				<td align="center">1</td>
			</tr>
		</tbody>
	</table>
	<p>返回成功示例：</p>
	<p>
		<img src="http://i.imgur.com/aArDiwy.png" alt="aaa" />
	</p>
	<p>
		返回失败示例： <br />
	</p>
	<p>
		<img src="http://i.imgur.com/T5uH6CN.png" alt="bbb" />
	</p>
	<p>返回参数说明：</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">code</td>
				<td align="center">int</td>
				<td align="center">是</td>
				<td align="center">状态码</td>
				<td align="center">200</td>
			</tr>
			<tr>
				<td align="left">message</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">友情错误提示</td>
				<td align="center">参数格式不合法</td>
			</tr>
			<tr>
				<td align="left">result</td>
				<td align="center">T</td>
				<td align="center">是</td>
				<td align="center">返回结果集，以下为结果集参数</td>
				<td align="center">--</td>
			</tr>
			<tr>
				<td align="left">pgid</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">永久群ID</td>
				<td align="center">100011330</td>
			</tr>
		</tbody>
	</table>
	<h4>
		15：<span id="5.15">踢人出群</span>
	</h4>
	<p>
		接口描述：实现群主将群成员踢出群的功能。 <br /> 接口地址：http://<strong><em>ip</em></strong>/chatting/crowd/kick
		<br /> 请求参数说明：
	</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">tgid</td>
				<td align="center">long</td>
				<td align="center">是</td>
				<td align="center">永久群的ID，群唯一标识</td>
				<td align="center">100011330</td>
			</tr>
			<tr>
				<td align="left">pid</td>
				<td align="center">long</td>
				<td align="center">是</td>
				<td align="center">群主嗒嗒ID</td>
				<td align="center">1000001</td>
			</tr>
			<tr>
				<td align="left">members</td>
				<td align="center">String[]</td>
				<td align="center">是</td>
				<td align="center">被移除群成员的嗒嗒ID</td>
				<td align="center">[1000002,1000003]</td>
			</tr>
		</tbody>
	</table>
	<p>
		返回成功示例： <br />
	</p>
	<p>
		<img src="http://i.imgur.com/vnc0jTg.png" alt="aaa" />
	</p>
	<p>
		返回失败示例： <br />
	</p>
	<p>
		<img src="http://i.imgur.com/T5uH6CN.png" alt="bbb" />
	</p>
	<p>返回参数说明：</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">code</td>
				<td align="center">int</td>
				<td align="center">是</td>
				<td align="center">状态码</td>
				<td align="center">200</td>
			</tr>
			<tr>
				<td align="left">message</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">友情错误提示</td>
				<td align="center">参数格式不合法</td>
			</tr>
		</tbody>
	</table>
	<h4>
		16：<span id="5.16">解散群</span>
	</h4>
	<p>
		接口描述：实现群主将群解散的功能。 <br /> 接口地址：http://<strong><em>ip</em></strong>/chatting/crowd/remove
		<br /> 请求参数说明：
	</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">tgid</td>
				<td align="center">long</td>
				<td align="center">是</td>
				<td align="center">永久群的ID，群唯一标识</td>
				<td align="center">100011330</td>
			</tr>
			<tr>
				<td align="left">pid</td>
				<td align="center">long</td>
				<td align="center">是</td>
				<td align="center">群主嗒嗒ID</td>
				<td align="center">1000001</td>
			</tr>
		</tbody>
	</table>
	<p>返回成功示例：</p>
	<p>
		<img src="http://i.imgur.com/vnc0jTg.png" alt="aaa" />
	</p>
	<p>返回失败示例：</p>
	<p>
		<img src="http://i.imgur.com/T5uH6CN.png" alt="bbb" />
	</p>
	<p>返回参数说明：</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">code</td>
				<td align="center">int</td>
				<td align="center">是</td>
				<td align="center">状态码</td>
				<td align="center">200</td>
			</tr>
			<tr>
				<td align="left">message</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">友情错误提示</td>
				<td align="center">参数格式不合法</td>
			</tr>
		</tbody>
	</table>
	<h4>
		17：<span id="5.17">编辑群资料</span>
	</h4>
	<p>
		接口描述：实现群主修改群名称、群公告、群描述、群头像的功能。 <br /> 接口地址：http://<strong><em>ip</em></strong>/chatting/crowd/update
		<br /> 请求参数说明：
	</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">tgname</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">群名称</td>
				<td align="center">聊天群</td>
			</tr>
			<tr>
				<td align="left">tgid</td>
				<td align="center">long</td>
				<td align="center">是</td>
				<td align="center">永久群的ID，群唯一标识</td>
				<td align="center">100011330</td>
			</tr>
			<tr>
				<td align="left">pid</td>
				<td align="center">long</td>
				<td align="center">是</td>
				<td align="center">群主的嗒嗒ID</td>
				<td align="center">1000001</td>
			</tr>
			<tr>
				<td align="left">announcement</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">群公告</td>
				<td align="center">welcome</td>
			</tr>
			<tr>
				<td align="left">intro</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">群描述</td>
				<td align="center">聊天群</td>
			</tr>
			<tr>
				<td align="left">icon</td>
				<td align="center">int</td>
				<td align="center">是</td>
				<td align="center">群头像</td>
				<td align="center">1</td>
			</tr>
		</tbody>
	</table>
	<p>返回成功示例：</p>
	<p>
		<img src="http://i.imgur.com/vnc0jTg.png" alt="aaa" />
	</p>
	<p>返回失败示例：</p>
	<p>
		<img src="http://i.imgur.com/T5uH6CN.png" alt="bbb" />
	</p>
	<p>返回参数说明：</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">code</td>
				<td align="center">int</td>
				<td align="center">是</td>
				<td align="center">状态码</td>
				<td align="center">200</td>
			</tr>
			<tr>
				<td align="left">message</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">友情错误提示</td>
				<td align="center">参数格式不合法</td>
			</tr>
		</tbody>
	</table>
	<h4>
		18：<span id="5.18">主动退群</span>
	</h4>
	<p>
		接口描述：实现用户退群的功能。 <br /> 接口地址：http://<strong><em>ip</em></strong>/chatting/crowd/leave
		<br /> 请求参数说明：
	</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">tgid</td>
				<td align="center">long</td>
				<td align="center">是</td>
				<td align="center">永久群的ID，群唯一标识</td>
				<td align="center">100011330</td>
			</tr>
			<tr>
				<td align="left">members</td>
				<td align="center">String[]</td>
				<td align="center">是</td>
				<td align="center">用户的嗒嗒ID</td>
				<td align="center">[1000002,1000003]</td>
			</tr>
		</tbody>
	</table>
	<p>返回成功示例：</p>
	<p>
		<img src="http://i.imgur.com/vnc0jTg.png" alt="aaa" />
	</p>
	<p>返回失败示例：</p>
	<p>
		<img src="http://i.imgur.com/T5uH6CN.png" alt="bbb" />
	</p>
	<p>返回参数说明：</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">code</td>
				<td align="center">int</td>
				<td align="center">是</td>
				<td align="center">状态码</td>
				<td align="center">200</td>
			</tr>
			<tr>
				<td align="left">message</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">友情错误提示</td>
				<td align="center">参数格式不合法</td>
			</tr>
		</tbody>
	</table>
	<h4>
		19：<span id="5.19">我的群</span>
	</h4>
	<p>
		接口描述：用于实现群成员从云端同步群信息到本地的功能。 <br /> 接口地址：http://<strong><em>ip</em></strong>/chatting/group_data/get_groups
		<br /> 请求参数说明：
	</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">pid</td>
				<td align="center">pid</td>
				<td align="center">是</td>
				<td align="center">用户的嗒嗒ID</td>
				<td align="center">1000002</td>
			</tr>
		</tbody>
	</table>
	<p>返回成功示例：</p>
	<p>
		<img src="http://i.imgur.com/m0xlerp.png" alt="aaa" />
	</p>
	<p>返回失败示例：</p>
	<p>
		<img src="http://i.imgur.com/T5uH6CN.png" alt="bbb" />
	</p>
	<p>返回参数说明：</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">code</td>
				<td align="center">int</td>
				<td align="center">是</td>
				<td align="center">状态码</td>
				<td align="center">200</td>
			</tr>
			<tr>
				<td align="left">message</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">友情错误提示</td>
				<td align="center">参数格式不合法</td>
			</tr>
			<tr>
				<td align="left">result</td>
				<td align="center">T</td>
				<td align="center">是</td>
				<td align="center">返回结果集，由set集合转为json的字符串</td>
				<td align="center">--</td>
			</tr>
			<tr>
				<td align="left">Set<GroupInfo></td>
				<td align="center">Set</td>
				<td align="center">是</td>
				<td align="center">用户所有的群信息集合</td>
				<td align="center">200</td>
			</tr>
			<tr>
				<td align="left">pgid</td>
				<td align="center">long</td>
				<td align="center">是</td>
				<td align="center">群ID</td>
				<td align="center">100011340</td>
			</tr>
			<tr>
				<td align="left">groupInfo</td>
				<td align="center">GroupPermanentPO</td>
				<td align="center">是</td>
				<td align="center">群信息</td>
				<td align="center">--</td>
			</tr>
			<tr>
				<td align="left">usersMap</td>
				<td align="center">Map&lt;String,GroupUserInfoPO&gt;</td>
				<td align="center">是</td>
				<td align="center">群用户信息集合</td>
				<td align="center">--</td>
			</tr>
		</tbody>
	</table>
	<h4>
		20：<span id="5.20">上传群聊记录</span>
	</h4>
	<p>
		接口描述：将存在本地的聊天记录上传至云服务器。 <br /> 接口地址：http://<strong><em>ip</em></strong>/chatting/group_message/sync
		<br /> 请求参数说明：
	</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">pid</td>
				<td align="center">Long</td>
				<td align="center">是</td>
				<td align="center">用户的嗒嗒ID</td>
				<td align="center">1000002</td>
			</tr>
			<tr>
				<td align="left">tgid</td>
				<td align="center">Long</td>
				<td align="center">是</td>
				<td align="center">永久群的ID，群唯一标识</td>
				<td align="center">2000001</td>
			</tr>
			<tr>
				<td align="left">createTime</td>
				<td align="center">Long</td>
				<td align="center">是</td>
				<td align="center">开始聊天时的时间</td>
				<td align="center">1489909978734</td>
			</tr>
			<tr>
				<td align="left">message</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">本地的聊天记录</td>
				<td align="center">--</td>
			</tr>
			<tr>
				<td align="left">fileId</td>
				<td align="center">String</td>
				<td align="center">否</td>
				<td align="center">上传文件时由服务器返回的ID</td>
				<td align="center">590c46e442b0b8efd473412f</td>
			</tr>
			<tr>
				<td align="left">fileName</td>
				<td align="center">String</td>
				<td align="center">否</td>
				<td align="center">文件的名称</td>
				<td align="center">nihao.zip</td>
			</tr>
			<tr>
				<td align="left">longitude</td>
				<td align="center">double</td>
				<td align="center">是</td>
				<td align="center">开始聊天时的经度</td>
				<td align="center">12.345678</td>
			</tr>
			<tr>
				<td align="left">latitude</td>
				<td align="center">double</td>
				<td align="center">是</td>
				<td align="center">开始聊天时的纬度</td>
				<td align="center">12.345678</td>
			</tr>
			<tr>
				<td align="left">address</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">开始聊天的地址</td>
				<td align="center">悠乐汇E</td>
			</tr>
		</tbody>
	</table>
	<p>返回成功示例：</p>
	<p>
		<img src="http://i.imgur.com/vnc0jTg.png" alt="aaa" />
	</p>
	<p>返回失败示例：</p>
	<p>
		<img src="http://i.imgur.com/T5uH6CN.png" alt="bbb" />
	</p>
	<p>返回参数说明：</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">code</td>
				<td align="center">int</td>
				<td align="center">是</td>
				<td align="center">状态码</td>
				<td align="center">200</td>
			</tr>
			<tr>
				<td align="left">message</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">友情错误提示</td>
				<td align="center">参数格式不合法</td>
			</tr>
		</tbody>
	</table>
	<h4>
		21：<span id="5.21">获取足迹列表</span>
	</h4>
	<p>
		接口描述：获取我在云端的足迹列表。 <br /> 接口地址：http://<strong><em>ip</em></strong>/chatting/group_message/list
		<br /> 请求参数说明：
	</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">pid</td>
				<td align="center">long</td>
				<td align="center">是</td>
				<td align="center">用户的嗒嗒ID</td>
				<td align="center">1000001</td>
			</tr>
		</tbody>
	</table>
	<p>返回成功示例：</p>
	<p>
		<img src="http://i.imgur.com/JqsYgSM.png" alt="aaa" />
	</p>
	<p>返回失败示例：</p>
	<p>
		<img src="http://i.imgur.com/T5uH6CN.png" alt="bbb" />
	</p>
	<p>返回参数说明：</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">code</td>
				<td align="center">int</td>
				<td align="center">是</td>
				<td align="center">状态码</td>
				<td align="center">200</td>
			</tr>
			<tr>
				<td align="left">message</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">友情错误提示</td>
				<td align="center">21002</td>
			</tr>
			<tr>
				<td align="left">result</td>
				<td align="center">T</td>
				<td align="center">是</td>
				<td align="center">返回集</td>
				<td align="center">--</td>
			</tr>
			<tr>
				<td align="left">pid</td>
				<td align="center">long</td>
				<td align="center">是</td>
				<td align="center">群主id</td>
				<td align="center">1000001</td>
			</tr>
			<tr>
				<td align="left">pgid</td>
				<td align="center">long</td>
				<td align="center">是</td>
				<td align="center">群id</td>
				<td align="center">2015454</td>
			</tr>
			<tr>
				<td align="left">createTime</td>
				<td align="center">long</td>
				<td align="center">是</td>
				<td align="center">创建时间</td>
				<td align="center">1490076234000</td>
			</tr>
			<tr>
				<td align="left">address</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">详细地址</td>
				<td align="center">北京市朝阳区</td>
			</tr>
			<tr>
				<td align="left">latitude</td>
				<td align="center">double</td>
				<td align="center">是</td>
				<td align="center">纬度</td>
				<td align="center">39.98953</td>
			</tr>
			<tr>
				<td align="left">longitude</td>
				<td align="center">double</td>
				<td align="center">是</td>
				<td align="center">经度</td>
				<td align="center">116.479315</td>
			</tr>
			<tr>
				<td align="left">messageInfo</td>
				<td align="center">int</td>
				<td align="center">是</td>
				<td align="center">消息id</td>
				<td align="center">79</td>
			</tr>
			<tr>
				<td align="left">status</td>
				<td align="center">int</td>
				<td align="center">是</td>
				<td align="center">记录状态1：正常</td>
				<td align="center">1</td>
			</tr>
		</tbody>
	</table>
	<h4>
		22：<span id="5.22">获取足迹信息</span>
	</h4>
	<p>
		接口描述：实现用户从云端获取某个足迹内容的功能。 <br /> 接口地址：http://<strong><em>ip</em></strong>/chatting/group_message/info
		<br /> 请求参数说明：
	</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">id</td>
				<td align="center">long</td>
				<td align="center">是</td>
				<td align="center">messageInfo的id</td>
				<td align="center">123</td>
			</tr>
		</tbody>
	</table>
	<p>返回成功示例：</p>
	<p>
		<img src="http://i.imgur.com/4xUAbT7.png" alt="aaa" />
	</p>
	<p>返回失败示例：</p>
	<p>
		<img src="http://i.imgur.com/T5uH6CN.png" alt="bbb" />
	</p>
	<p>返回参数说明：</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">code</td>
				<td align="center">int</td>
				<td align="center">是</td>
				<td align="center">状态码</td>
				<td align="center">200</td>
			</tr>
			<tr>
				<td align="left">message</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">友情错误提示</td>
				<td align="center">参数格式不合法</td>
			</tr>
			<tr>
				<td align="left">result</td>
				<td align="center">T</td>
				<td align="center">是</td>
				<td align="center">返回结果集</td>
				<td align="center">--</td>
			</tr>
			<tr>
				<td align="left">id</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">足迹id</td>
				<td align="center">--</td>
			</tr>
			<tr>
				<td align="left">messageInfo</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">足迹内容</td>
				<td align="center">--</td>
			</tr>
		</tbody>
	</table>
	<h4>
		23：<span id="5.23">消息通知</span>
	</h4>
	<p>
		接口描述：获取关于自己的系统通知（私信数量）。 <br /> 接口地址：http://<strong><em>ip</em></strong>/chatting/notice/dada_notice
		<br /> 请求参数说明：
	</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">pid</td>
				<td align="center">long</td>
				<td align="center">是</td>
				<td align="center">用户的嗒嗒ID</td>
				<td align="center">1000002</td>
			</tr>
		</tbody>
	</table>
	<p>返回成功示例：</p>
	<p>
		<img src="http://i.imgur.com/GWknRXP.png" alt="dada" />
	</p>
	<p>返回失败示例：</p>
	<p>
		<img src="http://i.imgur.com/T5uH6CN.png" alt="bbb" />
	</p>
	<p>返回参数说明：</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">code</td>
				<td align="center">int</td>
				<td align="center">是</td>
				<td align="center">状态码</td>
				<td align="center">200</td>
			</tr>
			<tr>
				<td align="left">result</td>
				<td align="center">T</td>
				<td align="center">是</td>
				<td align="center">返回集，由map转为的jsonArray</td>
				<td align="center">--</td>
			</tr>
			<tr>
				<td align="left">map键值对</td>
				<td align="center">String</td>
				<td align="center">是</td>
				<td align="center">key为发送私信的pid，value为私信数量</td>
				<td align="center">--</td>
			</tr>
		</tbody>
	</table>
	<h4>
		24：<span id="5.24">上传文件</span>
	</h4>
	<p>
		接口描述：可以将需要想存留的信息或记录以文件方式上传到云端。 <br /> 接口地址：http://<strong><em>ip</em></strong>/chatting/file/upload
		<br /> 请求参数说明：
	</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">pid</td>
				<td align="center">long</td>
				<td align="center">是</td>
				<td align="center">用户的嗒嗒ID</td>
				<td align="center">1000002</td>
			</tr>
			<tr>
				<td align="left">files</td>
				<td align="center">file</td>
				<td align="center">是</td>
				<td align="center">上传文件的二进制流</td>
				<td align="center">--</td>
			</tr>
		</tbody>
	</table>
	<h6>请求头Content-Type:multipart/form-data</h6>
	<p>返回成功示例：</p>
	<p>
		<img src="http://i.imgur.com/6Jfx4we.png" alt="aaa" />
	</p>
	<p>返回失败示例：</p>
	<p>
		<img src="http://i.imgur.com/T5uH6CN.png" alt="bbb" />
	</p>
	<p>返回参数说明：</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">code</td>
				<td align="center">int</td>
				<td align="center">是</td>
				<td align="center">状态码</td>
				<td align="center">200</td>
			</tr>
			<tr>
				<td align="left">result</td>
				<td align="center">string</td>
				<td align="center">是</td>
				<td align="center">上传文件的ID（用来放入群聊信息中）</td>
				<td align="center">--</td>
			</tr>
		</tbody>
	</table>
	<h4>
		25：<span id="5.25">下载文件</span>
	</h4>
	<p>
		请求方式：GET请求 接口描述：可以将存留在云端的文件下载到本地。 <br /> 接口地址：http://<strong><em>ip</em></strong>/chatting/file/download/(fileId)?fileName=nihao.zip
		<br /> 请求参数说明：
	</p>
	<table>
		<thead>
			<tr>
				<th align="left">参数名</th>
				<th align="center">类型</th>
				<th align="center">是否必填</th>
				<th align="center">描述</th>
				<th align="center">示例</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td align="left">fileId</td>
				<td align="center">string</td>
				<td align="center">是</td>
				<td align="center">文件的id，从群聊信息中获取</td>
				<td align="center">590c622242b0b881bc2f2648</td>
			</tr>
			<tr>
				<td align="left">fileName</td>
				<td align="center">string</td>
				<td align="center">是</td>
				<td align="center">上传文件的名称，从群聊信息中获取</td>
				<td align="center">nihao.zip</td>
			</tr>
		</tbody>
	</table>
	<p>返回成功示例：</p>
	<p>
		<img src="http://i.imgur.com/fsFI6Vz.png" alt="aaa" />
	</p>
	<p>
		<img src="http://i.imgur.com/wNjM3I2.png" alt="qq" />
	</p>
	<p>返回失败示例：</p>
	<p>
		<img src="http://i.imgur.com/T5uH6CN.png" alt="bbb" />
	</p>
	<p>返回参数说明：GET请求，成功则为下载文件。</p>
	<h4>
		26：<span id="5.26">版本更新</span>
	</h4>
	<p>
		请求方式：GET请求 接口描述：可以将最新的版本文件下载到本地。 <br /> 接口地址：http://<strong><em>ip</em></strong>/chatting/file/(action)/(device)
		<br />
	</p>
	<h6>acction:find或get；</h6>
	<h6>device:dacat或gacat或tomcat；</h6>
	<p>请求参数说明：无</p>
	<p>返回成功示例：</p>
	<p>
		<img src="http://i.imgur.com/wNjM3I2.png" alt="qq" />
	</p>
	<p>返回失败示例：</p>
	<p>
		<img src="http://i.imgur.com/T5uH6CN.png" alt="bbb" />
	</p>
	<p>返回参数说明：GET请求，成功则为下载文件。</p>
	<h2>
		六、<span id="6"> 正则文档 </span>
	</h2>
	<h5>验证手机：REGEX_CHINESE = &quot;^[\u4e00-\u9fa5],{0,}$&quot;;</h5>
	<h5>
		验证邮箱：REGEX_EMAIL = &quot;^\w+([-+.]\w+)<em>@\w+([-.]\w+)</em>\.[a-z]+([-.][a-z]+)*$&quot;;
	</h5>
	<h5>校验验证码：REGEX_CHECKCODE = &quot;^\d{6}$&quot;;</h5>
	<h5>校验密码：REGEX_PASSWORD = &quot;^[a-z0-9]{32}$&quot;;</h5>
	<h5>校验手机唯一识别码：REGEX_IMEI = &quot;(\d{15})|([0-9A-Z]{32})&quot;;</h5>
	<h5>校验真实姓名：REGEX_NAME =
		&quot;^[\u4e00-\u9fa5]+(·[\u4e00-\u9fa5]+)*$&quot;;</h5>
	<h5>校验身份证号：REGEX_IDCERD = &quot;^(\d{14}|\d{17})[0-9|x]{1}$&quot;;</h5>
	<h5>校验昵称：12个字节长度</h5>
	<p>
		<br />
	</p>
	<h3>
		<div class="rTop" id="rTop"
			style="margin-left: auto; margin-right: 2px; width: 45px; background-color: #E9E9E9;">
			<a href="#0"> 回到<br />顶部
			</a>
		</div>
	</h3>

</body>
</html>
