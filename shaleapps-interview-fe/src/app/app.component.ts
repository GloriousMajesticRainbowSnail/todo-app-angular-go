import { Component } from '@angular/core';
import { TodoModel } from './todo.model';
import { TodoService } from './todo.service';
import { Observable, of } from 'rxjs';
import { switchMap } from 'rxjs/operators';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent {
  todo$: Observable<TodoModel[]>;

  public constructor(private service: TodoService) {
    this.refreshTodos(of(null));
  }

  refreshTodos(change: Observable<any>) {
    this.todo$ = change.pipe(switchMap(() => this.service.getAllTodos()));
  }

  onTodoCreated(description: string) {
    this.refreshTodos(
      this.service.createTodo(new TodoModel({ Description: description }))
    );
  }

  onTodoToggled(todo: TodoModel) {
    this.refreshTodos(this.service.updateTodo(todo));
  }
}
