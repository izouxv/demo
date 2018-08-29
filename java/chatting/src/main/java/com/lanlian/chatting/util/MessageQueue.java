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
package com.lanlian.chatting.util;

import java.util.concurrent.BlockingQueue;
import java.util.concurrent.LinkedBlockingQueue;

import javax.jms.JMSException;
import javax.jms.Message;
import javax.jms.MessageListener;
import javax.jms.TextMessage;

import com.lanlian.chatting.logger.LogUtil;
import com.lanlian.chatting.vo.UserLetter;

/** 
 * @Title MessageQueue.java
 * @Package com.lanlian.chatting.util
 * @author 王东阳
 * @version V1.0
 * @date 2017年4月18日 下午6:46:48
 * @explain 信息队列;
 */

public class MessageQueue {
	
	/**
	 * 队列,默认时为int的max
	 */
	private static BlockingQueue<UserLetter> basket = new LinkedBlockingQueue<>(20000);
	
	/**
	 * 生产信息，放入队列
	 * @param pm
	 * @throws InterruptedException
	 */
	public static void produce(int pid, int opid, String type, String info) throws InterruptedException{
		System.err.println("=======信息进入队列！！！");
		UserLetter letter = new UserLetter();
		letter.setUid(DataFinals.DADA);
		letter.setTouid(opid);
		letter.setType(type);
		letter.setLetter(info);
		basket.put(letter);
	}
	
	/**
	 * 消费信息，从队列中取走
	 * @return
	 * @throws InterruptedException
	 */
	public static UserLetter consume() throws InterruptedException {
		// take方法取出一个消息，若basket为空，等到basket有消息为止(获取并移除此队列的头部)
		System.err.println("=======信息放出队列！！！");
		return basket.take();
	}
	
	
}

class QueueMessageListener implements MessageListener {
	@Override
    public void onMessage(Message message) {
        TextMessage tm = (TextMessage) message;
        try {
            System.out.println("QueueMessageListener监听到了文本消息：\t"
                    + tm.getText());
            //do something ...
        } catch (JMSException e) {
        	LogUtil.error(e);
        }
    }
}
