/** 
 *<p>开发公司 :		           蓝涟科技<p>
 *<p>版权所有 :		           蓝涟科技<p>
 *<p>责任人     :		               王东阳<p> 
 *<p>网址         :    www.radacat.com<p>
 *<p>邮箱         : wangdy@radacat.com<p>
 */

package testVoice;

import java.io.File;
import java.io.InputStreamReader;
import java.io.LineNumberReader;
import java.util.logging.Logger;

import java.io.FileInputStream;
import java.io.IOException;
import java.util.ArrayList;
import java.util.Scanner;

import com.iflytek.cloud.speech.RecognizerListener;
import com.iflytek.cloud.speech.RecognizerResult;
import com.iflytek.cloud.speech.Setting;
import com.iflytek.cloud.speech.SpeechConstant;
import com.iflytek.cloud.speech.SpeechError;
import com.iflytek.cloud.speech.SpeechRecognizer;
import com.iflytek.cloud.speech.SpeechUtility;

/**
 * @author 王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2018年4月2日 上午11:41:37
 * @wangdyq
 * @explain 在视频文件中提取音频文件，调用科大讯飞jar，返回文字识别
 */

public class TestFile {

	private static Logger logger = Logger.getAnonymousLogger();

	public static void main(String[] args) {
		try {
			transform(args[0]);
		} catch (Exception e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
	}

	/**
	 * 提取视频中的音频文件
	 * 
	 * @param fileName
	 * @return 音频文件名
	 * @throws Exception
	 */
	public static String transform(String fileName) throws Exception {

		File file = new File(fileName);
		if (!file.exists()) {
			logger.info("文件不存在：" + fileName);
			throw new RuntimeException("文件不存在：" + fileName);
		}

		// 讯飞现在支持pcm，wav的语音流文件
		String name = fileName.substring(0, fileName.lastIndexOf(".")) + ".wav";
		logger.info("获取到的音频文件：" + name);

		// 提取视频中的音频文件。 根据讯飞要求设置采样率， 位数，
		String cmd = "ffmpeg -i " + fileName + " -f s16le -ar 16000 " + name;
		Process process = Runtime.getRuntime().exec(cmd);// 执行命令

		//
		InputStreamReader ir = new InputStreamReader(process.getInputStream());
		LineNumberReader input = new LineNumberReader(ir);
		String line;
		// 输出结果，需要有这部分代码， 否则不能生产抽取的音频文件
		while ((line = input.readLine()) != null) {
			System.out.println(line);
		}
		process.destroy();
		return name;
	}
}

class Test {
	private static String fileName = "1481023006148.wav";
	public static void main(String[] args) throws InterruptedException {
		Scanner scanner = new Scanner(System.in);
		MscTests mObject = new MscTests();
		int a = 1;
		boolean flag = true;
		for (int i = 0; i < 200; i++) {
			// MscTest mObject = new MscTest();
			if (flag) {
				System.out.print("place input num:");
				int num = scanner.nextInt();
				if (num > 10) {
					flag = false;
				}
			}
			// 语音识别， 并给mResult赋值, 在获取音频文件中的文字代码中添加while就是为了确保该行代码执行完成时，
			// 语音识别解析工作已经完成，否则可能获取不到识别结果， 或仅仅是获取到识别结果的一部分
			String result = mObject.voice2words(fileName);
			System.out.println(a + "--->MscTest.main()--" + result);
			a++;
		}
	}
}

class MscTests {

	private static final String APPID = "5860ec57";

	private StringBuffer mResult = new StringBuffer();

	/** 最大等待时间， 单位ms */
	private int maxWaitTime = 500;
	/** 每次等待时间 */
	private int perWaitTime = 100;
	/** 出现异常时最多重复次数 */
	private int maxQueueTimes = 3;
	/** 音频文件 */
	private String fileName = "";

	static {
		Setting.setShowLog(false);
		SpeechUtility.createUtility("appid=" + APPID);
	}

	public String voice2words(String fileName) throws InterruptedException {
		return voice2words(fileName, true);
	}

	/**
	 * 
	 * @desc: 工具类，在应用中有一个实例即可， 但是该实例是有状态的， 因此要消除其他调用对状态的修改，所以提供一个init变量
	 * @auth: zona 2017年1月4日 下午4:38:45
	 * @param fileName
	 * @param init
	 *            是否初始化最大等待时间。
	 * @return
	 * @throws InterruptedException
	 */
	public String voice2words(String fileName, boolean init) throws InterruptedException {
		if (init) {
			maxWaitTime = 500;
			maxQueueTimes = 3;
		}
		if (maxQueueTimes <= 0) {
			mResult.setLength(0);
			mResult.append("解析异常！");
			return mResult.toString();
		}
		this.fileName = fileName;

		return recognize();
	}

	// *************************************音频流听写*************************************

	/**
	 * 听写
	 * 
	 * @return
	 * @throws InterruptedException
	 */
	private String recognize() throws InterruptedException {
		if (SpeechRecognizer.getRecognizer() == null)
			SpeechRecognizer.createRecognizer();
		return RecognizePcmfileByte();
	}

	/**
	 * 自动化测试注意要点 如果直接从音频文件识别，需要模拟真实的音速，防止音频队列的堵塞
	 * 
	 * @throws InterruptedException
	 */
	private String RecognizePcmfileByte() throws InterruptedException {
		// 1、读取音频文件
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
			// 解析之前将存出结果置为空
			mResult.setLength(0);
			SpeechRecognizer recognizer = SpeechRecognizer.getRecognizer();
			recognizer.setParameter(SpeechConstant.DOMAIN, "iat");
			recognizer.setParameter(SpeechConstant.LANGUAGE, "zh_cn");
			recognizer.setParameter(SpeechConstant.AUDIO_SOURCE, "-1");
			// 写音频流时，文件是应用层已有的，不必再保存
			// recognizer.setParameter(SpeechConstant.ASR_AUDIO_PATH,
			// "./iflytek.pcm");
			recognizer.setParameter(SpeechConstant.RESULT_TYPE, "plain");
			recognizer.startListening(recListener);
			ArrayList<byte[]> buffers = splitBuffer(voiceBuffer, voiceBuffer.length, 4800);
			for (int i = 0; i < buffers.size(); i++) {
				// 每次写入msc数据4.8K,相当150ms录音数据
				recognizer.writeAudio(buffers.get(i), 0, buffers.get(i).length);
				try {
					Thread.sleep(150);
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
			mResult.append(result.getResultString());
		}

		public void onError(SpeechError error) {
			try {
				voice2words(fileName);
				maxQueueTimes--;
			} catch (InterruptedException e) {
				Thread.currentThread().interrupt();
				throw new RuntimeException(e);
			}
		}

		public void onEvent(int eventType, int arg1, int agr2, String msg) {
		}

	};

}
