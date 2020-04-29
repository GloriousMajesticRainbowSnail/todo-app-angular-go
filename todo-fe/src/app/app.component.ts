import { Component } from '@angular/core';
import { TodoModel } from './todo.model';
import { TodoService } from './todo.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent {
  todoList: TodoModel[] = [];

  public constructor(private service: TodoService) {
    this.getAllTodos();
  }

  getAllTodos() {
    this.service.getAllTodos().subscribe((todos) => (this.todoList = todos));
  }

  onTodoCreated(todo: TodoModel) {
    this.service.createTodo(todo).subscribe(() => this.getAllTodos());
  }

  onTodoToggled(todo: TodoModel) {
    this.service.updateTodo(todo).subscribe(() => this.getAllTodos());
  }

  onTodoRemoved(todo: TodoModel) {
    this.service.removeTodo(todo).subscribe(() => this.getAllTodos());
  }
}
