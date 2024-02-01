### Complexity task1

Определите самый маленький и самый большой элемент

```
v1[]={3,5,8,6,2,7,10,11}

min=2
max=11
```

в этом решении получается O(n)

# пример реализации этой версии задачи на Cobol

```cobol
IDENTIFICATION DIVISION.
PROGRAM-ID. ArrayMinMax.
DATA DIVISION.
WORKING-STORAGE SECTION.
01 ARRAY-SIZE PIC 9(4).
01 ARRAY OCCURS 100 TIMES PIC 9(4).
01 MIN-VALUE PIC 9(4).
01 MAX-VALUE PIC 9(4).
01 I PIC 9(4).

PROCEDURE DIVISION.
    ACCEPT ARRAY-SIZE.
    PERFORM VARYING I FROM 1 BY 1 UNTIL I > ARRAY-SIZE
        ACCEPT ARRAY(I)
    END-PERFORM.

    MOVE ARRAY(1) TO MIN-VALUE.
    MOVE ARRAY(1) TO MAX-VALUE.

    PERFORM VARYING I FROM 2 BY 1 UNTIL I > ARRAY-SIZE
        IF ARRAY(I) < MIN-VALUE
            MOVE ARRAY(I) TO MIN-VALUE
        END-IF.
        IF ARRAY(I) > MAX-VALUE
            MOVE ARRAY(I) TO MAX-VALUE
        END-IF.
    END-PERFORM.

    DISPLAY 'Min Value: ' MIN-VALUE.
    DISPLAY 'Max Value: ' MAX-VALUE.

    STOP RUN.

```

# объяснения для себя:

В COBOL, `PERFORM VARYING` является структурой цикла, которая позволяет выполнять набор операций или команд некоторое количество раз с изменением значения переменной в каждой итерации. Давайте разберем конструкцию `PERFORM VARYING` в вашем коде:

```cobol
PERFORM VARYING I FROM 1 BY 1 UNTIL I > ARRAY-SIZE
    ACCEPT ARRAY(I)
END-PERFORM.
```

1. `VARYING I FROM 1 BY 1 UNTIL I > ARRAY-SIZE`: Эта часть устанавливает начальное значение переменной `I` в 1, и в каждой итерации увеличивает `I` на 1, пока `I` не превысит значение, указанное в `ARRAY-SIZE`. Таким образом, цикл будет выполняться от 1 до `ARRAY-SIZE`.

2. `ACCEPT ARRAY(I)`: В каждой итерации цикла происходит операция `ACCEPT`, которая предназначена для ввода данных с терминала (клавиатуры). Здесь `ARRAY(I)` используется для ввода значения и сохранения его в `I`-том элементе массива `ARRAY`.

3. `END-PERFORM`: Это ключевое слово завершает блок цикла `PERFORM VARYING`.

Итак, данная конструкция означает, что вы вводите значения с клавиатуры и сохраняете их в массив `ARRAY`. Цикл выполняется столько раз, сколько указано в переменной `ARRAY-SIZE`.
