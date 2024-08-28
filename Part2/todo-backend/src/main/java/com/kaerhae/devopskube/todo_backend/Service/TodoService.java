package com.kaerhae.devopskube.todo_backend.Service;

import java.util.ArrayList;
import java.util.List;
import java.util.UUID;

import org.springframework.stereotype.Service;

import com.kaerhae.devopskube.todo_backend.Model.Todo;

@Service
public class TodoService {
    private List<Todo> todos = new ArrayList<>();
    public List<Todo> GetTodos() {
        return todos;
    }

    public Todo AddTodo(Todo todo) {
        if(todo.getContent() != null) {
            todo.setId(GenUUID());
            todos.add(todo);
        }

        return null;
    }

    private String GenUUID() {
        return UUID.randomUUID().toString();
    }
}
