import { Component, OnInit, EventEmitter, Output } from '@angular/core';
import { TodoModel } from '../todo.model';

@Component({
  selector: 'app-todo-creator',
  templateUrl: './todo-creator.component.html',
  styleUrls: ['./todo-creator.component.scss'],
})
export class TodoCreatorComponent implements OnInit {
  @Output()
  created = new EventEmitter<TodoModel>();

  description = '';

  constructor() {}

  ngOnInit(): void {}

  todoCreated() {
    this.created.emit(new TodoModel({ Description: this.description }));
    this.description = '';
  }
}
