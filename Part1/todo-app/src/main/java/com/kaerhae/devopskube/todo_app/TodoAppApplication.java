package com.kaerhae.devopskube.todo_app;


import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.ApplicationArguments;
import org.springframework.boot.ApplicationRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.scheduling.annotation.EnableScheduling;
import org.springframework.stereotype.Component;

import com.kaerhae.devopskube.todo_app.Service.ImageService;
import com.kaerhae.devopskube.todo_app.Task.ScheduledImageTask;

@SpringBootApplication
@EnableScheduling
public class TodoAppApplication {
	private static final Logger logger = LoggerFactory.getLogger(TodoAppApplication.class);

	@Component
	public class AppStartupRunner implements ApplicationRunner {
		@Autowired
		private ImageService imageService;
		private ScheduledImageTask scheduledImageTask;
		@Override
		public void run(ApplicationArguments args) throws Exception {
			if(System.getenv("PORT") == null || System.getenv("IMAGE_FILE_PATH") == null) {
				throw new Exception("Environment variables missing");
			}

			/* Check if image exists */
			boolean doesExists = imageService.CheckIfImageExists();
			if(!doesExists) {
				/* Run image service task, only if no previous */
				scheduledImageTask.FetchNewImageToDestination();
			}
		}
	}
	public static void main(String[] args) throws Exception {
		SpringApplication.run(TodoAppApplication.class, args);
		
		logger.info(String.format("Server started in port %s", System.getenv("PORT")));
	}

}
