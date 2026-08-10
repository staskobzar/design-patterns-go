public class PrepareBeverages {
  public static void main(String[] args) {
    TeaBeverage tea = new TeaBeverage();
    CoffeeBeverage coffee = new CoffeeBeverage();

    System.out.println("Making tea...");
    tea.prepareRecipe();

    System.out.println("\nMaking coffee...");
    coffee.prepareRecipe();
  }
}
