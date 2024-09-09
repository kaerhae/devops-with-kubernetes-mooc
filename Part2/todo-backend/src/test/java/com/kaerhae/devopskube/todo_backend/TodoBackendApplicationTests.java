package com.kaerhae.devopskube.todo_backend;

import static org.junit.jupiter.api.Assertions.assertNotNull;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.ActiveProfiles;
import org.springframework.test.context.TestPropertySource;

import com.kaerhae.devopskube.todo_backend.Controller.TodoController;

@SpringBootTest
@TestPropertySource(properties = {
    "spring.datasource.url=jdbc:h2:mem:testdb",
    "spring.datasource.driverClassName=org.h2.Driver",
    "spring.datasource.username=sa",
    "spring.datasource.password=password"
})
class TodoBackendApplicationTests {
	@Autowired
	private TodoController todoController;
	@Test
	void contextLoads() {
		assertNotNull(todoController);
	}

}
