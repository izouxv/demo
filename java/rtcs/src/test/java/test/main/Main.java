/** 
 *<p>开发公司 :		           蓝涟科技<p>
 *<p>版权所有 :		           蓝涟科技<p>
 *<p>责任人     :		               王东阳<p> 
 *<p>网址         :    www.radacat.com<p>
 *<p>邮箱         : wangdy@radacat.com<p>
 */

package test.main;


import com.plys.rtcs.po.AbsException;
import com.plys.rtcs.po.Proto;
import com.plys.rtcs.po.ProtoMsg;

/**
 * @author  王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年10月18日 上午10:24:19
 * @$
 * @Administrator
 * @explain 
 */

public class Main {
	
	public static void main(String[] args) throws AbsException {
//		protoMsg.setHead(new ProtoHead());
		for (int i = 0; i < 3; i++) {
			ProtoMsg protoMsg = new ProtoMsg();
			System.out.println(i+":"+System.currentTimeMillis());
			String msg = "{" + 
					"\"head\":{"+
						"\"pv\":"+"\"1.1."+i+"\"," + 
						"\"client\":"+"\"12345"+i+"\"," + 
//						"\"type\":"+"1," + 
//						"\"subType\":"+"1," + 
//						"\"flage\":"+"1," + 
//						"\"scMulitca\":"+"1," + 
//						"\"desID\":"+"[1,2,3]," + 
//						"\"sourceID\":"+"10," + 
//						"\"length\":"+"1234," + 
					"}," + 
					"\"body\":{"+
						"\"aaa\":"+"\"aaa\"," + 
						"\"bbb\":"+"\"bbb\"," + 
						"\"ccc\":"+"\"ccc\"," + 
						"\"ddd\":"+"\"ddd\"," + 
						"\"eee\":"+"\"eee\"," + 
						"\"fff\":"+"\"fff\"," + 
						"\"ggg\":"+"\"ggg\"," + 
					"}" + 
				"}";
			
			protoMsg = (ProtoMsg)Proto.jsonToBean(msg , protoMsg);
			String string = Proto.beanToJson(protoMsg);
			System.out.println(i+":"+string);
			System.out.println(i+":"+System.currentTimeMillis());
			System.out.println(i+":"+protoMsg);
		}
	}


}

