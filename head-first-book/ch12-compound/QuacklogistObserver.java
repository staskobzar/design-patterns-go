public class QuacklogistObserver implements Observer {
  public void update(QuackObservable duck) {
    System.out.println("Quacklogist: " + duck + " just quacked.");
  }
}
