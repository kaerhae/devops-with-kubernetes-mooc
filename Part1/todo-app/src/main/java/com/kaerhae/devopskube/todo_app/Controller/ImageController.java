package com.kaerhae.devopskube.todo_app.Controller;

import org.apache.catalina.connector.Response;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.core.io.UrlResource;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;

import com.kaerhae.devopskube.todo_app.Service.ImageService;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;


@Controller
public class ImageController {
    @Autowired
    private ImageService imageService;

    @GetMapping("/image/{filename}")
    public ResponseEntity<UrlResource> getImage(@PathVariable String filename) throws Exception {
        UrlResource imgUrl = imageService.GetImage(filename);
        if(imgUrl != null) {
            return ResponseEntity.ok().body(imgUrl);
        }

        return ResponseEntity.notFound().build();
    }
    
}
