public abstract class Character {
  WeaponBehavior weapon;

  public Character() {
  }

  public void fight() {
    weapon.useWeapon();
  }
}
