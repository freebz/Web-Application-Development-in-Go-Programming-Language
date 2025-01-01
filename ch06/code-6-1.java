// 코드 6.1 자바로 표현한 Person 클래스를 상속하는 Korea 클래스

class Person {
  String name;
  int age;
}

// Person 클래스를 상속한 Korean 클래스
class Korean extends Person {
  int myNumber;
}

class Main {
  // Person 클래스를 인수로 사용하는 메서드
  public static void hello(Person p) {
    System.out.println("Hello " + p.name);
  }
}