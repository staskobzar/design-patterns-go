public class WeatherStation {
  public static void main(String[] args) {
    WeatherData weatherData = new WeatherData();
    CurrentConditionsDisplay currentDisplay = new CurrentConditionsDisplay(weatherData);

    currentDisplay.update(80, 65, 30.4f);
    currentDisplay.update(82, 70, 29.2f);
    currentDisplay.update(78, 90, 29.2f);
  }
}
