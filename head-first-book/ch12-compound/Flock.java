import java.util.*;

// Composite pattern: Flock is a composite of Quackable objects
public class Flock implements Quackable {
  ArrayList<Quackable> quackers = new ArrayList<>();

  public void add(Quackable quacker) {
    quackers.add(quacker);
  }

  public void quack() {
    Iterator<Quackable> iterator = quackers.iterator();

    while (iterator.hasNext()) {
      Quackable quacker = iterator.next();
      quacker.quack();
    }

    notifyObservers();
  }

  public void registerObserver(Observer observer) {
    Iterator<Quackable> iterator = quackers.iterator();
    while (iterator.hasNext()) {
      Quackable duck = iterator.next();
      duck.registerObserver(observer);
    }
  }

  public void notifyObservers() {
  }
}
