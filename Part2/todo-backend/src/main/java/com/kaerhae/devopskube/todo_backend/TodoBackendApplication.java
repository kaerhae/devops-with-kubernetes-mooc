package com.kaerhae.devopskube.todo_backend;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.ApplicationArguments;
import org.springframework.boot.ApplicationRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.stereotype.Component;

@SpringBootApplication
public class TodoBackendApplication {
	private static final Logger logger = LoggerFactory.getLogger(TodoBackendApplication.class);

	@Component
	public class AppStartupRunner implements ApplicationRunner {
		@Override
		public void run(ApplicationArguments args) throws Exception {
			if(System.getenv("PORT") == null) {
				throw new Exception("Environment variables missing");
			}
		}
	}
	public static void main(String[] args) throws Exception {
		SpringApplication.run(TodoBackendApplication.class, args);
		
		logger.info(String.format("Server started in port %s", System.getenv("PORT")));
	}

}
