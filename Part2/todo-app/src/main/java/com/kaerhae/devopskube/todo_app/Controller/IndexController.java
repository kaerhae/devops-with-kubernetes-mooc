package com.kaerhae.devopskube.todo_app.Controller;

import java.net.http.HttpClient;
import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.info.BuildProperties;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;

import com.kaerhae.devopskube.todo_app.Model.Todo;
import com.kaerhae.devopskube.todo_app.Service.HttpClientService;




@Controller
public class IndexController {
    @Autowired
    private BuildProperties buildProperties;
    @Autowired
    private HttpClientService httpClientService;


    @GetMapping("")
    public String GetServerInfo(Model model) throws Exception {
        //return String.format("REST API Server version %s", buildProperties.getVersion());
        Todo todo = new Todo();
        List<Todo> todos = httpClientService.getTodos();
        model.addAttribute("VERSION", buildProperties.getVersion());
        model.addAttribute("IMAGE_URL", String.format("%s/image.jpg", System.getenv("IMAGE_FILE_PATH")));
        model.addAttribute("todo", todo);
        model.addAttribute("todos", todos);
        return "index.html";
    }
    
}
