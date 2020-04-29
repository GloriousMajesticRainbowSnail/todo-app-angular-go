import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { TodoModel } from '../todo.model';

@Component({
  selector: 'app-todo-item',
  templateUrl: './todo-item.component.html',
  styleUrls: ['./todo-item.component.scss'],
})
export class TodoItemComponent implements OnInit {
  @Input()
  todo: TodoModel;

  @Output()
  toggled = new EventEmitter<TodoModel>();

  @Output()
  removed = new EventEmitter<TodoModel>();

  constructor() {}

  ngOnInit(): void {}

  toggleCompleted() {
    const todo = new TodoModel(this.todo);
    todo.IsComplete = !todo.IsComplete;
    this.toggled.emit(todo);
  }

  removeTodo() {
    this.removed.emit(this.todo);
  }
}
