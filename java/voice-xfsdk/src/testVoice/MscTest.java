/** 
 *<p>开发公司 :		           蓝涟科技<p>
 *<p>版权所有 :		           蓝涟科技<p>
 *<p>责任人     :		               王东阳<p> 
 *<p>网址         :    www.radacat.com<p>
 *<p>邮箱         : wangdy@radacat.com<p>
 */

package testVoice;

import java.io.File;
import java.io.FileInputStream;
import java.io.IOException;
import java.util.ArrayList;
import java.util.Scanner;

import com.iflytek.cloud.speech.DataUploader;
import com.iflytek.cloud.speech.LexiconListener;
import com.iflytek.cloud.speech.RecognizerListener;
import com.iflytek.cloud.speech.RecognizerResult;
import com.iflytek.cloud.speech.Setting;
import com.iflytek.cloud.speech.SpeechConstant;
import com.iflytek.cloud.speech.SpeechError;
import com.iflytek.cloud.speech.SpeechListener;
import com.iflytek.cloud.speech.SpeechRecognizer;
import com.iflytek.cloud.speech.SpeechSynthesizer;
import com.iflytek.cloud.speech.SpeechUtility;
import com.iflytek.cloud.speech.SynthesizeToUriListener;
import com.iflytek.cloud.speech.UserWords;

/**
 * @author 王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2018年4月2日 上午11:49:48
 * @wangdyq
 * @explain
 */

public class MscTest {
	
	/**
	 * 音频文件格式：采样率 16k,采样精度16bit,识别的音频长度最大为60S
	 * @param args
	 * @throws InterruptedException
	 */
	public static void main(String[] args) throws InterruptedException {
		MscTest mObject = new MscTest();
		mObject.recognize();
		//读取语音文件
		System.out.println(System.currentTimeMillis());
		for (int i = 0; i < 1; i++) {
			mObject.fileName="E:\\1.wav";
			String result = mObject.RecognizePcmfileByte();
			System.out.println("--->Msc:" + result);
			System.out.println(System.currentTimeMillis());
		}
		
		//上传用户词表
//		mObject.json="[{\"name\": \"111\",\"words\":\"你好\",\"你好好\"]},{\"name\": \"222\",\"words\": [\"喂\",\"喂喂\"]}]";
//		SpeechRecognizer.createRecognizer();
//		mObject.uploadUserWords();
	}

	private static final String APPID = "5ac18c2d";
	private StringBuffer mResult = new StringBuffer();
	/** 最大等待时间， 单位ms */
	private int maxWaitTime = 2000;
	/** 每次等待时间 */
	private static int perWaitTime = 3;
	/** 音频文件 */
	public String fileName = "";

	static {
		Setting.setShowLog(false);
		SpeechUtility.createUtility("appid=" + APPID);
	}
	/**
	 * 创建识别类
	 * @return
	 * @throws InterruptedException
	 */
	public void recognize() {
		if (SpeechRecognizer.getRecognizer() == null)
			SpeechRecognizer.createRecognizer();
	}

