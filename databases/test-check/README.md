# Description

Your professor gave the class a bonus task: Write a program that will check the answers for the latest test. The program will be given a table **answers** with the following columns:

- **id**: the unique ID of the question;
- **correct_answer**: the correct answer to the question, given as a string;
- **given_answer**: the answer given to the question, which can be NULL.

Your task is to return the table with a column `id` and a column `checks`, where for each **answers** `id` the following string should be returned:

- `"no answer"` if the `given_answer` is empty;
- `"correct"` if the `given_answer` is the same as the correct_answer;
- `"incorrect"` if the `given_answer` is not empty and is incorrect.
Order the records in the answer table by `id`.

### Example

For the following table **answers**

| **id** | **correct_answer** | **given_answer** |
|:------:|:------------------:|:----------------:|
|    1   |          a         |         a        |
|    2   |          b         |       NULL       |
|    3   |          c         |         b        |

the output should be:

| **id** | **correct_answer** | **given_answer** |
|:------:|:------------------:|:----------------:|
|    1   |          a         |         a        |
|    2   |          b         |       NULL       |
|    3   |          c         |         b        |