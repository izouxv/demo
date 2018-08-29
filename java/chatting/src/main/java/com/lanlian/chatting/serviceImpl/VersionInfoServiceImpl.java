/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 * @author wangdyqxx
 * @email wangdy@radacat.com
 */

/**
 * 
 */
package com.lanlian.chatting.serviceImpl;

import java.io.File;
import java.text.SimpleDateFormat;
import java.util.Date;

import javax.annotation.Resource;
import javax.servlet.http.HttpServletResponse;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Propagation;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.web.multipart.MultipartFile;

import com.lanlian.chatting.dao.NewVersionInfoDao;
import com.lanlian.chatting.po.VersionPO;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.rpc.VersionClient;
import com.lanlian.chatting.service.VersionInfoService;
import com.lanlian.chatting.util.PublicMethod;
import com.lanlian.chatting.util.DataFinals.*;
import com.lanlian.chatting.vo.Version;
import com.lanlian.chatting.vo.Version2;
import com.lanlian.rpc.version.GetLatestVersionRequest.Builder;
import com.lanlian.rpc.version.GetLatestVersionResponse;

/**
 * @Title VersionInfoServiceImpl.java
 * @Package cn.lanlian.ccat.serviceImpl
 * @author 王东阳
 * @version V1.0
 * @date 2017年3月28日 下午3:50:57
 * @explain 用于存取版本文件的信息
 */
@Service("newVersionInfoService")
public class VersionInfoServiceImpl implements VersionInfoService {

	@Autowired
	NewVersionInfoDao newVersionInfoDao;
	
	@Resource
	VersionClient versionClient;
	
	public Builder getBuilder() {
		return versionClient.getBuilder();
	}

	/**
	 * 保存文件与信息
	 */
	@Transactional(propagation = Propagation.SUPPORTS, readOnly = true)
	@Override
	public void saveVersionInfo(Version version, MultipartFile file) throws Parameter_Exception {
		// 得到文件保存的路径
		SimpleDateFormat sdf = new SimpleDateFormat("yyyy_MM_dd_HH-mm");
		String time = sdf.format(new Date(System.currentTimeMillis()));
		String savePath = "/var/log/chatting/upload/" + version.getDevice() + File.separator;
		//request.getSession().getServletContext().getRealPath("E:\\CCAT_File" + random) + File.separator;
		// 得到上传的文件名
		String filename = file.getOriginalFilename();
//		System.out.println("目录:" + savePath);
		File tmpFile = new File(savePath);
		// 判断上传文件的保存目录是否存在
		if (!tmpFile.exists()) {
//			System.out.println(savePath + "目录不存在，创建目录-----");
			tmpFile.mkdirs();
		}
		File files = new File(filename);
		if (files.exists() && files.isFile()) {
			if (files.delete()) {
//				System.out.println("删除文件" + filename + "成功！");
			} else {
//				System.out.println("删除文件" + filename + "失败！");
			}
		}
		// 保存完整路径
		String path = savePath + time + "_" + filename;
		// 把文件上传至path的路径
		File localFile = new File(path);
		String md5 = "";
		try {
			file.transferTo(localFile);
			// 生成上传文件的MD5
			md5 = PublicMethod.getMd5ByFile(localFile);
//			System.out.println("md5:"+md5);
		} catch (Exception e) {
			throw new Parameter_Exception(10002);
		}
		if (version.getMd5() != null) {
			if (!md5.equals(version.getMd5())) {
				// MD5不一致
				throw new Parameter_Exception(21002);
			}
		}
		VersionPO ver = new VersionPO();
		ver.setDevice(version.getDevice());
		ver.setVersionName(version.getVersionName());
		ver.setVersionCode(version.getVersionCode());
		ver.setDescription(version.getDescription());
		ver.setDescription1(version.getDescription1());
		ver.setLength(version.getLength());
		ver.setMd5(md5);
		ver.setFilename(filename);
		ver.setPath(path);
//		ver.setAdminPid(version.getUid());
		int lines = newVersionInfoDao.saveVersionInfo(ver);
		if (lines != 1) {
			throw new Parameter_Exception(10002);
		}
	}

	/**
	 * 获取文件与信息
	 */
	@Transactional(propagation = Propagation.SUPPORTS, readOnly = true)
	public Version findVersionInfo(String action, Version version, HttpServletResponse response)
			throws Parameter_Exception {
		VersionPO verpo = new VersionPO();
		verpo.setDevice(version.getDevice());
		verpo = newVersionInfoDao.findVersionInfo(verpo);
		if (verpo == null || verpo.getStatus() != 1) {
			return null;
		}
		version.setDevice(verpo.getDevice());
		version.setVersionName(verpo.getVersionName());
		version.setVersionCode(verpo.getVersionCode());
		version.setFilename(verpo.getFilename());
		version.setMd5(verpo.getMd5());
		SimpleDateFormat sdf = new SimpleDateFormat("yyyy/MM/dd");
		version.setTime(sdf.format(verpo.getTime()));
		version.setDescription(verpo.getDescription());
		version.setDescription1(verpo.getDescription1());
		version.setLength(verpo.getLength());
		if (Actions.equals(0, action)) {
			return version;
		}
		// 开始下载文件
		if (Actions.equals(1, action)) {
			PublicMethod.downloadFile(response, verpo.getFilename(),verpo.getPath(), 1);
		}
		// 返回结果
		return version;
	}

	/**
	 * 获取新版本信息
	 */
	public Version2 findNewVersion(Builder builder) throws Parameter_Exception {
		GetLatestVersionResponse versionResponse = versionClient.getNewVersion(builder);
		Version2 version = new Version2();
		if (versionResponse.getErrorCode() != 10000) {
			if (versionResponse.getErrorCode() == 37031) {
				throw new Parameter_Exception(21008);
			}
			throw new Parameter_Exception(32001);
		}
		version.setDevice(versionResponse.getVersion().getDevice());
		version.setVersionName(versionResponse.getVersion().getVersionName());
		version.setVersionCode(versionResponse.getVersion().getVersionCode());
		version.setFilename(versionResponse.getVersion().getFileName());
		version.setMd5(versionResponse.getVersion().getMd5());
		version.setUrl(versionResponse.getVersion().getPath());
		SimpleDateFormat sdf = new SimpleDateFormat("yyyy/MM/dd");
		version.setTime(sdf.format(versionResponse.getVersion().getCreateTime() * 1000));
		version.setDescriptionCn(versionResponse.getVersion().getDescriptionCn());
		version.setDescriptionEn(versionResponse.getVersion().getDescriptionEn());
		version.setLength(versionResponse.getVersion().getLength());
		return version;
	}


}
