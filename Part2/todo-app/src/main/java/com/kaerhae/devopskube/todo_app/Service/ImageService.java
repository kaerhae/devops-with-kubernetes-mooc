package com.kaerhae.devopskube.todo_app.Service;

import java.awt.image.BufferedImage;
import java.io.File;
import java.net.URL;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;

import javax.imageio.ImageIO;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.core.io.UrlResource;
import org.springframework.stereotype.Service;

@Service
public class ImageService {
    	private static final Logger logger = LoggerFactory.getLogger(ImageService.class);


    public void RetrieveImageToFile() {
        String url = "https://picsum.photos/1200";
        try {
            URL imageUrl = new URL(url);
            BufferedImage image = ImageIO.read(imageUrl);
            ImageIO.write(image, "jpg", new File(String.format("%s/image.jpg", System.getenv("IMAGE_FILE_PATH"))));
        } catch (Exception e) {
            logger.error("Error while retrieving image to file", e);
        }
    }

    public UrlResource GetImage(String filename) throws Exception {
        /* Get image from file path and create UrlResource to serve it */
        try {
            Path file = Paths.get(System.getenv("IMAGE_FILE_PATH")).resolve(filename);
            UrlResource resource = new UrlResource(file.toUri());

            if (resource.exists() || resource.isReadable()) {
                return resource;
            } else {
                logger.info("Image not found");
                return null;
            }
        } catch (Exception e) {
            logger.error("Error on fetching image", e);
            return null;
        }
    }

    public boolean CheckIfImageExists() {
        Path path = Paths.get(String.format("%s/image.jpg", System.getenv("IMAGE_FILE_PATH")));
        return Files.exists(path);
    }
}
