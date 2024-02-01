### Task2

Определите сумму и среднее значение массива

```
v1[]={5,2,7,6,5,3}

sum=28
avg=4.67
```

# пример реализации задачи на Cobol

```cobol
IDENTIFICATION DIVISION.
PROGRAM-ID. ArraySumAndAvg.
DATA DIVISION.
WORKING-STORAGE SECTION.
01 ARRAY-SIZE     PIC 9(4).
01 ARRAY          OCCURS 100 TIMES PIC 9(4).
01 SUMMA          PIC 9(6).
01 AVERAGE        PIC 9(6)V99.
01 I              PIC 9(4).

PROCEDURE DIVISION.
    ACCEPT ARRAY-SIZE.

    PERFORM VARYING I FROM 1 BY 1 UNTIL I > ARRAY-SIZE
        ACCEPT ARRAY(I)
    END-PERFORM.

    PERFORM VARYING I FROM 1 BY 1 UNTIL I > ARRAY-SIZE
        ADD ARRAY(I) TO SUMMA
    END-PERFORM.

    COMPUTE AVERAGE = SUMMA / ARRAY-SIZE.

    DISPLAY 'sum: ' SUMMA.
    DISPLAY 'avg: ' AVERAGE.

    STOP RUN.

```

# объяснения для себя:

1. В строке `AVERAGE = SUMMA / ARRAY-SIZE` была добавлена операция `COMPUTE` для выполнения вычислений с использованием `SUMMA` и `ARRAY-SIZE`.

2. Тип данных для `SUMMA` изменен на `PIC 9(6)`, чтобы вместить сумму значений массива.

3. Тип данных для `AVERAGE` изменен на `PIC 9(6)V99`, чтобы вместить десятичные доли в среднем значении.

4. В COBOL фраза `ADD ARRAY(I) TO SUMMA` означает, что значение, содержащееся в элементе массива `ARRAY(I)`, будет добавлено к переменной `SUMMA`.
