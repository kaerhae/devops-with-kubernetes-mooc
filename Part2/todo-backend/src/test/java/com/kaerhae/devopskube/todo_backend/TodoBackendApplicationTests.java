package com.kaerhae.devopskube.todo_backend;

import static org.junit.jupiter.api.Assertions.assertNotNull;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import com.kaerhae.devopskube.todo_backend.Controller.TodoController;

@SpringBootTest
class TodoBackendApplicationTests {
	@Autowired
	private TodoController todoController;

	@Test
	void contextLoads() {
		assertNotNull(todoController);
	}

}