	/**
	 * 自动化测试注意要点 如果直接从音频文件识别，需要模拟真实的音速，防止音频队列的堵塞
	 * 
	 * @throws InterruptedException
	 */
	public String RecognizePcmfileByte() throws InterruptedException {
		System.out.println("1、读取音频文件");
		FileInputStream fis = null;
		byte[] voiceBuffer = null;
		try {
			fis = new FileInputStream(new File(fileName));
			voiceBuffer = new byte[fis.available()];
			fis.read(voiceBuffer);
		} catch (Exception e) {
			e.printStackTrace();
		} finally {
			try {
				if (null != fis) {
					fis.close();
					fis = null;
				}
			} catch (IOException e) {
				e.printStackTrace();
			}
		}
		// 2、音频流听写
		if (0 == voiceBuffer.length) {
			mResult.append("no audio avaible!");
		} else {
			System.out.println("length:"+voiceBuffer.length);
			// 解析之前将存出结果置为空
			mResult.setLength(0);
			SpeechRecognizer recognizer = SpeechRecognizer.getRecognizer();
			recognizer.setParameter(SpeechConstant.DOMAIN, "iat");
			recognizer.setParameter(SpeechConstant.LANGUAGE, "zh_cn");
			recognizer.setParameter(SpeechConstant.AUDIO_SOURCE, "-1");
			recognizer.setParameter (SpeechConstant.ACCENT, "mandarin ");
			//音频流保存
//			recognizer.setParameter(SpeechConstant.ASR_AUDIO_PATH,"./iflytek.pcm");
			recognizer.setParameter(SpeechConstant.RESULT_TYPE, "plain");
			recognizer.startListening(recListener);
			ArrayList<byte[]> buffers = splitBuffer(voiceBuffer, voiceBuffer.length, 4800);
			for (int i = 0; i < buffers.size(); i++) {
				// 每次写入msc数据4.8K,相当150ms录音数据
				recognizer.writeAudio(buffers.get(i), 0, buffers.get(i).length);
				try {
					Thread.sleep(50);
				} catch (InterruptedException e) {
					e.printStackTrace();
				}
			}
			recognizer.stopListening();
			// 在原有的代码基础上主要添加了这个while代码等待音频解析完成，recognizer.isListening()返回true，说明解析工作还在进行
			while (recognizer.isListening()) {
				if (maxWaitTime < 0) {
					mResult.setLength(0);
					mResult.append("解析超时！");
					break;
				}
				Thread.sleep(perWaitTime);
				maxWaitTime -= perWaitTime;
			}
		}
		return mResult.toString();
	}

	/**
	 * 将字节缓冲区按照固定大小进行分割成数组
	 * 
	 * @param buffer
	 *            缓冲区
	 * @param length
	 *            缓冲区大小
	 * @param spsize
	 *            切割块大小
	 * @return
	 */
	private ArrayList<byte[]> splitBuffer(byte[] buffer, int length, int spsize) {
		ArrayList<byte[]> array = new ArrayList<byte[]>();
		if (spsize <= 0 || length <= 0 || buffer == null || buffer.length < length)
			return array;
		int size = 0;
		while (size < length) {
			int left = length - size;
			if (spsize < left) {
				byte[] sdata = new byte[spsize];
				System.arraycopy(buffer, size, sdata, 0, spsize);
				array.add(sdata);
				size += spsize;
			} else {
				byte[] sdata = new byte[left];
				System.arraycopy(buffer, size, sdata, 0, left);
				array.add(sdata);
				size += left;
			}
		}
		return array;
	}

	/**
	 * 听写监听器
	 */
	private RecognizerListener recListener = new RecognizerListener() {
		public void onBeginOfSpeech() {
		}
		public void onEndOfSpeech() {
		}
		public void onVolumeChanged(int volume) {
		}
		public void onResult(RecognizerResult result, boolean islast) {
//			System.out.println("result:"+result.getResultString());
			mResult.append(result.getResultString());
		}
		public void onError(SpeechError error) {
//			try {
				System.err.println("onError:"+error.getErrorCode());
//				voice2words(fileName);
//				maxQueueTimes--;
//			} catch (InterruptedException e) {
//				Thread.currentThread().interrupt();
//				throw new RuntimeException(e);
//			}
		}
		public void onEvent(int eventType, int arg1, int agr2, String msg) {
		}

	};

	private String json = "";
	
	private void uploadUserWords() {
	    SpeechRecognizer recognizer = SpeechRecognizer.getRecognizer();
	    UserWords userwords = new UserWords(json);
	    recognizer.setParameter(SpeechConstant.DATA_TYPE, "userword" );
	    recognizer.updateLexicon("userwords", userwords.toString(), lexiconListener);
	}
	
	/**
	 * 词表上传监听器
	 */
	LexiconListener lexiconListener = new LexiconListener() {
	    @Override
	    public void onLexiconUpdated(String lexiconId, SpeechError error) {
	    	System.out.println("lexiconId:"+lexiconId);
	        if (error == null)
	            System.out.println("*************上传成功*************");
	        else
	        	System.out.println("*************" + error.getErrorCode()+ "*************");
	    }
	};
}