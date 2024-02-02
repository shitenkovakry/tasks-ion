### Task3

Ищите наличие введенного с клавиатуры числа и определите, сколько раз оно повторяется в векторе.

```
v1[]={1,3,4,3,8}

число 3 в v[2], число 3 в v[4]
```

# пример реализации задачи на Cobol

```cobol
IDENTIFICATION DIVISION.
PROGRAM-ID. SimilarElements.

DATA DIVISION.
WORKING-STORAGE SECTION.
01 ARRAY-SIZE PIC 9(4).
01 ARRAY OCCURS 100 TIMES PIC 9(4).
01 SIM-VALUE PIC 9(4).
01 ARRAY-OF-INDEXES OCCURS 100 TIMES PIC 9(4).
01 INDEX PIC 9(4).
01 ARRAY-OF-RESULTS OCCURS 100 TIMES PIC 9(4).

PROCEDURE DIVISION.
PERFORM INPUT-ARRAY.
PERFORM FIND-SIMILAR-ELEMENT.
PERFORM GET-INDEXES-OF-SIMILAR-VALUES.
PERFORM DISPLAY-RESULTS.
STOP RUN.

INPUT-ARRAY.
DISPLAY "Enter the size of the array:".
ACCEPT ARRAY-SIZE.

    DISPLAY "Enter the elements of the array:".
    PERFORM VARYING INDEX FROM 1 BY 1 UNTIL INDEX > ARRAY-SIZE
        ACCEPT ARRAY(INDEX)
    END-PERFORM.

FIND-SIMILAR-ELEMENT.
PERFORM VARYING INDEX FROM 1 BY 1 UNTIL INDEX > ARRAY-SIZE
SET SIM-VALUE TO ARRAY(INDEX)
PERFORM VARYING NEXT-INDEX FROM INDEX + 1 BY 1 UNTIL NEXT-INDEX > ARRAY-SIZE
IF ARRAY(NEXT-INDEX) = SIM-VALUE
MOVE SIM-VALUE TO ARRAY-OF-RESULTS(1)
EXIT PERFORM
END-IF
END-PERFORM
END-PERFORM.

GET-INDEXES-OF-SIMILAR-VALUES.
PERFORM VARYING INDEX FROM 1 BY 1 UNTIL INDEX > ARRAY-SIZE
IF ARRAY(INDEX) = SIM-VALUE
ADD 1 TO ARRAY-OF-INDEXES(1)
MOVE INDEX TO ARRAY-OF-INDEXES(ARRAY-OF-INDEXES(1))
END-IF
END-PERFORM.

DISPLAY-RESULTS.
DISPLAY "Similar Value: " SIM-VALUE.
DISPLAY "Indexes of Similar Values: ".
PERFORM VARYING INDEX FROM 2 BY 2 UNTIL INDEX > ARRAY-OF-INDEXES(1)
DISPLAY ARRAY-OF-INDEXES(INDEX)
END-PERFORM.

```

# объяснения для себя:

В данном участке кода на COBOL происходит поиск одинаковых элементов в массиве. Давайте разберем каждую часть:

1. `PERFORM VARYING INDEX FROM 1 BY 1 UNTIL INDEX > ARRAY-SIZE`: Это цикл, который перебирает индексы массива от 1 до `ARRAY-SIZE`.

2. `SET SIM-VALUE TO ARRAY(INDEX)`: Устанавливает переменной `SIM-VALUE` значение элемента массива, находящегося по текущему индексу `INDEX`.

`SET SIM-VALUE TO ARRAY(INDEX)` в языке COBOL означает присвоение значения переменной `SIM-VALUE` элементу массива, находящемуся по индексу `INDEX`.

Давайте рассмотрим более подробно:

- `SIM-VALUE`: это переменная, которой присваивается значение.

- `ARRAY(INDEX)`: это выражение, которое обращается к элементу массива `ARRAY` с индексом `INDEX`.

Таким образом, операция `SET SIM-VALUE TO ARRAY(INDEX)` означает "присвоить переменной `SIM-VALUE` значение элемента массива `ARRAY`, который расположен по индексу `INDEX`".

3. `PERFORM VARYING NEXT-INDEX FROM INDEX + 1 BY 1 UNTIL NEXT-INDEX > ARRAY-SIZE`: Вложенный цикл, который начинается с индекса `INDEX + 1` и перебирает следующие индексы массива.

4. `IF ARRAY(NEXT-INDEX) = SIM-VALUE`: Проверка, равен ли элемент массива с индексом `NEXT-INDEX` значению `SIM-VALUE`.

5. `MOVE SIM-VALUE TO ARRAY-OF-RESULTS(1)`: Если элементы равны, то значение `SIM-VALUE` помещается в массив результатов `ARRAY-OF-RESULTS` (по всегда первому индексу).

6. `EXIT PERFORM`: Выход из вложенного цикла после нахождения одинаковых элементов. Таким образом, при первом совпадении значение `SIM-VALUE` помещается в `ARRAY-OF-RESULTS(1)` и происходит выход из цикла.
