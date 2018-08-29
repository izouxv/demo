/** 
 *<p>开发公司 :		           蓝涟科技<p>
 *<p>版权所有 :		           蓝涟科技<p>
 *<p>责任人     :		               王东阳<p> 
 *<p>网址         :    www.radacat.com<p>
 *<p>邮箱         : wangdy@radacat.com<p>
 */

package com.lanlian.chatting.serviceImpl;

import javax.annotation.Resource;

import org.springframework.stereotype.Service;

import com.lanlian.chatting.dao.FileDao;
import com.lanlian.chatting.po.AdverPo;
import com.lanlian.chatting.po.AdvertisementPo;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.rpc.AdverClient;
import com.lanlian.chatting.service.FileService;
import com.lanlian.chatting.util.PublicMethod;
import com.lanlian.rpc.adver.AdvertisementReply;
import com.lanlian.rpc.adver.AdvertisementRequest.Builder;

/**
 * @author  王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年11月15日 上午10:37:21
 * @Administrator
 * @explain 
 */
@Service(value="fileService")
public class FileServiceImpl implements FileService {

	@Resource
	FileDao fileDao;
	
	@Resource(name="adverClient")
	AdverClient adverClient;
	
	@Override
	public void setAdvertisement(AdvertisementPo advertisementPo) {
		fileDao.setAdvertisement(advertisementPo);
	}

	@Override
	public AdvertisementPo getAdvertisement(AdvertisementPo advertisementPo) throws Parameter_Exception {
		advertisementPo = fileDao.getAdvertisement(advertisementPo);
		if (advertisementPo == null) {
			throw new Parameter_Exception(21008);
		}
		advertisementPo.setAdvertiseUrl(PublicMethod.decode(advertisementPo.getAdvertiseUrl()));
		return advertisementPo;
	}
	
	@Override
	public AdverPo getNewAdver(AdverPo adverPo) throws Parameter_Exception {
		Builder builder = adverClient.getBuilder();
		builder.setSource(adverPo.getSource());
		AdvertisementReply adverRe = adverClient.getAdver(builder);
		if (adverRe.getErrorCode() == 37033) {
			throw new Parameter_Exception(21008);
		}
		if (adverRe.getErrorCode() != 10000) {
			throw new Parameter_Exception(10002);
		}
		adverPo.setCode(adverRe.getErrorCode());
		adverPo.setId(adverRe.getId());
		adverPo.setName(adverRe.getFileName());
		adverPo.setMd5(adverRe.getMd5());
		adverPo.setFileUrl(adverRe.getFileUrl());
		adverPo.setAdvertiseUrl(adverRe.getAdvUrl());
		adverPo.setStartTime(adverRe.getStartTime());
		adverPo.setEndTime(adverRe.getEndTime());
		return adverPo;
	}

	@Override
	public AdvertisementPo updateAdvertisement(AdvertisementPo advertisementPo) {
		return null;
	}

}

