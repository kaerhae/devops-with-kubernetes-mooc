package com.kaerhae.devopskube.todo_app.Task;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Component;

import com.kaerhae.devopskube.todo_app.Service.ImageService;

@Component
public class ScheduledImageTask {
    @Autowired
    private ImageService imageService;
    private static final Logger logger = LoggerFactory.getLogger(ScheduledImageTask.class);
    @Scheduled(cron = "@hourly")
    public void FetchNewImageToDestination() {
        imageService.RetrieveImageToFile();
        logger.info("New image retrieved");
    }
}
