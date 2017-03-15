# 시작하기

## 1. Go 언어 소개

### 생산성이 높은 이유
* 부분적이지만 편리한 자료형 추론을 제공하여 굳이 반복해서 자료형 이름을 쓰지 않아도 된다
* 소스 코드 형식을 자동으로 맞춰주는 도구 및 여러 편리한 도구가 기본으로 제공된다
* Example 테스트를 이용하여 쉽게 테스트 코드를 작성하면서 코드 문서화까지 할 수 있다
* 함수 리터럴 및 클로저를 자유자재로 사용할 수 있다
* 명시적으로 인터페이스를 지정하지 않아도 인터페이스 구현이 가능하여 기존에 있던 코드를 구치지 않고도 유연한 구현이 가능하다
* 채널을 이용하여 동시성 구현을 락 등을 이용하지 않고 간편하게 할 수 있으며 언어 고유의 지원으로 교착 상태나 경쟁 상태를 파악하기 쉽다
* 컴파일 속도가 빨라서 컴파일 및 테스트를 반복적으로 수행하면서 코드를 작성하기 용이하다
* 가비지 컬렉션 지원으로 메모리 관리에 대한 부담을 덜 수 있다
* 자료형 리터럴을 쉽게 쓸 수 있다

## 2. 첫 프로그램

### 1. Go 놀이터

* https://play.golang.org

### 2. 기본 예제 코드

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, playground")
}
```

## 3. 자료형 및 변수

* 자료형을 정적으로 검사하기에 변수에 자료형이 정해져있음
* 정적 자료형을 지원하지만 자료형을 미리 선언하고 값을 할당하는 번거로움을 벗어나게 해주는 자료형 추론 기능이 있다

### 1. 변수 선언

* 변수 선언

```
# C, C++, Java
int x;

# Go
var x int
```

* 배열 선언

```
# C, C++
int arr[5];

# Java
int[] arr = new int[5];

# Go
var arr [5]int
```

* 포인터

```
var p *int
```

* 변수 선언과 동시에 값 할당

```
var x int = 10
```

### 2. 자료형 추정

* 자료형이 명확한 경우 자료형을 사용하지 않아도 됨

```
var i = 10
var p = &i  // i의 주소값을 p에 저장
```
* `var` 키워드를 삭제하자

```
i := 10
p := &i
```

## 4. 함수와 간단한 제어구조

[ex1_4_1.go](./ex1_4_1.go)
[ex1_4_2.go](./ex1_4_2.go)
[ex1_4_3.go](./ex1_4_3.go)

* `package` : 네임 스페이스 계념으로 보면 될듯
    * 패키지는 하나의 디렉터리에 저장
    * 같은 디렉터리에 여러개의 패키지가 불가능
    * 하나의 패키지를 여러 디렉터리에 나눠 구현 할 수 없다
    * 같은 패키지 이름을 선언하는 모든 `.go` 파일은 하나의 디렉터리에 저장되어야한다

## 5. 마치며
* go 언어의 특징은 다음과 같다
    * 높은 생산성
        * 간결하게 표현 가능한 코드
        * 빠른 컴파일 속도
        * 등 몇 가지 이유
    * 쉬운 동시성 코드 작성 가능
* Go 놀이터 웹 사이트를 이용해보자
    * 간단한 접속으로 가능한 프로그래밍
    * `import` 및 형식을 맞춰주는 기능
* Go 언어의 기본 형식은 다음과 같다
    * 구문 분석기를 통하여 문장 끝에 자동으로 붙는 세미콜론
    * `Package main - import - func main()` 순서
    * 변수 이름 > 자료형 순서의 변수 선언
    * 다른 언어와 비교되는 조건문 반복문의 형식
        * 괄호가 필요 없음
        * 타 언어의 `while` 문과 `for` 문을 `for` 문 하나로 표현 가능