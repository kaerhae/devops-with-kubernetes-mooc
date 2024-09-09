package com.kaerhae.devopskube.todo_backend.Controller;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

import com.kaerhae.devopskube.todo_backend.Model.Todo;
import com.kaerhae.devopskube.todo_backend.Repository.TodoRepository;

@RestController
public class TodoController {

    @Autowired
    private TodoRepository todoRepository;

    @GetMapping("/api/todo")
    public List<Todo> getTodos() {
        return todoRepository.findAll();
    }

    @PostMapping("/api/todo")
    public ResponseEntity<String> postTodo(@RequestBody Todo entity) {
        Todo res = todoRepository.save(entity);
        if (res != null) {
            return new ResponseEntity<>("Todo successfully created", HttpStatus.CREATED);
        }

        return new ResponseEntity<>("Error while creating todo", HttpStatus.BAD_REQUEST);
        
    }
    
    
}
