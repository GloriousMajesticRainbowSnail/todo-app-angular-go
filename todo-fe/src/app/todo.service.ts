import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { TodoModel } from './todo.model';

@Injectable({
  providedIn: 'root',
})
export class TodoService {
  constructor(private http: HttpClient) {}

  getAllTodos(): Observable<TodoModel[]> {
    return this.http.get<TodoModel[]>('todos');
  }

  createTodo(todo: TodoModel) {
    return this.http.post<TodoModel>('todos', todo);
  }

  updateTodo(todo: TodoModel) {
    return this.http.put<TodoModel>('todos', todo);
  }

  removeTodo(todo: TodoModel) {
    return this.http.delete<TodoModel>(`todos?id=${todo.Id}`);
  }
}
