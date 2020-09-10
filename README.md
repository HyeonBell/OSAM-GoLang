# OSAM-GoLANG

# 개요
2019 OSAM Github Training에서 배운 Github를 다루는 것을 익히기 위해 GoLang Repo를 만듬.

- go 실행 방법
go run [실행 goLang].go

func : 함수 선언 키워드
콘솔 입출력 제공 : fmt
콘솔 출력 함수 : println, print
*fmt패키지를 import 하지 않아도 기본적으로 println과 print함수 지원*

var 변수이름 변수형 == 함수 밖에서(전역변수) 선언시 꼭 var 키워드 사용
var a, b int
별다른 형 선언 없이 타입 추론 == Short Assignment Satement ":="
이 방법은 func 내에서만 사용이 가능
*:= 선언에는 var 키워드가 들어가지 않음*

초기값을 설정하지 않으면 'Zero value'로 설정
Bool = false
숫자 타입 = 0
string 타입 = ""(빈문자열)
*Go언어에서는 선언만 하고 쓰지 않았다면 에러가 발생하여 컴파일 실패*
	* 
이 규칙은 변수, 패키지, 함수 등 모든 선언에서 동일
	* 
메모리를 아무 이유없이 차지하는 변수들에 대해 굉장히 단호
	* 
꼭 쓰이는 변수만 선언후 값을 지울 때 선언한 모든 부분을 지워야함.



주석
// 단일행 주석
/* */ 다중행 주석

변수 여러개 선언 가능 
 var a,b int = 10, 20
	* 
초기화하는 값의 갯수가 동일해야하며
	* 
초기화 하지 않는다면 모든 값을 초기화하지 않아야함.



상수 = 한번 초기화 되면 그 이후에 수정할 수 없음.
'const' 키워드 사용 == "const 상수이름 상수형"
= 꼭 선언과 동시에 초기화
= 초기화 이후에 사용하지 않아도 에러 x
= 타입추론을 사용할 수없어서 ':='용법사용 불가
변수의 고유값으로 초기화할 때 사용

= 괄호를 이용해 여러개의 값을 묶어서 초기화 가능 == 각각 상수는 개행해서 초기화하고 콤마를 입력하지 않음 
*iota라는 식별자를 값으로 초기화 하면 그 후에 초기화 하지 않고 이어지는 상수들은 순서(index)가 값으로 저장

수식 연산자
+ 더하기
- 빼기
* 곱하기
/ 나누기
% 나머지

증감 연산자
++ 1증가
-- 1감소

할당 연산자
= 변수나 상수에 값을 대입
:= 선언 및 대입
+= 더한 후 대입
-= 뺀후 대입
*= 곱한후 대입
/= 나눈후 대입
%= 나눗셈후 나머지를 대입
&= 값의 AND 비트 연산 후 대입
|= 값의 OR 비트 연산 후 대입
^= 값의 XOR 비트 연산 후 대입
&^= 값의 AND NOT 비트 연산 후 대입
<<= 비트를 왼쪽으로 이동 후 대입
>>= 비트를 오른쪽으로 이동 후 대입

논리 연산자
&& A B 모두 참이면 연산 결과 참 반환
|| AB  둘중하나라도 참이면 연산결과 참 반환
! A가 참이면 거짓 거짓이면 참 반환

연산자 우선순위


콘솔 입력함수 == fmt.Scanln(&num1, &num2, &num3)
*printf함수는 C언어처럼 포맷 출력변환기호를 사용하여 출력가능 print는 그냥 싹다 출력*

자료형의 종류와 특징
어떤 데이터를 저장할지 표현하는 것이 자료형
부울린 타입, 문자열 타입, 정수형 타입, 실수 타입, 복소수 타입, 기타 타입
':' 용법을 활용한 자료형 추론

몇바이트인지 확인 하기
import "unsafe"
unsafe.Sizeof(변수)

부울린 
bool 1

