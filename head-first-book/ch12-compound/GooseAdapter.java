public class GooseAdapter implements Quackable {
  Observable observable;
  Goose goose;

  public GooseAdapter(Goose goose) {
    this.goose = goose;
    observable = new Observable(this);
  }

  public void quack() {
    goose.honk();
    notifyObservers();
  }

  public void registerObserver(Observer observer) {
    observable.registerObserver(observer);
  }

  public void notifyObservers() {
    observable.notifyObservers();
  }
}
