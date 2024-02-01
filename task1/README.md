### Task1

Определите самый маленький и самый большой элемент

```
v1[]={3,5,8,6,2,7,10,11}

min=2
max=11
```

в этом решении получается O(2n), но это все равно что и O(n). Просто коэффициенты и константы умножения не учитываются. Поэтому O(2n) теоретически эквивалентно O(n).

# пример реализации этой версии задачи на Cobol

```cobol
       IDENTIFICATION DIVISION.
       PROGRAM-ID. ArrayMinMax.
       DATA DIVISION.
       WORKING-STORAGE SECTION.
       01  ARRAY-SIZE     PIC 9(4).
       01  ARRAY          OCCURS 100 TIMES PIC 9(4).
       01  MIN-VALUE      PIC 9(4).
       01  MAX-VALUE      PIC 9(4).

       PROCEDURE DIVISION.
           ACCEPT ARRAY-SIZE.

           PERFORM VARYING I FROM 1 BY 1 UNTIL I > ARRAY-SIZE
               ACCEPT ARRAY(I)
           END-PERFORM.

           PERFORM FIND-MIN-AND-MAX.

           DISPLAY 'Min Value: ' MIN-VALUE.
           DISPLAY 'Max Value: ' MAX-VALUE.

           STOP RUN.

       FIND-MIN-AND-MAX SECTION.
           MOVE ARRAY(1) TO MIN-VALUE.
           MOVE ARRAY(1) TO MAX-VALUE.

           PERFORM VARYING I FROM 2 BY 1 UNTIL I > ARRAY-SIZE
               IF ARRAY(I) < MIN-VALUE
                   MOVE ARRAY(I) TO MIN-VALUE
               END-IF
               IF ARRAY(I) > MAX-VALUE
                   MOVE ARRAY(I) TO MAX-VALUE
               END-IF
           END-PERFORM.
```

# объяснения для себя:

Переписывание кода с Go на COBOL будет включать в себя несколько изменений, так как структура и синтаксис COBOL отличаются от Go. Вот пример, как можно было бы переписать ваш код на COBOL:

1. `IDENTIFICATION DIVISION / PROGRAM-ID:` Определяют идентификацию программы.

2. `DATA DIVISION / WORKING-STORAGE SECTION:` Здесь определены переменные и рабочая память программы.

3. `PROCEDURE DIVISION:` Здесь расположен основной код программы.

4. `PERFORM VARYING:` Цикл для ввода значений в массив.

5. `FIND-MIN-AND-MAX SECTION:` Секция для поиска минимального и максимального значений в массиве.

6. `DISPLAY:` Вывод результатов.

7. `STOP RUN:` Остановка выполнения программы.

`ARRAY          OCCURS 100 TIMES PIC 9(4).` - Это объявление массива `ARRAY` с размером 100 элементов. Каждый элемент массива имеет формат `PIC 9(4)` и также может содержать четыре десятичных разряда.

```cobol
       ACCEPT ARRAY-SIZE.

       PERFORM VARYING I FROM 1 BY 1 UNTIL I > ARRAY-SIZE
           ACCEPT ARRAY(I)
       END-PERFORM.

       PERFORM FIND-MIN-AND-MAX.

       DISPLAY 'Min Value: ' MIN-VALUE.
       DISPLAY 'Max Value: ' MAX-VALUE.

       STOP RUN.

   FIND-MIN-AND-MAX SECTION.
       MOVE ARRAY(1) TO MIN-VALUE.
       MOVE ARRAY(1) TO MAX-VALUE.

       PERFORM VARYING I FROM 2 BY 1 UNTIL I > ARRAY-SIZE
           IF ARRAY(I) < MIN-VALUE
               MOVE ARRAY(I) TO MIN-VALUE
           END-IF
           IF ARRAY(I) > MAX-VALUE
               MOVE ARRAY(I) TO MAX-VALUE
           END-IF.
```

1. `ACCEPT ARRAY-SIZE.`: Здесь вводится размер массива (`ARRAY-SIZE`). `ACCEPT` используется для ввода данных с клавиатуры.

2. `PERFORM VARYING I FROM 1 BY 1 UNTIL I > ARRAY-SIZE`: Запускается цикл `PERFORM`, который будет повторяться столько раз, сколько указано в `ARRAY-SIZE`. В каждой итерации цикла выполняется команда `ACCEPT ARRAY(I)`, которая запрашивает ввод числа и сохраняет его в `ARRAY(I)`.

3. `PERFORM FIND-MIN-AND-MAX.`: Здесь вызывается подпрограмма (секция) `FIND-MIN-AND-MAX`, которая находит минимальное и максимальное значения в массиве.

4. `DISPLAY 'Min Value: ' MIN-VALUE.` и `DISPLAY 'Max Value: ' MAX-VALUE.`: Выводятся минимальное и максимальное значения на экран.

5. `STOP RUN.`: Завершает выполнение программы.

6. `FIND-MIN-AND-MAX SECTION.`: Это начало секции, где определены действия для поиска минимального и максимального значений в массиве.

7. `MOVE ARRAY(1) TO MIN-VALUE.` и `MOVE ARRAY(1) TO MAX-VALUE.`: Инициализируют `MIN-VALUE` и `MAX-VALUE` первым элементом массива.

8. `PERFORM VARYING I FROM 2 BY 1 UNTIL I > ARRAY-SIZE`: Запускает цикл, начиная со второго элемента массива. Для каждого элемента, если он меньше текущего `MIN-VALUE`, он становится новым `MIN-VALUE`. Если он больше текущего `MAX-VALUE`, он становится новым `MAX-VALUE`.

Таким образом, весь код написан для ввода массива, нахождения его минимального и максимального значений, а затем вывода этих значений на экран.
