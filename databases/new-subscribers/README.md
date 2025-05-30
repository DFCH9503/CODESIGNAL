# Description

You are managing a small newspaper subscription service. Anyone who uses it can subscribe to a large number of different newspapers for a full year or just a half year.

The information about subscriptions is stored in the **full_year** and **half_year** tables, which have the following structures:

- **full_year**:
    - **id**: the unique subscription ID;
    - **newspaper**: the newspaper’s name;
    - **subscriber**: the name of the subscriber.
- **half_year**:
    - **id**: the unique subscription ID;
    - **newspaper**: the newspaper’s name;
    - **subscriber**: the name of the subscriber.

Given the **full_year** and **half_year** tables, compose the result as follows: The resulting table should have one column `subscriber` that contains all the distinct names of anyone who is subscribed to a newspaper with the word `Daily` in its name. The table should be sorted in ascending order by the subscribers’ first names.

### Example

For the following tables **full_year**

| **id** |    **newspaper**   | **subscriber** |
|:------:|:------------------:|:--------------:|
|    1   | The Paragon Herald |   Crissy Sepe  |
|    2   | The Daily Reporter |  Tonie Moreton |
|    3   |  Morningtide Daily |  Erwin Chitty  |
|    4   |   Daily Breakfast  |  Tonie Moreton |
|    5   | Independent Weekly |   Lavelle Phu  |

and **half_year**

| **id** |    **newspaper**   | **subscriber** |
|:------:|:------------------:|:--------------:|
|    1   | The Paragon Herald |   Crissy Sepe  |
|    2   | The Daily Reporter |  Tonie Moreton |
|    3   |  Morningtide Daily |  Erwin Chitty  |
|    4   |   Daily Breakfast  |  Tonie Moreton |
|    5   | Independent Weekly |   Lavelle Phu  |

the output should be:

| **subscriber** |
|:--------------:|
|  Erwin Chitty  |
|    Joya Buss   |
|   Lavelle Phu  |
|    Lia Cover   |
|  Tonie Moreton |