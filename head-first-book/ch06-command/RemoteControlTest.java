public class RemoteControlTest { // client
  void main(String[] args) {
    SimpleRemoteControl remote = new SimpleRemoteControl(); // invoker

    Light light = new Light("Living room"); // receiver

    LightOnCommand lightOn = new LightOnCommand(light);

    remote.setCommand(lightOn);
    remote.buttonWasPressed();
  }
}
