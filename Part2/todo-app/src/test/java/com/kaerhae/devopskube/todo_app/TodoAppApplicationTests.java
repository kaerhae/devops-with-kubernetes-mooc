package com.kaerhae.devopskube.todo_app;

import static org.junit.jupiter.api.Assertions.assertNotNull;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.info.BuildProperties;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.TestPropertySource;

import com.kaerhae.devopskube.todo_app.Controller.ImageController;
import com.kaerhae.devopskube.todo_app.Controller.IndexController;

@SpringBootTest
@TestPropertySource(properties = { "server.url=9000" })
class TodoAppApplicationTests {

	static {
		System.setProperty("PORT", "9000");
		System.setProperty("TODO_BACKEND_BASEURL", "http://foobar");
		System.setProperty("IMAGE_FILE_PATH", "/opt");
	}
	@Autowired
	private IndexController indexController;
	@Autowired
	private ImageController imageController;
	@Autowired
	private BuildProperties buildProperties;
	
	@Test
	public void contextLoads() throws Exception {
		assertNotNull(indexController);
		assertNotNull(imageController);
	}

}
