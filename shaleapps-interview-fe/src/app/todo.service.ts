import { Injectable } from '@angular/core';
import { TodoModel } from './todo.model';
import { Observable, of } from 'rxjs';
import { switchMap } from 'rxjs/operators';
import { HttpClient } from '@angular/common/http';

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
    return this.http.put('todos', todo);
  }
}
