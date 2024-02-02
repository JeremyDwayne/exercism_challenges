public class Lasagna {
  public int expectedMinutesInOven() {
    return 40;
  }

  public int remainingMinutesInOven(int cookedTime) {
    return expectedMinutesInOven() - cookedTime;
  }

  public int preparationTimeInMinutes(int layers) {
    return layers * 2;
  }

  public int totalTimeInMinutes(int layers, int cookedTime) {
    return preparationTimeInMinutes(layers) + (expectedMinutesInOven() - remainingMinutesInOven(cookedTime));
  }
}
