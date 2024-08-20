package com.kaerhae.devopskube.todo_app.Controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.info.BuildProperties;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.GetMapping;


@RestController
public class IndexController {
    @Autowired
    private BuildProperties buildProperties;
    @GetMapping("")
    public String GetServerInfo() {
        return String.format("REST API Server version %s", buildProperties.getVersion());
    }

    
}
