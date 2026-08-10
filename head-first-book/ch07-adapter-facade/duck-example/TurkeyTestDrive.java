public class TurkeyTestDrive {
  public static void main(String[] args) {
    MallardDuck mallard = new MallardDuck();

    WildTurkey turkey = new WildTurkey();
    Duck turkeyAdapter = new TurkeyAdapter(turkey);

    System.out.println("The Turkey says...");
    turkey.gobble();
    turkey.fly();

    System.out.println("\nThe Duck says...");
    testDuck(mallard);

    System.out.println("\nThe TurkeyAdapter says...");
    testDuck(turkeyAdapter);
  }

  static void testDuck(Duck duck) {
    duck.quack();
    duck.fly();
  }
}
