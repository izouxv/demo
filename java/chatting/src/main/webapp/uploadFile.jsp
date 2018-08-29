<%@ page language="java" contentType="text/html; charset=utf-8"
	pageEncoding="utf-8"%>
<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<title>文件上传</title>
</head>

<body>
	<form action="${pageContext.request.contextPath}/file/upload"
		enctype="multipart/form-data" method="POST">
		上传用户：<input type="text" name="username"><br /> 上传文件1：<input
			type="file" name="files"><br /> 上传文件2：<input type="file"
			name="files"><br /> <input type="submit" value="提交...">
	</form>
</body>
</html>