# Description

Mr. Cash wants to keep track of his expenses, so he has prepared a list of all the products he bought this month. Now he is interested in finding the product on which he spent the largest amount of money. If there are products that cost the same amount of money, he’d like to find the one with the *lexicographically smallest* name.

*Note: String A is lexicographically smaller than string B either if A is a prefix of B (and A ≠ B), or if there exists such index i (0 ≤ i < min(A.length, B.length)), that Ai < Bi, and for any j (0 ≤ j < i) Aj = Bj. The lexicographic comparison of strings is implemented by operator < in modern programming languages.*

The list of expenses is stored in a table **Products** which has the following columns:

- **id**: unique product id
- **name**: the unique name of the product
- **price**: the price for one item
- **quantity**: the number of items bought


The resulting table should contain one row with a single column: the product with the *lexicographically smallest* name on which Mr. Cash spent the largest amount of money.

The total amount of money spent on a product is calculated as
> price \* quantity.

### Example

- For given table **Products**

| **id** |    **name**   | **price** | **quantity** |
|:------:|:-------------:|:---------:|:------------:|
|    1   |  MacBook Air  |    1500   |       1      |
|    2   |  Magic Mouse  |     79    |       1      |
|    3   | Spray cleaner |     10    |       3      |

the output should be:

|   **name**  |
|:-----------:|
| MacBook Air |

- For given table **Products**

| **id** |  **name**  | **price** | **quantity** |
|:------:|:----------:|:---------:|:------------:|
|    1   |   Tomato   |     10    |       4      |
|    2   |  Cucumber  |     8     |       5      |
|    3   | Red Pepper |     20    |       2      |
|    4   |    Feta    |     40    |       1      |

the output should be:

| **name** |
|:--------:|
| Cucumber |

While the total cost for each product was 40, Cucumber has the lexicographically smallest name.