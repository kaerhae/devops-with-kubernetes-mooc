package com.kaerhae.devopskube.todo_backend.Service;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.ArrayList;
import java.util.List;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.TestPropertySource;
import org.springframework.test.util.ReflectionTestUtils;

import com.kaerhae.devopskube.todo_backend.Model.Todo;

@SpringBootTest
@TestPropertySource(properties = {
    "spring.datasource.url=jdbc:h2:mem:testdb",
    "spring.datasource.driverClassName=org.h2.Driver",
    "spring.datasource.username=sa",
    "spring.datasource.password=password"
})
public class TodoServiceTests {
    @Autowired
    private TodoService todoService;

    @Test
    public void getTodos_ShouldReturnTodos() 
    {
        Todo t = new Todo("Test Content");
        List<Todo> todos = new ArrayList<>();
        todos.add(t);
        ReflectionTestUtils.setField(todoService, "todos", todos);

        List<Todo> res = todoService.GetTodos();

        assertEquals(todos, res);
    }

    @Test
    public void addTodo_ShouldNotThrowError() throws Exception
    {
        Todo t = new Todo("Test Content");
        List<Todo> todos = new ArrayList<>();
        ReflectionTestUtils.setField(todoService, "todos", todos);

        todoService.AddTodo(t);
    }

    @Test
    public void addTodo_ShouldAddNewTodo() {
        Todo t = new Todo("Test Content");
        List<Todo> expected = new ArrayList<>();
        expected.add(t);
        List<Todo> todos = new ArrayList<>();
        ReflectionTestUtils.setField(todoService, "todos", todos);

        todoService.AddTodo(t);

        List<Todo> result = todoService.GetTodos();
        assertEquals(expected, result);
    }
}
