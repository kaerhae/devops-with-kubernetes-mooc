package com.kaerhae.devopskube.todo_backend.Repository;

import java.util.List;

import org.springframework.data.repository.CrudRepository;

import com.kaerhae.devopskube.todo_backend.Model.Todo;

public interface TodoRepository extends CrudRepository<Todo, Long> {

    List<Todo> findAll();
    Todo findById(long id);
    Todo save(Todo todo);
}