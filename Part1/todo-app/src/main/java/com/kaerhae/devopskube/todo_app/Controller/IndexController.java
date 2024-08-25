package com.kaerhae.devopskube.todo_app.Controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.info.BuildProperties;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;

import com.kaerhae.devopskube.todo_app.Service.ImageService;



@Controller
public class IndexController {
    @Autowired
    private BuildProperties buildProperties;


    @GetMapping("")
    public String GetServerInfo(Model model) {
        //return String.format("REST API Server version %s", buildProperties.getVersion());
        model.addAttribute("VERSION", buildProperties.getVersion());
        model.addAttribute("IMAGE_URL", String.format("%s/image.jpg", System.getenv("IMAGE_FILE_PATH")));
        return "index.html";
    }
    
}
