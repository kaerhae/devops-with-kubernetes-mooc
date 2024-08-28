package com.kaerhae.devopskube.todo_app.Controller;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import com.kaerhae.devopskube.todo_app.Model.Todo;
import com.kaerhae.devopskube.todo_app.Service.HttpClientService;

import org.springframework.web.bind.annotation.PostMapping;



@Controller
public class TodoController {
    @Autowired
    private HttpClientService httpClientService;
    @GetMapping("/action/todo")
    public ResponseEntity<List<Todo>> getTodos() throws Exception {
        return new ResponseEntity<>(httpClientService.getTodos(), HttpStatus.OK);
    }

    @PostMapping("/action/todo")
    public ResponseEntity<String> postMethodName(Todo entity) throws Exception {
        httpClientService.postTodos(entity);
        
        return new ResponseEntity<>("todo created",HttpStatus.CREATED);
    }
    
    
}
