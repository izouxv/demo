/** 
 *<p>开发公司：	鹏联优思 <p>
 *<p>版权所有：	鹏联优思 <p>
 *<p>责任人：	王东阳    <p> 
 *<p>网址：www.penslink.com <p>
 */

package baidusdk;

import java.util.HashMap;
import java.util.Map;

import org.json.JSONObject;

import com.baidu.aip.speech.AipSpeech;

/**
 * @author  wangdyq
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2018年4月3日 上午10:44:24
 * @explain 
 */

public class Sample {
	//设置APPID/AK/SK
    public static final String APP_ID = "11039452";
    public static final String API_KEY = "8mEqgxHppCR5o1cdQCZyeLHY";
    public static final String SECRET_KEY = "a016b9a5c4489989c68f304c3803e2c0";
    public static void main(String[] args) {
        // 初始化一个AipSpeech
    	System.out.println(System.currentTimeMillis());
        AipSpeech client = new AipSpeech(APP_ID, API_KEY, SECRET_KEY);
        // 可选：设置网络连接参数
        client.setConnectionTimeoutInMillis(2000);
        client.setSocketTimeoutInMillis(60000);
        // 可选：设置代理服务器地址, http和socket二选一，或者均不设置
//        client.setHttpProxy("proxy_host", proxy_port);  // 设置http代理
//        client.setSocketProxy("proxy_host", proxy_port);  // 设置socket代理
        // 可选：设置log4j日志输出格式，若不设置，则使用默认配置
        // 也可以直接通过jvm启动参数设置此环境变量
//        System.setProperty("aip.log4j.conf", "log4j.properties");
        //可选参数
        HashMap<String, Object> options = new HashMap<String, Object>();
        options.put("dev_pid", "1536");
        // 调用接口
        System.out.println("---"+System.currentTimeMillis());
        for (int i = 0; i < 3; i++) {
            JSONObject res = client.asr("E:\\1.wav", "wav", 16000, options);
            System.out.println(res.toString(1));
            System.out.println(System.currentTimeMillis());
		}
    }
}
