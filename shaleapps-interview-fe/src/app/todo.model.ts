export class TodoModel {
  public Id: string;
  public Description: string;
  public IsComplete = false;

  // Allows a Todo to be initialized with a partial object
  public constructor(init?: Partial<TodoModel>) {
    Object.assign(this, init);
  }
}
