public class Soy extends CondimentDecorator {
  Beverage beverage;

  public Soy(Beverage beverage) {
    this.beverage = beverage;
  }

  public double cost() {
    return this.beverage.cost() + 0.15;
  }

  public String getDescription() {
    return this.beverage.getDescription() + ", Soy";
  }
}
