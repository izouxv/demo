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
package com.lanlian.chatting.controller;

import java.io.File;
import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;
import java.net.URLEncoder;
import java.sql.Timestamp;
import java.text.SimpleDateFormat;
import java.util.Date;

import javax.annotation.Resource;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.apache.log4j.Logger;
import org.bson.types.ObjectId;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.multipart.MultipartFile;

import com.lanlian.chatting.logger.LogUtil;
import com.lanlian.chatting.mongo.Mongo;
import com.lanlian.chatting.po.AdverPo;
import com.lanlian.chatting.po.AdvertisementPo;
import com.lanlian.chatting.po.SsoPo;
import com.lanlian.chatting.result.JSON;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.service.FileService;
import com.lanlian.chatting.util.Base64s;
import com.lanlian.chatting.util.ParameterVerify;
import com.lanlian.chatting.util.PublicMethod;
import com.lanlian.chatting.util.RequestSetting;
import com.mongodb.MongoGridFSException;
import com.mongodb.client.gridfs.GridFSBucket;
import com.mongodb.client.gridfs.GridFSBuckets;
import com.mongodb.client.gridfs.GridFSFindIterable;

/**
 * @Title FileController.java
 * @Package com.lanlian.chatting.controller
 * @author 王东阳
 * @version V1.0
 * @date 2017年3月27日 下午8:47:55
 * @explain
 */
@Controller
@RequestMapping(value = RequestSetting.FILE_PARENT_F, produces = RequestSetting.PRODUCES)
public class FilesController extends MyAbstractController {

	private static Logger log = Logger.getLogger(FilesController.class);

	@Resource(name = "fileService")
	FileService fileService;

