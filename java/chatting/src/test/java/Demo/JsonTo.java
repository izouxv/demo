package Demo;

import java.io.IOException;
import java.util.List;

import org.codehaus.jackson.map.ObjectMapper;
import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;

/**
 * 实体类和JSON对象之间相互转化（依赖包jackson-all-1.7.6.jar、jsoup-1.5.2.jar）
 * 
 * @author wck
 *
 */
public class JsonTo {
	/**
	 * 将json转化为实体POJO
	 * 
	 * @param jsonStr
	 * @param obj
	 * @return
	 */
	public static <T> Object JSONToObj(String jsonStr, Class<T> obj) {
		T t = null;
		try {
			ObjectMapper objectMapper = new ObjectMapper();
			t = objectMapper.readValue(jsonStr, obj);
		} catch (Exception e) {
			e.printStackTrace();
		}
		return t;
	}

	/**
	 * 将实体POJO转化为JSON
	 * 
	 * @param obj
	 * @return
	 * @throws JSONException
	 * @throws IOException
	 */
	public static <T> JSONObject objectToJson(T obj) throws JSONException, IOException {
		ObjectMapper mapper = new ObjectMapper();
		// Convert object to JSON string
		String jsonStr = "";
		try {
			jsonStr = mapper.writeValueAsString(obj);
		} catch (IOException e) {
			throw e;
		}
		return new JSONObject(jsonStr);
	}

	public static void main(String[] args) throws JSONException, IOException {
		JSONArray array = new JSONArray();
		JSONObject obj = null;

		obj = new JSONObject();
		obj.put("name", "213");
		obj.put("age", 27);
		array.put(obj);
		
		obj = new JSONObject();
		obj.put("name", "214");
		obj.put("age", 28);
		array.put(obj);
		
		Student stu = (Student) JSONToObj(obj.toString(), Student.class);
		JSONObject objList = new JSONObject();
		objList.put("student", array);
		System.out.println("objList:" + objList);
		StudentList stuList = (StudentList) JSONToObj(objList.toString(), StudentList.class);
		System.out.println("student:" + stu);
		System.out.println("stuList:" + stuList);
		System.out.println("#####################################");
		JSONObject getObj = objectToJson(stu);
		System.out.println(getObj);
	}
}

class Student {
	private String name;
	private int age;

	// private StudentClass studentClass;
	public String getName() {
		return name;
	}

	public void setName(String name) {
		this.name = name;
	}

	public int getAge() {
		return age;
	}

	public void setAge(int age) {
		this.age = age;
	}

	@Override
	public String toString() {
		return "Student [name=" + name + ", age=" + age + "]";
	}
}

class StudentList {
	List<Student> student;

	public List<Student> getStudent() {
		return student;
	}

	public void setStudent(List<Student> student) {
		this.student = student;
	}

	@Override
	public String toString() {
		return "StudentList [student=" + student + "]";
	}
}