정수형(음수포함) *32비트 시스템에서 4바이트(32비트) 64시스템에서 8바이트(64비트)
int n비트 시스템에서 n비트 크기
int8 1
int16 2
int32 4
int64 8

정수형(0, 양수 포함)
uint n비트 시스템에서 n비트
uint8 1
uint16 2
uint32 4
uint64 8
uintptr 8

실수
float32 4
float64 8

복소수
complex64 8
complex128 16

문자열 *string으로 선언한 문자열 타입은 immutable 타입으로서 값을 수정할 수 없다.*
ex) var str string = "hello"로 선언하고 str[2] = 'a'로 수정불가
string 16

기타 타입
byte 1 == uint8과 똑같은 자료형 == 부호 없는 정수 값과 구별하는 데 사용
rune  4 == int32와 똑같은 자료형 == 관례상 문자 값을 정수 값과 구별하기 위해 사용

강제 형변환
타입(변수)

콘솔출력함수 : Println, Print, Printf, Fprintln,Fprint, Fprintf, Sprintln, Sprint, Sprintf
*Pirint 같은 함수는 반환값을 갖지않고 S함수는 string 형으로 출력*

서식 문자
원하는 출력폭 == %(폭)d
출력 폭에 0을 채워넣고 싶으면 == %0(폭)d
소수점 자리수를 설정할 때 == %.(자릿수)f


Scanln 함수 입력을 공백으로만 구분 == 개행이 되면 입력이 완료
 = Scanf는 개발자가 만들어놓은 형식으로 입력 가능 
ex) Scanf("%d-%d", &num1, &num2)로 만들었을때 940101-000000이라고입력하면 num1은 940101 num2는 00000이됨


Go언어에서는 while문을 제공하지 않아 for문만 사용


Go언어의 문법과 다르게 조건문은 몇 가지 엄격하다.
	* 
조건문의 조건식은 반드시 Boolean 형으로 표현
	* 
bool 형은 false와 true만 지원
	* 
if 조건식의 괄호는 생략가능
	* 
조건문의 중괄호는 필수



switch문
= 조건에 따른 흐름 분기
= 태그나 표현식이 어느 조건에도 맞지 않는다면 default문을 사용해 해당 구문을 수행
= break를 따로 입력하지 않아도 해당되는 case만 수행
= 쓰임새
	* 
switch에 전달되는 인자로 태그 사용
	* 
switch에 전달되는 인자로 표현식 사용
	* 
switch에 전달되는 인자 없이 case에 표현식 사용(참/거짓 판별)



break, continue, goto

두개이상의 변수를 모아 놓은 것을 '컬렉션' 이라고함
배열 == 정적 == 고정된 배열 크기안에 동일한 데이터를 연속적으로 저장
배열의 선언 "var 배열이름 [배열크기]자료형"
*Go언어에서는 배열의 크기는 자료형을 구성하는 한 요소"
= 선언 후 크기를 바꿀 수 있는 것이 아님.

슬라이스(Slice) == 고정된 크기를 미리 지정하지 않고 이후에 필요에 따라 크기를 동적으로 변경, 부분 발췌 가능
	* 
다차원 선언
	* 
배열과 똑같은 구현 가능
	* 
배열의 여러 제약점들을 넘어 여러 값을 다룰 때 개발자에게 주로 쓰임
	* 
슬라이스는 지금까지 배운 자료형과 내부적인 구조가 다르기 때문에 선언 및 초기화를 할 때 주의



var a []int를 선언한다면 배열의 일부분을 가리키는 포인터만 만듬
= 선언만 하고 초기화하지 않아서 슬라이스 정보만 있는 배열만 생성
= 실질적으로 어떠한 변수가 들어갈 공간은 생성되지 않음
= 슬라이스는 크기를 미리 지정하지 않아 컴퓨터가 어디서부터 어디까지 0이나 ""로 채워야하는지 알 수 없음.
Nil silce == 초기값을 지정하지 않고 선언만한 슬라이스 == 용량도 크기도 없는 상태

 = ptr 배열의 위치를 가리킴
 = len 배열의 길이
 = cap 전체 크기
