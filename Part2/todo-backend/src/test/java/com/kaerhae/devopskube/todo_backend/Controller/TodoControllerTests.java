package com.kaerhae.devopskube.todo_backend.Controller;

import java.util.ArrayList;
import java.util.List;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.Mockito;
import org.mockito.MockitoAnnotations;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders;
import org.springframework.test.web.servlet.result.MockMvcResultMatchers;
import org.springframework.test.web.servlet.setup.MockMvcBuilders;

import com.kaerhae.devopskube.todo_backend.Model.Todo;
import com.kaerhae.devopskube.todo_backend.Repository.TodoRepository;

public class TodoControllerTests {
    @Mock
    private TodoRepository todoRepository;

    @InjectMocks
    private TodoController todoController;

    private MockMvc mockMvc;

    @BeforeEach
	void setup() {
        MockitoAnnotations.openMocks(this);
		mockMvc = MockMvcBuilders.standaloneSetup(todoController).build();
	}
    @Test
    public void TodoController_GetReturnsTodos() throws Exception {
        Todo t = new Todo("Test Content");
        List<Todo> todos = new ArrayList<>();
        todos.add(t);
        Mockito.when(todoRepository.findAll()).thenReturn(todos);

        this.mockMvc
        .perform(MockMvcRequestBuilders.get("/api/todo").accept(MediaType.APPLICATION_JSON))
        .andExpect(MockMvcResultMatchers.status().isOk());
    }
}
