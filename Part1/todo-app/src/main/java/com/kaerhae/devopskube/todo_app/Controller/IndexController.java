package com.kaerhae.devopskube.todo_app.Controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.info.BuildProperties;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;


@Controller
public class IndexController {
    @Autowired
    private BuildProperties buildProperties;
    @GetMapping("")
    public String GetServerInfo(Model model) {
        //return String.format("REST API Server version %s", buildProperties.getVersion());
        model.addAttribute("VERSION", buildProperties.getVersion());
        return "index.html";
    }

    
}
