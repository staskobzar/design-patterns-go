public class Main {
  public static void main(String[] args) {
    Character queen = new QueenCharacter();
    Character king = new KingCharacter();
    Character troll = new TrollCharacter();
    Character knight = new KnightCharacter();

    System.out.println("Queen:");
    queen.fight();

    System.out.println("\nTroll:");
    troll.fight();

    System.out.println("\nKnight:");
    knight.fight();

    System.out.println("\nTroll:");
    troll.fight();

    System.out.println("\nKing:");
    king.fight();

    System.out.println("\nTroll is defeated and now has a bow!");
  }
}
