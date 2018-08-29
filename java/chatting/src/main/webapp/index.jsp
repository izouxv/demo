<%@ page language="java" contentType="text/html; charset=UTF-8"
	pageEncoding="UTF-8"%>
<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
<title>文件上传</title>
</head>
<body>

	<form action="http://192.168.0.54/chatting/file/upload"
		method="POST" 
		enctype="multipart/form-data">
		账号：<input type="text" name="userName" id="userName" /> <br />
		<br /> 密码：<input type="password" name="password" id="password" /> <br />
		<br /> 选择文件：<input type="file" name="files" id="files" /> <br />
		<br /> <input type="submit" value="提交" />
	</form>
	
	<br/><br/><br/><br/><br/><br/>
	
	<form action="http://192.168.0.54/chatting/file/upload" 
		method="POST" 
		enctype="multipart/form-data">
		账号：<input type="text" name="pid" id="pid" /> <br />
		文件: <input type="file" name="files" />
		<br/>
		<input type="submit" value="上传"/>
	</form>
	<br/><br/><br/>
	<form action="http://192.168.0.54/chatting/file/put/dacat" 
		method="POST" 
		enctype="multipart/form-data">
		名字：<input type="text" name="pid" id="name" /> <br />
		文件: <input type="file" name="file" />
		<br/>
		<input type="submit" value="上传"/>
	</form>

</body>
</html>