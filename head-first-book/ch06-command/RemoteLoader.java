public class RemoteLoader {
  public static void main(String[] args) {
    RemoteControl remoteControl = new RemoteControl();

    CeilingFan ceilingFan = new CeilingFan("Living Room");

    CeilingFanMediumCommand ceilingFanMedium = new CeilingFanMediumCommand(ceilingFan);
    CeilingFanHighCommand ceilingFanHigh = new CeilingFanHighCommand(ceilingFan);
    CeilingFanOffCommand ceilingFanOff = new CeilingFanOffCommand(ceilingFan);

    remoteControl.setCommand(0, ceilingFanMedium, ceilingFanOff);
    remoteControl.setCommand(1, ceilingFanHigh, ceilingFanOff);

    remoteControl.onButtonWasPushed(0);
    remoteControl.offButtonWasPushed(0);
    System.out.println(remoteControl);
    remoteControl.undoButtonWasPushed();

    remoteControl.onButtonWasPushed(1);
    System.out.println(remoteControl);
    remoteControl.undoButtonWasPushed();
  }

  public static void main0(String[] args) {
    RemoteControl remoteControl = new RemoteControl();

    Light livingRoomLight = new Light("Living Room");
    Light kitchenLight = new Light("Kitchen");
    CeilingFan ceilingFan = new CeilingFan("Living Room");
    GarageDoor garageDoor = new GarageDoor("");
    Stereo stereo = new Stereo("Living Room");

    LightOnCommand livingRoomLightOn = new LightOnCommand(livingRoomLight);
    LightOffCommand livingRoomLightOff = new LightOffCommand(livingRoomLight);
    LightOnCommand kitchenLightOn = new LightOnCommand(kitchenLight);
    LightOffCommand kitchenLightOff = new LightOffCommand(kitchenLight);

    CeilingFanOnCommand ceilingFanOn = new CeilingFanOnCommand(ceilingFan);
    CeilingFanOffCommand ceilingFanOff = new CeilingFanOffCommand(ceilingFan);

    GarageDoorUpCommand garageDoorUp = new GarageDoorUpCommand(garageDoor);
    GarageDoorDownCommand garageDoorDown = new GarageDoorDownCommand(garageDoor);

    StereoOnWithCDCommand stereoOnWithCD = new StereoOnWithCDCommand(stereo);
    StereoOffCommand stereoOff = new StereoOffCommand(stereo);

    remoteControl.setCommand(0, garageDoorUp, garageDoorDown);
    remoteControl.setCommand(1, livingRoomLightOn, livingRoomLightOff);
    remoteControl.setCommand(2, kitchenLightOn, kitchenLightOff);
    remoteControl.setCommand(3, ceilingFanOn, ceilingFanOff);
    remoteControl.setCommand(4, stereoOnWithCD, stereoOff);

    System.out.println(remoteControl);

    remoteControl.onButtonWasPushed(1);
    remoteControl.offButtonWasPushed(1);
    System.out.println(remoteControl);
    remoteControl.undoButtonWasPushed();
    remoteControl.offButtonWasPushed(1);
    remoteControl.onButtonWasPushed(1);
    System.out.println(remoteControl);
    remoteControl.undoButtonWasPushed();

    remoteControl.onButtonWasPushed(1);
    remoteControl.offButtonWasPushed(1);
    remoteControl.onButtonWasPushed(2);
    remoteControl.offButtonWasPushed(2);
    remoteControl.onButtonWasPushed(3);
    remoteControl.offButtonWasPushed(3);
    remoteControl.onButtonWasPushed(4);
    remoteControl.offButtonWasPushed(4);
    remoteControl.onButtonWasPushed(5);
  }
}
