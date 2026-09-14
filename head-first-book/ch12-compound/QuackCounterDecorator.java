public class QuackCounterDecorator implements Quackable {
  Quackable duck;
  static int numberOfQuacks;

  public QuackCounterDecorator(Quackable duck) {
    this.duck = duck;
  }

  public void quack() {
    duck.quack();
    numberOfQuacks++;
  }

  static int getQuacks() {
    return numberOfQuacks;
  }

  public void registerObserver(Observer observer) {
    duck.registerObserver(observer);
  }

  public void notifyObservers() {
    duck.notifyObservers();
  }
}
