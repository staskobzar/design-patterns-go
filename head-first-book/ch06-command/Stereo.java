public class Stereo {
  String location;

  public Stereo(String location) {
    this.location = location;
  }

  public void on() {
    System.out.println(location + " Stereo is ON");
  }

  public void setCD() {
    System.out.println(location + " Stereo is set for CD input");
  }

  public void setVolume(int volume) {
    System.out.println(location + " Stereo volume set to " + volume);
  }

  public void off() {
    System.out.println(location + " Stereo is OFF");
  }
}
