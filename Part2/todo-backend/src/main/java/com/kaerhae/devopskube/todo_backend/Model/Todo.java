package com.kaerhae.devopskube.todo_backend.Model;

import com.fasterxml.jackson.annotation.JsonProperty;

import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;

@Entity
public class Todo {
    @Id
    @GeneratedValue(strategy=GenerationType.AUTO)
    private long id;
    @JsonProperty("content")
    private String content;

    /* FOR JPA */
    public Todo() {
        super();
    }

    public Todo(@JsonProperty("content") String content) {
        this.content = content;
    }

    public long getId() {
        return id;
    }

    public String getContent() {
        return content;
    }

    public void setContent(String content) {
        this.content = content;
    }
}
