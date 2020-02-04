# Machine Learning - how Skynet was born.

The Epic of Gilgamesh.

"Life, which you look for, you will never find. For when the gods created man, they let death be his share, and life withheld in their own hands"

## 11/23/2019: Dimensionally reduction: 

It is often a good idea to try to reduce the dimension of your training data using a "Dimensionality Reduction algorithm" before you feed it to another Machine Learning algorithm (such as a supervised learning algorithm). It will run much faster, the data will take up less disk and memory space, and in some cases it may also perform better. 

A simple example might be associating truck age with truck mileage. Truck age and truck mileage would collapse into a single attribute composed of age, mileage, or some other transformation algorithm that accepts both as input.

Workbook Example 1:
How would you simplify this data with these attributes: 

Truck Info:
* vin
* engine 
* model
* make
* category - e.g. box truck, van, pickup
* year
* mileage
* total hours borrowered 
* average speed
* accident count
* failure count
* service count

Potiential Solution:
The service count, accident count, and failure count may be correlated.  Combine and transform the 3 attributes into a single attribute. Name it accident-failure-service-count this is just the concatenation of service, accident, and failure count into a simple string value.

### Before you begin the coding exercises.
* https://www.tensorflow.org/install/lang_c

### Dependencies 
* go - v1.13.3
* tensorflow

### Build command
```
$ export SSH_PRIVATE_KEY="$(cat ../id_rsa)"
$ make build
$ make push
```