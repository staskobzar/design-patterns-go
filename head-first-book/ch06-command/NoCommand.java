public class NoCommand implements Command {
  public void execute() {
    System.out.println("No Command set");
  }

  public void undo() {
    System.out.println("No Command set");
  }
}