	/**
	 * mongo上传文件
	 * 
	 * @param uid
	 * @param device
	 * @param imei
	 * @param files
	 * @return
	 * @throws Parameter_Exception
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.FILE_BODY_UPLOAD, method = RequestMethod.POST)
	public String uploadFile(@RequestParam(value = "uid", required = true) int uid,
			@RequestParam(value = "device", required = true) String device,
			@RequestParam(value = "imei", required = true) String imei,
			@RequestParam(value = "files", required = true) MultipartFile files) throws Parameter_Exception {
		InputStream ins = null;
		// for (MultipartFile file : files) {
		if (!(files instanceof MultipartFile)) {
			throw new Parameter_Exception(21002);
		}
		try {
			ins = files.getInputStream();
		} catch (IOException e) {
			throw new Parameter_Exception(10002);
		}
		// }
		GridFSBucket bucket = GridFSBuckets.create(Mongo.connect());
		// String json = JSONObject.toJSON(pid).toString();
		// BsonValue id = BsonDocument.parse(json);
		// bucket.uploadFromStream(id, "filename", ins);
		String fileName = files.getOriginalFilename();
		ObjectId objectId = bucket.uploadFromStream(uid + "_" + device + "_" + imei + "_" + fileName, ins);
		LogUtil.info("objectId:" + objectId.toString() + ",id:" + uid + "_" + device + "_" + imei + "_" + fileName);
		return JSON.toJson(objectId.toString());
	}

	/**
	 * mongo下载文件
	 * 
	 * @param device
	 * @param action
	 * @param response
	 * @return
	 * @throws Parameter_Exception
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.FILE_BODY_DOWNLOAD, method = RequestMethod.GET)
	public String download(@PathVariable String id, String name, HttpServletResponse response)
			throws Parameter_Exception {
		LogUtil.info("download-id:" + id + ",filename:" + name);
		// 下载文件
		downloadFile(id, name, response);
		response.setStatus(404);
		return null;
	}

	/**
	 * app回传日志文件
	 * 
	 * @param uid
	 * @param device
	 * @param imei
	 * @param files
	 * @param request
	 * @return
	 * @throws Parameter_Exception
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.FILE_BODY_UPLOAD_LOG, method = RequestMethod.POST)
	public String uploadFileLog(@PathVariable(value = "source", required = true) String source,
			@RequestParam(value = "files", required = true) MultipartFile files, HttpServletRequest request)
			throws Parameter_Exception {
		try {
			LogUtil.info(">>>>> upload_log >>>>source:" + "source");
			// 此处单文件MultipartFile
			// System.out.println(files.getSize());
			if (files == null || files.isEmpty()) {
				// System.err.println("文件:" + files + ",为空!");
				return JSON.toJson(21002);
			}
			SimpleDateFormat sdf = new SimpleDateFormat("yyyy-MM-dd-HH-mm-ss");
			String time = sdf.format(new Date(System.currentTimeMillis()));
			// 得到文件保存的路径 "/usr/share/tomcat/upload/logs/"
			String savePath = "/var/log/chatting/app/" + source + File.separator;
			// String savePath = request.getSession().getServletContext().getRealPath("/") +
			// File.separator;
			File tmpFile = new File(savePath);
			// 判断上传文件的保存目录是否存在
			if (!tmpFile.exists()) {
				System.out.println(savePath + "目录不存在，创建目录-----");
				tmpFile.mkdirs();
			}
			// 得到上传的文件名
			String filename = files.getOriginalFilename();
			// 保存完整路径
			String path = savePath + "_" + time + "_" + filename;
			// 查看文件上传路径,方便查找
			// 把文件上传至path的路径
			files.transferTo(new File(path));
			LogUtil.info(JSON.toJson());
			return JSON.toJson();
		} catch (IOException e) {
			LogUtil.error(e);
			return JSON.toJson(10002);
		}
	}

	/**
	 * 上传广告
	 * 
	 * @param code
	 * @param name
	 * @param topic
	 * @param file
	 * @return
	 * @throws Parameter_Exception
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.FILE_BODY_ADVERTISEMENT, method = RequestMethod.POST)
	public String uploadAdvertisement(@PathVariable(value = "name", required = true) String name,
			@RequestParam(value = "md5", required = true) String md5,
			@RequestParam(value = "advertiseUrl", required = true) String advertiseUrl,
			@RequestParam(value = "start", required = true) Long start,
			@RequestParam(value = "end", required = true) Long end,
			@RequestParam(value = "code", required = true) Integer code,
			@RequestParam(value = "file", required = true) MultipartFile file, HttpServletRequest request,
			HttpServletResponse response) throws Parameter_Exception {
		if (code != 123456789) {
			return JSON.toJson();
		}
		ParameterVerify.verifyNull(name, md5);
		if (file.isEmpty()) {
			throw new Parameter_Exception(21001);
		}
		SsoPo ssoPo = getContextSsoPo(request, response);
		Timestamp timeNow = getTime();
		SimpleDateFormat sdf = new SimpleDateFormat("yyyy_MM_dd_HH_mm");
		String time = sdf.format(timeNow);
		String savePath = "/var/log/chatting/advertisement/" + name + "/";// + File.separator;
		String path = PublicMethod.saveFile(savePath, name, md5, file, time);
		AdvertisementPo advertisementPo = new AdvertisementPo(name.trim(), ssoPo.getSource(), md5.trim(),
				Base64s.encryptBASE64(path), PublicMethod.encode(advertiseUrl), new Timestamp(start),
				new Timestamp(end), timeNow, timeNow, 1);
		log.info("FilesController-uploadAdvertisement:" + advertisementPo);
		fileService.setAdvertisement(advertisementPo);
		return JSON.toJson();
	}

	/**
	 * 获取新广告信息
	 * 
	 * @param request
	 * @param response
	 * @return
	 * @throws Parameter_Exception
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.FILE_BODY_GET_ADVER, method = RequestMethod.GET)
	public String getNewAdver(HttpServletRequest request, HttpServletResponse response) throws Parameter_Exception {
		log.info("getNewAdver......");
		SsoPo ssoPo = getContextSsoPo(request, response);
		AdverPo adverPo = new AdverPo();
		adverPo.setSource(ssoPo.getSource());
		return JSON.toJson(fileService.getNewAdver(adverPo));
	}

	/**
	 * 获取广告信息
	 * 
	 * @param data
	 * @param request
	 * @param response
	 * @return
	 * @throws Parameter_Exception
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.FILE_BODY_GET_ADVERTISEMENT, method = RequestMethod.GET)
	public String getAdvertisementInfo(HttpServletRequest request, HttpServletResponse response)
			throws Parameter_Exception {
		log.info("getAdvertisementInfo......");
		SsoPo ssoPo = getContextSsoPo(request, response);
		AdvertisementPo advertisementPo = new AdvertisementPo();
		advertisementPo.setSource(ssoPo.getSource());
		advertisementPo = fileService.getAdvertisement(advertisementPo);
		if (advertisementPo.getCode() != 10000) {
			JSON.toJson(21008);
		}
		return JSON.toJson(advertisementPo);
	}

	/**
	 * 下载广告
	 * 
	 * @param data
	 * @param request
	 * @param response
	 * @return
	 * @throws Parameter_Exception
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.FILE_BODY_DOWNLOAD_ADVERTISEMENT, method = RequestMethod.GET)
	public String downAdvertisement(@PathVariable(value = "file", required = true) String file,
			HttpServletRequest request, HttpServletResponse response) throws Parameter_Exception {
		ParameterVerify.verifyNull(file);
		SsoPo ssoPo = getContextSsoPo(request, response);
		log.info("downAdvertisement-ssoPo:" + ssoPo);
		PublicMethod.downloadFile(response, getTimes().toString() + ".jpg", Base64s.decryptBASE64(file), 1);
		return null;
	}

	/***********************************************************************************************/

	/**
	 * 根据文件id，使用mongo下载下载文件
	 * 
	 * @param fileId
	 * @param fileName
	 * @param response
	 * @throws Parameter_Exception
	 */
	private void downloadFile(String fileId, String fileName, HttpServletResponse response) throws Parameter_Exception {
		// 创建输出流
		OutputStream os = null;
		try {
			os = response.getOutputStream();
			// 创建一个容器，传入一个MongoDatabase类实例db
			GridFSBucket bucket = GridFSBuckets.create(Mongo.connect());
			GridFSFindIterable aaa = bucket.find();
			LogUtil.info("aaa:" + aaa.first().toString());
			bucket.downloadToStream(new ObjectId(fileId), os);
			// 下载
			response.reset();
			response.setHeader("Content-Disposition", "attachment;filename=" + URLEncoder.encode(fileName, "utf-8"));
			response.setCharacterEncoding("text/plain;charset=utf-8");
			response.setContentType("text/plain;charset=utf-8");
		} catch (MongoGridFSException e) {
			response.reset();
			response.setContentType("application/json;charset=utf-8");
			response.setCharacterEncoding("application/json;charset=utf-8");
			// System.out.println("Content-Disposition:"+response.getHeader("Content-Disposition"));
			throw new Parameter_Exception(21007);
		} catch (IOException e) {
			throw new Parameter_Exception(10002);
		} catch (Exception e) {
			throw new Parameter_Exception(10002);
		} finally {
			// 这里主要关闭。
			try {
				if (os != null) {
					os.flush();
					os.close();
				}
			} catch (IOException e) {
				LogUtil.error(e);
			}
		}
	}

}
