public class CountingDuckFactory extends AbstractDuckFactory {
  public Quackable createMallardDuck() {
    return new QuackCounterDecorator(new MallardDuck());
  }

  public Quackable createRedheadDuck() {
    return new QuackCounterDecorator(new RedheadDuck());
  }

  public Quackable createDuckCall() {
    return new QuackCounterDecorator(new DuckCall());
  }

  public Quackable createRubberDuck() {
    return new QuackCounterDecorator(new RubberDuck());
  }
}