var a []int = []int{1,2,3,4} 는 슬라이스를 선언함과 동시에 1,2,3,4를 위한 메모리를 만든다는 뜻

슬라이스를 복사해온다는 것은 사실 같은 주소를 참조한다는 것.
l = a[2:5]를 입력하는것은 a의 인덱스2요소부터 4요소까지 참조하는 것


*슬라이스의 길이와 용량을 지정하지 않고 슬라이스만 선언하면 nil과 if문으로 비교 가능*


make()함수를 이용한 슬라이스 선언 == 값을 저장할 수 있는 메모리를 선언만 함으로써 생성하는법
 = make(슬라이스 타입, 슬라이스 길이, 슬라이스 용량)
 = 용량은 생략해서 선언 가능
	* 
길이 : 초기화된 슬라이스의 요소 개수 == len(컬렉션이름)
	* 
용량 : 슬라이스는 배열의 길이가 동적으로 늘어날 수 있기 때문에 길이와 용량을 구분 == 용량은 승객을 태울수 있는 량 cap(컬렉션이름)
	* 
주의할 점은 make로 생성한후 []int{1,2,3,4}와 같은 형식으로 초기화하면 이전 값은 새로운 메모리를 할당하면서 사라짐


package main import "fmt" func main() { s := make([]int, 0, 3) // len=0, cap=3 인 슬라이스 선언 for i := 1; i <= 10; i++ { // 1부터 차례대로 한 요소씩 추가 s = append(s, i) fmt.Println(len(s), cap(s)) // 슬라이스 길이와 용량 확인 } fmt.Println(s) // 최종 슬라이스 출력 }

슬라이스 추가, 병합, 복사
append()함수를 이용해서 슬라이스에 데이터 추가
= 용량이 남는경우 슬라이스의 길이를 변경하여 데이터 추가
= 용량이 초과하는 경우 설정한 용량만큼 새로운 배열을 생성하고 기존 배열 값들을 모두 새 배열의 복제한 후 다시 슬라이스를 할당
= copy(붙여넣을 슬라이스, 복사할 슬라이스)
*슬라이스에 슬라이스를 추가해서 붙일 수 있음.*

슬라이스의 부분만 자르기
붙여넣을 슬라이스 := 복사할 슬라이스[복사할 첫 인덱스: 복사할 마지막 인덱스+1]
	* 
마지막 요소는 복사하지 않음



슬라이스는 배열과 다르게 값을 복사해오는 것이 아니라 슬라이스 자체가 참조하고 있는 주소값을 같이 참조하는 것을 의미

Map == Slice와 배열 공통점은 인덱스 번호를 임의대로 바꿀 수 없음 == map은 key와 value를 매핑해 저장하는 hash table
 = Ruby의 hash와 python의 dict과 같은 기능
 = key: value 형식으로 저장
 = 슬라이스와 맵의 공통점 == 값을 저장하지 않고 참조하는 타입

맵의 선언 == "var 맵이름 map[key자료형]value자료형 "
ex) var a map[int]string == 아무것도 초기화하지 않으면 Nil map

map 변수의 추가, 갱신, 삭제
 = 할당(make()함수 혹은 사용할 값 초기화)됐으면 값을 추가, 갱신, 삭제 가능
추가방식 == 맵이름[key] = value
*이미있는 key값에 다시 다른 value값을 저장한다면 최근 저장한 값으로 갱신*
삭제 == delete(맵이름,key) == 해당 eky값에 해당하는 value값도 같이삭제

Map의 key 체크와 value 읽기
맵은 맵이름[key]의 key에 저장돼 있는 value 값을 반환할 뿐만 아니라, 해당 키에 값이 존재하는지 안 하는지 판별하는 true/false로 반환 == 키값이 사용되고 있다면 true 아니라면 false 반환 == 존재하지 않는 key값을 입력했다면 0이나 ""

