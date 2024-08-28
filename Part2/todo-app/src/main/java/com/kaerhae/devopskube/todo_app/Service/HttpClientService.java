package com.kaerhae.devopskube.todo_app.Service;

import java.util.ArrayList;
import java.util.List;


import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.core.ParameterizedTypeReference;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestClient;

import com.kaerhae.devopskube.todo_app.Model.Todo;

@Service
public class HttpClientService {

    public List<Todo> getTodos() throws Exception {
        String url = String.format("http://%s/api/todo", System.getenv("TODO_BACKEND_BASEURL"));
        RestClient restClient = RestClient.create();
        try {
            List<Todo> todos = restClient.get()
                .uri(url)
                .retrieve()
                .body(new ParameterizedTypeReference<>() {});
                return todos;
        } catch (Exception e) {
            return new ArrayList<>();
        }                
    }

    public void postTodos(Todo todo) throws Exception {
        String url = String.format("http://%s/api/todo", System.getenv("TODO_BACKEND_BASEURL"));

        RestClient restClient = RestClient.create();
        try {
            restClient.post()
                .uri(url)
                .body(todo)
                .retrieve()
                .toBodilessEntity();
        } catch (Exception e) {
            System.out.println(e.getMessage());
        }                
    }



}
