Тестовое задание по Go
===

## Текст задания
Тестовое задание
Сделать сервис на ​golang,​ который будет выполнять арифметические операции над числами.
У http сервиса должно быть 4 api метода (add, sub, mul, div), каждый из которых выполняет соответствующую арифметическую операцию (сложения, вычитания, умножения и деления).
Каждый API метод на вход получает 2 get параметра - числа, с которыми надо выполнить операцию.
Оба параметра должны валидироваться - что они числа, что присутствуют только 2 параметра и они правильно названы
Код сервиса не должен падать (например деление на 0)

## Запуск

* Запуск тестов:  
``
make test
``
* Запуск через docker:    
``
make build && make run
``
* Запуск локально:    
``
make run_local
``

## Примеры запросов
Request
```
http://0.0.0.0:8080/api/mul?a=9&b=2
```
Response
```
{
  "Success": true,
  "ErrCode": "",
  "Value": 18
}
```

## Примечания к решению
 * Решение выполнение в соотвествии с принципами Clean Architecture
 * Настроен CI, в котором запускаются тесты, используя GitHub Actions
 * Используется context для возможности ограничения времени исполнения операций, а также для передачи в нем данных впоследствии (если бы оно понадобилось)