값 반환받기. true/false 반환받기
val,exist := 맵이름[key]형식으로 입력해야 val에는 value값이 exist에는 true/false 값이 초기화
value 값만 입력 받고 싶다면 val:= 맵이름[key]형식 == 꼭 두개의 값을 반환하는 것은 아님
= true/false 값만 반환 받고 싶다면, "_,bool변수" 형식을 선언해 할당 받기
	* 
ex) _,exist := 맵이름[key]



프로그래밍에 있어서 설계(Design)가 상당히 중요
= 함수는 특정 기능을 위해 만든 여러 문장을 묶어서 실행하는 코드 블록 단위
= 프로그램의 특정 기능들을 기능별로 묶어 구현해 놓은 것
선언 형식 : func 함수이름 (매개변수이름 매개변수형) 반환형
	* 
반환형이 괄호 뒤에 명시, 매개변수형도 매개변수이름 뒤에 명시
	* 
함수는 패키지안에서 정의되고 호출되는 함수가 꼭 호출하는 함수 앞에 있을 필요는 없다
	* 
반환 값이 여러개 일 수 있다 == (반환형1, 반환형2) == 두형이 같더라도 두번써야함
	* 
블록 시작 브레이스가 함수선언과 동시에 첫줄에 위치해야한다



함수 호출 == 함수이름(전달자이름)

전역변수와 지역변수
매개변수 = 값자체를 전달하는 방식 = 값의 주소를 전달하는 방식
매개변수에 전달하려는 변수가 어떤 유형의 변수이냐에 따라 사용 방법과 결과가 다름

지역변수 == 중괄호 {} 안에서 선언된 변수 == 선언된 지역내에서 유효 = 코드가 지역을 벗어나면 삭제
전역변수 == 중괄호 밖에서 선언된 변수 == 코드가 시작되어 선언되는 순간 메모리가 생성되고 코드 전체가 끝날때까지 메모리를 차지.
두가지 차이점
	* 
메모리에 존재하는 시간
	* 
변수에 접근할 수 있는 범위



매개변수 == 언어를 불문하고 개발자들이 개발을 하면서 많이 헷갈려하는 부분
= 코드가 길어지고 많은 변수를 사용하다보면 변수가 전역변수인지 지역변수인지 헷갈리게됨 == 에러가 발생하면 그 이유가 매개변수 사용에서의 실수일 가능성이 높음
*변수의 사용이 헷갈리는 것은 변수의 개념이 헷갈린다는 뜻이고 결국 함수 활용을 제대로 할 수 없는것*

Pass by value, Pass by reference, 가변인자
Pass by value == 인자 값을 복사해서 전달하는 방식

Pass by reference == Go언어에서는 C/C++ 언어에서 핵심 개념인 '포인터'라는 개념을 지원
 = & 는 주소
 = *는 직접 참조
C언어에서는 배열이름 자체가 첫번째 인덱스 요소의 주솟값 == 주솟값은 어떤 변수 앞에 &를 붙이는 것
C언어에서는 *(배열이름 + 인덱스)는 배열이름[인덱스]와 같은 기능 == Go에는 없음 == 직접 참조를 원하면 포인터 변수앞에 *를 붙이는 것
*함수를 호출할 때 주솟값 전달을위해 함수이름(&변수이름), 매개변수이름을 입력할 때는 직접 참조하기 위해 *를 매개변수형 앞에 붙임 == 함수 안에서는 매개변수앞에 모두 *를 붙임

가변 인자 함수 == 전달하는 매개변수의 개수를 고정한 함수가 아니라 함수를 호출할 때마다 매개변수의 개수를 다르게 전달할 수 있도록 만든 함수
*기능이 다르지만 비슷한 개념으로 Java에서 메소드 '오버로딩' 이 있음*
Go언어의 가변 함수는 동일한 형의 매개변수를 n개 전달할 수 있음. == 숫자를 더하는 함수를 만들 때 더하는 값의 개수에 따라 각각 함수를 만들 필요가 없음
	* 
n개의 동일한 형의 매개변수를 전달
	* 
