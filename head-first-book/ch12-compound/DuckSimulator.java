public class DuckSimulator {
  public static void main(String[] args) {
    DuckSimulator simulator = new DuckSimulator();
    AbstractDuckFactory duckFactory = new CountingDuckFactory();

    simulator.simulate(duckFactory);
  }

  void simulate(AbstractDuckFactory duckFactory) {
    Quackable redheadDuck = duckFactory.createRedheadDuck();
    Quackable duckCall = duckFactory.createDuckCall();
    Quackable rubberDuck = duckFactory.createRubberDuck();
    Quackable gooseDuck = new GooseAdapter(new Goose());

    System.out.println("\nDuck Simulator: With Composite - Flocks");

    Flock flock = new Flock();
    flock.add(redheadDuck);
    flock.add(duckCall);
    flock.add(rubberDuck);
    flock.add(gooseDuck);

    Flock flockOfMallards = new Flock();
    Quackable mallardOne = duckFactory.createMallardDuck();
    Quackable mallardTwo = duckFactory.createMallardDuck();
    Quackable mallardThree = duckFactory.createMallardDuck();
    Quackable mallardFour = duckFactory.createMallardDuck();

    flockOfMallards.add(mallardOne);
    flockOfMallards.add(mallardTwo);
    flockOfMallards.add(mallardThree);
    flockOfMallards.add(mallardFour);

    flock.add(flockOfMallards);

    System.out.println("\nDuck Simulator: Whole Flock Simulation");
    simulate(flock);

    System.out.println("\nDuck Simulator: Mallard Flock Simulation");
    simulate(flockOfMallards);

    System.out.println("\nDuck Simulator: With Observer");
    QuacklogistObserver observer = new QuacklogistObserver();
    flock.registerObserver(observer);

    simulate(flock);

    System.out.println("\nThe ducks quacked " + QuackCounterDecorator.getQuacks() + " times");
  }

  void simulate(Quackable duck) {
    duck.quack();
  }
}
