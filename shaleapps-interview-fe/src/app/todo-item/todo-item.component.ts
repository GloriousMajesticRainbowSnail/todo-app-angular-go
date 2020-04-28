import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';
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

  constructor() {}

  ngOnInit(): void {}

  toggleCompleted() {
    const todo = new TodoModel(this.todo);
    todo.IsComplete = !todo.IsComplete;
    this.toggled.emit(todo);
  }
}