전달된 변수들은 슬라이스 형태이다. 변수를 다룰 때 슬라이스를 다루는 방법과 동일
	* 
선언 형식 : func 함수이름(매개변수이름 ...매개변수형) 반환형
	* 
매개변수로 슬라이스 전달 가능 == 함수이름(슬라이스이름...) == 슬라이스이름에 ...붙이기



반환값 == Go언어에서는 복수개의  반환값을 반환 할 수 있음
	* 
반환값의 개수만큼 반환형을 명시해야함 == 2개 이상의 반환형을 입력할 때 괄호()에 명시
	* 
동일한 반환형이더라도 모두 병시 == (int, int, int)



Named Return Parameter == 이름이 붙여진 반환 인자
= (반환값이름1 반환형1, 반환값이름2 반환형2, 반환값이름3 반환형3, ...) 형식으로 입력
"반환값이름 반환형" 자체가 변수 선언
	* 
returnd을 생략하면 안됨 반드시 return 명시
	* 
반환 값이 하나라도 반환 값이름을 명시했다면 괄호안에 써야함.



익명 함수 == 이름이 없는 함수
*코드를 작성할 때 아무런 규칙 없이 마구잡이로 작성하는 것보다 코드의 기능별 '함수화'하는 것이 굉장히 중요*
	* 
단점은 프로그램 속도 저하
	* 
함수 선언 자체가 프로그래밍 전역으로 초기화되면서 메모리를 잡아먹기 때문
	* 
기능을 수행할 때마다 함수를 찾아가서 호출해야함


= 이러한 단점을 보완하기 위해 '익명 함수'가 필요
= 함수의 이름만 없고 그 외 형태는 동일 == 함수의 마지막 브레이스 } 뒤에 괄호 ()를 사용해 함수 바로 호출 == 안에 매개변수를 넣을 수 있음
	* 
함수의 기능적인 요소만 빼서 가볍게 활용하기 위해 사용



변수에 초기화한 익명 함수는 변수 이름을 함수의 이름처럼 사용 가능
선언함수와 익명함수는 프로그램의 내부적으로 읽는 순서가 다름
	* 
선언함수는 프로그램이 시작됨과 동시에 읽힘
	* 
익명함수는 그 자리에서 실행되기 떄문에 해당 함수가 실행되는 곳에서 읽음
	* 
선언 함수보다 익명 함수가 나중에 읽힘



일급 함수 (First-Class Function)
익명함수의 사용은 Go언어에서의 함수가 '일급 함수'이기 때문
	* 
일급함수라는 의미는 함수를 기본 타입과 동일하게 사용할 수 있어 함수 자체를 다른 함수의 매개변수로 전달하거나 다른 함수의 반환 값으로 사용될 수 있다는것 == 함수는 다른 타입들과 비교했을 때 높은 수준의 용법이 아니라 같은 객체로서 사용 가능
	* 
매개변수로 사용, 변수에 초기화 가능, 함수가 함수의 반환값으로 사용되는 것은 클로저 에서
	* 
매개변수형으로 사용될 때 형식 : 매개변수함수이름 func(전달받는함수의매개변수형) 전달받는함수의반환형



type문을 사용한 함수 원형 정의
	* 
매개변수로 전달하는 함수의 가독성일 떨어지는 문제를 해결
	* 
type문을 통해 원형을 정의하고 정의한 이름을 사용



외부 변수 접근 : 클로저 = 함수 안에서 익명 함수를 정의해서 바깥쪽 함수에 선언한 변수에도 접근할 수 있는 함수
 함수 안에서 바깥 변수를 사용하려면 매개 변수를 사용해 Pass by value 형식이나 Pass by reference 형식으로 사용해야함 == 그런데 익명함수는 클로저 이기 때문에 외부 함수의 변수를 그냥 접근 가능
 = 함수안에서 함수를 반환하는 것

for 초기식; 조건식; 조건 변화식 { 반복 수행할 구문 }
for 초기식; 조건식; 조건 변화식 { 반복 수행할 구문 }
